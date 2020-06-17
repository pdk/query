# query

This is a very simple SQL query builder/helper, supporting in-place variable binding.

    import (
        q "github.com/pdk/query"
    )

    qry := q.Select("alpha", "sum(beta) as beta").
        From("foo").
        Where("gamma =", q.Bind(17)).
        GroupBy("alpha").
        Having("sum(beta) >", q.Bind(1000)).
        OrderBy("sum(beta) desc")

    sql = qry.SQL()
    // select alpha, sum(beta) as beta from foo where gamma = $1 group by alpha having sum(beta) > $2 order by sum(beta) desc

    vals = qry.BindValues()
    // [ 17, 1000 ]
