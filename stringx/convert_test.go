package stringx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCamelToSnake(t *testing.T) {
	testCases := []struct {
		name      string
		camelCase string

		want string
	}{
		{
			name:      "empty string",
			camelCase: "",
			want:      "",
		},
		{
			name:      "not camel string",
			camelCase: "User",
			want:      "user",
		},
		{
			name:      "camel string",
			camelCase: "UserName",
			want:      "user_name",
		},
		{
			name:      "complex camel string",
			camelCase: "ThisUserName",
			want:      "this_user_name",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			snake := CamelToSnake(tc.camelCase)
			assert.Equal(t, tc.want, snake)
		})
	}
}

func TestBigCamelToSmallCamel(t *testing.T) {
	testCases := []struct {
		name     string
		bigCamel string

		want string
	}{
		{
			name:     "empty string",
			bigCamel: "",
			want:     "",
		},
		{
			name:     "not camel string",
			bigCamel: "User",
			want:     "user",
		},
		{
			name:     "small camel string",
			bigCamel: "userName",
			want:     "userName",
		},
		{
			name:     "big camel string",
			bigCamel: "UserName",
			want:     "userName",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			smallCamel := BigCamelToSmallCamel(tc.bigCamel)
			assert.Equal(t, tc.want, smallCamel)
		})
	}
}

func TestCapitalizeFirstLetter(t *testing.T) {
	testCases := []struct {
		name  string
		input string

		want string
	}{
		{
			name:  "empty string",
			input: "",
			want:  "",
		},
		{
			name:  "not capitalized string",
			input: "user",
			want:  "User",
		},
		{
			name:  "already capitalized string",
			input: "User",
			want:  "User",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			capitalized := CapitalizeFirstLetter(tc.input)
			assert.Equal(t, tc.want, capitalized)
		})
	}
}
