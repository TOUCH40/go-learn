package main

import (
	"math"
	"strings"

)

func main(){
	strToInt("42")
}

func strToInt(str string) int {
	str = strings.Trim(str," ")
	if(len(str)==0){
		return 0
	}
	out,start,sign,boundary:=0,1,1,math.MaxInt32/10
	if(str[0]=='-'){
		sign = -1
	}else if(str[0]!='+'){
		start = 0
	}
	for i:=start;i<len(str);i++{
		if str[i]<'0'||str[i]>'9' {
			break
		}
		if out>boundary||(out==boundary&&str[i]>'7'){// key point
			if(sign==1){
				return math.MaxInt32
			}else{
				return math.MinInt32
			}
		}

			out=out*10+int(str[i]-'0')
	}
	return out*sign
}