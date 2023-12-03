package matcher

import (
	"testing"

	"github.com/sei40kr/zsh-fast-alias-tips/model"
	"github.com/stretchr/testify/assert"
)

var mockAliasDefs = []model.AliasDef{
	{Name: "dk", Abbr: "docker"},
	{Name: "gb", Abbr: "git branch"},
	{Name: "gco", Abbr: "git checkout"},
	{Name: "gcb", Abbr: "git checkout -b"},
	{Name: "ls", Abbr: "ls -G"},
	{Name: "ll", Abbr: "ls -lh"},
	{Name: "yst", Abbr: "yarn start"},
}

func TestMatch_NoMatches(t *testing.T) {
	aliasTips := Match(mockAliasDefs, "cd ..")
	assert.Equal(t, aliasTips, "cd ..")
}

func TestMatch_NoMatches2(t *testing.T) {
	aliasTips := Match(mockAliasDefs, "yarn start :ja")
	assert.Equal(t, aliasTips, "yarn start:ja")
}

func TestMatch_SingleToken(t *testing.T) {
	aliasTips := Match(mockAliasDefs, "docker")
	assert.Equal(t, aliasTips, "dk")
}

func TestMatch_MultipleTokens(t *testing.T) {
	aliasTips := Match(mockAliasDefs, "git branch")
	assert.Equal(t, aliasTips, "gb")
}

func TestMatch_MultipleMatches(t *testing.T) {
	aliasTips := Match(mockAliasDefs, "git checkout -b")
	assert.Equal(t, aliasTips, "gcb",
		"should return the alias definition that has the longest abbreviation when multiple matches found")
}

func TestMatch_RecursiveDefs(t *testing.T) {
	aliasTips := Match(mockAliasDefs, "ls -G -lh")
	assert.Equal(t, aliasTips, "ll", "should apply aliases recursively")
}

func TestMatch_RecursiveDefsWithPartialMatch(t *testing.T) {
	aliasTips := Match(mockAliasDefs, "ls -G -lh -a")
	assert.Equal(t, aliasTips, "ll -a", "should apply aliases recursively")
}
