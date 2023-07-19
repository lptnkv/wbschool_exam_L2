package pattern

import (
	"testing"
)

func TestFacade(t *testing.T) {
	expect := "Build house\nTree grow\nChild born"

	man := NewMan()

	result := man.Todo()

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}

func TestBuilder(t *testing.T) {
	expect := "<header>Header</header>" +
		"<article>Body</article>" +
		"<footer>Footer</footer>"

	product := &Product{}

	director := Director{&ConcreteBuilder{product}}
	director.Construct()

	result := product.Show()

	if result != expect {
		t.Errorf("Expect result to %s, but %s", result, expect)
	}
}

func TestVisitorHuman(t *testing.T) {
	expect := "Buy sushi...Buy pizza...Buy burger..."

	city := &City{}

	city.Add(&SushiBar{})
	city.Add(&Pizzeria{})
	city.Add(&BurgerBar{})

	result := city.Accept(&Human{})

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}
