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
