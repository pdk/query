package query

import (
	"strconv"
	"strings"
)

// query.Builder is basically just a collection of QueryElements

type QueryElement interface {
	sql(*strings.Builder, *int)
	bindValues([]interface{}) []interface{}
}

type QueryElements []QueryElement

type Builder struct {
	selects QueryElements
	froms   QueryElements
	wheres  QueryElements
	groups  QueryElements
	havings QueryElements
	orders  QueryElements
	limit   int
	offset  int
}

func (b Builder) Merge(other Builder) Builder {
	return Builder{
		selects: b.selects.appendAll(other.selects),
		froms:   b.froms.appendAll(other.froms),
		wheres:  b.wheres.appendAll(other.wheres),
		groups:  b.groups.appendAll(other.groups),
		havings: b.havings.appendAll(other.havings),
		orders:  b.orders.appendAll(other.orders),
		limit:   minInt(b.limit, other.limit),
		offset:  maxInt(b.offset, other.offset),
	}
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
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

	if b.limit > 0 {
		sb.WriteString(" limit ")
		sb.WriteString(strconv.Itoa(b.limit))
	}

	if b.offset > 0 {
		sb.WriteString(" offset ")
		sb.WriteString(strconv.Itoa(b.offset))
	}

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
