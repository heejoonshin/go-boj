package main
const MAX_N = 100001
var G [MAX_N][]int
type RMQ struct{
	n int
	tree []int


}
func NewRMQ(n int)*RMQ{
	ret := new(RMQ)
	ret.n = n
	ret.tree = make([]int,n*4)

	return ret
}
func(t *RMQ)update(index, newvalue, node,nodeleft,noderight int) int{
	if index < nodeleft || noderight < index{
		return t.tree[node]
	}
	if nodeleft == noderight{
		t.tree[node] = newvalue
		return t.tree[node]
	}
	var mid int
	mid = int((nodeleft + nodeleft)/2)
	leftvalue := t.update(index,newvalue,node *2 ,nodeleft,mid)
	rightvalue := t.update(index,newvalue,node *2 + 1, mid +1, noderight)
	if leftvalue < rightvalue{
		t.tree[node] = rightvalue
	}else{
		t.tree[node] = leftvalue
	}
	return t.tree[node]
}
func(t *RMQ)query(left,right, node,nodeleft,noderight int)int{
	if right < nodeleft || noderight < left{
		return -987654321
	}
	if left <= nodeleft && noderight <= right{
		return t.tree[node]
	}
	mid := int((nodeleft + noderight) / 2)
	leftvalue := t.query(left,right,node * 2,nodeleft,mid)
	rightvalue := t.query(left,right,node *2 + 1,mid+1,noderight)
	if leftvalue < rightvalue{
		return rightvalue
	}else{
		return leftvalue
	}
}
func main() {
	
}
