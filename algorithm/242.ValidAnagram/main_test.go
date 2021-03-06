package isAnagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one string
	two string
}

type ans struct {
	one bool
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
				one: "anagram",
				two: "nagaram",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "a",
				two: "b",
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: "aa",
				two: "b",
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: "",
				two: "",
			},
			a: ans{
				one: true,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isAnagram(p.one, p.two), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
