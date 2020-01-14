package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RMQ struct{
	size int
	tree []int64
}
func NewRMQ(data *[]int64)*RMQ{
	r := new(RMQ)
	r.size = len(*data)
	r.tree = make([]int64, 4 * r.size)
	r.init(data,0,r.size-1,1)
	return r


}
func(r *RMQ)init(data *[]int64,leftnode,rightnode,node int)int64{
	if leftnode == rightnode{
		r.tree[node] = (*data)[leftnode]
		return r.tree[node]
	}
	mid := int((leftnode + rightnode)/2)
	leftvalue := r.init(data,leftnode,mid,2*node)
	rightvalue := r.init(data,mid + 1,rightnode,2 * node +1)
	if leftvalue < rightvalue{
		r.tree[node] = leftvalue
	}else{
		r.tree[node] = rightvalue
	}
	return r.tree[node]
}
func (r *RMQ)query(left,right,node,nodeleft,noderight int)int64{
	if noderight < left || right < nodeleft{
		return 1000000001
	}
	if left <= nodeleft && noderight <= right{
		return r.tree[node]
	}
	mid := int((nodeleft+noderight)/2)
	leftvalue := r.query(left,right,node*2,nodeleft,mid)
	rightvalue := r.query(left,right,node *2 + 1, mid+1,noderight)
	if leftvalue < rightvalue{
		return leftvalue
	}else{
		return rightvalue
	}
}
func (r *RMQ)Qeury(left,right int)int64{

	return r.query(left,right,1,0,r.size-1)

}
func main() {
	reader := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	reader.Scan()
	line := reader.Text()
	s := strings.Split(line," ")
	N, _ := strconv.Atoi(s[0])
	M, _ := strconv.Atoi(s[1])
	data := make([]int64,N)
	for i := 0; i < N; i++{
		reader.Scan()
		a ,_ := strconv.ParseInt(reader.Text(),10,64)
		data[i] = a
	}
	minRMQ := NewRMQ(&data)
	for i := 0; i < N; i++{
		data[i] *= int64(-1)
	}
	maxRMQ := NewRMQ(&data)


	for ; M>0; M--{
		reader.Scan()
		line := reader.Text()
		s := strings.Split(line," ")
		l,_ := strconv.Atoi(s[0])
		r,_ := strconv.Atoi(s[1])
		l--
		r--
		m := minRMQ.Qeury(l,r)
		M := maxRMQ.Qeury(l,r)
		fmt.Fprintf(writer,"%d %d\n",m,-1 * M)
	}
	writer.Flush()
}
