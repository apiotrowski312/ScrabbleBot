package test_utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"testing"
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
	expected, err := ioutil.ReadFile(fileName)
	if err != nil {
		t.Fatal(err)
	}
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
	syscall.Umask(0)
	if _, err := os.Stat(prepareDirPath(fileName)); err != nil {
		if err := os.MkdirAll(prepareDirPath(fileName), 0777); err != nil {
			t.Fatalf("Error direcotry for golden files %s: %s", fileName, err)
		}
	}

	if err := ioutil.WriteFile(preparePath(fileName), bytes, 0666); err != nil {
		t.Fatalf("Error writing golden file for filename=%s: %s", preparePath(fileName), err)
	}
}

func preparePath(path string) string {
	return strings.ReplaceAll(path, " ", "_") + ".golden"
}

func prepareDirPath(path string) string {
	dir := strings.Split(strings.ReplaceAll(path, " ", "_"), "/")
	return strings.Join(dir[:len(dir)-1], "/")
}
