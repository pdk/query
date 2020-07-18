package query

// Select starts a new query builder with some selected columns
func Select(columnNames ...string) Builder {
	return (Builder{}).Select(columnNames...)
}

// Select add column names (string) to the select clause
func (b Builder) Select(columnNames ...string) Builder {

	for _, name := range columnNames {
		b = b.addSelect(queryElementOf(name))
	}

	return b
}

// SelectExpr adds a single expression to the select clause
func (b Builder) SelectExpr(expression ...interface{}) Builder {
	return b.addSelect(Expr(expression...))
}

// From adds from clauses to the query
func (b Builder) From(tableNames ...string) Builder {

	for _, t := range tableNames {
		b = b.addFrom(queryElementOf(t))
	}

	return b
}

// Where starts a new builder with a where clause
func Where(expression ...interface{}) Builder {
	return Builder{}.Where(expression...)
}

// Where adds a where clause, eg:
// .Where("(t.col2 =", Bind(42), "or t.col5 >", Bind(-1), ")")
func (b Builder) Where(expression ...interface{}) Builder {
	return b.addWhere(Expr(expression...))
}

// GroupBy adds order-by expressions
func (b Builder) GroupBy(groupBys ...string) Builder {

	for _, grp := range groupBys {
		b = b.addGroup(queryElementOf(grp))
	}

	return b
}

// OrderBy adds order-by expressions
func (b Builder) OrderBy(orderBys ...string) Builder {

	for _, ord := range orderBys {
		b = b.addOrder(queryElementOf(ord))
	}

	return b
}

// Having adds a single having conditions
func (b Builder) Having(expression ...interface{}) Builder {
	return b.addHaving(Expr(expression...))
}

// Bind allows inclusion on bind values in a query
func Bind(value interface{}) QueryElement {
	return bindValue{
		value: value,
	}
}

// Binder places an empty bind value
func Binder() QueryElement {
	return bindValue{}
}

// Expr builds a single expression from elements
func Expr(expression ...interface{}) QueryElement {

	e := expr{}

	for _, p := range expression {
		e = append(e, queryElementOf(p))
	}

	return e
}
