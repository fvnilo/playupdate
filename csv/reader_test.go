package csv_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/nylo-andry/playupdate/csv"
)

type testCase struct {
	in  string
	out []string
}

func TestReadFile(t *testing.T) {
	testCases := []testCase{
		{"", []string{}},
		{"mac_addresses, id1, id2, id3", []string{}},
		{`mac_addresses, id1, id2, id3
a1:bb:cc:dd:ee:ff, 1, 2, 3
a2:bb:cc:dd:ee:ff, 1, 2, 3`, []string{"a1:bb:cc:dd:ee:ff", "a2:bb:cc:dd:ee:ff"}},
	}

	for _, tt := range testCases {
		t.Run(tt.in, func(t *testing.T) {
			macs, _ := csv.ReadFile(strings.NewReader(tt.in))
			if !reflect.DeepEqual(tt.out, macs) {
				t.Fatalf("expected: %v, got: %v", tt.out, macs)
			}
		})
	}
}
