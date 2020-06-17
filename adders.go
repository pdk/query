package query

// adder funcs to add elements to particular clauses

func (b Builder) addSelect(item QueryElement) Builder {
	b.selects = append(b.selects, item)
	return b
}

func (b Builder) addFrom(item QueryElement) Builder {
	b.froms = append(b.froms, item)
	return b
}

func (b Builder) addWhere(item QueryElement) Builder {
	b.wheres = append(b.wheres, item)
	return b
}

func (b Builder) addGroup(item QueryElement) Builder {
	b.groups = append(b.groups, item)
	return b
}

func (b Builder) addHaving(item QueryElement) Builder {
	b.havings = append(b.havings, item)
	return b
}

func (b Builder) addOrder(item QueryElement) Builder {
	b.orders = append(b.orders, item)
	return b
}
