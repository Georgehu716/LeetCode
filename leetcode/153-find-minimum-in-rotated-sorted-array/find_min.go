package leetcode

func findMin(nums []int) int {
        left, right := -1, len(nums) - 1
        for left + 1 < right {
                mid := left + (right - left) / 2
                if nums[mid] > nums[len(nums) - 1] {
                        left = mid
                } else {
                        right = mid
                }
        }
        return nums[right]
}
