package query

import (
	"fmt"
	"strings"
)

// stringElement is a very basic element
type stringElement string

// sql of a stringElement is just th string itself
func (s stringElement) sql(sb *strings.Builder, pCounter *int) {
	sb.WriteString(string(s))
}

// bindValues of a string element adds no values
func (s stringElement) bindValues(values []interface{}) []interface{} {
	return values
}

// queryElementOf makes a best effort to discover/convert an item to a QueryElement
func queryElementOf(item interface{}) QueryElement {

	switch el := item.(type) {
	case QueryElement:
		return el
	case string:
		return stringElement(el)
	case fmt.Stringer:
		return stringElement(el.String())
	default:
		return stringElement(fmt.Sprintf("%v", el))
	}
}
