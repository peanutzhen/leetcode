package main

import "fmt"

func simplifyPath(path string) string {
	queue := []byte{}

	rtv := []byte{'/'} // 初始化

	p := []byte(path)
	p = append(p, '/') // 保证path结尾必是 /
	n := len(p)
	for i := 0; i < n; i++ {
		if p[i] != '/' { // 保存字符
			queue = append(queue, p[i])
			continue
		}
		dirName := string(queue) // 一次 / 的闭合
		if dirName == ".." {
			m := len(rtv)
			i := m - 2
			for i >= 0 && rtv[i] != '/' {
				i--
			}
			rtv = rtv[:i+1]    // 弹出一个目录名
			if len(rtv) == 0 { // 保证至少为根目录
				rtv = append(rtv, '/')
			}
		} else if dirName != "." && dirName != "" { // dirName是正常目录名时
			rtv = append(rtv, []byte(dirName)...)
			rtv = append(rtv, '/') // 保证结尾是 /
		}
		queue = queue[:0]
	}
	if len(rtv) == 1 {
		return "/"
	}
	return string(rtv[:len(rtv)-1]) // 忽略结尾 /
}

func main() {
	fmt.Println(simplifyPath("/a//b////c/d//././/.."))
}
