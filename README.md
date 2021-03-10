# Leetcode刷题记录

记录一下自己的刷题思路，题目来源很杂，不只leetcode，还有别的地方的，但是大部分是leetcode。

**题目思路我直接写在代码注释里，方便阅读理解。**

咨询问题可以联系邮箱或者知乎私信，每天都会看。

在下的邮箱：astzls213@163.com

知乎ID：钟海宁

## 目录

|      | 题目                                                         | 类型        | 代码                                                         | 难度 | 备注                               |
| ---- | ------------------------------------------------------------ | ----------- | ------------------------------------------------------------ | ---- | ---------------------------------- |
| 1    | [4. 寻找两个正序数组的中位数](https://leetcode-cn.com/problems/median-of-two-sorted-arrays/) | 二分查找    | [C](solutions/001_Median_Two_Sorted_Arrays.c)                | **   |                                    |
| 2    | [546. 移除盒子](https://leetcode-cn.com/problems/remove-boxes/) | 动态规划    | [C](solutions/002_Drop_Box.c)                                | ***  |                                    |
| 3    | [面试题 02.01 移除重复节点](https://leetcode-cn.com/problems/remove-duplicate-node-lcci/) | 益智        | [C++](solutions/003_Duplicated_Node.cpp)                     | *    |                                    |
| 4    | [ACM 洪水]()                                                 | 益智        | [C](solutions/004_Flooded.c)                                 | ***  | 很好玩的一题，抽象思维             |
| 5    | [26. 删除排序数组中的重复项](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/) | 数组        | [Python](solutions/005_Remove_Duplicates_Sorted_Array.py)    | *    |                                    |
| 6    | [29. 两数相除](https://leetcode-cn.com/problems/divide-two-integers/) | 益智        | [C](solutions/006_Divide_Two_Integers.c);[Go](solutions/006_divide_two_integers.go) | **   | 记住logN解法                       |
| 7    | [31. 下一个排列](https://leetcode-cn.com/problems/next-permutation/) | 数组        | [Python](solutions/007_Next_Permutation.py)                  | ***  | 用右边交换左边的数，使得增量足够小 |
| 8    | [33. 搜索旋转排序数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/) | 二分查找    | [C](solutions/008_Search_in_Rotated_Sorted_Array.c)          | *    |                                    |
| 9    | [2. 两数相加](https://leetcode-cn.com/problems/add-two-numbers/) | 链表        | [C](solutions/009_Add_Two_Sum.c)                             | *    |                                    |
| 10   | [36. 有效的数独](https://leetcode-cn.com/problems/valid-sudoku/) | 哈希表      | [Python](solutions/010_Valid_Sudoku.py)                      | *    |                                    |
| 11   | [101. 对称二叉树](https://leetcode-cn.com/problems/symmetric-tree/) | 树          | [C](solutions/011_Symmetric_Tree.c)                          | *    |                                    |
| 12   | [50. Pow(x, n)](https://leetcode-cn.com/problems/powx-n/)    | 二分查找    | [C](solutions/012_Pow.c)                                     | *    |                                    |
| 13   | [62. 不同路径](https://leetcode-cn.com/problems/unique-paths/) | 动态规划    | [C](solutions/013_Unique_Path.c)                             | *    |                                    |
| 14   | [74. 搜索二维矩阵](https://leetcode-cn.com/problems/search-a-2d-matrix/) | 双指针      | [Python](solutions/014_search_2d_matrix.py)                  | *    |                                    |
| 15   | [1. 两数之和](https://leetcode-cn.com/problems/two-sum/)     | 数组        | [Go](solutions/015_two_sum.go)                               | *    |                                    |
| 16   | [3. 无重复字符的最长子串](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/) | 字符串      | [Go](solutions/016_longest_substring_without_repeating_char.go) | **   | 很好玩的一题                       |
| 17   | [5. 最长回文子串](https://leetcode-cn.com/problems/longest-palindromic-substring/) | 动态规划    | [Go](solutions/017_longest_palindromic_substring.go)         | **   | 这题的O(n)解法有点难               |
| 18   | [6. Z 字形变换](https://leetcode-cn.com/problems/zigzag-conversion/) | 益智        | [Go](solutions/018_zigzag_conversion.go)                     | *    | 找规律题                           |
| 19   | [7. 整数反转](https://leetcode-cn.com/problems/reverse-integer/) | 数学        | [Go](solutions/019_reverse_integer.go)                       | *    |                                    |
| 20   | [9. 回文数](https://leetcode-cn.com/problems/palindrome-number/) | 数学        | [Go](solutions/020_palindrome_number.go)                     | *    |                                    |
| 21   | [10. 正则表达式匹配](https://leetcode-cn.com/problems/regular-expression-matching/) | 动态规划    | [Go](solutions/021_regular_expression_matching.go)           | ***  | 状态方程有点难想到                 |
| 22   | [11. 盛最多水的容器](https://leetcode-cn.com/problems/container-with-most-water/) | 双指针      | [Go](solutions/022_container_with_most_water.go)             | *    | 双指针经典题                       |
| 23   | [12. 整数转罗马数字](https://leetcode-cn.com/problems/integer-to-roman/) | 数学        | [Go](solutions/023_integer_to_roman.go)                      | *    | 很无聊的题..                       |
| 24   | [14. 最长公共前缀](https://leetcode-cn.com/problems/longest-common-prefix/) | 字符串      | [Go](solutions/024_longest_common_prefix.go)                 | *    | 暴力就完事了                       |
| 25   | [15. 三数之和](https://leetcode-cn.com/problems/3sum/)       | 双指针      | [Go](solutions/025_three_num_sum.go)                         | **   | 如何避免重复答案很重要             |
| 26   | [16. 最接近的三数之和](https://leetcode-cn.com/problems/3sum-closest/) | 双指针      | [Go](solutions/026_three_sum_closet.go)                      | *    | 比三数之和简单                     |
| 27   | [17. 电话号码的字母组合](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/) | DFS         | [Go](solutions/027_combinations_phone.go)                    | *    | 好久没写DFS，有点生疏              |
| 28   | [18. 四数之和](https://leetcode-cn.com/problems/4sum/)       | 双指针      | [Go](solutions/028_four_num_sum.go)                          | *    | 三数之和变形题                     |
| 29   | [19. 删除链表的倒数第 N 个结点](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/) | 链表        | [Go](solutions/029_rm_node.go)                               | *    |                                    |
| 30   | [20. 有效的括号](https://leetcode-cn.com/problems/valid-parentheses/) | 栈          | [Go](solutions/030_valid_parentheses.go)                     | *    |                                    |
| 31   | [21. 合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists/) | 链表        | [Go](solutions/031_merge_sorted_list.go)                     | *    |                                    |
| 32   | [22. 括号生成](https://leetcode-cn.com/problems/generate-parentheses/) | 回朔        | [Go](solutions/032_gen_parentheses.go)                       | **   | 回朔算法很适合这题                 |
| 33   | [23. 合并K个升序链表](https://leetcode-cn.com/problems/merge-k-sorted-lists/) | 堆          | [Go](solutions/033_merge_k_sorted_list.go)                   | *    | 构造数组+排序+构造链表             |
| 34   | [24. 两两交换链表中的节点](https://leetcode-cn.com/problems/swap-nodes-in-pairs/) | 递归        | [Go](solutions/034_swap_adjacent_node.go)                    | *    | 往递归想就简单了                   |
| 35   | [25. K 个一组翻转链表](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/) | 链表        | [Go](solutions/035_reverse_k_nodes.go)                       | ***  | 此题难在设计                       |
| 36   | [27. 移除元素](https://leetcode-cn.com/problems/remove-element/) | 数组        | [Go](solutions/036_remove_elements.go)                       | *    | 假移除真交换                       |
| 37   | [28. 实现 strStr()](https://leetcode-cn.com/problems/implement-strstr/) | 字符串      | [Go](solutions/037_strStr.go)                                | *    | 哈希法判断字符串相等有点意思       |
| 38   | [30. 串联所有单词的子串](https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/) | 字符串      | [Go](solutions/038_find_words_in_string.go)                  | **   | 词频+滑动窗口                      |
| 39   | [32. 最长有效括号](https://leetcode-cn.com/problems/longest-valid-parentheses/) | 动态规划/栈 | [Go](solutions/039_longest_valid_parentheses.go)             | ***  | 看了题解才会，草                   |
| 40   | [34. 在排序数组中查找元素的第一个和最后一个位置](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/) | 二分查找    | [Go](solutions/040_first_last_index_target.go)               | *    |                                    |
| 41   | [35. 搜索插入位置](https://leetcode-cn.com/problems/search-insert-position/) | 二分查找    | [Go](solutions/041_search_insert_index.go)                   | *    | 二分查找变形题                     |
| 42   | [37. 解数独](https://leetcode-cn.com/problems/sudoku-solver/) | 回朔算法    | [Go](solutions/042_solve_sudoku.go)                          | **   | 代码写的不够快还是                 |
| 43   | [38. 外观数列](https://leetcode-cn.com/problems/count-and-say/) | 递归        | [Go](solutions/043_count_and_say.go)                         | *    | 无聊                               |
| 44   | [39. 组合总和](https://leetcode-cn.com/problems/combination-sum/) | 回朔算法    | [Go](solutions/044_combination_sum.go)                       | *    |                                    |
| 45   | [40. 组合总和 II](https://leetcode-cn.com/problems/combination-sum-ii/) | 回朔算法    | [Go](solutions/045_combination_sum2.go)                      | *    | 044的变形题                        |
| 46   | [41. 缺失的第一个正数](https://leetcode-cn.com/problems/first-missing-positive/) | 原地哈希    | [Go](solutions/046_first_missing_positive.go)                | ***  | O(n)算法有点难想 操                |
| 47   | [Shopee 笔试题1.有效的括号字符串](https://www.nowcoder.com/questionTerminal/2a9453b8c4a74b11a360edce506df26d) | 回朔算法    | [Go](solutions/047_check_valid_string.go)                    | **   |                                    |
| 48   | [42. 接雨水](https://leetcode-cn.com/problems/trapping-rain-water/) | 双指针      | [Go](solutions/048_trap_water.go)                            | ***  | 又是看题解的一天                   |
| 49   | [43. 字符串相乘](https://leetcode-cn.com/problems/multiply-strings/) | 数学        | [Go](solutions/049_string_multiply.go)                       | ***  | NMSL                               |
| 50   | [415. 字符串相加](https://leetcode-cn.com/problems/add-strings/) | 字符串      | [Go](solutions/050_string_add.go)                            | *    |                                    |
| 51   | [44. 通配符匹配](https://leetcode-cn.com/problems/wildcard-matching/) | 动态规划    | [Go](solutions/051_wildcard_matching.go)                     | **   | NMSL                               |
| 52   | [45. 跳跃游戏 II](https://leetcode-cn.com/problems/jump-game-ii/) | 贪心算法    | [Go](solutions/052_jump_game.go)                             | *    | So easy                            |
| 53   | [46. 全排列](https://leetcode-cn.com/problems/permutations/) | 回朔算法    | [053_permute.go](solutions/053_permute.go)                   | **   | 减少空间复杂度值得思考             |
