package findAnagrams

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
	two string
}

type ans struct {
	one []int
}

type question struct {
	p para
	a ans
}

func Test_OK(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			p: para{
				one: "abab",
				two: "ab",
			},
			a: ans{
				one: []int{0, 1, 2},
			},
		},
		question{
			p: para{
				one: "cbaebabacd",
				two: "abc",
			},
			a: ans{
				one: []int{0, 6},
			},
		},
		question{
			p: para{
				one: "baa",
				two: "aa",
			},
			a: ans{
				one: []int{1},
			},
		},
		question{
			p: para{
				one: "acdcaeccde",
				two: "c",
			},
			a: ans{
				one: []int{1, 3, 6, 7},
			},
		},
		question{
			p: para{
				one: "ababababab",
				two: "aab",
			},
			a: ans{
				one: []int{0, 2, 4, 6},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, findAnagrams(p.one, p.two), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}