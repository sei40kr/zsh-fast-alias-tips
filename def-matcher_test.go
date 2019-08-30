package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseDef(t *testing.T) {
	assert.Equal(t, ParseDef("dk=docker"), Def{alias: "dk", abbr: "docker"})
	assert.Equal(t, ParseDef("gb='git branch'"), Def{alias: "gb", abbr: "git branch"})
	assert.Equal(t, ParseDef("'g cb'='git checkout -b'"), Def{alias: "g cb", abbr: "git checkout -b"})
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

	{
		candidate, _ := MatchDef(mockDefs, "docker")
		assert.Equal(t, candidate.alias, "dk")
	}

	{
		candidate, _ := MatchDef(mockDefs, "git branch")
		assert.Equal(t, candidate.alias, "gb")
	}

	{
		candidate, _ := MatchDef(mockDefs, "git checkout -b")
		assert.Equal(t, candidate.alias, "gcb", "should return the longest match when multiple matches found")
	}

	{
		candidate, _ := MatchDef(mockDefs, "cd ..")
		assert.Nil(t, candidate, "should return nil when no matches found")
	}

	{
		candidate, _ := MatchDef(mockDefs, "ls -G -lh")
		assert.Equal(t, candidate.alias, "ll", "should apply aliases recursively")
	}
}
