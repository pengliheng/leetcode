package fairCandySwap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one []int
	two []int
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
				one: []int{1, 1},
				two: []int{2, 2},
			},
			a: ans{
				one: []int{1, 2},
			},
		},
		question{
			p: para{
				one: []int{2},
				two: []int{1, 3},
			},
			a: ans{
				one: []int{2, 3},
			},
		},
		question{
			p: para{
				one: []int{222222},
				two: []int{1, 3},
			},
			a: ans{
				one: []int{},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, fairCandySwap(p.one, p.two), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
