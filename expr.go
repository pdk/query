package query

import "strings"

type expr []QueryElement

func (e expr) sql(sb *strings.Builder, pCounter *int) {
	for i, p := range e {
		if i > 0 {
			sb.WriteString(" ")
		}
		p.sql(sb, pCounter)
	}
}

func (e expr) bindValues(values []interface{}) []interface{} {
	for _, p := range e {
		values = p.bindValues(values)
	}
	return values
}
