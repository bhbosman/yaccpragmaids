package yaccpragmaids

type lexem struct {
	token int
	data  string
}

type YaccPragmaIdsLexerImpl struct {
	pos          int
	data         []lexem
	error        bool
	errorMessage string
	version      ITypeVersion
}

func (y *YaccPragmaIdsLexerImpl) GetTypeValue() ITypeVersion {
	return y.version
}

func (y *YaccPragmaIdsLexerImpl) SetTypeValue(version ITypeVersion) {
	y.version = version
}

func NewYaccYaccPragmaIdsLexerImpl(pos int, data []lexem) *YaccPragmaIdsLexerImpl {
	return &YaccPragmaIdsLexerImpl{
		pos:  pos,
		data: data,
	}
}

func (y *YaccPragmaIdsLexerImpl) Lex(lval *YaccPragmaIdsSymType) int {
	y.pos++
	if y.pos == len(y.data)+1 {
		return 0
	}
	switch y.data[y.pos-1].token {
	case Value:
		lval.Value = y.data[y.pos-1].data
		return y.data[y.pos-1].token
	}

	return y.data[y.pos-1].token

}

func (y *YaccPragmaIdsLexerImpl) Error(s string) {
	y.error = true
	y.errorMessage = s
}
