package gosql

type Ast struct {
	Statements []*Statement
}

type Statement struct {
	SelectStatement      *SelectStatement
	CreateTableStatement *CreateTableStatement
	InsertStatement      *InsertStatement
	Kind                 AstKind
}

type SelectStatement struct {
	item []*expression
	from token
}

type CreateTableStatement struct {
	name token
	cols *[]*columnDefinition
}

type columnDefinition struct {
	name     token
	datatype token
}

type InsertStatement struct {
	table  token
	values *[]*expression
}

type AstKind uint

const (
	SelectKind AstKind = iota
	CreateTableKind
	InsertKind
)

type expressionKind uint

const (
	literalKind expressionKind = iota
)

type expression struct {
	literal *token
	kind    expressionKind
}
