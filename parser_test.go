package main

import (
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []Link
	}{
		{
			name: "Single Link",
			input: `
			<a href="/dog">Dog</a>
			`,
			expected: []Link{
				{Href: "/dog", Text: "Dog"},
			},
		},
		{
			name: "Multiple Links",
			input: `
			<a href="/cat">Cat</a>
			<a href="/dog">Dog</a>
			`,
			expected: []Link{
				{Href: "/cat", Text: "Cat"},
				{Href: "/dog", Text: "Dog"},
			},
		},
		{
			name: "Nested Tags in Link",
			input: `
			<a href="/dog"><span>Dog</span> is great!</a>
			`,
			expected: []Link{
				{Href: "/dog", Text: "Dog is great!"},
			},
		},
		{
			name: "Invalid HTML",
			input: `
			<a href="/cat">Cat<a href="/dog">Dog</a>
			`,
			expected: []Link{
				{Href: "/cat", Text: "Cat"},
				{Href: "/dog", Text: "Dog"},
			},
		},
		{
			name: "No Links",
			input: `
			<div>No links here</div>
			`,
			expected: []Link{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := strings.NewReader(tt.input)
			links, err := Parse(r)
			if err != nil {
				t.Fatalf("Parse() error: %v", err)
			}

			if len(links) != len(tt.expected) {
				t.Fatalf("Expected %d links, got %d", len(tt.expected), len(links))
			}

			for i, link := range links {
				if link != tt.expected[i] {
					t.Errorf("Expected link %v, got %v", tt.expected[i], link)
				}
			}
		})
	}
}
