package rack_test

import (
	"errors"
	"flag"
	"testing"

	"github.com/apiotrowski312/scrabbleBot/rack"
	"github.com/apiotrowski312/scrabbleBot/utils/test_utils"
	"github.com/bmizerany/assert"
)

var update = flag.Bool("update", false, "update the golden files of this test")

func Test_CreateRack(t *testing.T) {
	createdRack := rack.CreateRack(5)

	var expected rack.Rack
	test_utils.GetGoldenFileJSON(t, createdRack, &expected, t.Name(), *update)

	assert.Equal(t, expected, createdRack)
}

func TestRack_RemoveFromRack(t *testing.T) {
	type args struct {
		in0 []rune
	}
	tests := []struct {
		name        string
		fixtureName string
		args        args
		wantErr     error
	}{
		{"empty_rack", "testdata/empty_rack.fixture", args{[]rune{'a'}}, errors.New("Empty rack, you cannot remove any letter from it")},
		{"full_rack", "testdata/full_rack.fixture", args{[]rune{'a'}}, nil},
		{"full_rack_error", "testdata/full_rack.fixture", args{[]rune{'z'}}, errors.New("There is no tile z in rack")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r rack.Rack
			var expected rack.Rack

			test_utils.LoadJSONFixture(t, tt.fixtureName, &r)

			err := r.RemoveFromRack(tt.args.in0)

			test_utils.GetGoldenFileJSON(t, r, &expected, t.Name()+"_"+tt.name, *update)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, expected, r)
		})
	}
}

func TestRack_AddFromRack(t *testing.T) {
	type args struct {
		in0 []rune
	}
	tests := []struct {
		name        string
		fixtureName string
		args        args
		wantErr     error
	}{
		{"empty_rack", "testdata/empty_rack.fixture", args{[]rune{'a'}}, nil},
		{"full_rack", "testdata/full_rack.fixture", args{[]rune{'a'}}, errors.New("Rack will be overfilled")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r rack.Rack
			var expected rack.Rack

			test_utils.LoadJSONFixture(t, tt.fixtureName, &r)

			err := r.AddToRack(tt.args.in0)

			test_utils.GetGoldenFileJSON(t, r, &expected, t.Name()+"_"+tt.name, *update)
			assert.Equal(t, tt.wantErr, err)
			assert.Equal(t, expected, r)
		})
	}
}

func TestRack_AreThereLetters(t *testing.T) {
	type args struct {
		in0 []rune
	}
	tests := []struct {
		name        string
		fixtureName string
		letters     []rune
		isOk        bool
		err         error
	}{
		{"empty_rack", "testdata/empty_rack.fixture", []rune{'a'}, false, errors.New("There is no such letters")},
		{"full_rack", "testdata/full_rack.fixture", []rune{'a'}, true, nil},
		{"full_rack", "testdata/full_rack.fixture", []rune{'a', 'z'}, false, errors.New("There is no such letters")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var r rack.Rack
			test_utils.LoadJSONFixture(t, tt.fixtureName, &r)
			isOk, err := r.AreThereLetters(tt.letters)
			assert.Equal(t, tt.isOk, isOk)
			assert.Equal(t, tt.err, err)
		})
	}
}
