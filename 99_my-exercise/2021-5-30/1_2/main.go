package main

func main() {
	print("&b", 6)
	print("&b", 5)
	println(6 & (5))
	// isSumEqual("j", "j", "bi")
}

// func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
// 	fv, sv, tv := 0, 0, 0
// 	for _, it := range firstWord {
// 		fv += (int(it) - 'a') + fv*10
// 	}
// 	for _, it := range secondWord {
// 		sv += (int(it) - 'a') + sv*10
// 	}
// 	for _, it := range targetWord {
// 		tv += (int(it) - 'a') + tv*10
// 	}
// 	return (fv + sv) == tv
// }

// 12341234  3

// func maxValue(n string, x int) string {
// 	l := len(n)
// 	if n[0] == '-' {
// 		for i := 1; i < l; i++ {
// 			iv, _ := strconv.Atoi(string(n[i]))
// 			if x < iv {
// 				n = n[:i] + strconv.Itoa(x) + n[i:]
// 				break
// 			} else if i == l-1 {
// 				n += strconv.Itoa(x)
// 			}
// 		}
// 	} else {
// 		for i := 0; i < l; i++ {
// 			iv, _ := strconv.Atoi(string(n[i]))
// 			if x > iv {
// 				n = n[:i] + strconv.Itoa(x) + n[i:]
// 				break
// 			} else if i == l-1 {
// 				n += strconv.Itoa(x)
// 			}
// 		}
// 	}
// 	return n
// }

// func assignTasks(servers []int, tasks []int) []int {
// 	sl := len(servers)

// 	// server index form weight
// 	status := [][]int{}
// 	for i,it:= range servers{
// 		status[it]=i
// 	}

// 	answer := make(map[int]int, len(tasks))
// 	for i, it := range tasks {
// 		for ist, itst := range status {
// 			if itst==0{
// 				answer[i]=
// 			}
// 		}
// 	}
// }
