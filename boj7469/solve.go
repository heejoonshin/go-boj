package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MergeSort struct{
	Tree [][]int
	N int

}
func NewMergetSort(data *[]int)*MergeSort{

	ret := new(MergeSort)
	ret.N = len(*data)
	ret.Tree = make([][]int,ret.N*4)
	ret.init(data,0,ret.N-1,1)
	return ret

}
func (mdata *MergeSort)merge(leftnode,rightnode,node int){
	i ,j:=0,0
	for; i < len(mdata.Tree[leftnode]) && j < len(mdata.Tree[rightnode]);{
		if mdata.Tree[leftnode][i] < mdata.Tree[rightnode][j]{
			mdata.Tree[node] = append(mdata.Tree[node],mdata.Tree[leftnode][i])
			i++

		}else{
			mdata.Tree[node] = append(mdata.Tree[node],mdata.Tree[rightnode][j])
			j++

		}
	}
	for ; i < len(mdata.Tree[leftnode]); i++{
		mdata.Tree[node] = append(mdata.Tree[node],mdata.Tree[leftnode][i])
	}
	for ; j< len(mdata.Tree[rightnode]); j++{
		mdata.Tree[node] = append(mdata.Tree[node],mdata.Tree[rightnode][j])
	}
}
func (mdata *MergeSort)init(data *[]int,start,end ,node int){

	if start == end{
		mdata.Tree[node] = append(mdata.Tree[node],(*data)[end])
		return
	}
	mid := int((start + end)/2)
	mdata.init(data,start,mid,2*node)
	mdata.init(data,mid+1,end,2*node+1)
	mdata.merge(2 * node,2 *node + 1,node)


}
func(mdata *MergeSort)query(start,end,node,left,right,value int) int{
	if left > end || right < start{
		return 0
	}
	if start <= left && right <= end{
		return upper_boud(&mdata.Tree[node],len(mdata.Tree[node]),value) - 1
	}
	mid := int((left+right)/2)
	return mdata.query(start,end,2 * node,left,mid,value) + mdata.query(start,end,2*node+1,mid+1,right,value)
}
func(mdata *MergeSort)qeuryanw(left,right,value int)int{
	lo , hi := -1000000003,1000000003
	for;lo<=hi;{
		mid := int((lo+hi)/2)
		rank := mdata.query(left,right,1,0,mdata.N-1,mid)
		if rank < value{
			lo = mid + 1
		}else{
			hi = mid -1

		}
	}
	return lo
}
func upper_boud(arr *[]int,end,value int)int{
	var start = 0
	var mid int
	for; end - start > 0;{
		mid = int((start+end)/2)


		if((*arr)[mid] <= value){
			start = mid + 1
		}else{
			end = mid
		}
	}
	return end +1
}

func main() {

	reader := bufio.NewScanner(os.Stdin)
	wrieter := bufio.NewWriter(os.Stdout)
	var N,Q int;

	step := 0
	var t *MergeSort
	var l,r,v int
	i :=0
	run := 0
	finish := false
	var data []int
	for ;finish != true; {
		reader.Scan()
		line := reader.Text()
		s := strings.Split(line, " ")

		for j := 0; j < len(s); j++ {
			if step == 0 {
				if i == 0 {
					N, _ = strconv.Atoi(s[j])
				}
				if i == 1 {
					Q, _ = strconv.Atoi(s[j])

				}
				i++
				if i == 2 {
					step++
					i = 0

				}
			} else if step == 1 {
				if i== 0{
					data = make([]int,N)
				}
				data[i], _ = strconv.Atoi(s[j])
				i++
				if i == N {
					i = 0
					t = NewMergetSort(&data)
					step++

				}
			} else if step == 2 {
				if i == 0 {
					l, _ = strconv.Atoi(s[j])
				}
				if i == 1 {
					r, _ = strconv.Atoi(s[j])
				}
				if i == 2 {
					v, _ = strconv.Atoi(s[j])
				}
				i++
				if i == 3 {

					i = 0
					l--
					r--

					fmt.Fprintf(wrieter, "%d\n", t.qeuryanw(l, r, v))
					run++
					if run == Q {
						finish = true
					}

				}
			}
		}
	}



	wrieter.Flush()
	
}
