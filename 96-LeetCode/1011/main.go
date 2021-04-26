package main

import (
	"fmt"
)

func main() {
	In := []int{1,2,3,1,1}
	ret := shipWithinDays(In,4)
	fmt.Println(ret)
}
func shipWithinDays(weights []int, D int) int {
	 check:=func(cap int)bool{
		ons:=0
		ndays:=1
		for _,_iw:=range weights{
			ons+=_iw
			if(ons>cap){
				ons=_iw
				ndays++
			}
		}
		return ndays<=D
	}
	//二分法 division***
	l:=weights[0]
	r:=0
	for _,_i:=range weights{
		if l<_i{
			l=_i
		}
		r+=_i
	}
	// templete
	for;l<r;{
		mid := (l + r) >> 1
		if check(mid){
			r=mid
		}else{ 
			l=mid+1
		}
	}
	return l

}

// func search(nums []int,left int,right int,target int) int {
// 	for ;left<right;{
// 		mid:=left+(right-left)/2
// 	}
// }
