package yaccpragmaids

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPragmaIdRmi(t *testing.T) {
	t.Run("", func(t *testing.T) {
		lexer := NewYaccYaccPragmaIdsLexerImpl(
			0,
			[]lexem{
				{token: RWRmi, data: ""},
				{token: ':', data: ""},
				{token: Value, data: "abc"},
				{token: ':', data: ""},
				{token: Value, data: "1234567812345678"},
			})

		parser := YaccPragmaIdsNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetTypeValue())
		assert.Equal(t, "RMI", lexer.GetTypeValue().VersionType())
		if !assert.Implements(t, (*IRmiVersion)(nil), lexer.GetTypeValue()) {
			return
		}
		node, ok1 := lexer.GetTypeValue().(IRmiVersion)
		assert.True(t, ok1)
		assert.Equal(t, "abc", node.ClassName())
		assert.Equal(t, uint64(0x1234567812345678), node.Hash())
	})
	t.Run("", func(t *testing.T) {
		lexer := NewYaccYaccPragmaIdsLexerImpl(
			0,
			[]lexem{
				{token: RWRmi, data: ""},
				{token: ':', data: ""},
				{token: Value, data: "abc"},
				{token: ':', data: ""},
				{token: Value, data: "0123456789ABCDEF"},
				{token: ':', data: ""},
				{token: Value, data: "FFFFFFFFFFFFFFFF"},
			})

		parser := YaccPragmaIdsNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetTypeValue())
		assert.Equal(t, "RMI", lexer.GetTypeValue().VersionType())
		if !assert.Implements(t, (*IRmiVersion)(nil), lexer.GetTypeValue()) {
			return
		}
		node, ok1 := lexer.GetTypeValue().(IRmiVersion)
		assert.True(t, ok1)
		assert.Equal(t, "abc", node.ClassName())
		assert.Equal(t, uint64(0x0123456789ABCDEF), node.Hash())
		assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFF), node.SerializationVersionUID())
	})
}
