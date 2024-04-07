package analyzer

import (
	"github.com/kotaroooo0/stalefish/morphology"
)

type Tokenizer interface {
	Tokenize(string) TokenStream
}

// 形態素解析機
type MorphologicalTokenizer struct {
	morphology morphology.Morphology
}

func NewMorphologicalTokenizer(morphology morphology.Morphology) MorphologicalTokenizer {
	return MorphologicalTokenizer{
		morphology: morphology,
	}
}

// Tokenize は与えられた文字列に対し、形態素解析を行う。返り値には、形態素ごとに分割された文字列を返す。
func (t MorphologicalTokenizer) Tokenize(s string) TokenStream {
	mTokens := t.morphology.Analyze(s)
	tokens := make([]Token, len(mTokens))
	for i, t := range mTokens {
		tokens[i] = NewToken(t.Term, SetKana(t.Kana))
	}
	return NewTokenStream(tokens)
}

type NgramTokenizer struct {
	n int
}

// Tokenizeは、引数に与えられた文字列に対してNgram解析を行う。NgramTokenizerが保持するN数ごとに分割された文字列を返す。
func (t NgramTokenizer) Tokenize(s string) TokenStream {
	count := len([]rune(s)) + 1 - t.n
	tokens := make([]Token, count)
	for i := 0; i < count; i++ {
		tokens[i] = NewToken(string([]rune(s)[i : i+t.n]))
	}
	return NewTokenStream(tokens)
}
