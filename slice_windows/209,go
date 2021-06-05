// 基本滑动窗口思路
// 遍历数组，每次往窗口加入当前元素，如果窗口总和大于等于target，则窗口左边界右移，直到窗口总和小于target为止
func minSubArrayLen(target int, nums []int) int {
    res := math.MaxInt32
    size := len(nums)
    if size == 0 {
        return res
    }
    left, sum := 0, 0
    for i := 0; i < size; i++ {
        sum += nums[i]
        for sum >= target {
            res = min(res, i - left + 1)
            sum -= nums[left]
            left++
        }
    }
    if res == math.MaxInt32 {
        res = 0
    }
    return res
}

func min(x, y int) int {
    if (x < y) {
        return x
    }
    return y
}