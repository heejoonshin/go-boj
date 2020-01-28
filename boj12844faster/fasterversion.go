package main

import (
	"bufio"
	"fmt"
	"os"
)
const maxsize = 2000000
var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }
var tree [maxsize]int
var lazy [maxsize]int
var data [500000]int
var N int
func init_tree(node, start,end int) int{
	if start == end{
		tree[node] = data[end]
		return tree[node]
	}
	mid := int((start + end)/2)
	tree[node] = init_tree(2 * node,start,mid) ^ init_tree(2 * node+1,mid+1,end)
	return tree[node]

}
func update_lazy(node,start,end int){
	if lazy[node] != 0{
		if (end-start +1) % 2==1 {
			tree[node] ^= lazy[node]
		}

		if start != end{
			lazy[node * 2] ^= lazy[node]
			lazy[node * 2 + 1] ^= lazy[node]
		}
		lazy[node] = 0
	}
}
func update_range(node, start,end, left,right int, diff int){
	update_lazy(node,start,end)
	if left > end || right < start{
		return
	}
	if left <= start && end <= right{
		if (end - start +1)%2 == 1{
			tree[node] ^= diff

		}

		if start != end{
			lazy[node *2] ^= diff
			lazy[node *2 + 1] ^= diff
		}
		return
	}
	mid := int((start + end) /2)
	update_range(2 * node,start,mid,left,right,diff)
	update_range(2 * node + 1,mid + 1,end, left, right,diff)
	tree[node] = tree[node * 2] ^ tree[node * 2 + 1]
}
func xor(node,start,end,left,right int)int{
	update_lazy(node,start,end)
	if left > end || right < start{
		return 0
	}
	if left <= start && end <= right{
		return tree[node]
	}
	mid := int((start + end) / 2)
	return xor(node * 2,start,mid,left,right) ^ xor(node * 2 + 1,mid+1,end,left,right)

}
func Update_range(start,end int,diff int){
	update_range(1,0,N-1,start,end,diff)

}
func Xor(start,end int)int{
	return xor(1,0,N-1,start,end)
}

func main() {
	defer writer.Flush()



	scanf("%d\n",&N)

	for i:=0; i <N; i++{
		if i == N-1{
			scanf("%d\n",&data[i])
		}else {
			scanf("%d",&data[i])
		}
	}
	init_tree(1,0,N-1)

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


			Update_range(l,r,diff)
		}
		if a==2{
			scanf("\n")


			ans := Xor(l,r)
			printf("%d\n",ans)

		}
	}




}
