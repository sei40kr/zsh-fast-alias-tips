package parser

import (
	"testing"

	"github.com/sei40kr/zsh-fast-alias-tips/model"
	"github.com/stretchr/testify/assert"
)

func TestParse_NoQuotesInRightExp(t *testing.T) {
	assert.Equal(t, Parse("dk=docker"), model.AliasDef{Name: "dk", Abbr: "docker"})
}

func TestParse_QuotesInRightExp(t *testing.T) {
	assert.Equal(t, Parse("gb='git branch'"), model.AliasDef{Name: "gb", Abbr: "git branch"})
}

func TestParse_QuotesInBothExps(t *testing.T) {
	assert.Equal(t, Parse("'g cb'='git checkout -b'"), model.AliasDef{Name: "g cb", Abbr: "git checkout -b"})
}
