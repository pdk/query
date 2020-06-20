package query

// adder funcs to add elements to particular clauses

func (b Builder) addSelect(item QueryElement) Builder {
	b.selects = b.selects.append(item)
	return b
}

func (b Builder) addFrom(item QueryElement) Builder {
	b.froms = b.froms.append(item)
	return b
}

func (b Builder) addWhere(item QueryElement) Builder {
	b.wheres = b.wheres.append(item)
	return b
}

func (b Builder) addGroup(item QueryElement) Builder {
	b.groups = b.groups.append(item)
	return b
}

func (b Builder) addHaving(item QueryElement) Builder {
	b.havings = b.havings.append(item)
	return b
}

func (b Builder) addOrder(item QueryElement) Builder {
	b.orders = b.orders.append(item)
	return b
}

// append does a "safe" append, making a new copy
func (e QueryElements) append(items ...QueryElement) QueryElements {

	size := len(e) + len(items)

	return QueryElements(
		append(append(make([]QueryElement, 0, size), []QueryElement(e)...), items...))
}

// appendAll appends all the other items
func (e QueryElements) appendAll(items QueryElements) QueryElements {
	return e.append(([]QueryElement(items))...)
}
