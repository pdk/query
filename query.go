package query

import (
	"strings"
)

// New constructs a new, empty, query.Builder
func New() Builder {
	return Builder{}
}

// Select starts a new query builder with some select items
func Select(exprs ...interface{}) Builder {
	return New().Select(exprs...)
}

// Select adds select clauses to the query
func (b Builder) Select(exprs ...interface{}) Builder {

	for _, e := range exprs {
		b = b.addSelect(queryElementOf(e))
	}

	return b
}

// From adds from clauses to the query
func (b Builder) From(tableNames ...string) Builder {

	for _, t := range tableNames {
		b = b.addFrom(queryElementOf(t))
	}

	return b
}

// Bind allows inclusion on bind values in a query
func Bind(value interface{}) QueryElement {
	return bindValue{
		value: value,
	}
}

// Expr builds a single expression from elements
func Expr(expression ...interface{}) QueryElement {

	e := expr{}

	for _, p := range expression {
		e = append(e, queryElementOf(p))
	}

	return e
}

// Where adds a where clause, eg:
// .Where("(t.col2 =", Bind(42), "or t.col5 >", Bind(-1), ")")
func (b Builder) Where(expression ...interface{}) Builder {
	return b.addWhere(Expr(expression...))
}

// GroupBy adds order-by expressions
func (b Builder) GroupBy(groupBys ...interface{}) Builder {

	for _, grp := range groupBys {
		b = b.addGroup(queryElementOf(grp))
	}

	return b
}

// OrderBy adds order-by expressions
func (b Builder) OrderBy(orderBys ...interface{}) Builder {

	for _, ord := range orderBys {
		b = b.addOrder(queryElementOf(ord))
	}

	return b
}

// Having adds having conditions
func (b Builder) Having(expression ...interface{}) Builder {
	return b.addHaving(Expr(expression...))
}

// SQL returns the complete SQL from the query
func (b Builder) SQL() string {

	var bindCount int

	sb := &strings.Builder{}

	writeSQL(sb, &bindCount, "select ", ", ", b.selects)
	writeSQL(sb, &bindCount, " from ", ", ", b.froms)
	writeSQL(sb, &bindCount, " where ", " and ", b.wheres)
	writeSQL(sb, &bindCount, " group by ", ", ", b.groups)
	writeSQL(sb, &bindCount, " having ", " and ", b.havings)
	writeSQL(sb, &bindCount, " order by ", ", ", b.orders)

	return sb.String()
}

func writeSQL(sb *strings.Builder, bindCount *int, init, sep string, elements []QueryElement) {

	for i, phrase := range elements {
		if i < 1 {
			sb.WriteString(init)
		} else {
			sb.WriteString(sep)
		}

		phrase.sql(sb, bindCount)
	}
}

// BindValues gathers all the bind values within the query
func (b Builder) BindValues() []interface{} {

	vals := getBindValues(b.selects, nil)
	vals = getBindValues(b.froms, vals)
	vals = getBindValues(b.wheres, vals)
	vals = getBindValues(b.groups, vals)
	vals = getBindValues(b.havings, vals)
	vals = getBindValues(b.orders, vals)

	return vals
}
