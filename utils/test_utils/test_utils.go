package test_utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

const (
	permission = 0666
)

func GetGoldenFile(t *testing.T, actual []byte, fileName string, shouldUpdate bool) []byte {

	golden := filepath.Join("testdata", fileName+".golden")
	if shouldUpdate {
		if err := ioutil.WriteFile(golden, actual, permission); err != nil {
			t.Fatalf("Error writing golden file for filename=%s: %s", fileName, err)
		}
	}
	expected, err := ioutil.ReadFile(golden)
	if err != nil {
		t.Fatal(err)
	}
	return expected
}

func GetGoldenFileJSON(t *testing.T, actual interface{}, fileName string, shouldUpdate bool) []byte {
	golden := filepath.Join("testdata", fileName+".golden")

	if shouldUpdate {
		bytes, err := json.Marshal(actual)

		if err != nil {
			t.Fatalf("Error while marshal a struct %v: %s", actual, err)
		}

		if err = ioutil.WriteFile(golden, bytes, permission); err != nil {
			t.Fatalf("Error writing golden file for filename=%s: %s", fileName, err)
		}
	}
	expected, err := ioutil.ReadFile(golden)
	if err != nil {
		t.Fatal(err)
	}
	return expected
}

func GetGoldenFileString(t *testing.T, actual string, fileName string, shouldUpdate bool) string {
	golden := filepath.Join("testdata", fileName+".golden")
	if shouldUpdate {
		if err := ioutil.WriteFile(golden, []byte(actual), permission); err != nil {
			t.Fatalf("Error writing golden file for filename=%s: %s", fileName, err)
		}
	}
	expected, err := ioutil.ReadFile(golden)
	if err != nil {
		t.Fatal(err)
	}
	return string(expected)
}

func LoadJSONFixture(t *testing.T, fileName string, structToLoad interface{}) {
	jsonFile, err := os.Open(fileName)

	if err != nil {
		t.Fatal(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &structToLoad)
}

func BytesToStruct(t *testing.T, bytes []byte, structToLoad interface{}) {
	err := json.Unmarshal(bytes, &structToLoad)

	if err != nil {
		t.Fatal(err)
	}
}
