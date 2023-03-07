package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// начало решения

// Handy предоставляет удобный интерфейс
// для выполнения HTTP-запросов
type Handy struct {
	url        string
	client     *http.Client
	headers    http.Header
	parameters url.Values
	body       []byte
	err        error
}

// NewHandy создает новый экземпляр Handy
func NewHandy() *Handy {
	return &Handy{
		client:     http.DefaultClient,
		headers:    make(http.Header),
		parameters: make(url.Values),
	}
}

// URL устанавливает URL, на который пойдет запрос
func (h *Handy) URL(uri string) *Handy {
	h.url = uri
	return h
}

// Client устанавливает HTTP-клиента
// вместо умолчательного http.DefaultClient
func (h *Handy) Client(client *http.Client) *Handy {
	h.client = client
	return h
}

// Header устанавливает значение заголовка
func (h *Handy) Header(key, value string) *Handy {
	h.headers.Set(key, value)
	return h
}

// Param устанавливает значение URL-параметра
func (h *Handy) Param(key, value string) *Handy {
	h.parameters.Set(key, value)
	return h
}

// Form устанавливает данные, которые будут закодированы
// как application/x-www-form-urlencoded и отправлены в теле запроса
// с соответствующим content-type
func (h *Handy) Form(form map[string]string) *Handy {
	data := make(url.Values)
	for key, value := range form {
		data.Set(key, value)
	}
	h.body = []byte(data.Encode())
	h.headers.Set("Content-Type", "application/x-www-form-urlencoded")
	return h
}

// JSON устанавливает данные, которые будут закодированы
// как application/json и отправлены в теле запроса
// с соответствующим content-type
func (h *Handy) JSON(v interface{}) *Handy {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	if err := encoder.Encode(v); err != nil {
		h.body = nil
	} else {
		h.body = buf.Bytes()
	}
	// Проверка валидности JSON
	if !json.Valid(h.body) {
		h.err = fmt.Errorf("invalid JSON: %s", string(h.body))
		return h
	}
	h.headers.Set("Content-Type", "application/json")
	return h
}

// Get выполняет GET-запрос с настроенными ранее параметрами
func (h *Handy) Get() *HandyResponse {
	req, err := http.NewRequest("GET", h.url, nil)
	if err != nil {
		return &HandyResponse{err: err}
	}
	if h.err != nil {
		return &HandyResponse{err: h.err}
	}
	req.URL.RawQuery = h.parameters.Encode()
	for key, values := range h.headers {
		for _, value := range values {
			req.Header.Add(key, value)
		}
	}
	resp, err := h.client.Do(req)
	if err != nil {
		return &HandyResponse{err: err}
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &HandyResponse{err: err}
	}
	return &HandyResponse{
		StatusCode: resp.StatusCode,
		Body:       body,
	}
}

// Post выполняет POST-запрос с настроенными ранее параметрами
func (h *Handy) Post() *HandyResponse {
	req, err := http.NewRequest("POST", h.url, nil)
	if err != nil {
		return &HandyResponse{err: err}
	}
	if h.err != nil {
		return &HandyResponse{err: h.err}
	}

	req.Header = h.headers
	req.URL.RawQuery = h.parameters.Encode()

	// устанавливаем тело запроса
	if len(h.body) > 0 {
		req.Body = io.NopCloser(bytes.NewReader(h.body))
	} else {
		formData := []byte(h.parameters.Encode())
		req.Body = io.NopCloser(bytes.NewReader(formData))
	}

	// отправляем запрос
	resp, err := h.client.Do(req)
	if err != nil {
		return &HandyResponse{err: err}
	}
	defer resp.Body.Close()

	// читаем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &HandyResponse{err: err}
	}

	return &HandyResponse{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       body,
	}
}

// HandyResponse представляет ответ на HTTP-запрос
type HandyResponse struct {
	Body       []byte
	StatusCode int
	Headers    http.Header
	err        error
}

// OK возвращает true, если во время выполнения запроса
// не произошло ошибок, а код HTTP-статуса ответа равен 200
func (r *HandyResponse) OK() bool {
	return r.StatusCode == http.StatusOK && r.err == nil
}

// Bytes возвращает тело ответа как срез байт
func (r *HandyResponse) Bytes() []byte {
	return r.Body
}

// String возвращает тело ответа как строку
func (r *HandyResponse) String() string {
	return string(r.Body)
}

// JSON декодирует тело ответа из JSON и сохраняет
// результат по адресу, на который указывает v
func (r *HandyResponse) JSON(v interface{}) error {
	err := json.Unmarshal(r.Body, v)
	if err != nil {
		r.err = err
	}
	return err
}

// Err возвращает ошибку, которая возникла при выполнении запроса
// или обработке ответа
func (r *HandyResponse) Err() error {
	return r.err
}

// конец решения

func main() {
	{
		// примеры запросов

		// GET-запрос с параметрами
		NewHandy().URL("https://httpbingo.org/get").Param("id", "42").Get()

		// HTTP-заголовки
		NewHandy().
			URL("https://httpbingo.org/get").
			Header("Accept", "text/html").
			Header("Authorization", "Bearer 1234567890").
			Get()

		// POST формы
		params := map[string]string{
			"brand":    "lg",
			"category": "tv",
		}
		NewHandy().URL("https://httpbingo.org/post").Form(params).Post()

		// POST JSON-документа
		NewHandy().URL("https://httpbingo.org/post").JSON(params).Post()
	}

	{
		// пример обработки ответа

		// отправляем GET-запрос с параметрами
		resp := NewHandy().URL("https://httpbingo.org/get").Param("id", "42").Get()
		if !resp.OK() {
			panic(resp.String())
		}

		// декодируем ответ в JSON
		var data map[string]any
		_ = resp.JSON(&data)

		fmt.Println(data["url"])
		// "https://httpbingo.org/get"
		fmt.Println(data["args"])
		// map[id:[42]]
	}
}
