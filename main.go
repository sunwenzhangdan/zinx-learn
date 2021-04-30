/*
 * @Author: seven
 * @Date: 2021-04-22 11:39:36
 * @LastEditTime: 2021-04-30 11:40:31
 * @LastEditors: seven sun
 * @Description: In User Settings Edit
 * @FilePath: /zinx-learn/main.go
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

/**
 * @description:
 * @param {*}
 * @return {*}
 */
func main() {
	fmt.Println("hello")
	var mutex sync.Mutex
	count := 0
	for r := 0; r < 50; r++ {
		go func() {
			mutex.Lock()
			count += 1
			mutex.Unlock()
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("the count is : ", count)
	fmt.Println(devideTwo([]int{1, 2, 3, 5, 6, 7, 8, 9, 10, 11, 23, 55, 66, 77}, 11))
	fmt.Println(devidetwo1([]int{1, 2, 3, 5, 6, 7, 8, 9, 10, 11, 23, 55, 66, 77}, 11))
	fmt.Println(StealRoom([]int{2, 11, 8, 6, 10, 2, 6, 7, 8, 2}))
	fmt.Println(merger_sorted_list([]int{1, 3, 5, 7, 845, 869}, []int{12, 34, 56, 80, 89, 789}))
	fmt.Println(StealRoomdy([]int{2, 11, 8, 6, 10, 2, 6, 7, 8, 2}))
	//for i := 0; i < 100; i++ {
	//	fmt.Println(<-rand())
	//	}
	for j := range xrange() {
		fmt.Println(j)
	}
	time.Sleep(5 * time.Second)

}

func CoinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		return -1
	}
	return 0
}

/**
 * @description:
 * @param {[]int} moneys  总共的房子及金额
 * @param {int}
 * @return {*}
 */
func StealRoom(moneys []int) int {
	k := len(moneys)
	if k == 1 {
		return moneys[0]
	} else if k == 2 {
		return Max(moneys[0], moneys[1])
	} else {
		return Max(StealRoom(moneys[0:k-1]), (StealRoom(moneys[0:k-2]) + moneys[k-1]))
	}
}

func StealRoomdy(moneys []int) int {
	k := len((moneys))
	FF := make([]int, k+1, k+1)
	if k == 0 {
		return moneys[0]
	} else if k == 1 {
		return Max(moneys[0], moneys[1])
	} else {
		FF[0] = moneys[0]
		FF[1] = Max(moneys[0], moneys[1])
		for i := 2; i < k; i++ {
			FF[i] = Max(FF[i-1], FF[i-2]+moneys[i])
		}

	}
	return FF[k-1]
}

/**
 * @description:两数中取最大值
 * @param {*} x
 * @param {int} y
 * @return {*}
 */
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/**
 * @description: 背包问题
 * @param {*}
 * @return {*}
 */

// 原问题  假如有n间房子可以偷窃
// 子问题

// 我们再回顾整个流程，如何尝试去处理一个动态规划问题？

// 首先分析这个问题符不符合动态规划最重要的前两个性质（重叠子问题，最优子结构）；
// 如果满足前两个性质，那么我们尝试写出状态转移方程，也即递归式；
// 优化：将自顶向下的递归式（函数调用）改为自底向上的动态规划（for loop）
// 好了，今天的动态规划问题到此就结束啦，当然动态规划的威力还远不止于此，关于动态规划更多的内容，我们后续再见～
func devideTwo(s1 []int, target int) int {
	return devide(s1, 0, len(s1)-1, target)
}

func devide(s []int, start, end, target int) int {
	mid := (start + end) / 2
	if s[mid] == target {
		return mid
	} else if s[mid] < target {
		start = mid + 1
		return devide(s, start, end, target)
	} else if s[mid] > target {
		end = mid - 1
		return devide(s, start, end, target)
	} else {
		return -1
	}
	return -1
}

func devidetwo1(s []int, target int) int {
	//第一个从0开始
	//在[start end]区间里面找target
	start := 0
	end := len(s) - 1
	for end >= start {
		mid := (start + end) / 2
		if s[mid] == target {
			return mid
		} else if target < s[mid] {
			end = mid - 1
		} else if target > s[mid] {
			start = mid + 1
		} else {
			return -1
		}
	}
	return -1
}

type BinaryTree struct {
	root *TreeNode
}

type TreeNode struct {
	data  interface{}
	left  *TreeNode
	right *TreeNode
}

func merger_sorted_list(l1, l2 []int) []int {
	l1_lenth := len(l1)
	l2_lenth := len(l2)
	l3 := make([]int, 0, l1_lenth+l2_lenth)
	l1_index := 0
	l2_index := 0
	for l1_index < l1_lenth && l2_index < l2_lenth {
		if l1[l1_index] < l2[l2_index] {
			l3 = append(l3, l1[l1_index])
			l1_index++
		} else {
			l3 = append(l3, l2[l2_index])
			l2_index++
		}
	}
	for l1_index < len(l1) {
		l3 = append(l3, l1[l1_index])
		l1_index++
	}
	for l2_index < len(l2) {
		l3 = append(l3, l2[l2_index])
		l1_index++
	}

	return l3
}

/*
随机数设计模式
*/

func rand() chan int {
	ch := make(chan int)
	go func() {
		for {
			select {
			case ch <- 0:
			case ch <- 1:
			case ch <- 2:
			default:
				fmt.Println("dfd")
			}
		}
	}()
	return ch
}

/*
生成器模式
*/

func xrange() chan int {
	var ch chan int = make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			ch <- i
		}
	}()
	return ch
}
