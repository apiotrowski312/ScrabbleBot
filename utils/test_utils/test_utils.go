package test_utils

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func LoadJSONFixture(t *testing.T, fileName string, structToLoad interface{}) {
	expected, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Fatal(err)
	}
	json.Unmarshal(expected, &structToLoad)
}
