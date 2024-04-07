package char_filter_test

import (
	"demo-stalefish/char_filter"
	"testing"

	gocmp "github.com/google/go-cmp/cmp"
)

func Test_MappingCharFilter_Filter(t *testing.T) {
	cases := map[string]struct {
		text     string
		mapper   char_filter.MappingCharFilter
		expected string
	}{
		"no change": {
			text: "I feel TIRED :(",
			mapper: char_filter.MappingCharFilter{
				Mapper: map[string]string{
					":(": "sad",
				},
			},
			expected: "I feel TIRED sad",
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			if diff := gocmp.Diff(c.mapper.Filter(c.text), c.expected); diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			}
		})
	}
}
