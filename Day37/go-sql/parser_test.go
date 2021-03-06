package gosql

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		source string
		ast    *Ast
	}{
		{
			source: "INSERT INTO users VALUES (105, 233)",
			ast: &Ast{
				Statements: []*Statement{
					{
						Kind: InsertKind,
						InsertStatement: &InsertStatement{
							table: token{
								loc:   location{col: 12, line: 0},
								kind:  identifierKind,
								value: "users",
							},
							values: &[]*expression{
								{
									literal: &token{
										loc:   location{col: 26, line: 0},
										kind:  numericKind,
										value: "105",
									},
									kind: literalKind,
								},
								{
									literal: &token{
										loc:   location{col: 32, line: 0},
										kind:  numericKind,
										value: "233",
									},
									kind: literalKind,
								},
							},
						},
					},
				},
			},
		},
		{
			source: "CREATE TABLE users (id INT, name TEXT)",
			ast: &Ast{
				Statements: []*Statement{
					{
						Kind: CreateTableKind,
						CreateTableStatement: &CreateTableStatement{
							name: token{
								loc:   location{col: 13, line: 0},
								kind:  identifierKind,
								value: "users",
							},
							cols: &[]*columnDefinition{
								{
									name: token{
										loc:   location{col: 20, line: 0},
										kind:  identifierKind,
										value: "id",
									},
									datatype: token{
										loc:   location{col: 23, line: 0},
										kind:  keywordKind,
										value: "int",
									},
								},
								{
									name: token{
										loc:   location{col: 28, line: 0},
										kind:  identifierKind,
										value: "name",
									},
									datatype: token{
										loc:   location{col: 33, line: 0},
										kind:  keywordKind,
										value: "text",
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		ast, err := Parse(test.source)
		assert.Nil(t, err, test.source)
		assert.Equal(t, test.ast, ast, test.source)
	}
}
