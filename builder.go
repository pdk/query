package query

import "strings"

// query.Builder is basically just a collection of QueryElements

type QueryElement interface {
	sql(*strings.Builder, *int)
	bindValues([]interface{}) []interface{}
}

type Builder struct {
	selects []QueryElement
	froms   []QueryElement
	wheres  []QueryElement
	groups  []QueryElement
	havings []QueryElement
	orders  []QueryElement
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
