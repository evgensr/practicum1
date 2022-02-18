package handler_test

import (
	"bytes"
	"github.com/evgensr/practicum1/cmd/shortener/handler"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_getHash(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "pass md5 of 1",
			args: args{"1"},
			want: "c4ca4238a0b923820dcc509a6f75849b",
		},
		{
			name: "pass md5 of string example",
			args: args{"example"},
			want: "1a79a4d60de6718e8e5b326e338ae533",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.GetHash(tt.args.text); got != tt.want {
				t.Errorf("getHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {

	// определяем структуру теста
	type want struct {
		code     int
		request  string
		response string
	}
	baseURL := "http://localhost:8080/"

	// создаём массив тестов: имя и желаемый результат
	tests := []struct {
		name string
		want want
	}{
		// определяем все тесты
		{
			name: "positive test #1",
			want: want{
				code:     201,
				request:  `https://ya.ru`,
				response: baseURL + "e98192e19505472476a49f10388428ab",
			},
		},
		{
			name: "positive test #2",
			want: want{
				code:     201,
				request:  `https://habr.ru`,
				response: baseURL + "98981d87735b7f871c516eaf9b6bf461",
			},
		},
	}
	for _, tt := range tests {
		// запускаем каждый тест
		t.Run(tt.name, func(t *testing.T) {

			// создаём новый Recorder
			w := httptest.NewRecorder()

			reqBody := []byte(tt.want.request)
			request := httptest.NewRequest(http.MethodPost, "/", bytes.NewReader(reqBody))

			// определяем хендлер
			h := http.HandlerFunc(handler.HandlerPOST)
			// запускаем сервер
			h.ServeHTTP(w, request)
			res := w.Result()

			// проверяем код ответа
			if res.StatusCode != tt.want.code {
				t.Errorf("Expected status code %d, got %d", tt.want.code, w.Code)
			}

			// получаем и проверяем тело запроса
			defer res.Body.Close()
			resBody, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}
			if string(resBody) != tt.want.response {
				t.Errorf("Expected body %s, got %s", tt.want.response, w.Body.String())
			}

		})
	}

}
