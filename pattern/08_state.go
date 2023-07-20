package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

// Интерфейс состояния вендингового автомата
type State interface {
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}

// Вендинговый автомат
type VendingMachine struct {
	hasItem       State
	itemRequested State
	hasMoney      State
	noItem        State

	currentState State

	itemCount int
	itemPrice int
}

// Конструктор Вендингового автомата
func newVendingMachine(itemCount, itemPrice int) *VendingMachine {
	v := &VendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &HasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &ItemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &HasMoneyState{
		vendingMachine: v,
	}
	noItemState := &NoItemState{
		vendingMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

// Метод запроса товара
func (v *VendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

// Метод добавления товара
func (v *VendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

// Метод внесения денег
func (v *VendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

// Метод выдачи товара
func (v *VendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

// Метод установки состояния автомата
func (v *VendingMachine) setState(s State) {
	v.currentState = s
}

// Метод увеличения кол-ва товаром
func (v *VendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}

// Реализация конкретного состояния - Нет нужного товара
type NoItemState struct {
	vendingMachine *VendingMachine
}

// Метод запроса товара при его отсутствии
func (i *NoItemState) requestItem() error {
	return fmt.Errorf("item out of stock")
}

// Метод добавления товара при его отсутствии
func (i *NoItemState) addItem(count int) error {
	i.vendingMachine.incrementItemCount(count)
	i.vendingMachine.setState(i.vendingMachine.hasItem)
	return nil
}

// Метод оплаты при отсутствии товара
func (i *NoItemState) insertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

// Метод выдачи товара при его отсутствии
func (i *NoItemState) dispenseItem() error {
	return fmt.Errorf("item out of stock")
}

// Реализация конкретного состояния - Товар присутствует
type HasItemState struct {
	vendingMachine *VendingMachine
}

// Метод запроса товара при его наличии
func (i *HasItemState) requestItem() error {
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
		return fmt.Errorf("no item present")
	}
	fmt.Printf("Item requestd\n")
	i.vendingMachine.setState(i.vendingMachine.itemRequested)
	return nil
}

// Метод увеличения кол-ва товара при его наличии
func (i *HasItemState) addItem(count int) error {
	fmt.Printf("%d items added\n", count)
	i.vendingMachine.incrementItemCount(count)
	return nil
}

// Метод оплаты при наличии товара
func (i *HasItemState) insertMoney(money int) error {
	return fmt.Errorf("please select item first")
}

// Метод выдачи товара при его наличии
func (i *HasItemState) dispenseItem() error {
	return fmt.Errorf("please select item first")
}

// Реализация конкретного состояния - Товар выбран
type ItemRequestedState struct {
	vendingMachine *VendingMachine
}

// Метод запроса уже запрошенного товара
func (i *ItemRequestedState) requestItem() error {
	return fmt.Errorf("item already requested")
}

// Метод добавления запрошенного товара
func (i *ItemRequestedState) addItem(count int) error {
	return fmt.Errorf("item Dispense in progress")
}

// Метод оплаты запрошенного товара
func (i *ItemRequestedState) insertMoney(money int) error {
	if money < i.vendingMachine.itemPrice {
		return fmt.Errorf("inserted money is less. Please insert %d", i.vendingMachine.itemPrice)
	}
	fmt.Println("money entered is ok")
	i.vendingMachine.setState(i.vendingMachine.hasMoney)
	return nil
}

// Метод выдачи запрошенного, но неоплаченного товара
func (i *ItemRequestedState) dispenseItem() error {
	return fmt.Errorf("please insert money first")
}

// Реализация конкретного состояния - Товар оплачен и выдается
type HasMoneyState struct {
	vendingMachine *VendingMachine
}

// Метод запроса выдаваемого оплаченного товара
func (i *HasMoneyState) requestItem() error {
	return fmt.Errorf("item dispense in progress")
}

// Метод добавления оплаченного товара
func (i *HasMoneyState) addItem(count int) error {
	return fmt.Errorf("item dispense in progress")
}

// Метод оплаты оплаченного товара
func (i *HasMoneyState) insertMoney(money int) error {
	return fmt.Errorf("item out of stock")
}

// Метод выдачи оплаченного товара
func (i *HasMoneyState) dispenseItem() error {
	fmt.Println("Dispensing Item")
	i.vendingMachine.itemCount = i.vendingMachine.itemCount - 1
	if i.vendingMachine.itemCount == 0 {
		i.vendingMachine.setState(i.vendingMachine.noItem)
	} else {
		i.vendingMachine.setState(i.vendingMachine.hasItem)
	}
	return nil
}
