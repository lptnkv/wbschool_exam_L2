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

func TestCommand(t *testing.T) {

	expect := "Toggle On\n" +
		"Toggle Off\n"

	invoker := &Invoker{}
	receiver := &Receiver{}

	invoker.AddCommand(&ToggleOnCommand{receiver: receiver})
	invoker.AddCommand(&ToggleOffCommand{receiver: receiver})

	result := invoker.Execute()

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}

func TestChainOfResp(t *testing.T) {
	expect := "Reception registering patient\nDoctor checking patient\nMedical giving medicine to patient\nCashier getting money from patient"

	cashier := &Cashier{}

	medical := &Medical{}
	medical.setNext(cashier)

	doctor := &Doctor{}
	doctor.setNext(medical)

	reception := &Reception{}
	reception.setNext(doctor)

	patient := &Patient{name: "Kirill"}
	result := reception.execute(patient)

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.\n", expect, result)
	}
}

func TestFactory(t *testing.T) {
	expect := []string{
		"shot from ak47",
		"shot from musket",
		"shot from simple gun",
	}
	factory := &GunFactory{}
	guns := []IGun{
		factory.CreateGun("ak47"),
		factory.CreateGun("musket"),
		factory.CreateGun("gun"),
	}

	for i, gun := range guns {
		if res := gun.shoot(); res != expect[i] {
			t.Errorf("Expected %s, but got %s\n", expect[i], res)
		}
	}

}
