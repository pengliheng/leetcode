package plusOne

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two int
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
				one: []int{1, 2, 4},
			},
			a: ans{
				one: []int{1, 2, 5},
			},
		},
		question{
			p: para{
				one: []int{9},
			},
			a: ans{
				one: []int{1, 0},
			},
		},
		question{
			p: para{
				one: []int{9, 9},
			},
			a: ans{
				one: []int{1, 0, 0},
			},
		},
		question{
			p: para{
				one: []int{1, 9},
			},
			a: ans{
				one: []int{2, 0},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, plusOne(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
