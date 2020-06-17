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
