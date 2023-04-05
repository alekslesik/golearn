package main

import (
	"reflect"
	"testing"
)

func TestMain(t *testing.T) {
	testCases := []struct {
		number string
		desc   string
		input  string
		want   string
	}{
		{number: "1", desc: "()[]}", input: "()[]}", want: "5"},
		{number: "2", desc: "{{[()]]", input: "{{[()]]", want: "7"},
		{number: "3", desc: "([](){([])})", input: "([](){([])})", want: "Success"},
		{number: "4", desc: "{{{[][][]", input: "{{{[][][]", want: "3"},
		{number: "5", desc: "{*{{}", input: "{*{{}", want: "3"},
		{number: "6", desc: "[[*", input: "[[*", want: "2"},
		{number: "7", desc: "{*}", input: "{*}", want: "Success"},
		{number: "8", desc: "{{", input: "{{", want: "2"},
		{number: "9", desc: "{}", input: "{}", want: "Success"},
		{number: "10", desc: "", input: "", want: "Success"},
		{number: "11", desc: "}", input: "}", want: "1"},
		{number: "12", desc: "*{}", input: "*{}", want: "Success"},
		{number: "13", desc: "{{{**[][][])", input: "{{{**[][][])", want: "3"},
		{number: "14", desc: "()({}", input: "()({}", want: "3"},
		{number: "15", desc: "{{[()]}", input: "{{[()]}", want: "1"},
		{number: "16", desc: "[]", input: "[]", want: "Success"},
		{number: "17", desc: "{}[]", input: "{}[]", want: "Success"},
		{number: "18", desc: "[()]", input: "[()]", want: "Success"},
		{number: "19", desc: "(())", input: "(())", want: "Success"},
		{number: "20", desc: "{[]}()", input: "{[]}()", want: "Success"},
		{number: "21", desc: "([](){([])})", input: "([](){([])})", want: "Success"},
		{number: "22", desc: "foo(bar);'", input: "foo(bar);'", want: "Success"},
		{number: "23", desc: "{", input: "{", want: "1"},
		{number: "24", desc: "{[}", input: "{[}", want: "3"},
		{number: "25", desc: "()[]}", input: "()[]}", want: "5"},
		{number: "26", desc: "{{[()]]", input: "{{[()]]", want: "7"},
		{number: "27", desc: "foo(bar[i);'", input: "foo(bar[i);'", want: "10"},
		{number: "28", desc: "[]([]", input: "[]([]", want: "3"},
	}

	for _, tC := range testCases {
		got := isBalanced(tC.input)
		if !reflect.DeepEqual(tC.want, got) {
			t.Fatalf("number : %v, decs: %v expected: %v, got: %v", tC.number, tC.desc, tC.want, got)
		}
	}
}
