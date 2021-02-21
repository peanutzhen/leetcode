package main

// 找规律题 找出下一个字符的间隔即可 小心别越界
// O(n)时间复杂度
func convert(s string, numRows int) string {
	// 简单情形
	if numRows == 1 {
		return s
	}

	length := len(s)
	offset := 2*numRows - 2       // 间隔公式1
	bytes := make([]byte, length) // O(n)空间复杂度

	var k = 0 // bytes数组下标
	for i := 0; i < length && i < numRows; i++ {
		for j := i; j < length; j = j + offset {
			bytes[k] = s[j]
			k++
			// 非首和末尾行 都要补充一个字符
			// P       I    N
			// A   (L) S  I G
			// Y (A)   H R
			// P       I
			// 例如上面的L与A
			if i != 0 && i != numRows-1 && j+2*(numRows-1-i) < length {
				bytes[k] = s[j+2*(numRows-1-i)] // 间隔公式2
				k++
			}
		}
	}
	return string(bytes)
}
