package tree_sitter_qmljs_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_qmljs "git+github.com/yuja/tree-sitter-qmljs.git/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_qmljs.Language())
	if language == nil {
		t.Errorf("Error loading Qmljs grammar")
	}
}
