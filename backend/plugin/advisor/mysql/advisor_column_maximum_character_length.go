package mysql

// Framework code is generated by the generator.

import (
	"fmt"

	"github.com/pingcap/tidb/parser/ast"
	"github.com/pingcap/tidb/parser/mysql"
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/advisor"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

var (
	_ advisor.Advisor = (*ColumnMaximumCharacterLengthAdvisor)(nil)
	_ ast.Visitor     = (*columnMaximumCharacterLengthChecker)(nil)
)

func init() {
	advisor.Register(storepb.Engine_MYSQL, advisor.MySQLColumnMaximumCharacterLength, &ColumnMaximumCharacterLengthAdvisor{})
	advisor.Register(storepb.Engine_MARIADB, advisor.MySQLColumnMaximumCharacterLength, &ColumnMaximumCharacterLengthAdvisor{})
	advisor.Register(storepb.Engine_OCEANBASE, advisor.MySQLColumnMaximumCharacterLength, &ColumnMaximumCharacterLengthAdvisor{})
}

// ColumnMaximumCharacterLengthAdvisor is the advisor checking for max character length.
type ColumnMaximumCharacterLengthAdvisor struct {
}

// Check checks for maximum character length.
func (*ColumnMaximumCharacterLengthAdvisor) Check(ctx advisor.Context, _ string) ([]advisor.Advice, error) {
	stmtList, ok := ctx.AST.([]ast.StmtNode)
	if !ok {
		return nil, errors.Errorf("failed to convert to StmtNode")
	}

	level, err := advisor.NewStatusBySQLReviewRuleLevel(ctx.Rule.Level)
	if err != nil {
		return nil, err
	}
	payload, err := advisor.UnmarshalNumberTypeRulePayload(ctx.Rule.Payload)
	if err != nil {
		return nil, err
	}
	checker := &columnMaximumCharacterLengthChecker{
		level:   level,
		title:   string(ctx.Rule.Type),
		maximum: payload.Number,
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

type columnMaximumCharacterLengthChecker struct {
	adviceList []advisor.Advice
	level      advisor.Status
	title      string
	text       string
	line       int
	maximum    int
}

// Enter implements the ast.Visitor interface.
func (checker *columnMaximumCharacterLengthChecker) Enter(in ast.Node) (ast.Node, bool) {
	var tableName, columnName string
	var line int
	switch node := in.(type) {
	case *ast.CreateTableStmt:
		for _, column := range node.Cols {
			charLength := getCharLength(column)
			if checker.maximum > 0 && charLength > checker.maximum {
				tableName = node.Table.Name.O
				columnName = column.Name.Name.O
				line = column.OriginTextPosition()
				break
			}
		}
	case *ast.AlterTableStmt:
		for _, spec := range node.Specs {
			switch spec.Tp {
			case ast.AlterTableAddColumns:
				for _, column := range spec.NewColumns {
					charLength := getCharLength(column)
					if checker.maximum > 0 && charLength > checker.maximum {
						tableName = node.Table.Name.O
						columnName = column.Name.Name.O
						line = node.OriginTextPosition()
					}
				}
			case ast.AlterTableChangeColumn, ast.AlterTableModifyColumn:
				charLength := getCharLength(spec.NewColumns[0])
				if checker.maximum > 0 && charLength > checker.maximum {
					tableName = node.Table.Name.O
					columnName = spec.NewColumns[0].Name.Name.O
					line = node.OriginTextPosition()
				}
			}
			if tableName != "" {
				break
			}
		}
	}

	if tableName != "" {
		checker.adviceList = append(checker.adviceList, advisor.Advice{
			Status:  checker.level,
			Code:    advisor.CharLengthExceedsLimit,
			Title:   checker.title,
			Content: fmt.Sprintf("The length of the CHAR column `%s` is bigger than %d, please use VARCHAR instead", columnName, checker.maximum),
			Line:    line,
		})
	}

	return in, false
}

// Leave implements the ast.Visitor interface.
func (*columnMaximumCharacterLengthChecker) Leave(in ast.Node) (ast.Node, bool) {
	return in, true
}

func getCharLength(column *ast.ColumnDef) int {
	if column.Tp.GetType() == mysql.TypeString {
		return column.Tp.GetFlen()
	}
	return 0
}
