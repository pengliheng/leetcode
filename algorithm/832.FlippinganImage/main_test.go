package flipAndInvertImage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one [][]int
}

type ans struct {
	one [][]int
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
				one: [][]int{
					[]int{1, 1, 0},
					[]int{1, 0, 1},
					[]int{0, 0, 0},
				},
			},
			a: ans{
				one: [][]int{
					[]int{1, 0, 0},
					[]int{0, 1, 0},
					[]int{1, 1, 1},
				},
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, flipAndInvertImage(p.one), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
