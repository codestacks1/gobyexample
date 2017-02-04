// 42collectionfunctions
/*
	组合函数
	我们经常需要在程序集上进行操作，比如选择满足给定条件的所有项，
	或者将所有的项通过一个自定义函数映射到一个新的集合上

	在C#中用泛型来完成，Go不支持泛型
*/
package xiaowenbasic

import (
	"fmt"
	"strings"
)

func CollectionFunctionsMain() {
	fmt.Println("Collection Functions main")

	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println(index(strs, "pear"))
	fmt.Println(include(strs, "grape"))
	fmt.Println("prefix has P > ", any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println("prefix has P > ", all(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println("contains e > ", filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	fmt.Println(imap(strs, strings.ToUpper))
}

func index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func include(vs []string, t string) bool {
	return index(vs, t) >= 0
}

func any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func all(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0) // 初始化数组容器
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func imap(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
