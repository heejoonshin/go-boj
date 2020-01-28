package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }

type SegmentTree_lazy struct{
	tree []int
	lazy []int
	N int
}
func NewSegmentTree_lazy(data *[]int) *SegmentTree_lazy{
	tree := new(SegmentTree_lazy)
	tree.N = len(*data)
	tree.tree = make([]int,4 *tree.N)
	tree.lazy = make([]int,4 * tree.N)
	tree.init(data,1,0,tree.N-1)
	return tree



}
func (t *SegmentTree_lazy)init(data *[]int,node, start,end int) int{
	if start == end{
		t.tree[node] = (*data)[end]
		return t.tree[node]
	}
	mid := int((start + end)/2)
	t.tree[node] = t.init(data,2 * node,start,mid) ^ t.init(data,2 * node+1,mid+1,end)
	return t.tree[node]

}
func (t *SegmentTree_lazy)update_lazy(node,start,end int){
	if t.lazy[node] != 0{
		if (end-start +1) % 2==1 {
			t.tree[node] ^= t.lazy[node]
		}

		if start != end{
			t.lazy[node * 2] ^= t.lazy[node]
			t.lazy[node * 2 + 1] ^= t.lazy[node]
		}
		t.lazy[node] = 0
	}
}
func (t *SegmentTree_lazy)update_range(node, start,end, left,right int, diff int){
	t.update_lazy(node,start,end)
	if left > end || right < start{
		return
	}
	if left <= start && end <= right{
		if (end - start +1)%2 == 1{
			t.tree[node] ^= diff

		}

		if start != end{
			t.lazy[node *2] ^= diff
			t.lazy[node *2 + 1] ^= diff
		}
		return
	}
	mid := int((start + end) /2)
	t.update_range(2 * node,start,mid,left,right,diff)
	t.update_range(2 * node + 1,mid + 1,end, left, right,diff)
	t.tree[node] = t.tree[node * 2] ^ t.tree[node * 2 + 1]
}
func (t *SegmentTree_lazy)xor(node,start,end,left,right int)int{
	t.update_lazy(node,start,end)
	if left > end || right < start{
		return 0
	}
	if left <= start && end <= right{
		return t.tree[node]
	}
	mid := int((start + end) / 2)
	return t.xor(node * 2,start,mid,left,right) ^ t.xor(node * 2 + 1,mid+1,end,left,right)

}
func (t *SegmentTree_lazy)Update_range(start,end int,diff int){
	t.update_range(1,0,t.N-1,start,end,diff)

}
func (t *SegmentTree_lazy)Xor(start,end int)int{
	return t.xor(1,0,t.N-1,start,end)
}

func main() {
	defer writer.Flush()


	var N int
	scanf("%d\n",&N)
	data := make([]int,N)
	for i:=0; i <N; i++{
		if i == N-1{
			scanf("%d\n",&data[i])
		}else {
			scanf("%d",&data[i])
		}
	}
	segment := NewSegmentTree_lazy(&data)
	var M int
	scanf("%d\n",&M)
	for i := 0; i < M ;i++{
		var a,l,r int
		var diff int
		scanf("%d%d%d",&a,&l,&r)
		if l > r{
			temp := r
			r = l
			l = temp
		}
		if a==1{

			scanf("%d\n",&diff)


			segment.Update_range(l,r,diff)
		}
		if a==2{
			scanf("\n")


			ans := segment.Xor(l,r)
			printf("%d\n",ans)

		}
	}




}
