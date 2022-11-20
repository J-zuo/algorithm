package main

import "fmt"

/**

给定一个n个元素有序的（升序）整型数组nums 和一个目标值target 　，写一个函数搜索nums中的 target，如果目标值存在返回下标，否则返回 -1。


示例 1:

输入: nums = [-1,0,3,5,9,12], target = 9
输出: 4
解释: 9 出现在 nums 中并且下标为 4
示例2:

输入: nums = [-1,0,3,5,9,12], target = 2
输出: -1
解释: 2 不存在 nums 中因此返回 -1


提示：

你可以假设 nums中的所有元素是不重复的。
n将在[1, 10000]之间。
nums的每个元素都将在[-9999, 9999]之间。

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/binary-search
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//算法之二分查找
//二分查找的基本前提：1，线性表中的记录必须是关键码有序 2，它是一个顺序结构
//二分发的go语言实现查找切片中是否存在指定元素

/**
此种注释内容高亮显示
*/
func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(IsElementTrueAndReturnIndex(s, 10))

}

//给定一个切片和一个指定元素，判断该元素是否在此切片中，如果在，请返回该元素最小下标,如果不再请返回-1

/**
解体过程：借鉴别人的想法，确定好需要查找的区间，根据数组下标，有两种，一种是 闭区间[0,len(s)],一种是开区间[0,len(s))
注意点：闭区间在for循环时候的条件和替换边界是是否包含
*/

func IsElementTrueAndReturnIndex(s []int, e int) (idx int) {
	left := 0
	right := len(s) - 1

	for left <= right {
		mid := (left + right) >> 1
		if s[mid] > e {
			right = mid - 1
		} else if s[mid] < e {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
