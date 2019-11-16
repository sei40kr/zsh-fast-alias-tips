package matcher

import (
	"testing"

	"github.com/sei40kr/zsh-fast-alias-tips/model"
	"github.com/stretchr/testify/assert"
)

var mockAliasDefs = []model.AliasDef{
	{
		Name: "dk",
		Abbr: "docker",
	},
	{
		Name: "gb",
		Abbr: "git branch",
	},
	{
		Name: "gco",
		Abbr: "git checkout",
	},
	{
		Name: "gcb",
		Abbr: "git checkout -b",
	},
	{
		Name: "ls",
		Abbr: "ls -G",
	},
	{
		Name: "ll",
		Abbr: "ls -lh",
	},
}

func TestMatch_NoMatches(t *testing.T) {
	candidate, _ := Match(mockAliasDefs, "cd ..")
	assert.Nil(t, candidate, "should return nil when no matches found")
}

func TestMatch_SingleToken(t *testing.T) {
	candidate, _ := Match(mockAliasDefs, "docker")
	assert.Equal(t, candidate.Name, "dk")
}

func TestMatch_MultipleTokens(t *testing.T) {
	candidate, _ := Match(mockAliasDefs, "git branch")
	assert.Equal(t, candidate.Name, "gb")
}

func TestMatch_MultipleMatches(t *testing.T) {
	candidate, _ := Match(mockAliasDefs, "git checkout -b")
	assert.Equal(t, candidate.Name, "gcb",
		"should return the alias definition that has the longest abbreviation when multiple matches found")
}

func TestMatch_RecursiveDefs(t *testing.T) {
	candidate, _ := Match(mockAliasDefs, "ls -G -lh")
	assert.Equal(t, candidate.Name, "ll", "should apply aliases recursively")
}
