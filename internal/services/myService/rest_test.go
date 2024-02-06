package services

import (
	"testing"
)

func TestExternalAPI(t *testing.T) {

	testTable := []struct {
		testNum int
		name    string
		age     int
		gender  string
		nation  string
	}{
		{
			testNum: 1,
			name:    "alex",
			age:     47,
			gender:  "male",
			nation:  "CZ",
		},
		{
			testNum: 2,
			name:    "anna",
			age:     52,
			gender:  "female",
			nation:  "PL",
		},
		{
			testNum: 3,
			name:    "jojo",
			age:     58,
			gender:  "male",
			nation:  "SA",
		},
	}

	for _, tc := range testTable {

		chann1 := make(chan int)
		chann2 := make(chan string)
		chann3 := make(chan string)

		go getAgeFromName(tc.name, chann1)
		go getSexFromName(tc.name, chann2)
		go getNationalityFromName(tc.name, chann3)

		res1 := <-chann1
		res2 := <-chann2
		res3 := <-chann3

		if res1 != tc.age {
			t.Errorf("result wrong at test #%v, got [%v] want [%v]", tc.testNum, res1, tc.age)
		}
		if res2 != tc.gender {
			t.Errorf("result wrong at test #%v, got [%v] want [%v]", tc.testNum, res2, tc.gender)
		}
		if res3 != tc.nation {
			t.Errorf("result wrong at test #%v, got [%v] want [%v]", tc.testNum, res3, tc.nation)
		}

	}
}

func TestWebConnection(t *testing.T) {

	testTable := []struct {
		testNum int
		name    string
		site    string
		param   string
		result  string
	}{
		{
			testNum: 1,
			name:    "agify",
			site:    "https://api.agify.io/",
			param:   "!2#%6^gf999ы",
			result:  "{\"count\":0,\"name\":\"!2#%6^gf999ы\",\"age\":null}",
		},
		{
			testNum: 2,
			name:    "agify",
			site:    "https://api.genderize.io/",
			param:   "!2#%6^gf999ы",
			result:  "{\"count\":0,\"name\":\"!2#%6^gf999ы\",\"gender\":null,\"probability\":0.0}",
		},
		{
			testNum: 3,
			name:    "agify",
			site:    "https://api.nationalize.io/",
			param:   "!2#%6^gf999ы",
			result:  "{\"count\":0,\"name\":\"!2#%6^gf999ы\",\"country\":[]}",
		},
	}

	for _, tc := range testTable {

		res := sendRequestToGet(tc.site, tc.param)

		if res != tc.result {
			t.Errorf("result wrong at test #%v, got [%v] want [%v]", tc.testNum, res, tc.result)
		}

	}
}
