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

	return readFile(t, golden)
}

func GetGoldenFileJSON(t *testing.T, actual interface{}, expected interface{}, fileName string, shouldUpdate bool) {
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

	err := json.Unmarshal(readFile(t, golden), &expected)

	if err != nil {
		t.Fatal(err)
	}

}

func GetGoldenFileString(t *testing.T, actual string, fileName string, shouldUpdate bool) string {
	golden := filepath.Join("testdata", fileName+".golden")
	if shouldUpdate {
		if err := ioutil.WriteFile(golden, []byte(actual), permission); err != nil {
			t.Fatalf("Error writing golden file for filename=%s: %s", fileName, err)
		}
	}

	return string(readFile(t, golden))
}

func LoadJSONFixture(t *testing.T, fileName string, structToLoad interface{}) {
	expected, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Fatal(err)
	}

	json.Unmarshal(expected, &structToLoad)
}

func readFile(t *testing.T, fileName string) []byte {
	expected, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Fatal(err)
	}
	return expected
}
