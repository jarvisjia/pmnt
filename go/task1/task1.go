package task1

import (
	"fmt"
	"sort"
	"strconv"
)

/*
1.只出现一次的数字
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，
结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := make(map[int]int)
	for _, v := range nums {
		if result[v] == 0 {
			result[v] = 1
		} else {
			delete(result, v) //如果出现过，就从map中删除
		}
	}
	for key := range result { //map中剩余的就是只出现一次的
		return key
	}
	return 0
}

/*
2.回文数
判断一个整数是否是回文数
*/
func isPalindrome(x int) bool {
	var str string = strconv.Itoa(x)
	var length = len(str)
	var mid = length / 2
	for i := 0; i <= mid; i++ {
		if str[i] != str[length-1-i] { //前半部分和后半部分逐个对比
			return false
		}
	}
	return true
}

/*
3.字符串，有效的括号
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效
*/
// '('，')'，'{'，'}'，'['，']'
func isValid(s string) bool {
	if len(s) < 2 { //字符串长度小于2，直接返回false
		return false
	}
	var m = map[string]string{
		"{": "}",
		"[": "]",
		"(": ")",
	}
	var p []string
	for i, v := range s {
		var sv = string(v)
		if sv == "{" || sv == "[" || sv == "(" { //左边的符号，先入栈
			p = append(p, sv)
		} else {
			//右边的符号，i>0从第二个字符开始算；先确保栈中有值（防止后面的栈溢出），再跟map中的配对值对比
			if i > 0 && len(p) > 0 && m[p[len(p)-1]] == sv {
				p = p[:len(p)-1]
			} else {
				return false
			}
		}
	}
	return len(p) == 0
}

/*
4.最长公共前缀
查找字符串数组中的最长公共前缀
*/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var num = len(strs[0])
	for i := range strs {
		if num > len(strs[i]) {
			num = len(strs[i]) //求出最短字符串的长度
		}
	}
	var sameStr string
	var flag bool = false
	for m := range num {
		var sameS byte = strs[0][m]
		for n := range len(strs) { //循环对比每个字符串中对应位置的值是否相同
			if sameS != strs[n][m] {
				flag = true
				break
			}
		}

		if flag { //字符不同，则退出判断
			break
		} else {
			sameStr += string(strs[0][m]) //字符相同，则入sameStr
		}
	}
	return sameStr
}

/*
5.加一
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一
*/
func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}
	//加一不进位情况
	if digits[len(digits)-1] < 9 {
		digits[len(digits)-1] += 1
		return digits
	}
	//加一进位情况
	var i int = len(digits) - 1
	for ; i >= 0; i-- {
		if digits[i] == 9 {
			digits[i] = 0
		} else {
			digits[i] += 1
			break
		}
	}
	//每个数字都为9的情况
	if i == -1 {
		digits = append(digits, 0)
		copy(digits[1:], digits[:])
		digits[0] = 1
	}
	return digits
}

/*
6.删除有序数组中的重复项
给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。不要使用额外的数组空间，
你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，
一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
*/
func removeDuplicates(nums []int) int {
	var n = len(nums)
	if n == 0 {
		return 0
	}
	var i int = 1
	for j := 1; j < n; j++ {
		if nums[j] != nums[j-1] {
			nums[i] = nums[j]
			i++
		}
	}
	return i
}

/*
7.合并区间
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。请你合并所有重叠的区间，
并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。可以先对区间数组按照区间的起始位置进行排序，
然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；
如果没有重叠，则将当前区间添加到切片中。
*/
func merge(intervals [][]int) [][]int {
	var n = len(intervals)
	if n == 0 {
		return intervals
	}
	//给二维数组排序
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1] //第一列相同，按照第二列排序
		} else {
			return intervals[i][0] < intervals[j][0]
		}
	})
	var nArr [][]int
	var bigin int = intervals[0][0]
	var end int = intervals[0][1]
	for i := 1; i < n; i++ {
		if intervals[i][0] <= end {
			if intervals[i][0] < bigin {
				bigin = intervals[i][0]
			}
			if intervals[i][1] > end {
				end = intervals[i][1]
			}
		} else {
			var tmp = [2]int{bigin, end}
			nArr = append(nArr, tmp[:])
			bigin = intervals[i][0]
			end = intervals[i][1]
		}
	}
	var tmp = [2]int{bigin, end}
	nArr = append(nArr, tmp[:])
	return nArr
}

/*
8.两个数之和
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
*/
func twoSum(nums []int, target int) []int {
	// 记录遍历过的数和其下标
	idxMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		remain := target - nums[i]
		if idx, ok := idxMap[remain]; ok {
			return []int{idx, i}
		}
		idxMap[nums[i]] = i
	}
	return nil
}

func Task1() {
	//只出现一次的数字
	nums := []int{1, 2, 3, 1, 4, 4, 3, 7, 7}
	var result = singleNumber(nums)
	fmt.Println("1.只出现一次的数字", result)

	//回文数
	fmt.Println("2.回文数", isPalindrome(102301))

	//字符串，有效的括号
	fmt.Println(isValid("(){}}{"))
	fmt.Println("3.字符串，有效的括号", isValid("(])"))

	//最长公共前缀
	var strs = [3]string{"flower", "flow", "flight"}
	fmt.Println("4.最长公共前缀", longestCommonPrefix(strs[:]))

	//加一
	// digits := []int{4, 3, 2, 1}
	digits := []int{9, 2, 9, 9}
	fmt.Println("5.加一", plusOne(digits))

	//删除有序数组中的重复项
	numInt := []int{1, 2, 2, 3, 4, 4, 4, 4, 5, 6, 6, 6, 7, 9, 10}
	count := removeDuplicates(numInt)
	fmt.Println("6.删除有序数组中的重复项", "nums=", numInt, "length=", count)

	//合并区间
	internals := [][]int{{1, 2}, {12, 15}, {4, 7}, {20, 25}, {6, 10}, {11, 17}}
	fmt.Println("7.合并区间", merge(internals))

	//两个数之和
	nnum := []int{1, 3, 6, 8, 2, 7, 9}
	fmt.Println("8.两个数之和", twoSum(nnum, 5))
}
