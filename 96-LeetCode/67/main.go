package main

import (
	"fmt"
	"strconv"
)

func main(){
	a, b := "1010" , "1011"
 fmt.Println(addBinary(a,b))
}

func addBinary(a string, b string) string {
	var out string
	lenA, lenB := len(a), len(b)
	n:=lenA
	if(lenB>n){
		n=lenB
	}
	carry:=0
	for i:=0;i<n;i++{
		if(i<lenA){
			// carry += int(a[lenA-i-1] - '0')
			tNum,error:=strconv.Atoi(string(a[lenA-i-1]))
			if(error==nil){
				carry+=tNum
			}
		}
		if(i<lenB){
		  carry += int(b[lenB-i-1] - '0')
		
		}
		out=strconv.Itoa(carry%2)+out
			carry/=2
	}
	if(carry>0){
		out="1"+out
	}
	 

	return out
}