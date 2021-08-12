package main

import (
	"fmt"
	"sort"
)

type info struct {
	name  byte
	score int
}

func rank(infoList []*info) {
	sort.Slice(infoList, func(i, j int) bool {
		return infoList[i].score > infoList[j].score
	})
}

func input() []*info {
	hashmap := map[byte]int{}
	var i, j byte
	var m, n int
	for {
		_, err := fmt.Scanf("%c-%c %d:%d\n", &i, &j, &m, &n)
		if err != nil {
			break
		}
		// 解析输入字符串
		if m > n {
			hashmap[i] += 3
			hashmap[j] += 0
		} else if m == n {
			hashmap[i] += 1
			hashmap[j] += 1
		} else {
			hashmap[j] += 3
			hashmap[j] += 0
		}
	}
	// 构造info slice
	rtv := make([]*info, len(hashmap))
	idx := 0

	keys := []byte{}
	for k := range hashmap {
		keys = append(keys, k)
	}
	// 字典序排列
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for _, k := range keys {
		rtv[idx] = &info{k, hashmap[k]}
		idx++
	}
	return rtv
}

func main() {
	infoList := input()
	rank(infoList)
	// output
	for i := 0; i < len(infoList)-1; i++ {
		fmt.Printf("%c %d,", infoList[i].name, infoList[i].score)
	}
	fmt.Printf("%c %d\n", infoList[len(infoList)-1].name, infoList[len(infoList)-1].score)
}
