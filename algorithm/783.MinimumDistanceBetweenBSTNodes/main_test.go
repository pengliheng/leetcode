package minDiffInBST

import (
	"testing"

	"github.com/pengliheng/leetcode/Helper"

	"github.com/stretchr/testify/assert"
)

var tcs = []struct {
	N1  *TreeNode
	ans int
}{
	{
		Helper.Ints2TreeNode([]int{4, 2, 6, 1, 3}),
		1,
	},
}

func Test_bitwiseComplement(t *testing.T) {
	ast := assert.New(t)
	for _, tc := range tcs {
		ast.Equal(tc.ans, minDiffInBST(tc.N1), "输入:%v", tc)
	}
}
