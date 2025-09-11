package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestCase struct {
	InputData int
	Answer    int
	Expected  int
}

var Cases []TestCase = []TestCase{
	{InputData: 0,
		Expected: 1},
	{InputData: 1,
		Expected: 1},
	{InputData: 5,
		Expected: 120},
	{InputData: 3,
		Expected: 6},
}

func TestFactorial(t *testing.T) {
	for id, test := range Cases {
		if res := factorial(test.InputData); res != test.Expected {
			t.Errorf("test case %d failed: result %v expected %v", id, res, test.Expected)
		}
	}
}

type HttpTestCase struct {
	Name     string
	Numeric  int
	Expected []byte
}

var HttpCase = []HttpTestCase{
	{Name: "first test",
		Numeric:  1,
		Expected: []byte("1")},
	{Name: "second test",
		Numeric:  2,
		Expected: []byte("2")},
	{Name: "Third test",
		Numeric:  3,
		Expected: []byte("6")},
	{Name: "Foured test",
		Numeric:  5,
		Expected: []byte("120")},
}

func TestHandleFactorial(t *testing.T) {
	handler := http.HandlerFunc(HandleFactorial)
	for _, test := range HttpCase {
		t.Run(test.Name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			handlerData := fmt.Sprintf("/factorial?num=%d", test.Numeric)
			request, err := http.NewRequest("GET", handlerData, nil)
			// data := io.Reader([]byte(`{"num" : 5}`))
			//request, err := http.POST("http://localhost8080/factorial?num=5", "application/json", data)
			if err != nil {
				t.Error(err)
			}
			handler.ServeHTTP(recorder, request)
			if string(recorder.Body.Bytes()) != string(test.Expected) {
				t.Errorf("test %s failed: input: %v result: %v expected: %v", test.Name, test.Numeric, string(recorder.Body.Bytes()), string(test.Expected))
			}
		})
	}
}
