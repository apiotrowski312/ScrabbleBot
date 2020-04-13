package test_utils

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"testing"
)

const (
	permission = 0666
)

// TODO: Create dir if not exist
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
		bytes, err := json.MarshalIndent(actual, "", "\t")

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
	expected, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Fatal(err)
	}

	json.Unmarshal(expected, &structToLoad)
}

func BytesToStruct(t *testing.T, bytes []byte, structToLoad interface{}) {
	err := json.Unmarshal(bytes, &structToLoad)

	if err != nil {
		t.Fatal(err)
	}
}
