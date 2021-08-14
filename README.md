# Leetcode刷题记录

记录一下自己的刷题思路，题目来源很杂，不只leetcode，还有别的地方(Codeforces/UVA)的，但是大部分是leetcode。

### Tips

- 序号用<u>**下划线加粗**</u>的题目是用来复习某个**特定算法**的极佳例子。（如53:回朔算法/109:LRU算法/149:排序算法）
- 题目**思路**在**代码注释**里，方便阅读理解。

### Contact

咨询问题可以联系邮箱或者知乎私信，每天都会看。

- 邮箱：astzls213@163.com

- 知乎ID：钟海宁


## 目录

|                | 题目                                                         | 类型          | 代码                                                         | 难度 | 备注                                                 |
| -------------- | ------------------------------------------------------------ | ------------- | ------------------------------------------------------------ | ---- | ---------------------------------------------------- |
| 1              | [4. 寻找两个正序数组的中位数](https://leetcode-cn.com/problems/median-of-two-sorted-arrays/) | 二分查找      | [C](solutions/__001_Median_Two_Sorted_Arrays.c)              | **   |                                                      |
| 2              | [546. 移除盒子](https://leetcode-cn.com/problems/remove-boxes/) | 动态规划      | [C](solutions/__002_Drop_Box.c)                              | ***  |                                                      |
| 3              | [面试题 02.01 移除重复节点](https://leetcode-cn.com/problems/remove-duplicate-node-lcci/) | 益智          | [C++](solutions/__003_Duplicated_Node.cpp)                   | *    |                                                      |
| 4              | [ACM 洪水]()                                                 | 益智          | [C](solutions/__004_Flooded.c)                               | ***  | 很好玩的一题，抽象思维                               |
| 5              | [26. 删除排序数组中的重复项](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/) | 数组          | [Python](solutions/__005_Remove_Duplicates_Sorted_Array.py)  | *    |                                                      |
| 6              | [29. 两数相除](https://leetcode-cn.com/problems/divide-two-integers/) | 益智          | [C](solutions/__006_Divide_Two_Integers.c);[Go](solutions/__006_divide_two_integers.go) | **   | 记住logN解法                                         |
| 7              | [31. 下一个排列](https://leetcode-cn.com/problems/next-permutation/) | 数组          | [Python](solutions/__007_Next_Permutation.py)                | ***  | 用右边交换左边的数，使得增量足够小                   |
| 8              | [33. 搜索旋转排序数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/) | 二分查找      | [C](solutions/__008_Search_in_Rotated_Sorted_Array.c)        | *    |                                                      |
| 9              | [2. 两数相加](https://leetcode-cn.com/problems/add-two-numbers/) | 链表          | [C](solutions/__009_Add_Two_Sum.c)                           | *    |                                                      |
| 10             | [36. 有效的数独](https://leetcode-cn.com/problems/valid-sudoku/) | 哈希表        | [Python](solutions/__010_Valid_Sudoku.py)                    | *    |                                                      |
| 11             | [101. 对称二叉树](https://leetcode-cn.com/problems/symmetric-tree/) | 树            | [C](solutions/__011_Symmetric_Tree.c)                        | *    |                                                      |
| 12             | [50. Pow(x, n)](https://leetcode-cn.com/problems/powx-n/)    | 二分查找      | [C](solutions/__012_Pow.c)                                   | *    |                                                      |
| 13             | [62. 不同路径](https://leetcode-cn.com/problems/unique-paths/) | 动态规划      | [C](solutions/__013_Unique_Path.c)                           | *    |                                                      |
| 14             | [74. 搜索二维矩阵](https://leetcode-cn.com/problems/search-a-2d-matrix/) | 双指针        | [Python](solutions/__014_search_2d_matrix.py)                | *    |                                                      |
| 15             | [1. 两数之和](https://leetcode-cn.com/problems/two-sum/)     | 数组          | [Go](solutions/__015_two_sum.go)                             | *    |                                                      |
| 16             | [3. 无重复字符的最长子串](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/) | 字符串        | [Go](solutions/__016_longest_substring_without_repeating_char.go) | **   | 很好玩的一题                                         |
| 17             | [5. 最长回文子串](https://leetcode-cn.com/problems/longest-palindromic-substring/) | 动态规划      | [Go](solutions/__017_longest_palindromic_substring.go)       | **   | 这题的O(n)解法有点难                                 |
| 18             | [6. Z 字形变换](https://leetcode-cn.com/problems/zigzag-conversion/) | 益智          | [Go](solutions/__018_zigzag_conversion.go)                   | *    | 找规律题                                             |
| 19             | [7. 整数反转](https://leetcode-cn.com/problems/reverse-integer/) | 数学          | [Go](solutions/__019_reverse_integer.go)                     | *    |                                                      |
| 20             | [9. 回文数](https://leetcode-cn.com/problems/palindrome-number/) | 数学          | [Go](solutions/__020_palindrome_number.go)                   | *    |                                                      |
| 21             | [10. 正则表达式匹配](https://leetcode-cn.com/problems/regular-expression-matching/) | 动态规划      | [Go](solutions/__021_regular_expression_matching.go)         | ***  | 状态方程有点难想到                                   |
| 22             | [11. 盛最多水的容器](https://leetcode-cn.com/problems/container-with-most-water/) | 双指针        | [Go](solutions/__022_container_with_most_water.go)           | *    | 双指针经典题                                         |
| 23             | [12. 整数转罗马数字](https://leetcode-cn.com/problems/integer-to-roman/) | 数学          | [Go](solutions/__023_integer_to_roman.go)                    | *    | 很无聊的题..                                         |
| 24             | [14. 最长公共前缀](https://leetcode-cn.com/problems/longest-common-prefix/) | 字符串        | [Go](solutions/__024_longest_common_prefix.go)               | *    | 暴力就完事了                                         |
| 25             | [15. 三数之和](https://leetcode-cn.com/problems/3sum/)       | 双指针        | [Go](solutions/__025_three_num_sum.go)                       | **   | 如何避免重复答案很重要                               |
| 26             | [16. 最接近的三数之和](https://leetcode-cn.com/problems/3sum-closest/) | 双指针        | [Go](solutions/__026_three_sum_closet.go)                    | *    | 比三数之和简单                                       |
| 27             | [17. 电话号码的字母组合](https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number/) | DFS           | [Go](solutions/__027_combinations_phone.go)                  | *    | 好久没写DFS，有点生疏                                |
| 28             | [18. 四数之和](https://leetcode-cn.com/problems/4sum/)       | 双指针        | [Go](solutions/__028_four_num_sum.go)                        | *    | 三数之和变形题                                       |
| 29             | [19. 删除链表的倒数第 N 个结点](https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/) | 链表          | [Go](solutions/__029_rm_node.go)                             | *    |                                                      |
| 30             | [20. 有效的括号](https://leetcode-cn.com/problems/valid-parentheses/) | 栈            | [Go](solutions/__030_valid_parentheses.go)                   | *    |                                                      |
| 31             | [21. 合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists/) | 链表          | [Go](solutions/__031_merge_sorted_list.go)                   | *    | O(1)空间复杂度要学会                                 |
| 32             | [22. 括号生成](https://leetcode-cn.com/problems/generate-parentheses/) | 回朔          | [Go](solutions/__032_gen_parentheses.go)                     | **   | 回朔算法很适合这题                                   |
| 33             | [23. 合并K个升序链表](https://leetcode-cn.com/problems/merge-k-sorted-lists/) | 堆            | [Go](solutions/__033_merge_k_sorted_list.go)                 | *    | 构造数组+排序+构造链表                               |
| 34             | [24. 两两交换链表中的节点](https://leetcode-cn.com/problems/swap-nodes-in-pairs/) | 递归          | [Go](solutions/__034_swap_adjacent_node.go)                  | *    | 往递归想就简单了                                     |
| 35             | [25. K 个一组翻转链表](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/) | 链表          | [Go](solutions/__035_reverse_k_nodes.go)                     | ***  | 此题难在设计                                         |
| 36             | [27. 移除元素](https://leetcode-cn.com/problems/remove-element/) | 数组          | [Go](solutions/__036_remove_elements.go)                     | *    | 假移除真交换                                         |
| 37             | [28. 实现 strStr()](https://leetcode-cn.com/problems/implement-strstr/) | 字符串        | [Go](solutions/__037_strStr.go)                              | *    | 哈希法判断字符串相等有点意思                         |
| 38             | [30. 串联所有单词的子串](https://leetcode-cn.com/problems/substring-with-concatenation-of-all-words/) | 字符串        | [Go](solutions/__038_find_words_in_string.go)                | **   | 词频+滑动窗口                                        |
| 39             | [32. 最长有效括号](https://leetcode-cn.com/problems/longest-valid-parentheses/) | 动态规划/栈   | [Go](solutions/__039_longest_valid_parentheses.go)           | ***  | 看了题解才会，草                                     |
| 40             | [34. 在排序数组中查找元素的第一个和最后一个位置](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array/) | 二分查找      | [Go](solutions/__040_first_last_index_target.go)             | *    |                                                      |
| 41             | [35. 搜索插入位置](https://leetcode-cn.com/problems/search-insert-position/) | 二分查找      | [Go](solutions/__041_search_insert_index.go)                 | *    | 二分查找变形题                                       |
| 42             | [37. 解数独](https://leetcode-cn.com/problems/sudoku-solver/) | 回朔算法      | [Go](solutions/__042_solve_sudoku.go)                        | **   | 代码写的不够快还是                                   |
| 43             | [38. 外观数列](https://leetcode-cn.com/problems/count-and-say/) | 递归          | [Go](solutions/__043_count_and_say.go)                       | *    | 无聊                                                 |
| 44             | [39. 组合总和](https://leetcode-cn.com/problems/combination-sum/) | 回朔算法      | [Go](solutions/__044_combination_sum.go)                     | *    |                                                      |
| 45             | [40. 组合总和 II](https://leetcode-cn.com/problems/combination-sum-ii/) | 回朔算法      | [Go](solutions/__045_combination_sum2.go)                    | *    | 044的变形题                                          |
| 46             | [41. 缺失的第一个正数](https://leetcode-cn.com/problems/first-missing-positive/) | 原地哈希      | [Go](solutions/__046_first_missing_positive.go)              | ***  | O(n)算法有点难想 操                                  |
| 47             | [Shopee 笔试题1.有效的括号字符串](https://www.nowcoder.com/questionTerminal/2a9453b8c4a74b11a360edce506df26d) | 回朔算法      | [Go](solutions/__047_check_valid_string.go)                  | **   |                                                      |
| 48             | [42. 接雨水](https://leetcode-cn.com/problems/trapping-rain-water/) | 双指针        | [Go](solutions/__048_trap_water.go)                          | ***  | 又是看题解的一天                                     |
| 49             | [43. 字符串相乘](https://leetcode-cn.com/problems/multiply-strings/) | 数学          | [Go](solutions/__049_string_multiply.go)                     | ***  | NMSL                                                 |
| 50             | [415. 字符串相加](https://leetcode-cn.com/problems/add-strings/) | 字符串        | [Go](solutions/__050_string_add.go)                          | *    |                                                      |
| 51             | [44. 通配符匹配](https://leetcode-cn.com/problems/wildcard-matching/) | 动态规划      | [Go](solutions/__051_wildcard_matching.go)                   | **   | NMSL                                                 |
| 52             | [45. 跳跃游戏 II](https://leetcode-cn.com/problems/jump-game-ii/) | 贪心算法      | [Go](solutions/__052_jump_game.go)                           | *    | So easy                                              |
| <u>**53**</u>  | [46. 全排列](https://leetcode-cn.com/problems/permutations/) | 回朔算法      | [Go](solutions/__053_permute.go)                             | **   | 减少空间复杂度值得思考                               |
| 54             | [47. 全排列 II](https://leetcode-cn.com/problems/permutations-ii/) | 回朔算法      | [Go](solutions/__054_permute_unique.go)                      | *    | 避免重复答案                                         |
| 55             | [48. 旋转图像](https://leetcode-cn.com/problems/rotate-image/) | 益智          | [Go](solutions/__055_rotate_photo.go)                        | ***  | 没想到O(1)空间复杂度解法                             |
| 56             | [49. 字母异位词分组](https://leetcode-cn.com/problems/group-anagrams/) | 字符串        | [Go](solutions/__056_group_anagrams.go)                      | **   | 词频+哈希表                                          |
| 57             | [51. N 皇后](https://leetcode-cn.com/problems/n-queens/)     | 回朔算法      | [Go](solutions/__057_n_queens.go)                            | *    |                                                      |
| 58             | [52. N皇后 II](https://leetcode-cn.com/problems/n-queens-ii/) | 回朔算法      | [Go](solutions/__058_n_queens_2.go)                          | *    |                                                      |
| 59             | [53. 最大子序和](https://leetcode-cn.com/problems/maximum-subarray/) | 动态规划/分治 | [Go](solutions/__059_max_sub_array.go) /[Java](solutions/__Q059_max_sup_array.java) | *    | 注意分治法！                                         |
| 60             | [54. 螺旋矩阵](https://leetcode-cn.com/problems/spiral-matrix/) | 代码设计      | [Go](solutions/__060_spiral_order.go)                        | **   | 妙在标记已读数字解决边界问题                         |
| 61             | [55. 跳跃游戏](https://leetcode-cn.com/problems/jump-game/)  | 贪心算法      | [Go](solutions/__061_can_jump.go)                            | *    |                                                      |
| 62             | [56. 合并区间](https://leetcode-cn.com/problems/merge-intervals/) | 益智          | [Go](solutions/__062_merge_intervals.go)                     | *    |                                                      |
| 63             | [78. 子集](https://leetcode-cn.com/problems/subsets/)        | 回朔算法      | [Go](solutions/__063_subsets.go)                             | *    |                                                      |
| 64             | [57. 插入区间](https://leetcode-cn.com/problems/insert-interval/) | 益智          | [Go](solutions/__064_insert_intervals.go)                    | *    |                                                      |
| 65             | [58. 最后一个单词的长度](https://leetcode-cn.com/problems/length-of-last-word/) | 字符串        | [Go](solutions/__065_length_last_word.go)                    | *    | XSWL                                                 |
| 66             | [59. 螺旋矩阵 II](https://leetcode-cn.com/problems/spiral-matrix-ii/) | 代码设计      | [Go](solutions/__066_gen_matrix.go)                          | *    | 格局小了                                             |
| 67             | [60. 排列序列](https://leetcode-cn.com/problems/permutation-sequence/) | 数学          | [Go](solutions/__067_get_k_permute.go)                       | **   | 找出规律即可                                         |
| 68             | [61. 旋转链表](https://leetcode-cn.com/problems/rotate-list/) | 链表          | [Go](solutions/__068_right_shift_list.go)                    | *    |                                                      |
| 69             | [63. 不同路径 II](https://leetcode-cn.com/problems/unique-paths-ii/) | 动态规划      | [Go](solutions/__069_unique_path_obstacle.go)                | *    | 注意滚动数组思想                                     |
| 70             | [64. 最小路径和](https://leetcode-cn.com/problems/minimum-path-sum/) | 动态规划      | [Go](solutions/__070_min_path_sum.go)                        | *    | 滚动数组                                             |
| 71             | [65. 有效数字](https://leetcode-cn.com/problems/valid-number/) | 有穷自动机    | [Go](solutions/__071_is_number.go)                           | ***  | 编译原理的有穷自动机模型                             |
| 72             | [66. 加一](https://leetcode-cn.com/problems/plus-one/)       | 数组          | [Go](solutions/__072_plus_one.go)                            | *    |                                                      |
| 73             | [67. 二进制求和](https://leetcode-cn.com/problems/add-binary/) | 字符串        | [Go](solutions/__073_binary_add.go)                          | *    |                                                      |
| 74             | [69. x 的平方根](https://leetcode-cn.com/problems/sqrtx/)    | 二分查找      | [Go](solutions/__073_binary_add.go)                          | *    |                                                      |
| 75             | [70. 爬楼梯](https://leetcode-cn.com/problems/climbing-stairs/) | 动态规划      | [Go](solutions/__075_climb_stairs.go)                        | *    |                                                      |
| 76             | [71. 简化路径](https://leetcode-cn.com/problems/simplify-path/) | 字符串        | [Go](solutions/__076_simplify_path.go)                       | **   | 分割+状态机                                          |
| 77             | [72. 编辑距离](https://leetcode-cn.com/problems/edit-distance/) | 动态规划      | [Go](solutions/__077_min_edit_distance.go)                   | ***  | 又是看题解的一天                                     |
| 78             | [73. 矩阵置零](https://leetcode-cn.com/problems/set-matrix-zeroes/) | 益智          | [Go](solutions/__078_set_matrix_zeros.go)                    | *    | 注意O(1)空间复杂度的解法                             |
| 79             | [75. 颜色分类](https://leetcode-cn.com/problems/sort-colors/) | 数组          | [Go](solutions/__079_sort_color.go)                          | *    |                                                      |
| 80             | [76. 最小覆盖子串](https://leetcode-cn.com/problems/minimum-window-substring/) | 双指针        | [Go](solutions/__080_min_window_substr.go)                   | ***  | 词频+滑动窗口                                        |
| 81             | [77. 组合](https://leetcode-cn.com/problems/combinations/)   | 回朔算法      | [Go](solutions/__081_combine.go)                             | *    |                                                      |
| 82             | [79. 单词搜索](https://leetcode-cn.com/problems/word-search/) | 回朔算法      | [Go](solutions/__082_exist_word.go) /[Java](solutions/__Q082_search_word.java) | **   |                                                      |
| 83             | [80. 删除排序数组中的重复项 II](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii/) | 双指针        | [Go](solutions/__083_remove_duplicated_item2.go)             | ***  | O(n) O(1)难想到                                      |
| 84             | [81. 搜索旋转排序数组 II](https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii/) | DFS           | [Go](solutions/__084_search_rotated_array2.go)               | *    | 这个有重复元素                                       |
| 85             | [82. 删除排序链表中的重复元素 II](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/) | 双指针        | [Go](solutions/__085_remove_duplicate_listnode.go)           | *    |                                                      |
| 86             | [83. 删除排序链表中的重复元素](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/) | 双指针        | [Go](solutions/__086_remove_duplicate_listnode2.go)          | *    |                                                      |
| 87             | [84. 柱状图中最大的矩形](https://leetcode-cn.com/problems/largest-rectangle-in-histogram/) | 单调栈        | [Go](solutions/__087_largest_rectangle_in_histogram.go)      | ***  | 学到了学到了                                         |
| 88             | [85. 最大矩形](https://leetcode-cn.com/problems/maximal-rectangle/) | 单调栈        | [Go](solutions/__088_maxsize_one_matrix.go)                  | ***  | 上一题的变形题                                       |
| 89             | [86. 分隔链表](https://leetcode-cn.com/problems/partition-list/) | 双指针        | [Go](solutions/__089_partition_list.go)                      | *    |                                                      |
| 90             | [88. 合并两个有序数组](https://leetcode-cn.com/problems/merge-sorted-array/) | 双指针        | [Go](solutions/__090_merge_sort_array.go)                    | *    |                                                      |
| 91             | [89. 格雷编码](https://leetcode-cn.com/problems/gray-code/)  | 数学          | [Go](solutions/__091_gray_code.go)                           | *    | 找规律题                                             |
| 92             | [90. 子集 II](https://leetcode-cn.com/problems/subsets-ii/)  | 回朔算法      | [Go](solutions/__092_subsets2.go)                            | *    | 三数之和避免重复解                                   |
| 93             | [91. 解码方法](https://leetcode-cn.com/problems/decode-ways/) | 动态规划      | [Go](solutions/__093_num_decoding.go)                        | *    |                                                      |
| 94             | [92. 反转链表 II](https://leetcode-cn.com/problems/reverse-linked-list-ii/) | 链表          | [Go](solutions/__094_reverse_list_m_to_n.go)                 | *    | 难在设计                                             |
| 95             | [94. 二叉树的中序遍历](https://leetcode-cn.com/problems/binary-tree-inorder-traversal/) | 二叉树        | [Go](solutions/__095_inorder_traversal.go)                   | *    |                                                      |
| 96             | [95. 不同的二叉搜索树 II](https://leetcode-cn.com/problems/unique-binary-search-trees-ii/) | 递归          | [Go](solutions/__096_different_binary_trees.go)              | ***  | 又是看题解的一天                                     |
| 97             | [96. 不同的二叉搜索树](https://leetcode-cn.com/problems/unique-binary-search-trees/) | 动态规划      | [Go](solutions/__097_num_trees.go)                           | *    | 卡特兰数列有点东西                                   |
| 98             | [98. 验证二叉搜索树](https://leetcode-cn.com/problems/validate-binary-search-tree/) | 二叉树        | [Go](solutions/__098_valid_BST.go)                           | *    | BST中序遍历的性质                                    |
| 99             | [97. 交错字符串](https://leetcode-cn.com/problems/interleaving-string/) | 动态规划      | [Go](solutions/__099_interleaving_str.go)                    | *    |                                                      |
| 100            | [99. 恢复二叉搜索树](https://leetcode-cn.com/problems/recover-binary-search-tree/) | 二叉树        | [Go](solutions/__100_recover_BST.go)                         | *    | 利用BST中序遍历的性质                                |
| 101            | [100. 相同的树](https://leetcode-cn.com/problems/same-tree/) | 二叉树        | [Go](solutions/__101_same_tree.go)                           | *    | 还是中序遍历或递归                                   |
| 102            | [102. 二叉树的层序遍历](https://leetcode-cn.com/problems/binary-tree-level-order-traversal/) | BFS           | [Go](solutions/__102_level_traversal.go)                     | *    |                                                      |
| 103            | [103. 二叉树的锯齿形层序遍历](https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/) | BFS           | [Go](solutions/__103_zigzag_order.go)                        | *    |                                                      |
| 104            | [104. 二叉树的最大深度](https://leetcode-cn.com/problems/maximum-depth-of-binary-tree/) | 递归          | [Go](solutions/__104_depth_tree.go)                          | *    |                                                      |
| 105            | [105. 从前序与中序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/) | 递归          | [Go](solutions/__105_build_tree.go)                          | *    | 先序遍历中序遍历性质                                 |
| 106            | [106. 从中序与后序遍历序列构造二叉树](https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/) | 递归          | [Go](solutions/__106_build_tree2.go)                         | *    |                                                      |
| 107            | [107. 二叉树的层序遍历 II](https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/) | BFS           | [Go](solutions/__107_level_traversal2.go)                    | *    |                                                      |
| 108            | [108. 将有序数组转换为二叉搜索树](https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/) | 递归          | [Go](solutions/__108_array_to_BST.go)                        | *    | 中序遍历性质                                         |
| <u>**109**</u> | [146. LRU 缓存机制](https://leetcode-cn.com/problems/lru-cache/) | 设计          | [Go](solutions/__109_lru.go)                                 | **   | 哈希表+双向链表                                      |
| 110            | [134. 加油站](https://leetcode-cn.com/problems/gas-station/) | 贪心算法      | [Go](solutions/__110_petrol_station.go)                      | *    |                                                      |
| 111            | [109. 有序链表转换二叉搜索树](https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree/) | 链表          | [Go](solutions/__111_list_to_tree.go)                        | *    | list->array->tree                                    |
| 112            | [110. 平衡二叉树](https://leetcode-cn.com/problems/balanced-binary-tree/) | 递归          | [Go](solutions/__112_is_balance_tree.go)                     | *    |                                                      |
| 113            | [111. 二叉树的最小深度](https://leetcode-cn.com/problems/minimum-depth-of-binary-tree/) | 递归          | [Go](solutions/__113_min_depth.go)                           | *    |                                                      |
| 114            | [112. 路径总和](https://leetcode-cn.com/problems/path-sum/)  | 回朔算法      | [Go](solutions/__114_path_sum.go)                            | *    |                                                      |
| 115            | [113. 路径总和 II](https://leetcode-cn.com/problems/path-sum-ii/) | 回朔算法      | [Go](solutions/__115_path_sum2.go)                           | **   |                                                      |
| 116            | [152. 乘积最大子数组](https://leetcode-cn.com/problems/maximum-product-subarray/) | 动态规划      | [Go](solutions/__116_max_product.go)                         | ***  | 又是看题解的一天                                     |
| 117            | [189. 旋转数组](https://leetcode-cn.com/problems/rotate-array/) | 数组          | [Go](solutions/__117_rotate_array.go)                        | *    | 翻转法有意思                                         |
| 118            | [114. 二叉树展开为链表](https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/) | 二叉树        | [Go](solutions/__118_flatten_tree.go)                        | **   | 思路有却写不出代码草                                 |
| 119            | [121. 买卖股票的最佳时机](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/) | 益智          | [Go](solutions/__119_max_profit.go)                          | *    |                                                      |
| 120            | [124. 二叉树中的最大路径和](https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/) | 递归          | [Go](solutions/__120_max_path_sum.go)                        | **** | 这谁想到得到啊草                                     |
| 121            | [128. 最长连续序列](https://leetcode-cn.com/problems/longest-consecutive-sequence/) | 哈希表        | [Go](solutions/__121_longest_consecutive_sequence.go)        | ***  | 又是哈希表记录数字出现情况                           |
| 122            | [136. 只出现一次的数字](https://leetcode-cn.com/problems/single-number/) | 数学          | [Go](solutions/__122_only_one_element.go)                    | *    | 异或运算的妙用                                       |
| 123            | [141. 环形链表](https://leetcode-cn.com/problems/linked-list-cycle/) | 链表          | [Go](solutions/__123_is_cycle_list.go)                       | *    | 快慢指针很秀                                         |
| 124            | [139. 单词拆分](https://leetcode-cn.com/problems/word-break/) | 动态规划      | [Go](solutions/__124_word_split.go) /[Java](solutions/__124_word_split.java) | ***  | 哈希表+转移方程有点难想                              |
| 125            | [142. 环形链表 II](https://leetcode-cn.com/problems/linked-list-cycle-ii/) | 双指针/数学   | [Go](solutions/__125_is_cycle_list2.go)                      | ***  | O(1)空间复杂度解法难想                               |
| 126            | [148. 排序链表](https://leetcode-cn.com/problems/sort-list/) | 排序          | [Go](solutions/__126_sort_list.go)                           | **** | 迭代版归并排序代码难写                               |
| 127            | [155. 最小栈](https://leetcode-cn.com/problems/min-stack/)   | 栈/设计       | [Go](solutions/__127_min_stack.go)                           | *    | 双栈即可                                             |
| 128            | [198. 打家劫舍](https://leetcode-cn.com/problems/house-robber/) | 动态规划      | [Go](solutions/__128_robber.go)                              | *    |                                                      |
| 129            | [200. 岛屿数量](https://leetcode-cn.com/problems/number-of-islands/) | DFS/BFS       | [Go](solutions/__129_num_islands.go)                         | ***  | 没想到cnm                                            |
| 130            | [239. 滑动窗口最大值](https://leetcode-cn.com/problems/sliding-window-maximum/) | 堆/双端队列   | [Go](solutions/__130_slide_window_maxvalue.go)               | ***  | 队列是我没想到的                                     |
| 131            | [221. 最大正方形](https://leetcode-cn.com/problems/maximal-square/) | 动态规划      | [Go](solutions/__131_max_square.go)                          | ***  | 有时dp保存的不是答案                                 |
| 132            | [739. 每日温度](https://leetcode-cn.com/problems/daily-temperatures/) | 单调栈        | [Go](solutions/__132_daily_temperature.go)                   | *    | 单调栈用于找左边/右边第一个小于/大于当前位置数的位置 |
| 133            | [279. 完全平方数](https://leetcode-cn.com/problems/perfect-squares/) | 动态规划      | [Go](solutions/__133_num_squares.go)                         | *    |                                                      |
| 134            | [122. 买卖股票的最佳时机 II](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/) | 贪心算法      | [Go](solutions/__134_max_profit2.go)                         | *    |                                                      |
| 135            | [668. 乘法表中第k小的数](https://leetcode-cn.com/problems/kth-smallest-number-in-multiplication-table/) | 二分查找      | [Go](solutions/__135_kth_smallest_in_multiplication_table.go) | **** | 二分查找还能这样用@_@                                |
| 136            | [378. 有序矩阵中第 K 小的元素](https://leetcode-cn.com/problems/kth-smallest-element-in-a-sorted-matrix/) | 二分查找      | [Go](solutions/__136_kth_smallest_matrix.go)                 | **   | 135变形题                                            |
| 137            | [125. 验证回文串](https://leetcode-cn.com/problems/valid-palindrome/) | 双指针        | [Go](solutions/__137_palindrome_str.go)                      | *    |                                                      |
| 138            | [343. 整数拆分](https://leetcode-cn.com/problems/integer-break/) | 动态规划      | [Go](solutions/__138_int_break.go)                           | ***  | O(n)的动态规划难证明                                 |
| 139            | [322. 零钱兑换](https://leetcode-cn.com/problems/coin-change/) | 动态规划      | [Go](solutions/__139_coin.go)                                | **   | 贪心可用仅当零钱面值为常数c的幂                      |
| 140            | [199. 二叉树的右视图](https://leetcode-cn.com/problems/binary-tree-right-side-view/) | BFS           | [Go](solutions/__140_tree_right_side_view.go)                | *    |                                                      |
| 141            | [字节笔试C题. 从0至N的最小代价](https://codeforces.com/contest/710/problem/E) | 动态规划      | [Go](solutions/__141_min_time_to_n.go)                       | ***  | NMSL                                                 |
| 142            | 百度笔试A题. 牛牛最小的进食次数                              | 贪心算法      | [Go](solutions/__142_min_eat.go)                             | **   | 发现了Go的一个坑/题目简介in注释                      |
| 143            | [153. 寻找旋转排序数组中的最小值](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array/) | 二分搜索      | [Go](solutions/__143_rotate_array_min.go)                    | *    |                                                      |
| 144            | [470. 用 Rand7() 实现 Rand10()](https://leetcode-cn.com/problems/implement-rand10-using-rand7/) | 随机算法      | [Go](solutions/__144_rand7_to_rand10.go)                     | **   | 又是看题解的一天                                     |
| 145            | [剑指 Offer 43. 1～n 整数中 1 出现的次数](https://leetcode-cn.com/problems/1nzheng-shu-zhong-1chu-xian-de-ci-shu-lcof/) | 数学          | [Go](solutions/__145_count_k_beetween1_and_n.go)             | ***  | NMSL                                                 |
| 146            | [300. 最长递增子序列](https://leetcode-cn.com/problems/longest-increasing-subsequence/) | 动态规划      | [Go](solutions/__146_max_inc_seq.go)                         | ***  | dp[i]指以i为子序列最后元素的最长长度                 |
| 147            | [206. 反转链表](https://leetcode-cn.com/problems/reverse-linked-list/) | 链表          | [Go](solutions/__147_reverse_list.go)                        | *    | 三指针反转，头节点很重要                             |
| 148            | [598. 范围求和 II](https://leetcode-cn.com/problems/range-addition-ii/) | 益智          | [Go](solutions/__148_max_count.go)                           | *    |                                                      |
| <u>**149**</u> | [912. 排序数组](https://leetcode-cn.com/problems/sort-an-array/) | 排序          | [Go](solutions/__149_sort.go)                                | *    | 要求做到面试手撕代码                                 |
| **150**        | [264. 丑数 II](https://leetcode-cn.com/problems/ugly-number-ii/) | 动态规划      | [Go](solutions/__150_ugly_num.go)                            | ***  | 提前批字节商业化三面原题                             |
| 151            | 华为机试A题. 球队积分排行                                    | 排序          | [Go](solutions/__151_rank_team.go)                           | *    |                                                      |
| 152            | 华为机试B题. 最小员工数目                                    | 益智          | [Python](solutions/__152_min_employee.py)                    | **   | 输入是真的蛋疼                                       |
| 153            | 华为机试C题. 最小移动步数                                    | 动态规划      | [Go](solutions/__153_min_step_match.go)                      | ***  | 通过95%cases，最后超时了                             |
| 154            | [154. 寻找旋转排序数组中的最小值 II](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/) | 二分查找      | [Go](solutions/__154_find_min_rotated_ar2.go)                | **   | 注意不能用low和mid比较                               |
| 155            | 腾讯笔试B题. 消除数字                                        | 栈            | [Go](solutions/__155_remove_numbers.go)                      | *    | 考场PTSD了 以为是动态规划                            |
| 156            | 腾讯笔试C题. 过钢索                                          | 动态规划      | [Go](solutions/__156_walk_tightrope.go)                      | ***  | 转移方程有点难想                                     |
| 157            | [215. 数组中的第K个最大元素](https://leetcode-cn.com/problems/kth-largest-element-in-an-array/) | 排序          | [C](solutions/__157_find_kth_element.c)                      | *    | 代码写的不够熟练                                     |
| 158            | [Shortest Path with Obstacle](https://codeforces.com/contest/1547/problem/A) | 数学          | [Go](solutions/__158_shortest_path_with_obstacle.go)         | *    | 回归刷题day1                                         |
| 159            | [Alphabetical Strings](https://codeforces.com/contest/1547/problem/B) | 益智          | [Go](solutions/__159_alphabetical_strings.go)                | *    | 思路简单，关注怎么实现                               |
| 160            | [680. 验证回文字符串 Ⅱ](https://leetcode-cn.com/problems/valid-palindrome-ii/) | 双指针/递归   | [Go](solutions/__160_valid_palindrome2.go)                   | *    |                                                      |
| 161            | [Kefa and First Steps](https://codeforces.com/problemset/problem/580/A) | 益智          | [C](solutions/__161_kefa_and_first_steps.c)                  | *    | Go会TLE，草                                          |
| 162            | [美团2021第一场B](https://www.nowcoder.com/test/28665243/summary) | 浮点数        | [Go](solutions/__162_meituan2021r1_B.go)                     | *    |                                                      |
| 163            | [美团2021第一场C](https://www.nowcoder.com/test/28665243/summary) | 贪心          | [Go](solutions/__163_meituan2021r1_C.go)                     | *    |                                                      |
| 164            | [美团2021第一场A](https://www.nowcoder.com/test/28665243/summary) | 益智          | [Go](solutions/__164_meituan2021r1_A.go)                     | **   | 找到规律就很简单                                     |
| 165            | 美团2022第一场C                                              | 二分插入/堆   | [Java](solutions/__Q165_prev_sum.java) / [C++](solutions/__165_prev_sum.cpp) | **   | TreeSet yyds                                         |
| **166**        | [169. 多数元素](https://leetcode-cn.com/problems/majority-element/) | 投票算法      | [Java](solutions/__Q166_majority_number.java)                | **   | 好巧妙的算法                                         |
| 167            | [1555C - Coin Rows](https://codeforces.com/problemset/problem/1555/C) | 动态规划      | [C++](solutions/__167_coins_row.cpp)                         | 1300 |                                                      |
| 168            | [492B - Vanya and Lanterns](https://codeforces.com/problemset/problem/492/B) | 益智          | [C++](solutions/__168_lantern.cpp)                           | 1200 |                                                      |
| 169            | [1553D - Backspace](https://codeforces.com/problemset/problem/1553/D) | 贪心          | [Java](solutions/__169_backspace.java)                       | 1500 |                                                      |
| 170            | [1368B - Codeforces Subsequences](https://codeforces.com/problemset/problem/1368/B) | 数学          | [C++](solutions/__170_codeforces_subseq.cpp)                 | 1500 | 直接打表 秀死你                                      |
| 171            | 阿里20220813A                                                | 动态规划      | [Java](solutions/__171_alibaba20220813_A.java)               | *    | 特点在于大数                                         |
