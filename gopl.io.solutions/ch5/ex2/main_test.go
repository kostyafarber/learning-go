package main

import (
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func parseHtml(s string) *html.Node {
	r := strings.NewReader(s)
	html, _ := html.Parse(r)
	return html
}

func TestTraverse(t *testing.T) {
	type tagCount map[string]int
	type test struct {
		h    string
		got  tagCount
		want tagCount
	}

	tests := []test{
		{h: `<h1>hey there </h1>
				<p>
				This is a paragraph
				</p>`,
			got: tagCount{},
			want: tagCount{"body": 1,
				"html": 1,
				"h1":   1,
				"p":    1,
				"head": 1,
			}},
	}

	for _, tc := range tests {
		parsed := parseHtml(tc.h)
		traverse(parsed, tc.got)

		if !reflect.DeepEqual(tc.got, tc.want) {
			t.Fatalf("expected: %v, got: %v", tc.want, tc.got)
		}
	}
}
