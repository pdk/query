package query

import (
	"strconv"
	"strings"
)

type bindValue struct {
	value interface{}
}

func (bv bindValue) sql(sb *strings.Builder, pCounter *int) {
	sb.WriteString("$")
	*pCounter++
	sb.WriteString(strconv.Itoa(*pCounter))
}

func (bv bindValue) bindValues(values []interface{}) []interface{} {
	return append(values, bv.value)
}

// getBindValues gets bind values from all the elements
func getBindValues(stuffs []QueryElement, values []interface{}) []interface{} {

	for _, each := range stuffs {
		values = each.bindValues(values)
	}

	return values
}
