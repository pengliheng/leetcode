package isLongPressedName

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
				one: "alex",
				two: "aaleex",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "leelee",
				two: "lleeelee",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "vtkgn",
				two: "vttkgnn",
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: "vttkgnn",
				two: "vttkgnn",
			},
			a: ans{
				one: true,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isLongPressedName(p.one, p.two), "输入:%v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
