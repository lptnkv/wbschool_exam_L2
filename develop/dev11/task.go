package main

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	events    map[int]Event
	idCounter = 0
)

// Конфигурация сервера
type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

// Получить полный адрес сервера (хост+порт)
func (conf *Config) GetAdr() string {
	return conf.Host + ":" + conf.Port
}

// Конструктор config, читающий данные из файла config.json
func GetConfig() *Config {
	rawData, err := os.ReadFile("config.json")
	if err != nil {
		log.Printf("Can't open config file: %v\n", err)
	}
	var config Config
	err = json.Unmarshal(rawData, &config)
	if err != nil {
		log.Printf("Can't unmarshal config json: %v\n", err)
	}
	return &config
}

// Событие
type Event struct {
	ID   int       `json:"id"`
	Date time.Time `json:"date"`
	Name string    `json:"name"`
}

// Сообщение об ошибке
type ErrorResp struct {
	Content string `json:"error"`
}

// Успешный ответ
type SuccessResp struct {
	Content string `json:"result"`
}

// Middleware для логирования
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, req)
		log.Printf("%s %s %s", req.Method, req.RequestURI, time.Since(start))
	})
}

// Получение события из тела POST запроса
func parseBody(r *http.Request) (Event, error) {
	var e Event
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		return e, err
	}
	return e, nil
}

// Отправка ответа в формате json
func jsonResponse(w http.ResponseWriter, code int, msg interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(msg)
}

// Обработчик создания события
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errResp := ErrorResp{"Method is not allowed"}
		jsonResponse(w, http.StatusMethodNotAllowed, errResp)
		return
	}
	event, err := parseBody(r)
	if err != nil {
		errResp := ErrorResp{"Internal server error"}
		jsonResponse(w, http.StatusBadRequest, errResp)
		log.Printf("Could not parse POST request body: %v\n", err)
		return
	}
	idCounter++
	event.ID = idCounter
	events[event.ID] = event
	resp := SuccessResp{"Event added successfuly"}
	jsonResponse(w, http.StatusOK, resp)
}

// Обработчик обновления события
func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errResp := ErrorResp{"Method is not allowed"}
		jsonResponse(w, http.StatusMethodNotAllowed, errResp)
		return
	}
	event, err := parseBody(r)
	if err != nil {
		errResp := ErrorResp{"Bad request"}
		jsonResponse(w, http.StatusBadRequest, errResp)
		log.Printf("Could not parse POST request body: %v\n", err)
		return
	}
	if _, found := events[event.ID]; found {
		events[event.ID] = event
		resp := SuccessResp{"Event updated successfuly"}
		jsonResponse(w, http.StatusOK, resp)
		return
	} else {
		resp := SuccessResp{"Event not found"}
		jsonResponse(w, http.StatusNotFound, resp)
		return
	}
}

// Обработчик удаления события
func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		errResp := ErrorResp{"Method is not allowed"}
		jsonResponse(w, http.StatusMethodNotAllowed, errResp)
		return
	}
	event, err := parseBody(r)
	if err != nil {
		errResp := ErrorResp{"Bad request"}
		jsonResponse(w, http.StatusBadRequest, errResp)
		log.Printf("Could not parse POST request body: %v\n", err)
		return
	}
	if _, found := events[event.ID]; found {
		delete(events, event.ID)
		resp := SuccessResp{"Event deleted successfuly"}
		jsonResponse(w, http.StatusOK, resp)
		return
	} else {
		resp := SuccessResp{"Event not found"}
		jsonResponse(w, http.StatusNotFound, resp)
		return
	}
}

// Получение событий за день
func dayEventsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		errResp := ErrorResp{"Method is not allowed"}
		jsonResponse(w, http.StatusMethodNotAllowed, errResp)
		return
	}
	date, err := time.Parse("2006-01-02", r.URL.Query().Get("date"))
	if err != nil {
		errResp := ErrorResp{"Bad request"}
		jsonResponse(w, http.StatusBadRequest, errResp)
		log.Printf("Could not parse url params: %v\n", err)
		return
	}
	var res []Event
	for _, event := range events {
		if event.Date.Day() == date.Day() {
			res = append(res, event)
		}
	}
	if len(res) != 0 {
		jsonEvents, err := json.Marshal(res)
		if err != nil {
			errResp := ErrorResp{"Internal server error"}
			jsonResponse(w, http.StatusInternalServerError, errResp)
			log.Printf("Could not marshal json events: %v\n", err)
			return
		}
		resp := SuccessResp{string(jsonEvents)}
		jsonResponse(w, http.StatusOK, resp)
		return
	} else {
		errResp := ErrorResp{"Not found"}
		jsonResponse(w, http.StatusNotFound, errResp)
		return
	}
}

func main() {
	// Получение конфигурации
	config := GetConfig()

	// Создание роутера
	mux := http.NewServeMux()

	// Обработка эндпоинта create_event для создания событий
	mux.HandleFunc("/create_event", createEventHandler)
	mux.HandleFunc("/update_event", updateEventHandler)
	mux.HandleFunc("/delete_event", deleteEventHandler)
	mux.HandleFunc("/events_for_day", dayEventsHandler)

	// Применение middleware к нашему роутеру
	loggingHandler := Logging(mux)

	// Создание сервера по указанному адресу
	server := http.Server{
		Addr:    config.GetAdr(),
		Handler: loggingHandler,
	}
	// Канал для получения сигналов от ОС о прерывании работы
	c := make(chan os.Signal, 1)

	// Канал для получения ошибок
	errs := make(chan error)
	signal.Notify(c, os.Interrupt)

	// Запуск сервера в отдельной горутине
	go func() {
		errs <- server.ListenAndServe()
	}()

	// Получение сигнала от ОС, либо сигнала об ошибке для завершения сервера
	select {
	case sig := <-c:
		log.Printf("Got signal: %v\n", sig)
		log.Println("Shutting down server")
		server.Shutdown(context.Background())
	case e := <-errs:
		log.Printf("Got error: %v.\nShutting down server", e)
		server.Shutdown(context.Background())
	}
}
