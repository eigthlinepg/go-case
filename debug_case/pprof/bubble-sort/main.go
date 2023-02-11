/*
 * @ File: file_name
 * @ Author: author_name
 * @ Date: 2023年 02月 11日 星期六 14:21:55 CST
 * @ Description: file_description
 */

package main

import (
	"fmt"
	"net/http"
	"time"

	_ "net/http/pprof"

	"math/rand"
)

func generate(n int) []int {

	seed := time.Now().UnixNano()
	rand.Seed(seed)
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())

	}
	return nums
}

/*func bubbbleSort(nums []int) {*/
//for i := 0; i < len(nums); i++ {
//for j := 0; j < len(nums)-i; j++ {
//if nums[j] < nums[j-1] {

//nums[j], nums[j-1] = nums[j-1], nums[j]
//}
//}
//}
//}

/*
* Function: function_name
* Description: function_description
* Parameters:
  - : param1_name - param1_description

* Returns: return_value
*/
func main() {

	// 启动wbe
	go func() {
		http.ListenAndServe("0.0.0.0:8899", nil)
	}()

	for {
		n := 10
		for i := 0; i < 5; i++ {
			nums := generate(n)

			fmt.Println("Before bubbleSort:", nums)
			gots := bubbleSort(nums)
			fmt.Println("After bubbleSort:", gots)

			n *= 10
		}

	}

}
