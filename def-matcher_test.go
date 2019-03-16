package main

import (
	"fmt"
	"testing"
)

func TestParseDef(t *testing.T) {
	mockArgs := []struct {
		subject  string
		line     string
		expected Def
	}{
		{
			subject: "when neither of the alias nor the abbreviation are quoted",
			line:    "dk=docker",
			expected: Def{
				alias: "dk",
				abbr:  "docker",
			},
		},
		{
			subject: "when the abbreviation is quoted",
			line:    "gb='git branch'",
			expected: Def{
				alias: "gb",
				abbr:  "git branch",
			},
		},
		{
			subject: "when both of the alias and the abbreviation are quoted",
			line:    "'g cb'='git checkout -b'",
			expected: Def{
				alias: "g cb",
				abbr:  "git checkout -b",
			},
		},
	}

	for _, mockArg := range mockArgs {
		fmt.Printf("%s - ", mockArg.subject)

		actual := ParseDef(mockArg.line)
		expected := mockArg.expected
		if actual == expected {
			fmt.Println("ok")
		} else {
			fmt.Println("ng")
			t.Fatalf("expected=%s, aParseDef(mockArg.line)ctual=%s\n", expected, actual)
		}
	}
}

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
		{
			alias: "ls",
			abbr:  "ls -G",
		},
		{
			alias: "ll",
			abbr:  "ls -lh",
		},
	}

	mockArgs := []struct {
		subject  string
		command  string
		expected *Def
	}{
		{
			subject:  "when the command has single token",
			command:  "docker",
			expected: &Def{alias: "dk"},
		},
		{
			subject:  "when the command has multiple tokens",
			command:  "git branch",
			expected: &Def{alias: "gb"},
		},
		{
			subject:  "when it has more than 2 matches, then return the longest one",
			command:  "git checkout -b",
			expected: &Def{alias: "gcb"},
		},
		{
			subject:  "when it has no matches, then return a empty string",
			command:  "cd ..",
			expected: nil,
		},
		{
			subject:  "when it was expanded recursively from >1 aliases, then reduce it fully",
			command:  "ls -G -lh", // ll expands to that with ls='ls -G' and ll='ls -lh' aliases defined
			expected: &Def{alias: "ll"},
		},
	}

	for _, mockArg := range mockArgs {
		fmt.Printf("%s - ", mockArg.subject)

		expected := mockArg.expected
		actual, _ := MatchDef(mockDefs, mockArg.command)

		if (actual == nil && expected == nil) || actual.alias == expected.alias {
			fmt.Println("ok")
		} else {
			fmt.Println("ng")

			if actual != nil {
				t.Fatalf("expected=%s, actual=%s\n", expected.alias, actual.alias)
			} else {
				t.Fatalf("expected=%s, actual=nil\n", expected.alias)
			}
		}
	}
}
