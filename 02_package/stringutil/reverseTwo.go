package stringutil

// 数组：类型[n]T需指定个数([...]T 定义赋值时可忽略元素个数)
// 切片: 类型[]T不需要指定个数 引用类型 指向数组的指针

func reverseTwo(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i] // 相当于(r[i], r[j]) = (r[j], r[i])
	}
	return string(r)
}

// this demonstrates how an unexported function
// can be used by an exported function in the same Packagepackage stringutil

// MyName will be exported because it starts with a capital letter.
var MyName = "Todd"
