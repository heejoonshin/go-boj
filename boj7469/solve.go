package main

import (
	"bufio"
	"fmt"
	"os"
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
	i ,j,k:=0,0,0
	lsize := len(mdata.Tree[leftnode])
	rsize := len(mdata.Tree[rightnode])
	mdata.Tree[node] = make([]int,lsize+rsize)

	for; i < len(mdata.Tree[leftnode]) && j < len(mdata.Tree[rightnode]);{
		if mdata.Tree[leftnode][i] < mdata.Tree[rightnode][j]{
			mdata.Tree[node][k]=mdata.Tree[leftnode][i]
			i++
			k++

		}else{
			mdata.Tree[node][k] = mdata.Tree[rightnode][j]
			j++
			k++

		}
	}
	for ; i < len(mdata.Tree[leftnode]); i++{
		mdata.Tree[node][k] = mdata.Tree[leftnode][i]
		k++
	}
	for ; j< len(mdata.Tree[rightnode]); j++{
		mdata.Tree[node][k] =mdata.Tree[rightnode][j]
		k++
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
func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}
var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }
func main() {
	defer writer.Flush()



	var N,Q int
	scanf("%d%d\n",&N,&Q)
	data := make([]int,N)

	for i := 0; i < N; i++{
		//fmt.Sscanf(reader.Text(),"%d",&data[i])
		scanf("%d",&data[i])
	}
	scanf("\n")
	t := NewMergetSort(&data)


	for;Q>0;Q--{

		//s = strings.Split(line," ")
		//l,_ := strconv.Atoi(s[0])
		//r,_ := strconv.Atoi(s[1])
		//v,_ := strconv.Atoi(s[2])
		var l,r,v int
		scanf("%d%d%d\n",&l,&r,&v)
		l--
		r--

		printf("%d\n", t.qeuryanw(l,r,v))


	}
	//wrieter.Flush()

}
