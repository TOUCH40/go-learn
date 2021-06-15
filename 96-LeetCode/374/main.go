package main

import "fmt"

func main(){
	fmt.Println(guessNumber(10))
}
func guess(i int)int{
	if(i>6){
		return 1
	}else if(i<6){
		return -1
	}else{
		return 0
	}
}
func guessNumber(n int) int {
	l,r:=1,n
	for l <= r{
			mid := l + ( r-l )>>1
			if(guess(mid)==0){
					return mid
			}else if(guess(mid)==1){
					r=mid-1
			}else{
					l=mid+1
			}
	}
	return -1
}