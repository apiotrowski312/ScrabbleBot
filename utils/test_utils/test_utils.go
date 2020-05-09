package test_utils

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

const (
	permission = 0666
)

func GetGoldenFileJSON(t *testing.T, actual interface{}, expected interface{}, fileName string, shouldUpdate bool) {
	golden := filepath.Join("testdata", fileName)

	if shouldUpdate {
		bytes, err := json.MarshalIndent(actual, "", "\t")

		if err != nil {
			t.Fatalf("Error while marshal a struct %v: %s", actual, err)
		}
		writeFile(t, golden, bytes)
	}

	err := json.Unmarshal(readFile(t, golden), &expected)

	if err != nil {
		t.Fatal(err)
	}

}

func GetGoldenFileString(t *testing.T, actual string, fileName string, shouldUpdate bool) string {
	golden := filepath.Join("testdata", fileName)
	if shouldUpdate {
		writeFile(t, golden, []byte(actual))
	}

	return string(readFile(t, golden))
}

func LoadJSONFixture(t *testing.T, fileName string, structToLoad interface{}) {
	expected := readFile(t, fileName)
	json.Unmarshal(expected, &structToLoad)
}

func readFile(t *testing.T, fileName string) []byte {
	expected, err := ioutil.ReadFile(preparePath(fileName))
	if err != nil {
		t.Fatal(err)
	}
	return expected
}

func writeFile(t *testing.T, fileName string, bytes []byte) {
	if err := ioutil.WriteFile(preparePath(fileName), bytes, permission); err != nil {
		t.Fatalf("Error writing golden file for filename=%s: %s", fileName, err)
	}
}

func preparePath(path string) string {
	return strings.ReplaceAll(path, " ", "_") + ".golden"
}
