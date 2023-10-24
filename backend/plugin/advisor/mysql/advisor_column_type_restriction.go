package mysql

// Framework code is generated by the generator.

import (
	"fmt"
	"strings"

	"github.com/pingcap/tidb/parser/ast"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*ColumnTypeRestrictionAdvisor)(nil)
	_ ast.Visitor     = (*columnTypeRestrictionChecker)(nil)
)

func init() {
	advisor.Register(storepb.Engine_MYSQL, advisor.MySQLColumnTypeRestriction, &ColumnTypeRestrictionAdvisor{})
	advisor.Register(storepb.Engine_MARIADB, advisor.MySQLColumnTypeRestriction, &ColumnTypeRestrictionAdvisor{})
	advisor.Register(storepb.Engine_OCEANBASE, advisor.MySQLColumnTypeRestriction, &ColumnTypeRestrictionAdvisor{})
}

// ColumnTypeRestrictionAdvisor is the advisor checking for column type restriction.
type ColumnTypeRestrictionAdvisor struct {
}

// Check checks for column type restriction.
func (*ColumnTypeRestrictionAdvisor) Check(ctx advisor.Context, _ string) ([]advisor.Advice, error) {
	stmtList, ok := ctx.AST.([]ast.StmtNode)
	if !ok {
		return nil, errors.Errorf("failed to convert to StmtNode")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}
	paylaod, err := advisor.UnmarshalStringArrayTypeRulePayload(ctx.Rule.Payload)
	if err != nil {
		return nil, err
	}
	checker := &columnTypeRestrictionChecker{
		level:           level,
		title:           string(ctx.Rule.Type),
		typeRestriction: make(map[string]bool),
	}
	for _, tp := range paylaod.List {
		checker.typeRestriction[strings.ToUpper(tp)] = true
	}

	for _, stmt := range stmtList {
		checker.text = stmt.Text()
		checker.line = stmt.OriginTextPosition()
		(stmt).Accept(checker)
	}

	if len(checker.adviceList) == 0 {
		checker.adviceList = append(checker.adviceList, advisor.Advice{
			Status:  advisor.Success,
			Code:    advisor.Ok,
			Title:   "OK",
			Content: "",
		})
	}
	return checker.adviceList, nil
}

type columnTypeRestrictionChecker struct {
	adviceList      []advisor.Advice
	level           advisor.Status
	title           string
	text            string
	line            int
	typeRestriction map[string]bool
}

type columnTypeData struct {
	table  string
	column string
	tp     string
	line   int
}

// Enter implements the ast.Visitor interface.
func (checker *columnTypeRestrictionChecker) Enter(in ast.Node) (ast.Node, bool) {
	var columnList []columnTypeData
	switch node := in.(type) {
	case *ast.CreateTableStmt:
		for _, column := range node.Cols {
			if _, exist := checker.typeRestriction[strings.ToUpper(column.Tp.CompactStr())]; exist {
				columnList = append(columnList, columnTypeData{
					table:  node.Table.Name.O,
					column: column.Name.Name.O,
					tp:     strings.ToUpper(column.Tp.CompactStr()),
					line:   column.OriginTextPosition(),
				})
			}
		}
	case *ast.AlterTableStmt:
		for _, spec := range node.Specs {
			switch spec.Tp {
			case ast.AlterTableAddColumns:
				for _, column := range spec.NewColumns {
					if _, exist := checker.typeRestriction[strings.ToUpper(column.Tp.CompactStr())]; exist {
						columnList = append(columnList, columnTypeData{
							table:  node.Table.Name.O,
							column: column.Name.Name.O,
							tp:     strings.ToUpper(column.Tp.CompactStr()),
							line:   node.OriginTextPosition(),
						})
					}
				}
			case ast.AlterTableChangeColumn, ast.AlterTableModifyColumn:
				column := spec.NewColumns[0]
				if _, exist := checker.typeRestriction[strings.ToUpper(column.Tp.CompactStr())]; exist {
					columnList = append(columnList, columnTypeData{
						table:  node.Table.Name.O,
						column: column.Name.Name.O,
						tp:     strings.ToUpper(column.Tp.CompactStr()),
						line:   node.OriginTextPosition(),
					})
				}
			}
		}
	}

	for _, column := range columnList {
		checker.adviceList = append(checker.adviceList, advisor.Advice{
			Status:  checker.level,
			Code:    advisor.DisabledColumnType,
			Title:   checker.title,
			Content: fmt.Sprintf("Disallow column type %s but column `%s`.`%s` is", column.tp, column.table, column.column),
			Line:    column.line,
		})
	}

	return in, false
}

// Leave implements the ast.Visitor interface.
func (*columnTypeRestrictionChecker) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}
