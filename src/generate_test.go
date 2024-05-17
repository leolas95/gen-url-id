package main

import (
	"reflect"
	"testing"
)

func TestIntToBase62(t *testing.T) {
	// table-driven unit tests

	tests := map[string]struct {
		input    int64
		expected string
	}{
		"randomInput1": {input: 111111111111, expected: "1xHWtjL"},
		"randomInput2": {input: 101010101010, expected: "1mFw5Za"},
		"randomInput3": {input: 999999999999, expected: "HbXm5a3"},
		"randomInput4": {input: 123456789098, expected: "2Al26Xq"},
		"randomInput5": {input: 109238745665, expected: "1vEodwf"},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := IntToBase62(test.input)

			if !reflect.DeepEqual(got, test.expected) {
				t.Fatalf("expected: %s, got: %s", test.expected, got)
			}
		})
	}
}
