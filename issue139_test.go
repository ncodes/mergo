package mergo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIssue139(t *testing.T) {
	srcs := []map[string]interface{}{
		{
			"h": 10,
			"i": "i",
			"j": "j",
		},
		{
			"a": 1,
			"b": 2,
			"d": map[string]interface{}{
				"e": "four",
			},
			"g": []int{6, 7},
			"i": "aye",
			"j": "jay",
			"k": map[string]interface{}{
				"l": false,
			},
		},
	}
	dst := map[string]interface{}{
		"a": "one",
		"c": 3,
		"d": map[string]interface{}{
			"f": 5,
		},
		"g": []int{8, 9},
		"i": "eye",
		"k": map[string]interface{}{
			"l": true,
		},
	}

	for _, src := range srcs {
		if err := Merge(&dst, src); err != nil {
			t.Fatal(err)
		}
	}

	expected := map[string]interface{}{
		"a": "one", // key overridden
		"b": 2,     // merged from src1
		"c": 3,     // merged from dst
		"d": map[string]interface{}{ // deep merge
			"e": "four",
			"f": 5,
		},
		"g": []int{8, 9}, // overridden - arrays are not merged
		"h": 10,          // merged from src2
		"i": "eye",       // overridden twice
		"j": "jay",       // overridden and merged
		"k": map[string]interface{}{
			"l": true, // overridden
		},
	}
	assert.Equal(t, expected, dst)
}
