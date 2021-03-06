package isSymmetric

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

type para struct {
	one *TreeNode
	two *TreeNode
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
				one: Helper.Ints2TreeNode([]int{1, 2, 2, 3, 4, 4, 3}),
			},
			a: ans{
				one: true,
			},
		},
		question{
			p: para{
				one: Helper.Ints2TreeNode([]int{1, 2, 3}),
			},
			a: ans{
				one: false,
			},
		},
		question{
			p: para{
				one: Helper.Ints2TreeNode([]int{}),
			},
			a: ans{
				one: true,
			},
		},
	}

	for _, q := range qs {
		a, p := q.a, q.p
		ast.Equal(a.one, isSymmetric(p.one), "input: %v", p)
	}

	// ast.Panics(func() { longestPalindrome([]int{}, []int{}) }, "对空切片求中位数，却没有panic")

}
