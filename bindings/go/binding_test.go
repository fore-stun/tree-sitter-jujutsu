package tree_sitter_jujutsu_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/fore-stun/tree-sitter-jujutsu"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_jujutsu.Language())
	if language == nil {
		t.Errorf("Error loading Jujutsu grammar")
	}
}
