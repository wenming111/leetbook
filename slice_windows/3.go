// 基本滑动窗口题型
// 遍历数组，每次添加最右的元素，并判断当前窗口内是否包含此元素
// 如果包含 则一直右移窗口的左边界（一个while实现）

// go中set的实现，可以把map的value设置成void类型，不占用内存。void类型可以用struct{}来定义
type void struct{}
func lengthOfLongestSubstring(s string) int {
    look_up := make(map[byte]void)
	var member void
	size := len(s)
	left, right := 0, 0
	res := 0
	if size == 0 {
		return 0
	}
	for i := 0; i < size; ++i {
		for {
			_, ok := look_up[s[i]]
			if ok == true {
				delete(look_up, s[left])
				left++
			} else {
				break
			}
		}
		if res < (i - left + 1) {
			res = (i - left + 1)
		}
		look_up[s[i]] = member
	}
	return res
}
