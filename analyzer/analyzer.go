package analyzer

import (
	"demo-stalefish/char_filter"
)

type Analyzer struct {
	charFilters  []char_filter.CharFilter
	tokenizer    Tokenizer
	tokenFilters []TokenFilter
}

// 各トークンを識別するID
type TokenID uint64

type Token struct {
	ID   TokenID
	Term string // トークンの文字列。
	Kana string // トークンが感じだった場合の読み仮名。
}

type TokenOption func(*Token)

func NewToken(s string, options ...TokenOption) Token {
	token := Token{
		Term: s,
	}
	for _, option := range options {
		option(&token)
	}
	return token
}

func SetKana(k string) TokenOption {
	return func(t *Token) {
		t.Kana = k
	}
}

type TokenStream struct {
	Tokens []Token
}

func NewTokenStream(t []Token) TokenStream {
	return TokenStream{t}
}

func (ts TokenStream) Size() int {
	return len(ts.Tokens)
}

// Analyze は、引数に与えられた文字列に対してアナライズを実行し、
func (a Analyzer) Analyze(s string) TokenStream {
	for _, cf := range a.charFilters {
		s = cf.Filter(s)
	}
	tokenStream := a.tokenizer.Tokenize(s)
	for _, f := range a.tokenFilters {
		tokenStream = f.Filter(tokenStream)
	}
	return tokenStream
}
