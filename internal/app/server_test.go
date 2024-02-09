package app

import (
	"encoding/json"
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
			want:       `{"message":"operation failed, does not connect to database","code":400}` + "\n",
			statusCode: http.StatusOK,
		},
		{
			testNum:    2,
			name:       "empty id",
			method:     http.MethodDelete,
			param:      "id",
			value:      "",
			want:       `{"message":"operation failed, wrong id = ","code":400}` + "\n",
			statusCode: http.StatusOK,
		},
		{
			testNum:    3,
			name:       "string id",
			method:     http.MethodDelete,
			param:      "id",
			value:      "ddf",
			want:       `{"message":"operation failed, wrong id = ddf","code":400}` + "\n",
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
			want:       `{"message":"operation failed, does not connect to database","code":400}` + "\n",
			statusCode: http.StatusOK,
		},
		{
			testNum:    2,
			name:       "empty name",
			method:     http.MethodPost,
			param:      []string{"name", "surname", "patronymic"},
			value:      []string{"", "chehov", "ivanovich"},
			want:       `{"message":"operation failed, wrong users param","code":400}` + "\n",
			statusCode: http.StatusOK,
		},
		{
			testNum:    3,
			name:       "empty surname and patronymic",
			method:     http.MethodPost,
			param:      []string{"name", "surname", "patronymic"},
			value:      []string{"anton", "", ""},
			want:       `{"message":"operation failed, wrong users param","code":400}` + "\n",
			statusCode: http.StatusOK,
		},
		{
			testNum:    4,
			name:       "very long name",
			method:     http.MethodPost,
			param:      []string{"name", "surname", "patronymic"},
			value:      []string{"antonnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnn", "chehov", "ivanovich"},
			want:       `{"message":"operation failed, wrong users param","code":400}` + "\n",
			statusCode: http.StatusOK,
		},
		{
			testNum:    5,
			name:       "very long surname",
			method:     http.MethodPost,
			param:      []string{"name", "surname", "patronymic"},
			value:      []string{"ivan", "antonnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnnov", "nikolaevich"},
			want:       `{"message":"operation failed, wrong users param","code":400}` + "\n",
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range testTable {
		rr := httptest.NewRecorder()

		kvPairs := make(map[string]string)
		for i := 0; i < len(tc.param); i++ {
			kvPairs[tc.param[i]] = tc.value[i]
		}
		postJson, err := json.Marshal(kvPairs)
		if err != nil {
			panic(err)
		}

		req := httptest.NewRequest(tc.method, "/insert", strings.NewReader(string(postJson)))

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
			want:       `{"message":"operation failed, does not connect to database","code":400}` + "\n",
			statusCode: http.StatusOK,
		},
		{
			testNum:    2,
			name:       "null id",
			method:     http.MethodPost,
			param:      []string{"id", "name", "surname", "patronymic", "age", "sex", "nationality"},
			value:      []string{"", "anton", "chehov", "ivanovich", "45", "male", "RU"},
			want:       `{"message":"operation failed, does not connect to database","code":400}` + "\n",
			statusCode: http.StatusOK,
		},
		{
			testNum:    3,
			name:       "null name and surname",
			method:     http.MethodPost,
			param:      []string{"id", "name", "surname", "patronymic", "age", "sex", "nationality"},
			value:      []string{"9", "", "", "ivanovich", "45", "male", "RU"},
			want:       `{"message":"operation failed, does not connect to database","code":400}` + "\n",
			statusCode: http.StatusOK,
		},
		{
			testNum:    4,
			name:       "null nationality and age",
			method:     http.MethodPost,
			param:      []string{"id", "name", "surname", "patronymic", "age", "sex", "nationality"},
			value:      []string{"9", "anton", "chehov", "ivanovich", "", "male", ""},
			want:       `{"message":"operation failed, does not connect to database","code":400}` + "\n",
			statusCode: http.StatusOK,
		},
		{
			testNum:    5,
			name:       "string id",
			method:     http.MethodPost,
			param:      []string{"id", "name", "surname", "patronymic", "age", "sex", "nationality"},
			value:      []string{"gav", "anton", "chehov", "ivanovich", "44", "male", "RU"},
			want:       `{"message":"operation failed, wrong id = gav","code":400}` + "\n",
			statusCode: http.StatusOK,
		},
		{
			testNum:    6,
			name:       "change only age from id",
			method:     http.MethodPost,
			param:      []string{"id", "age"},
			value:      []string{"9", "56"},
			want:       `{"message":"operation failed, does not connect to database","code":400}` + "\n",
			statusCode: http.StatusOK,
		},
	}

	for _, tc := range testTable {
		rr := httptest.NewRecorder()

		kvPairs := make(map[string]string)
		for i := 0; i < len(tc.param); i++ {
			kvPairs[tc.param[i]] = tc.value[i]
		}
		postJson, err := json.Marshal(kvPairs)
		if err != nil {
			panic(err)
		}

		req := httptest.NewRequest(tc.method, "/update", strings.NewReader(string(postJson)))

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
			want:       `{"message":"operation failed, wrong id = dfhd","code":400}` + "\n",
			statusCode: http.StatusOK,
		},
		{
			testNum:    2,
			name:       "normal query",
			method:     http.MethodGet,
			param:      []string{"name", "limit", "sort"},
			value:      []string{"anton", "5", "id"},
			want:       `{"message":"operation failed, does not connect to database","code":400}` + "\n",
			statusCode: http.StatusOK,
		},
		{
			testNum:    3,
			name:       "show all query",
			method:     http.MethodGet,
			param:      []string{},
			value:      []string{},
			want:       `{"message":"operation failed, does not connect to database","code":400}` + "\n",
			statusCode: http.StatusOK,
		},
		{
			testNum:    4,
			name:       "wrong sort, show nothing",
			method:     http.MethodGet,
			param:      []string{"id", "name", "surname", "patronymic", "age", "sex", "nationality", "limit", "offset", "sort"},
			value:      []string{"5", "anton", "chehov", "ivanovich", "45", "male", "RU", "5", "5", "43634"},
			want:       `{"message":"operation failed, does not connect to database","code":400}` + "\n",
			statusCode: http.StatusOK,
		},
		{
			testNum:    5,
			name:       "show spec id",
			method:     http.MethodGet,
			param:      []string{"id"},
			value:      []string{"7"},
			want:       `{"message":"operation failed, does not connect to database","code":400}` + "\n",
			statusCode: http.StatusOK,
		},
		{
			testNum:    6,
			name:       "show spec name",
			method:     http.MethodGet,
			param:      []string{"name"},
			value:      []string{"anton"},
			want:       `{"message":"operation failed, does not connect to database","code":400}` + "\n",
			statusCode: http.StatusOK,
		},
		{
			testNum:    7,
			name:       "show spec age and gender",
			method:     http.MethodGet,
			param:      []string{"age", "sex"},
			value:      []string{"45", "female"},
			want:       `{"message":"operation failed, does not connect to database","code":400}` + "\n",
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

	req1, err := http.NewRequest("DELETE", "/delete", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr1 := newRequestRecorder(req1, "DELETE", "/delete", deleteGetRequest)
	if rr1.Code != 200 {
		t.Error("Expected response code to be 200")
	}

	req2, err := http.NewRequest("DELETE", "/delete", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr2 := newRequestRecorder(req2, "GET", "/show", showsSpecGetRequest)
	if rr2.Code != 200 {
		t.Error("Expected response code to be 200")
	}

	req3, err := http.NewRequest("DELETE", "/delete", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr3 := newRequestRecorder(req3, "POST", "/insert", insertGetRequest)
	if rr3.Code != 200 {
		t.Error("Expected response code to be 200")
	}

}

func newRequestRecorder(req *http.Request, method string, strPath string, fnHandler func(w http.ResponseWriter, r *http.Request)) *httptest.ResponseRecorder {
	router := routers()

	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)
	return rr
}
