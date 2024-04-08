package main

import "demo-stalefish/analyzer"

// 転置インデックス
type InvertedIndex map[analyzer.TokenID][]PostingList

type PostingList struct {
	// トークンごとのポスティングリスト
	Postings *Postings
}

// ポスティング
type Postings struct {
	// 各ドキュメントを識別するID
	DocumentID DocumentID
	// ドキュメント内のトークンの位置
	Positions []uint64
	Next      *Postings
}

// PushBack はとあるポスティングの直後に、引数のポスティングを追加する
func (p *Postings) PushBack(e *Postings) {
	e.Next = p.Next
	p.Next = e
}

// Size はポスティングリストのサイズを返す。
func (p PostingList) Size() int {
	size := 0
	pp := p.Postings
	for pp != nil {
		pp = pp.Next
		size++
	}
	return size
}

// AppearanceCountInDocument は、指定されたドキュメント内での出現回数を返す。
func (p PostingList) AppearanceCountInDocument(docID DocumentID) int {
	count := 0
	pp := p.Postings
	for pp != nil {
		if pp.DocumentID == docID {
			count = len(pp.Positions)
			break
		}
		pp = pp.Next
	}
	return count
}

type Indexer struct {
	storage Storage
	// 文章分割のためのアナライザ
	analyzer analyzer.Analyzer
	// メモリ上の転置インデックス
	InvertedIndex
	// メモリ上の転置インデックスをストレージ上の転置インデックスへマージする際の閾値
	indexSizeThreshold int
}
