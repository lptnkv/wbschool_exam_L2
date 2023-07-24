package main

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"

	"github.com/mitchellh/go-ps"
)

// Командная оболочка
type Shell struct {
	Out      io.Writer
	In       io.Reader
	Pipe     bool // Имеется ли пайп
	PipeBuff *bytes.Buffer
}

// Конструктор структуры Shell
func NewShell(w io.Writer, r io.Reader) *Shell {
	return &Shell{Out: w, In: r}
}

// Запуск оболочки shell
func (s *Shell) Run() error {
	if err := s.GetLines(); err != nil {
		if _, err := fmt.Fprintln(s.Out, err); err != nil {
			log.Fatalf("Error executing command: %v\n", err.Error())
		}
	}
	return nil
}

// Обработка команды cd - переход в указанную директорию
func (s *Shell) cd(arg string) error {
	err := os.Chdir(arg)
	if err != nil {
		return err
	}
	return nil
}

// Обработка команды pwd - вывод полного пути текущей директории
func (s *Shell) pwd() error {

	out := s.Out
	path, err := os.Getwd()

	if err != nil {
		return err
	}

	if s.Pipe {
		out = s.PipeBuff
	}
	_, err = fmt.Fprintln(out, path)
	if err != nil {
		return err
	}
	return nil
}

// Обработка команды echo - вывод переданных аргументов
func (s *Shell) echo(args []string, fullLine string) error {
	// Куда будем выводить
	printer := s.Out

	// Первый аргумент
	start := args[0]

	// Последний аргумент
	end := args[len(args)-1]

	if s.Pipe {
		printer = s.PipeBuff
	}

	//
	if start[0] == '"' && end[len(end)-1] == '"' {
		line := strings.TrimPrefix(fullLine, "echo ")
		line = strings.TrimLeft(line, `"`)
		line = strings.TrimRight(line, `"`)
		if _, err := fmt.Fprintln(printer, line); err != nil {
			return err
		}
	} else {
		if _, err := fmt.Fprintln(printer, strings.Join(args, " ")); err != nil {
			return err
		}
	}
	return nil
}

// Обработка команды kill - завершение процесса
func (s *Shell) kill(pids []string) error {
	for _, pid := range pids {
		if id, err := strconv.Atoi(pid); err != nil {
			return fmt.Errorf("could not convert to integer: %v", err.Error())
		} else {
			p, err := os.FindProcess(id)
			if err != nil {
				return fmt.Errorf("could not find process by pid: %v", err.Error())
			}
			err = p.Signal(syscall.SIGTERM)
			if err != nil {
				return fmt.Errorf("could not send kill sygnal: %v", err.Error())
			}
		}
	}
	return nil
}

// Обработка команды ps - вывод списка процессов
func (s *Shell) ps() error {
	processList, err := ps.Processes()
	if err != nil {
		return err
	}
	out := s.Out
	if s.Pipe {
		out = s.PipeBuff
	}
	for proc := range processList {
		process := processList[proc]
		_, err = fmt.Fprintf(out, "%v\t%v\t%v\n", process.Pid(), process.PPid(), process.Executable())
		if err != nil {
			return err
		}
	}
	return nil
}

// Метод сканирования строк
func (s *Shell) GetLines() error {
	scanner := bufio.NewScanner(s.In)
	fmt.Fprint(s.Out, "$ ")

	// Пока на вход поступают данные и не введена команда quit
	for scanner.Scan() && (scanner.Text() != `\quit`) {
		line := scanner.Text()
		err := s.Fork(line)
		if err != nil {
			return err
		}
		fmt.Fprint(s.Out, "$ ")
	}
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
		os.Exit(1)
	}
	return nil
}

// Обработка команды exec - замена текущего процесса новым
func (s *Shell) Exec(line []string) error {
	var cmd *exec.Cmd
	if len(line) == 1 {
		cmd = exec.Command(line[0])
	} else {
		cmd = exec.Command(line[0], line[1:]...)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if s.Pipe {
		cmd.Stdout = s.PipeBuff
	}
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

// CaseShell - выбор команд
func (s *Shell) CaseShell(line string) error {
	commandAndArgs := strings.Fields(line)
	if len(commandAndArgs) != 0 {
		switch commandAndArgs[0] {
		case "cd":
			if len(commandAndArgs) == 2 {
				err := s.cd(commandAndArgs[1])
				if err != nil {
					_, err := fmt.Fprintln(s.Out, err)
					if err != nil {
						return err
					}
				}
			} else {
				return fmt.Errorf("cd must have 1 argument")
			}
		case "ps":
			if len(commandAndArgs) == 1 {
				if err := s.ps(); err != nil {
					return err
				}
			} else {
				return fmt.Errorf("ps must have no arguments")
			}
		case "pwd":
			if len(commandAndArgs) != 1 {
				return fmt.Errorf("pwd must have no arguments")
			}
			err := s.pwd()
			if err != nil {
				return err
			}

		case "echo":
			if len(commandAndArgs) != 1 {
				err := s.echo(commandAndArgs[1:], line)
				if err != nil {
					return err
				}
			} else {
				return fmt.Errorf("echo must have 1+ arguments")
			}
		case "kill":
			if len(commandAndArgs) != 1 {
				err := s.kill(commandAndArgs[1:])
				if err != nil {
					fmt.Fprintln(s.Out, err.Error())
				}
			} else {
				return fmt.Errorf("kill must have 1+ arguments")
			}
		case "exec":
			if len(commandAndArgs) != 1 {
				err := s.Exec(commandAndArgs[1:])
				if err != nil {
					return err
				}
			} else {
				return fmt.Errorf("exec must have 1+ arguments")
			}
		default:
			if _, err := fmt.Fprintf(s.Out, "unknown command '%v'\n", commandAndArgs[0]); err != nil {
				return err
			}
		}
	}
	return nil
}

// Проверка на наличие пайпоа
func (s *Shell) CheckPipes(line string) error {
	strCmd := strings.Split(line, "|")
	if len(strCmd) > 1 {
		s.Pipe = true
		s.PipeBuff = new(bytes.Buffer)
		for index, value := range strCmd {
			if index != 0 {
				comm1 := strings.Fields(value)

				if len(comm1) > 1 {
					comm1New := make([]string, 2)
					comm1New[0], comm1New[1] = comm1[0], s.PipeBuff.String()
					comm1 = comm1New
				} else {
					comm1 = append(comm1, s.PipeBuff.String())
				}
				value = strings.Join(comm1, " ")
			}
			s.PipeBuff.Reset()
			if index == len(strCmd)-1 {
				s.Pipe = false
			}
			if err := s.CaseShell(value); err != nil {
				if _, err := fmt.Fprintln(s.Out, err); err != nil {
					return err
				}
			}
		}
	} else {
		if err := s.CaseShell(line); err != nil {
			if _, err = fmt.Fprintln(s.Out, err); err != nil {
				return err
			}
		}
	}
	return nil
}

// Подготовка строки к обработке и вызов метода обработки
func (s *Shell) Fork(str string) error {
	// Убираем лишние пробелы
	str = strings.TrimRight(str, " ")
	err := s.CheckPipes(str)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	s := NewShell(os.Stdin, os.Stdout)
	s.Run()
}
