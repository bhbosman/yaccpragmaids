package yaccpragmaids

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPragmaIdDCE(t *testing.T) {
	t.Run("", func(t *testing.T) {
		lexer := NewYaccYaccPragmaIdsLexerImpl(
			0,
			[]lexem{
				{token: RWDce, data: ""},
				{token: ':', data: ""},
				{token: Value, data: "167641ce-3194-4999-af0e-115fb5630204"},
				{token: ':', data: ""},
				{token: Value, data: "1"},
				{token: '.', data: ""},
				{token: Value, data: "0"},
			})

		parser := YaccPragmaIdsNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetTypeValue())
		assert.Equal(t, "DCE", lexer.GetTypeValue().VersionType())
		if !assert.Implements(t, (*IDceVersion)(nil), lexer.GetTypeValue()) {
			return
		}
		node, ok1 := lexer.GetTypeValue().(IDceVersion)
		assert.True(t, ok1)
		assert.Equal(t, "167641ce-3194-4999-af0e-115fb5630204", node.Identifier())
		assert.Equal(t, uint64(1), node.Value().Major())
		assert.Equal(t, uint64(0), node.Value().Minor())
	})
	t.Run("", func(t *testing.T) {
		lexer := NewYaccYaccPragmaIdsLexerImpl(
			0,
			[]lexem{
				{token: RWDce, data: ""},
				{token: ':', data: ""},
				{token: Value, data: "167641ce-3194-4999-af0e-115fb5630204"},
				{token: ':', data: ""},
				{token: Value, data: "1"},
				{token: '.', data: ""},
			})

		parser := YaccPragmaIdsNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetTypeValue())
		assert.Equal(t, "DCE", lexer.GetTypeValue().VersionType())
		if !assert.Implements(t, (*IDceVersion)(nil), lexer.GetTypeValue()) {
			return
		}
		node, ok1 := lexer.GetTypeValue().(IDceVersion)
		assert.True(t, ok1)
		assert.Equal(t, "167641ce-3194-4999-af0e-115fb5630204", node.Identifier())
		assert.Equal(t, uint64(1), node.Value().Major())
		assert.Equal(t, uint64(0), node.Value().Minor())
	})
	t.Run("", func(t *testing.T) {
		lexer := NewYaccYaccPragmaIdsLexerImpl(
			0,
			[]lexem{
				{token: RWDce, data: ""},
				{token: ':', data: ""},
				{token: Value, data: "167641ce-3194-4999-af0e-115fb5630204"},
				{token: ':', data: ""},
				{token: Value, data: "1"},
			})

		parser := YaccPragmaIdsNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetTypeValue())
		assert.Equal(t, "DCE", lexer.GetTypeValue().VersionType())
		if !assert.Implements(t, (*IDceVersion)(nil), lexer.GetTypeValue()) {
			return
		}
		node, ok1 := lexer.GetTypeValue().(IDceVersion)
		assert.True(t, ok1)
		assert.Equal(t, "167641ce-3194-4999-af0e-115fb5630204", node.Identifier())
		assert.Equal(t, uint64(1), node.Value().Major())
		assert.Equal(t, uint64(0), node.Value().Minor())
	})
}
