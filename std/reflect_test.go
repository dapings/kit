package std

import (
	"reflect"
	"testing"
)

func TestGetStructFieldValue(t *testing.T) {
	type inner struct {
		name    string
		Age     int
		country string
	}

	testCase := inner{
		name:    "Doe",
		Age:     29,
		country: "USA",
	}

	if val := getStructFieldValue(testCase, "name"); val != nil {
		t.Errorf("fileld: name, should be empty")
	}

	if val := getStructFieldValue(testCase, "Age"); val == nil {
		t.Errorf("fileld: Age, should not empty")
	}

	val := getStructFieldValue(testCase, "Age")
	t.Log("Age field value type: ", reflect.TypeOf(val))

	if val := getStructFieldValue(testCase, "country"); val != nil {
		t.Errorf("fileld: country, should be empty")
	}

	if val := getStructFieldValue(testCase, "unknown"); val != nil {
		t.Errorf("unknown fileld, should be empty")
	}
}
