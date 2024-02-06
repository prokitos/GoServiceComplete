package app

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestDelete(t *testing.T) {

	testTable := []struct {
		testNum    int
		name       string
		method     string
		param      string
		value      string
		body       string
		want       string
		statusCode int
	}{
		{
			testNum:    1,
			name:       "normal query",
			method:     http.MethodDelete,
			param:      "id",
			value:      "5",
			want:       `"message": "Delete request succes","status": "Null execute"`,
			statusCode: http.StatusOK,
		},
		{
			testNum:    2,
			name:       "empty id",
			method:     http.MethodDelete,
			param:      "id",
			value:      "",
			want:       `"message": "Delete request failed"`,
			statusCode: http.StatusOK,
		},
		{
			testNum:    3,
			name:       "string id",
			method:     http.MethodDelete,
			param:      "id",
			value:      "ddf",
			want:       `"message": "Delete request failed"`,
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range testTable {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(tc.method, "/delete", strings.NewReader(tc.body))

		params := url.Values{}
		params.Add(tc.param, tc.value)
		req.URL.RawQuery = params.Encode()

		handler := http.HandlerFunc(deleteGetRequest)
		handler.ServeHTTP(rr, req)

		// Обработчиик ошибок
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler wrong code at test #%v, got [%v] want [%v]", tc.testNum, status, http.StatusOK)
		}
		if rr.Body.String() != tc.want {
			t.Errorf("handler body wrong at test #%v, got [%v] want [%v]", tc.testNum, rr.Body.String(), tc.want)
		}

	}
}

func TestInsert(t *testing.T) {

	testTable := []struct {
		testNum    int
		name       string
		method     string
		param      []string
		value      []string
		body       string
		want       string
		statusCode int
	}{
		{
			testNum:    1,
			name:       "normal query",
			method:     http.MethodPost,
			param:      []string{"name", "surname", "patronymic"},
			value:      []string{"anton", "chehov", "ivanovich"},
			want:       `"message": "Insert request succes","status": "Null execute"`,
			statusCode: http.StatusOK,
		},
		{
			testNum:    2,
			name:       "empty name",
			method:     http.MethodPost,
			param:      []string{"name", "surname", "patronymic"},
			value:      []string{"", "chehov", "ivanovich"},
			want:       `"message": "Insert request failed"`,
			statusCode: http.StatusOK,
		},
		{
			testNum:    3,
			name:       "empty surname and patronymic",
			method:     http.MethodPost,
			param:      []string{"name", "surname", "patronymic"},
			value:      []string{"anton", "", ""},
			want:       `"message": "Insert request failed"`,
			statusCode: http.StatusOK,
		},
		{
			testNum:    4,
			name:       "very long name",
			method:     http.MethodPost,
			param:      []string{"name", "surname", "patronymic"},
			value:      []string{"antonnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnn", "chehov", "ivanovich"},
			want:       `"message": "Insert request failed"`,
			statusCode: http.StatusOK,
		},
		{
			testNum:    5,
			name:       "very long surname",
			method:     http.MethodPost,
			param:      []string{"name", "surname", "patronymic"},
			value:      []string{"ivan", "antonnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnov", "nikolaevich"},
			want:       `"message": "Insert request failed"`,
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range testTable {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(tc.method, "/insert", strings.NewReader(tc.body))

		params := url.Values{}
		for i := 0; i < len(tc.param); i++ {
			params.Add(tc.param[i], tc.value[i])
		}
		req.URL.RawQuery = params.Encode()

		handler := http.HandlerFunc(insertGetRequest)
		handler.ServeHTTP(rr, req)

		// Обработчиик ошибок
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler wrong code at test #%v, got [%v] want [%v]", tc.testNum, status, http.StatusOK)
		}
		if rr.Body.String() != tc.want {
			t.Errorf("handler body wrong at test #%v, got [%v] want [%v]", tc.testNum, rr.Body.String(), tc.want)
		}

	}
}

func TestUpdate(t *testing.T) {

	testTable := []struct {
		testNum    int
		name       string
		method     string
		param      []string
		value      []string
		body       string
		want       string
		statusCode int
	}{
		{
			testNum:    1,
			name:       "normal query",
			method:     http.MethodPost,
			param:      []string{"id", "name", "surname", "patronymic", "age", "sex", "nationality"},
			value:      []string{"9", "anton", "chehov", "ivanovich", "45", "male", "RU"},
			want:       `"message": "Update request succes","status": "Null execute"`,
			statusCode: http.StatusOK,
		},
		{
			testNum:    2,
			name:       "null id",
			method:     http.MethodPost,
			param:      []string{"id", "name", "surname", "patronymic", "age", "sex", "nationality"},
			value:      []string{"", "anton", "chehov", "ivanovich", "45", "male", "RU"},
			want:       `"message": "Update request failed"`,
			statusCode: http.StatusOK,
		},
		{
			testNum:    3,
			name:       "null name and surname",
			method:     http.MethodPost,
			param:      []string{"id", "name", "surname", "patronymic", "age", "sex", "nationality"},
			value:      []string{"9", "", "", "ivanovich", "45", "male", "RU"},
			want:       `"message": "Update request succes","status": "Null execute"`,
			statusCode: http.StatusOK,
		},
		{
			testNum:    4,
			name:       "null nationality and age",
			method:     http.MethodPost,
			param:      []string{"id", "name", "surname", "patronymic", "age", "sex", "nationality"},
			value:      []string{"9", "anton", "chehov", "ivanovich", "", "male", ""},
			want:       `"message": "Update request succes","status": "Null execute"`,
			statusCode: http.StatusOK,
		},
		{
			testNum:    5,
			name:       "string id",
			method:     http.MethodPost,
			param:      []string{"id", "name", "surname", "patronymic", "age", "sex", "nationality"},
			value:      []string{"gav", "anton", "chehov", "ivanovich", "44", "male", "RU"},
			want:       `"message": "Update request failed"`,
			statusCode: http.StatusOK,
		},
		{
			testNum:    6,
			name:       "change only age from id",
			method:     http.MethodPost,
			param:      []string{"id", "age"},
			value:      []string{"9", "56"},
			want:       `"message": "Update request succes","status": "Null execute"`,
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range testTable {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(tc.method, "/update", strings.NewReader(tc.body))

		params := url.Values{}
		for i := 0; i < len(tc.param); i++ {
			params.Add(tc.param[i], tc.value[i])
		}
		req.URL.RawQuery = params.Encode()

		handler := http.HandlerFunc(updateGetRequest)
		handler.ServeHTTP(rr, req)

		// Обработчиик ошибок
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler wrong code at test #%v, got [%v] want [%v]", tc.testNum, status, http.StatusOK)
		}
		if rr.Body.String() != tc.want {
			t.Errorf("handler body wrong at test #%v, got [%v] want [%v]", tc.testNum, rr.Body.String(), tc.want)
		}

	}
}

func TestShow(t *testing.T) {

	testTable := []struct {
		testNum    int
		name       string
		method     string
		param      []string
		value      []string
		body       string
		want       string
		statusCode int
	}{
		{
			testNum:    1,
			name:       "string id",
			method:     http.MethodGet,
			param:      []string{"id", "sort"},
			value:      []string{"dfhd", "id"},
			want:       `"message": "Show request failed"`,
			statusCode: http.StatusOK,
		},
		{
			testNum:    2,
			name:       "normal query",
			method:     http.MethodGet,
			param:      []string{"name", "limit", "sort"},
			value:      []string{"anton", "5", "id"},
			want:       `"message": "Show request succes","status": "Null execute"`,
			statusCode: http.StatusOK,
		},
		{
			testNum:    3,
			name:       "show all query",
			method:     http.MethodGet,
			param:      []string{},
			value:      []string{},
			want:       `"message": "Show request succes","status": "Null execute"`,
			statusCode: http.StatusOK,
		},
		{
			testNum:    4,
			name:       "wrong sort, show nothing",
			method:     http.MethodGet,
			param:      []string{"id", "name", "surname", "patronymic", "age", "sex", "nationality", "limit", "offset", "sort"},
			value:      []string{"5", "anton", "chehov", "ivanovich", "45", "male", "RU", "5", "5", "43634"},
			want:       `"message": "Show request succes","status": "Null execute"`,
			statusCode: http.StatusOK,
		},
		{
			testNum:    5,
			name:       "show spec id",
			method:     http.MethodGet,
			param:      []string{"id"},
			value:      []string{"7"},
			want:       `"message": "Show request succes","status": "Null execute"`,
			statusCode: http.StatusOK,
		},
		{
			testNum:    6,
			name:       "show spec name",
			method:     http.MethodGet,
			param:      []string{"name"},
			value:      []string{"anton"},
			want:       `"message": "Show request succes","status": "Null execute"`,
			statusCode: http.StatusOK,
		},
		{
			testNum:    7,
			name:       "show spec age and gender",
			method:     http.MethodGet,
			param:      []string{"age", "sex"},
			value:      []string{"45", "female"},
			want:       `"message": "Show request succes","status": "Null execute"`,
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range testTable {
		rr := httptest.NewRecorder()
		req := httptest.NewRequest(tc.method, "/show", strings.NewReader(tc.body))

		params := url.Values{}
		for i := 0; i < len(tc.param); i++ {
			params.Add(tc.param[i], tc.value[i])
		}
		req.URL.RawQuery = params.Encode()

		handler := http.HandlerFunc(showsSpecGetRequest)
		handler.ServeHTTP(rr, req)

		// Обработчиик ошибок
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler wrong code at test #%v, got [%v] want [%v]", tc.testNum, status, http.StatusOK)
		}
		if rr.Body.String() != tc.want {
			t.Errorf("handler body wrong at test #%v, got [%v] want [%v]", tc.testNum, rr.Body.String(), tc.want)
		}

	}
}

// valide http links test
func TestHttpConnect(t *testing.T) {

	testTable := []struct {
		testNum int
		method  string
		path    string
		functs  any
	}{
		{
			testNum: 1,
			method:  "POST",
			path:    "/insert",
			functs:  deleteGetRequest,
		},
		{
			testNum: 2,
			method:  "GET",
			path:    "/show",
			functs:  showsSpecGetRequest,
		},
	}

	for _, tc := range testTable {
		req, err := http.NewRequest(tc.method, tc.path, nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := newRequestRecorder(req, tc.method, tc.path, deleteGetRequest)
		if rr.Code != 200 {
			t.Error("Expected response code to be 200")
		}
	}

}

func newRequestRecorder(req *http.Request, method string, strPath string, fnHandler func(w http.ResponseWriter, r *http.Request)) *httptest.ResponseRecorder {
	router := routers()

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	return rr
}
