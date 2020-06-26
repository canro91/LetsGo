package gosql

import (
	"errors"
	"fmt"
)

func tokenFromKeyword(k keyword) token {
	return token{
		kind:  keywordKind,
		value: string(k),
	}
}

func tokenFromSymbol(s symbol) token {
	return token{
		kind:  symbolKind,
		value: string(s),
	}
}

func expectToken(tokens []*token, cursor uint, t token) bool {
	if cursor >= uint(len(tokens)) {
		return false
	}

	return t.equals(tokens[cursor])
}

func helpMessage(tokens []*token, cursor uint, msg string) {
	var c *token
	if cursor+1 < uint(len(tokens)) {
		c = tokens[cursor+1]
	} else {
		c = tokens[cursor]
	}

	fmt.Printf("[%d,%d]: %s, got: %s\n", c.loc.line, c.loc.col, msg, c.value)
}

func Parse(source string) (*Ast, error) {
	tokens, err := lex(source)
	if err != nil {
		return nil, err
	}

	semicolonToken := tokenFromSymbol(semicolonSymbol)
	if len(tokens) > 0 && !tokens[len(tokens)-1].equals(&semicolonToken) {
		tokens = append(tokens, &semicolonToken)
	}

	a := Ast{}
	cursor := uint(0)
	for cursor < uint(len(tokens)) {
		stmt, newCursor, ok := parseStatement(tokens, cursor, tokenFromSymbol(semicolonSymbol))
		if !ok {
			helpMessage(tokens, cursor, "Expected statement")
			return nil, errors.New("Failed to parse, expected statement")
		}
		cursor = newCursor

		a.Statements = append(a.Statements, stmt)

		atLeastOneSemicolon := false
		for expectToken(tokens, cursor, tokenFromSymbol(semicolonSymbol)) {
			cursor++
			atLeastOneSemicolon = true
		}

		if !atLeastOneSemicolon {
			helpMessage(tokens, cursor, "Expected semi-colon delimiter between statements")
			return nil, errors.New("Missing semi-colon between statements")
		}
	}

	return &a, nil
}

func parseStatement(tokens []*token, initialCursor uint, delimiter token) (*Statement, uint, bool) {
	cursor := initialCursor

	// Look for a SELECT statement
	semicolonToken := tokenFromSymbol(semicolonSymbol)
	slct, newCursor, ok := parseSelectStatement(tokens, cursor, semicolonToken)
	if ok {
		return &Statement{
			Kind:            SelectKind,
			SelectStatement: slct,
		}, newCursor, true
	}

	// Look for a INSERT statement
	inst, newCursor, ok := parseInsertStatement(tokens, cursor, semicolonToken)
	if ok {
		return &Statement{
			Kind:            InsertKind,
			InsertStatement: inst,
		}, newCursor, true
	}

	// Look for a CREATE statement
	crtTbl, newCursor, ok := parseCreateTableStatement(tokens, cursor, semicolonToken)
	if ok {
		return &Statement{
			Kind:                 CreateTableKind,
			CreateTableStatement: crtTbl,
		}, newCursor, true
	}

	return nil, initialCursor, false
}

func parseSelectStatement(tokens []*token, initialCursor uint, delimiter token) (*SelectStatement, uint, bool) {
	var ok bool
	cursor := initialCursor
	_, cursor, ok = parseToken(tokens, cursor, tokenFromKeyword(selectKeyword))
	if !ok {
		return nil, initialCursor, false
	}

	slct := SelectStatement{}

	fromToken := tokenFromKeyword(fromKeyword)
	item, newCursor, ok := parseSelectItem(tokens, cursor, []token{fromToken, delimiter})
	if !ok {
		return nil, initialCursor, false
	}

	slct.item = item
	cursor = newCursor

	whereToken := tokenFromKeyword(whereKeyword)

	_, cursor, ok = parseToken(tokens, cursor, fromToken)
	if ok {
		from, newCursor, ok := parseTokenKind(tokens, cursor, identifierKind)
		if !ok {
			helpMessage(tokens, cursor, "Expected FROM item")
			return nil, initialCursor, false
		}

		slct.from = &fromItem{table: from}
		cursor = newCursor
	}

	_, cursor, ok = parseToken(tokens, cursor, whereToken)
	if ok {
		where, newCursor, ok := parseExpression(tokens, cursor, []token{delimiter}, 0)
		if !ok {
			helpMessage(tokens, cursor, "Expected WHERE conditionals")
			return nil, initialCursor, false
		}

		slct.where = where
		cursor = newCursor
	}

	return &slct, cursor, true
}

func parseSelectItem(tokens []*token, initialCursor uint, delimiters []token) (*[]*selectItem, uint, bool) {
	cursor := initialCursor

	var s []*selectItem
outer:
	for {
		if cursor >= uint(len(tokens)) {
			return nil, initialCursor, false
		}

		current := tokens[cursor]
		for _, delimiter := range delimiters {
			if delimiter.equals(current) {
				break outer
			}
		}

		var ok bool
		if len(s) > 0 {
			_, cursor, ok = parseToken(tokens, cursor, tokenFromSymbol(commaSymbol))
			if !ok {
				helpMessage(tokens, cursor, "Expected comma")
				return nil, initialCursor, false
			}
		}

		var si selectItem
		_, cursor, ok = parseToken(tokens, cursor, tokenFromSymbol(asteriskSymbol))
		if ok {
			si = selectItem{asterisk: true}
		} else {
			asToken := tokenFromKeyword(asKeyword)
			delimiters := append(delimiters, tokenFromSymbol(commaSymbol), asToken)
			exp, newCursor, ok := parseExpression(tokens, cursor, delimiters, 0)
			if !ok {
				helpMessage(tokens, cursor, "Expected expression")
				return nil, initialCursor, false
			}

			cursor = newCursor
			si.exp = exp

			_, cursor, ok = parseToken(tokens, cursor, asToken)
			if ok {
				id, newCursor, ok := parseTokenKind(tokens, cursor, identifierKind)
				if !ok {
					helpMessage(tokens, cursor, "Expected identifier after AS")
					return nil, initialCursor, false
				}

				cursor = newCursor
				si.as = id
			}
		}

		s = append(s, &si)
	}

	return &s, cursor, true
}

func parseFromItem(tokens []*token, initialCursor uint, _ []token) (*fromItem, uint, bool) {
	ident, newCursor, ok := parseTokenKind(tokens, initialCursor, identifierKind)
	if !ok {
		return nil, initialCursor, false
	}

	return &fromItem{table: ident}, newCursor, true
}

/*
INSERT
INTO
$table-name
VALUES
(
$expression [, ...]
)
*/
func parseInsertStatement(tokens []*token, initialCursor uint, delimiter token) (*InsertStatement, uint, bool) {
	cursor := initialCursor
	ok := false

	_, cursor, ok = parseToken(tokens, cursor, tokenFromKeyword(insertKeyword))
	if !ok {
		return nil, initialCursor, false
	}

	_, cursor, ok = parseToken(tokens, cursor, tokenFromKeyword(intoKeyword))
	if !ok {
		helpMessage(tokens, cursor, "Expected into")
		return nil, initialCursor, false
	}

	table, newCursor, ok := parseTokenKind(tokens, cursor, identifierKind)
	if !ok {
		helpMessage(tokens, cursor, "Expected table name")
		return nil, initialCursor, false
	}
	cursor = newCursor

	_, cursor, ok = parseToken(tokens, cursor, tokenFromKeyword(valuesKeyword))
	if !ok {
		helpMessage(tokens, cursor, "Expected VALUES")
		return nil, initialCursor, false
	}

	_, cursor, ok = parseToken(tokens, cursor, tokenFromSymbol(leftParenSymbol))
	if !ok {
		helpMessage(tokens, cursor, "Expected left paren")
		return nil, initialCursor, false
	}

	values, newCursor, ok := parseExpressions(tokens, cursor, tokenFromSymbol(rightParenSymbol))
	if !ok {
		helpMessage(tokens, cursor, "Expected expressions")
		return nil, initialCursor, false
	}
	cursor = newCursor

	_, cursor, ok = parseToken(tokens, cursor, tokenFromSymbol(rightParenSymbol))
	if !ok {
		helpMessage(tokens, cursor, "Expected right paren")
		return nil, initialCursor, false
	}

	return &InsertStatement{
		table:  *table,
		values: values,
	}, cursor, true
}

/*
CREATE
$table-name
(
[$column-name $column-type [, ...]]
)
*/
func parseCreateTableStatement(tokens []*token, initialCursor uint, delimiter token) (*CreateTableStatement, uint, bool) {
	cursor := initialCursor

	if !expectToken(tokens, cursor, tokenFromKeyword(createKeyword)) {
		return nil, initialCursor, false
	}
	cursor++

	if !expectToken(tokens, cursor, tokenFromKeyword(tableKeyword)) {
		return nil, initialCursor, false
	}
	cursor++

	tableName, newCursor, ok := parseTokenKind(tokens, cursor, identifierKind)
	if !ok {
		helpMessage(tokens, cursor, "Expected tableName token")
		return nil, initialCursor, false
	}
	cursor = newCursor

	if !expectToken(tokens, cursor, tokenFromSymbol(leftParenSymbol)) {
		helpMessage(tokens, cursor, "Expected ( token")
		return nil, initialCursor, false
	}
	cursor++

	cols, newCursor, ok := parseColumnDefinitions(tokens, cursor, tokenFromSymbol(rightParenSymbol))
	if !ok {
		return nil, initialCursor, false
	}
	cursor = newCursor

	if !expectToken(tokens, cursor, tokenFromSymbol(rightParenSymbol)) {
		helpMessage(tokens, cursor, "Expected ) token")
		return nil, initialCursor, false
	}
	cursor++

	return &CreateTableStatement{
		name: *tableName,
		cols: cols,
	}, cursor, true
}

func parseExpressions(tokens []*token, initialCursor uint, delimiter token) (*[]*expression, uint, bool) {
	cursor := initialCursor

	var exps []*expression
	for {
		if cursor >= uint(len(tokens)) {
			return nil, initialCursor, false
		}

		current := tokens[cursor]
		if delimiter.equals(current) {
			break
		}

		if len(exps) > 0 {
			var ok bool
			_, cursor, ok = parseToken(tokens, cursor, tokenFromSymbol(commaSymbol))
			if !ok {
				helpMessage(tokens, cursor, "Expected comma")
				return nil, initialCursor, false
			}
		}

		exp, newCursor, ok := parseExpression(tokens, cursor, []token{tokenFromSymbol(commaSymbol), tokenFromSymbol(rightParenSymbol)}, 0)
		if !ok {
			helpMessage(tokens, cursor, "Expected expression")
			return nil, initialCursor, false
		}
		cursor = newCursor

		exps = append(exps, exp)
	}

	return &exps, cursor, true
}

func parseExpression(tokens []*token, initialCursor uint, delimiters []token, minBp uint) (*expression, uint, bool) {
	cursor := initialCursor

	var exp *expression
	_, newCursor, ok := parseToken(tokens, cursor, tokenFromSymbol(leftParenSymbol))
	if ok {
		cursor = newCursor
		rightParenToken := tokenFromSymbol(rightParenSymbol)

		exp, cursor, ok = parseExpression(tokens, cursor, append(delimiters, rightParenToken), minBp)
		if !ok {
			helpMessage(tokens, cursor, "Expected expression after opening paren")
			return nil, initialCursor, false
		}

		_, cursor, ok = parseToken(tokens, cursor, rightParenToken)
		if !ok {
			helpMessage(tokens, cursor, "Expected closing paren")
			return nil, initialCursor, false
		}
	} else {
		exp, cursor, ok = parseLiteralExpression(tokens, cursor)
		if !ok {
			return nil, initialCursor, false
		}
	}

	lastCursor := cursor
outer:
	for cursor < uint(len(tokens)) {
		for _, d := range delimiters {
			_, _, ok = parseToken(tokens, cursor, d)
			if ok {
				break outer
			}
		}

		binOps := []token{
			tokenFromKeyword(andKeyword),
			tokenFromKeyword(orKeyword),
			tokenFromSymbol(eqSymbol),
			tokenFromSymbol(neqSymbol),
			tokenFromSymbol(ltSymbol),
			tokenFromSymbol(lteSymbol),
			tokenFromSymbol(gtSymbol),
			tokenFromSymbol(gteSymbol),
			tokenFromSymbol(concatSymbol),
			tokenFromSymbol(plusSymbol),
		}

		var op *token
		for _, bo := range binOps {
			var t *token
			t, cursor, ok = parseToken(tokens, cursor, bo)
			if ok {
				op = t
				break
			}
		}

		if op == nil {
			helpMessage(tokens, cursor, "Expected binary operator")
			return nil, initialCursor, false
		}

		bp := op.bindingPower()
		if bp < minBp {
			cursor = lastCursor
			break
		}

		b, newCursor, ok := parseExpression(tokens, cursor, delimiters, bp)
		if !ok {
			helpMessage(tokens, cursor, "Expected right operand")
			return nil, initialCursor, false
		}
		exp = &expression{
			binary: &binaryExpression{
				*exp,
				*b,
				*op,
			},
			kind: binaryKind,
		}
		cursor = newCursor
		lastCursor = cursor
	}

	return exp, cursor, true
}

func parseToken(tokens []*token, initialCursor uint, t token) (*token, uint, bool) {
	cursor := initialCursor

	if cursor >= uint(len(tokens)) {
		return nil, initialCursor, false
	}

	if p := tokens[cursor]; t.equals(p) {
		return p, cursor + 1, true
	}

	return nil, initialCursor, false
}

func parseLiteralExpression(tokens []*token, initialCursor uint) (*expression, uint, bool) {
	cursor := initialCursor

	kinds := []tokenKind{identifierKind, numericKind, stringKind, boolKind, nullKind}
	for _, kind := range kinds {
		t, newCursor, ok := parseTokenKind(tokens, cursor, kind)
		if ok {
			return &expression{
				literal: t,
				kind:    literalKind,
			}, newCursor, true
		}
	}

	return nil, initialCursor, false
}

func parseTokenKind(tokens []*token, initialCursor uint, kind tokenKind) (*token, uint, bool) {
	cursor := initialCursor

	if cursor >= uint(len(tokens)) {
		return nil, initialCursor, false
	}

	current := tokens[cursor]
	if current.kind == kind {
		return current, cursor + 1, true
	}

	return nil, initialCursor, false
}

func parseColumnDefinitions(tokens []*token, initialCursor uint, delimiter token) (*[]*columnDefinition, uint, bool) {
	cursor := initialCursor

	cds := []*columnDefinition{}
	for {
		if cursor >= uint(len(tokens)) {
			return nil, initialCursor, false
		}

		// Look for a delimiter
		current := tokens[cursor]
		if delimiter.equals(current) {
			break
		}

		// Look for a comma
		if len(cds) > 0 {
			if !expectToken(tokens, cursor, tokenFromSymbol(commaSymbol)) {
				helpMessage(tokens, cursor, "Expected comma")
				return nil, initialCursor, false
			}

			cursor++
		}

		// Look for a column name
		id, newCursor, ok := parseTokenKind(tokens, cursor, identifierKind)
		if !ok {
			helpMessage(tokens, cursor, "Expected column name")
			return nil, initialCursor, false
		}
		cursor = newCursor

		// Look for a column type
		ty, newCursor, ok := parseTokenKind(tokens, cursor, keywordKind)
		if !ok {
			helpMessage(tokens, cursor, "Expected column type")
			return nil, initialCursor, false
		}
		cursor = newCursor

		cds = append(cds, &columnDefinition{
			name:     *id,
			datatype: *ty,
		})
	}

	return &cds, cursor, true
}
