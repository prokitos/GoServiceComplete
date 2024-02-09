package services

import "testing"

func TestJsonAge(t *testing.T) {

	TestTable := []struct {
		testNum    int
		name       string
		jsonString string
		expected   string
	}{
		{
			testNum:    1,
			name:       "test 1",
			jsonString: "{\"count\":128925,\"name\":\"eva\",\"age\":59}",
			expected:   "59",
		},
		{
			testNum:    2,
			name:       "test 2",
			jsonString: "{\"count\":128925,\"name\":\"alex\",\"age\":42}",
			expected:   "42",
		},
		{
			testNum:    3,
			name:       "test 3",
			jsonString: "{\"count\":128925,\"name\":\"gggdsgr\",\"age\":0}",
			expected:   "0",
		},
	}

	for i, testCase := range TestTable {

		result := AgeComputing(testCase.jsonString)

		t.Logf("calling jsonAgeGet №%d, get %s", i, result)

		if result != testCase.expected {
			t.Errorf("incorrecnt result. Expect %s, get %s", testCase.expected, result)
		}

	}

}

func TestJsonSex(t *testing.T) {

	TestTable := []struct {
		testNum    int
		name       string
		jsonString string
		expected   string
	}{
		{
			testNum:    1,
			name:       "male test",
			jsonString: "{\"count\":25459,\"name\":\"Dmitriy\",\"gender\":\"male\",\"probability\":1.0}",
			expected:   "male",
		},
		{
			testNum:    2,
			name:       "female test",
			jsonString: "{\"count\":16343,\"name\":\"Janna\",\"gender\":\"female\",\"probability\":0.99}",
			expected:   "female",
		},
	}

	for _, testCase := range TestTable {

		result := SexComputing(testCase.jsonString)

		// Обработчиик ошибок
		if result != testCase.expected {
			t.Errorf("result wrong at test #%v, got [%v] want [%v]", testCase.testNum, result, testCase.expected)
		}

	}

}

func TestJsonNation(t *testing.T) {

	TestTable := []struct {
		testNum    int
		name       string
		jsonString string
		expected   string
	}{
		{
			testNum:    1,
			name:       "much country in list",
			jsonString: "{\"count\":1079445,\"name\":\"Anna\",\"country\":[{\"country_id\":\"PL\",\"probability\":0.129},{\"country_id\":\"SE\",\"probability\":0.064},{\"country_id\":\"UA\",\"probability\":0.062},{\"country_id\":\"RU\",\"probability\":0.051},{\"country_id\":\"IT\",\"probability\":0.044}]}",
			expected:   "PL",
		},
		{
			testNum:    2,
			name:       "empty country, if name not exist id base",
			jsonString: "{\"count\":0,\"name\":\"jojojojojojojojojo\",\"country\":[]}",
			expected:   "NO",
		},
	}

	for _, testCase := range TestTable {

		result := NationComputing(testCase.jsonString)

		// Обработчиик ошибок
		if result != testCase.expected {
			t.Errorf("result wrong at test #%v, got [%v] want [%v]", testCase.testNum, result, testCase.expected)
		}

	}

}
