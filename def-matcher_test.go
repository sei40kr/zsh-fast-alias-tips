package main

import (
	"fmt"
	"testing"
)

func TestMatchDef(t *testing.T) {
	mockDefs := []Def{
		{
			alias: "dk",
			abbr:  "docker",
		},
		{
			alias: "gb",
			abbr:  "git branch",
		},
		{
			alias: "gco",
			abbr:  "git checkout",
		},
		{
			alias: "gcb",
			abbr:  "git checkout -b",
		},
	}

	mockArgs := []struct {
		subject  string
		command  string
		expected string
	}{
		{
			subject:  "when the command has single token",
			command:  "docker",
			expected: "dk",
		},
		{
			subject:  "when the command has multiple tokens",
			command:  "git branch",
			expected: "gb",
		},
		{
			subject:  "when it has more than 2 matches, then return the longest one",
			command:  "git checkout -b",
			expected: "gcb",
		},
		{
			subject:  "when it has no matches, then return a empty string",
			command:  "cd ..",
			expected: "",
		},
	}

	for _, mockArg := range mockArgs {
		fmt.Printf("%s - ", mockArg.subject)

		expected := mockArg.expected
		actual := MatchDef(mockDefs, mockArg.command)

		if actual == expected {
			fmt.Println("ok")
		} else {
			fmt.Println("ng")
			t.Fatalf("expected=%s, actual=%s\n", expected, actual)
		}
	}
}
