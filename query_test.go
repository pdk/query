package query_test

import (
	"testing"

	q "github.com/pdk/query"
)

func TestMakingQueries(t *testing.T) {

	b := q.Select("a", "b as boink", "c").
		From("foo").
		Where("a =", q.Bind(1)).
		Where("x >", q.Bind("fred"))

	sql := b.SQL()

	exp := "select a, b as boink, c from foo where a = $1 and x > $2"
	if sql != exp {
		t.Errorf("expected %s, got >>%s<<", exp, sql)
	}

	vals := b.BindValues()

	if len(vals) != 2 || vals[0].(int) != 1 || vals[1].(string) != "fred" {
		t.Errorf("unexpected bind values: %v", vals)
	}

	qry := q.Select("alpha", "sum(beta) as beta").
		From("foo").
		Where("gamma =", q.Bind(17)).
		GroupBy("alpha").
		Having("sum(beta) >", q.Bind(1000)).
		OrderBy("sum(beta) desc")

	sql = qry.SQL()

	exp = "select alpha, sum(beta) as beta from foo where gamma = $1 group by alpha having sum(beta) > $2 order by sum(beta) desc"

	if sql != exp {
		t.Errorf("expected %s, got %s", exp, sql)
	}

}

func TestWhereBuilding(t *testing.T) {

	b := q.Where("age =", q.Bind(42)).
		OrderBy("name", "city")

	a := q.Select("name", "age", "city").
		From("glorp_bliff")

	sql := a.Merge(b).SQL()

	exp := "select name, age, city from glorp_bliff where age = $1 order by name, city"

	if sql != exp {
		t.Errorf("expected %s, got %s", exp, sql)
	}

	b = q.Where("name like", q.Bind("x%")).
		Where("city like", q.Bind("san f%"))

	sql = a.Merge(b).SQL()

	exp = "select name, age, city from glorp_bliff where name like $1 and city like $2"

	if sql != exp {
		t.Errorf("expected %s, got %s", exp, sql)
	}
}
