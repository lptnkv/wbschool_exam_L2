package main

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Структура telnet
type TelnetClient struct {
	*Config
	Conn    net.Conn
	errChan chan error
	in      io.Reader
	out     io.Writer
}

// Конструктор клиента
func NewTelnet(c *Config, ind io.Reader, outd io.Writer) *TelnetClient {
	return &TelnetClient{
		Config:  c,
		errChan: make(chan error),
		in:      ind,
		out:     outd,
	}
}

// Метод получения полного адреса (хост:порт)
func (t *TelnetClient) getFullAddress() string {
	return net.JoinHostPort(t.Host, t.Port)
}

// Метод подключения по протоколу tcp к адресу
func (t *TelnetClient) connect() {
	conn, err := net.DialTimeout("tcp", t.getFullAddress(), t.TimeOutDuration)
	if err != nil {
		time.Sleep(t.TimeOutDuration)
		log.Fatalln("failed connection: ", err)
	}
	t.Conn = conn
	fmt.Println("Successfully connected!")
}

// Метод разрыва соединения
func (t *TelnetClient) disconnect() {
	if err := t.Conn.Close(); err != nil {
		log.Fatal("disconnect error")
	}
}

// Метод отправки данных
func (t *TelnetClient) send() error {
	if _, err := io.Copy(t.Conn, t.in); err != nil {
		return err
	}
	log.Println("EOF")
	return nil
}

// Метод получения данных
func (t *TelnetClient) receive() error {
	if _, err := io.Copy(t.out, t.Conn); err != nil {
		return err
	}
	log.Println("disconnect from server")
	return nil
}

// Метод запуска утилиты
func (t *TelnetClient) Run() {
	// Канал для получения сигнала от ОС
	sigint := make(chan os.Signal, 1)

	// Подписываемся на сигналы от ОС
	signal.Notify(sigint, syscall.SIGINT, syscall.SIGTERM) // ctrl+d
	t.connect()

	// Горутина для отправки данных
	go func() {
		if err := t.send(); err != nil {
			t.errChan <- err
			log.Println(err)
		}
	}()

	// Горутина для получения данных
	go func() {
		if err := t.receive(); err != nil {
			t.errChan <- err
			log.Println(err)
		}
	}()

	// Завершение работы утилиты при получении ошибки или сигнала от ОС
	select {
	case err := <-t.errChan:
		t.disconnect()
		log.Println(err)
	case <-sigint:
		t.disconnect()
		log.Println("telnet: exit")
	}
}

// Структура для конфигурации telnet-клиента
type Config struct {
	TimeOutDuration time.Duration
	Host            string
	Port            string
}

// Конструктор конфигурации
func NewConfig() *Config {
	return &Config{}
}

// Метод инициализации конфигурации
func (c *Config) InitConfig() {
	flag.DurationVar(&c.TimeOutDuration, "timeout", 10*time.Second, "timeout duration")
	flag.Parse()
	if len(flag.Args()) != 2 {
		log.Fatal("entered more or less then 2 arguments")
	}
	c.Host, c.Port = flag.Arg(0), flag.Arg(1)
}

func main() {
	c := NewConfig()
	c.InitConfig()
	t := NewTelnet(c, os.Stdin, os.Stdout)
	t.Run()
}
