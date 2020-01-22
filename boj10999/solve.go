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
	tree []int64
	lazy []int64
	N int
}
func NewSegmentTree_lazy(data *[]int64) *SegmentTree_lazy{
	tree := new(SegmentTree_lazy)
	tree.N = len(*data)
	tree.tree = make([]int64,4 *tree.N)
	tree.lazy = make([]int64,4 * tree.N)
	tree.init(data,1,0,tree.N-1)
	return tree



}
func (t *SegmentTree_lazy)init(data *[]int64,node, start,end int) int64{
	if start == end{
		t.tree[node] = (*data)[end]
		return t.tree[node]
	}
	mid := int((start + end)/2)
	t.tree[node] = t.init(data,2 * node,start,mid) + t.init(data,2 * node+1,mid+1,end)
	return t.tree[node]

}
func (t *SegmentTree_lazy)update_lazy(node,start,end int){
	if t.lazy[node] != 0{
		t.tree[node] += int64(end - start +1) * t.lazy[node]

		if start != end{
			t.lazy[node * 2] += t.lazy[node]
			t.lazy[node * 2 + 1] += t.lazy[node]
		}
		t.lazy[node] = 0
	}
}
func (t *SegmentTree_lazy)update_range(node, start,end, left,right int, diff int64){
	t.update_lazy(node,start,end)
	if left > end || right < start{
		return
	}
	if left <= start && end <= right{
		t.tree[node] += int64(end - start + 1) * diff
		if start != end{
			t.lazy[node *2] += diff
			t.lazy[node *2 + 1] += diff
		}
		return
	}
	mid := int((start + end) /2)
	t.update_range(2 * node,start,mid,left,right,diff)
	t.update_range(2 * node + 1,mid + 1,end, left, right,diff)
	t.tree[node] = t.tree[node * 2] + t.tree[node * 2 + 1]
}
func (t *SegmentTree_lazy)sum(node,start,end,left,right int)int64{
	t.update_lazy(node,start,end)
	if left > end || right < start{
		return 0
	}
	if left <= start && end <= right{
		return t.tree[node]
	}
	mid := int((start + end) / 2)
	return t.sum(node * 2,start,mid,left,right) + t.sum(node * 2 + 1,mid+1,end,left,right)

}
func (t *SegmentTree_lazy)Update_range(start,end int,diff int64){
	t.update_range(1,0,t.N-1,start,end,diff)

}
func (t *SegmentTree_lazy)Sum(start,end int)int64{
	return t.sum(1,0,t.N-1,start,end)
}

func main() {
	defer writer.Flush()
	var n,m,k int
	scanf("%d %d %d\n",&n,&m,&k)
	data := make([]int64,n)
	for i := 0; i < n; i++{
		scanf("%d\n",&data[i])
	}
	segment := NewSegmentTree_lazy(&data)

	for i :=0; i < m+k; i++{
		var c int
		var s,e int
		var v int64
		scanf("%d ",&c)
		if c == 1{
			scanf("%d %d %d",&s,&e,&v)
			scanf("\n")
			s--
			e--
			segment.Update_range(s,e,v)

		}
		if c == 2{
			scanf("%d %d",&s,&e)
			scanf("\n")
			s--
			e--
			ans :=segment.Sum(s,e)
			printf("%d\n",ans)
		}
	}
	
}
