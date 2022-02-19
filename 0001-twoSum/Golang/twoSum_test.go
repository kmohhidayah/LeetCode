package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func twoSum(nums []int, target int) []int {
	var ans []int
	for i := 1; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i-1]+nums[j] == target && i-1 != j {
				ans = append(ans, i-1, j)
				break
			}
		}
	}

	return ans
}

func TestTwoSum(t *testing.T) {
	assert.Equal(t, []int{1, 2}, twoSum([]int{2, 5, 5, 11}, 10))
}
