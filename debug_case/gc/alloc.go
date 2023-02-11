/*
 * @ File: file_name
 * @ Author: author_name
 * @ Date: 2023年 02月 11日 星期六 13:58:57 CST
 * @ Description: file_description
 */

package main

import (
	"fmt"
	"time"
)

/*
 * Function: function_name
 * Description:此代码将循环 100000000 次，并跟踪分配和释放的内存量。
			可以看到，内存分配和释放是周期性的，在代码的特定阶段内存被分配，
			在另一个阶段内存被释放。这是 Go 垃圾回收器的工作原理，它通过监控内存分配和使用情况，
			在必要时释放不再使用的内存。

 * Parameters:
    - : param1_name - param1_description
 * Returns: return_value
*/

func main() {
	var memAllocated uint64
	var memFreed uint64

	for i := 0; i < 100000000; i++ {
		memAllocated += 64
		makeSlice := make([]byte, 64)
		memFreed += (uint64)(64 * (i * 2))
		if i%2 == 0 {
			memFreed += 64
			continue
		}
		_ = makeSlice
	}
	fmt.Println("memory allocated:", memAllocated)
	fmt.Println("memory freed:", memFreed)

	time.Sleep(20 * time.Second)
}
