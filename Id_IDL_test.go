package yaccpragmaids

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPragmaIdIdl(t *testing.T) {
	t.Run("", func(t *testing.T) {
		lexer := NewYaccYaccPragmaIdsLexerImpl(
			0,
			[]lexem{
				{token: RWIdl, data: ""},
				{token: ':', data: ""},
				{token: Value, data: "A"},
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
		assert.Equal(t, "IDL", lexer.GetTypeValue().VersionType())
		if !assert.Implements(t, (*IIdlVersion)(nil), lexer.GetTypeValue()) {
			return
		}
		node, ok1 := lexer.GetTypeValue().(IIdlVersion)
		assert.True(t, ok1)
		assert.Equal(t, "A", node.Identifier())
		assert.Equal(t, uint64(1), node.Value().Major())
		assert.Equal(t, uint64(0), node.Value().Minor())
	})
	t.Run("", func(t *testing.T) {
		lexer := NewYaccYaccPragmaIdsLexerImpl(
			0,
			[]lexem{
				{token: RWIdl, data: ""},
				{token: ':', data: ""},
				{token: Value, data: "A"},
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
		assert.Equal(t, "IDL", lexer.GetTypeValue().VersionType())
		if !assert.Implements(t, (*IIdlVersion)(nil), lexer.GetTypeValue()) {
			return
		}
		node, ok1 := lexer.GetTypeValue().(IIdlVersion)
		assert.True(t, ok1)
		assert.Equal(t, "A", node.Identifier())
		assert.Equal(t, uint64(1), node.Value().Major())
		assert.Equal(t, uint64(0), node.Value().Minor())
	})
	t.Run("", func(t *testing.T) {
		lexer := NewYaccYaccPragmaIdsLexerImpl(
			0,
			[]lexem{
				{token: RWIdl, data: ""},
				{token: ':', data: ""},
				{token: Value, data: "A"},
				{token: ':', data: ""},
				{token: Value, data: "1"},
			})

		parser := YaccPragmaIdsNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetTypeValue())
		assert.Equal(t, "IDL", lexer.GetTypeValue().VersionType())
		if !assert.Implements(t, (*IIdlVersion)(nil), lexer.GetTypeValue()) {
			return
		}
		node, ok1 := lexer.GetTypeValue().(IIdlVersion)
		assert.True(t, ok1)
		assert.Equal(t, "A", node.Identifier())
		assert.Equal(t, uint64(1), node.Value().Major())
		assert.Equal(t, uint64(0), node.Value().Minor())
	})
}
