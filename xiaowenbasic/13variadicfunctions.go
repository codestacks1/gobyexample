// 13VariadicFunctions
// 可变参数的函数
// C# 下使用params
package xiaowenbasic

import (
	"fmt"
)

func VariadicFuncMain() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
