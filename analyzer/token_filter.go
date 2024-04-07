package analyzer

import (
	"strings"

	"github.com/kotaroooo0/gojaconv/jaconv"
)

type TokenFilter interface {
	Filter(TokenStream) TokenStream
}

type LowercaseFilter struct{}

func (f LowercaseFilter) Filter(tokenStream TokenStream) TokenStream {
	r := make([]Token, tokenStream.Size())
	for i, token := range tokenStream.Tokens {
		lower := strings.ToLower(token.Term)
		r[i] = NewToken(lower)
	}
	return NewTokenStream(r)
}

type StopWordFilter struct {
	stopWords []string
}

// Filter は、引数のTokenStreamから不要な文字列を削除します。
func (f StopWordFilter) Filter(tokenStream TokenStream) TokenStream {
	stopWords := make(map[string]struct{})
	for _, w := range f.stopWords {
		stopWords[w] = struct{}{}
	}
	r := make([]Token, tokenStream.Size())
	for _, token := range tokenStream.Tokens {
		if _, ok := stopWords[token.Term]; !ok {
			r = append(r, token)
		}
	}
	return NewTokenStream(r)
}

type RomajiReadingformFilter struct{}

// Filter は、引数のTokenStreamの文字列(漢字)をローマ字読みに変換します。変換するためには、形態素解析によって読み仮名をｓ取得する必要があります。
func (f RomajiReadingformFilter) Filter(tokenStream TokenStream) TokenStream {
	for i, token := range tokenStream.Tokens {
		tokenStream.Tokens[i].Term = jaconv.ToHebon(jaconv.KatakanaToHiragana(token.Term))
	}
	return tokenStream
}
