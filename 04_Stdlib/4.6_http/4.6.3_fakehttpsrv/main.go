package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
)

// начало решения

// statusHandler возвращает ответ с кодом, который передан
// в заголовке X-Status. Например:
//
//	X-Status = 200 -> вернет ответ с кодом 200
//	X-Status = 404 -> вернет ответ с кодом 404
//	X-Status = 503 -> вернет ответ с кодом 503
//
// Если заголовок отсутствует, возвращает ответ с кодом 200.
// Тело ответа пустое.
func statusHandler(w http.ResponseWriter, r *http.Request) {
	h := r.Header.Get("X-Status")
	w.Header().Set("Content-Type", "text/plain")
	var statusCode int
	if h == "" {
		statusCode = 200
	} else {
		i, err := strconv.Atoi(h)
		if err != nil {
			panic(err)
		}
		statusCode = i
	}

	w.WriteHeader(statusCode)
	_, _ = w.Write([]byte(""))
}

// echoHandler возвращает ответ с тем же телом
// и заголовком Content-Type, которые пришли в запросе
func echoHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем заголовок Content-Type из запроса
	contentType := r.Header.Get("Content-Type")

	// Устанавливаем заголовок Content-Type в ответе
	w.Header().Set("Content-Type", contentType)

	w.WriteHeader(http.StatusOK)
	_, err := io.Copy(w, r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// jsonHandler проверяет, что Content-Type = application/json,
// а в теле запроса пришел валидный JSON,
// после чего возвращает ответ с кодом 200.
// Если какая-то проверка не прошла — возвращает ответ с кодом 400.
// Тело ответа пустое.
func jsonHandler(w http.ResponseWriter, r *http.Request) {
	// Проверяем Content-Type
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Invalid Content-Type", http.StatusBadRequest)
		return
	}

	// Читаем содержимое тела запроса
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Проверяем валидность JSON из тела запроса
	if !json.Valid(body) {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Возвращаем ответ с кодом 200
	w.WriteHeader(http.StatusOK)
}

// конец решения

func startServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/status", statusHandler)
	mux.HandleFunc("/echo", echoHandler)
	mux.HandleFunc("/json", jsonHandler)
	return httptest.NewServer(mux)
}

func main() {
	server := startServer()
	defer server.Close()
	client := server.Client()

	{
		uri := server.URL + "/status"
		req, _ := http.NewRequest(http.MethodGet, uri, nil)
		req.Header.Add("X-Status", "201")
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		fmt.Println(resp.Status)
		// 201 Created
	}

	{
		uri := server.URL + "/echo"
		reqBody := []byte("hello world")
		resp, err := client.Post(uri, "text/plain", bytes.NewReader(reqBody))
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()
		respBody, _ := io.ReadAll(resp.Body)
		fmt.Println(resp.Status)
		fmt.Println(string(respBody))
		// 200 OK
		// hello world
	}

	{
		uri := server.URL + "/json"
		reqBody, _ := json.Marshal(map[string]bool{"ok": true})
		resp, err := client.Post(uri, "application/json", bytes.NewReader(reqBody))
		if err != nil {
			panic(err)
		}
		fmt.Println(resp.Status)
		// 200 OK
	}
}
