package pattern

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

// Интерфейс объекта-команды
type Command interface {
	Execute() string
}

// Получатель команды
type Receiver struct {
}

// Обработчик команды включения Получателем
func (r *Receiver) ToggleOn() string {
	return "Toggle On"
}

// Обработчик команды включения Получателем
func (r *Receiver) ToggleOff() string {
	return "Toggle Off"
}

// Конкретная команда включения
type ToggleOnCommand struct {
	receiver *Receiver
}

// Реализация выполнения команды включения
func (c *ToggleOnCommand) Execute() string {
	return c.receiver.ToggleOn()
}

// Команда выключения
type ToggleOffCommand struct {
	receiver *Receiver
}

func (c *ToggleOffCommand) Execute() string {
	return c.receiver.ToggleOff()
}

// Инициатор команд
type Invoker struct {
	commands []Command
}

// Метод добавления команды в список Инициатора
func (i *Invoker) AddCommand(command Command) {
	i.commands = append(i.commands, command)
}

// Метод удаления команды из списка Инициатора
func (i *Invoker) RemoveCommand() {
	if len(i.commands) != 0 {
		i.commands = i.commands[:len(i.commands)-1]
	}
}

// Метод выполнения всех команд
func (i *Invoker) Execute() string {
	res := ""
	for _, command := range i.commands {
		res += command.Execute() + "\n"
	}
	return res
}
