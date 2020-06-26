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
	item  *[]*selectItem
	from  *fromItem
	where *expression
}

type selectItem struct {
	exp      *expression
	asterisk bool
	as       *token
}

type fromItem struct {
	table *token
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
	binaryKind
)

type expression struct {
	literal *token
	binary  *binaryExpression
	kind    expressionKind
}

type binaryExpression struct {
	a  expression
	b  expression
	op token
}
