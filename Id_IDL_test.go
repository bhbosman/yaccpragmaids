package yaccpragmaids

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPragmaId(t *testing.T) {
	t.Run("", func(t *testing.T) {
		lexer := NewYaccYaccPragmaIdsLexerImpl(
			0,
			[]lexem{})

		parser := YaccPragmaIdsNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		//assert.NotNil(t, lexer.GetPragma())
		//assert.Equal(t, "Id", lexer.GetPragma().Type())
		//if !assert.Implements(t, (*IIdPragmaNode)(nil), lexer.GetPragma()) {
		//	return
		//}
		//idPragmaNode, ok1 := lexer.GetPragma().(IIdPragmaNode)
		//assert.True(t, ok1)
		//assert.Equal(t, "A", idPragmaNode.Identifier())
		//assert.Equal(t, "ABCDEF", idPragmaNode.Value())
	})
}
