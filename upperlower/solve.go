package main

import "fmt"

func lower_bound(arr []int,end ,value int)int {

	var start = 0
	var mid int
	for;end - start >0;{
		mid = int((start + end)/2)

		if(arr[mid] <value){
			start = mid + 1
		}else{
			end = mid
		}
	}

	return end + 1

}
func upper_boud(arr []int,end,value int)int{
	var start = 0
	var mid int
	for; end - start > 0;{
		mid = int((start+end)/2)

		if(arr[mid] <= value){
			start = mid + 1
		}else{
			end = mid
		}
	}
	return end +1
}

func main() {

	arr := []int{1,2,4,5,6,6,7,7,7,9}
	fmt.Println(upper_boud(arr,10,6))
}
