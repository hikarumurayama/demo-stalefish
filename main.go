package main

import (
	"fmt"

	"github.com/kotaroooo0/stalefish"
)

func main() {
	analyzer := stalefish.NewAnalyzer(
		[]stalefish.CharFilter{
			stalefish.NewMappingCharFilter(
				map[string]string{
					":(": "sad",
				},
			),
		},
		stalefish.NewStandardTokenizer(), // スペースで文章を区切る
		[]stalefish.TokenFilter{
			stalefish.NewLowercaseFilter(), // 小文字に変換
			stalefish.NewStemmerFilter(),   // 語幹抽出。
			stalefish.NewStopWordFilter( // 不要な単語を取り除く
				[]string{
					"i", "my", "me", "the", "a", "for",
				},
			),
		},
	)
	fmt.Println(analyzer.Analyze("I feel TIRED :(")) // {[{0 feel } {0 tire } {0 sad }]}
}
