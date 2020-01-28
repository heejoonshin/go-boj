package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const MAX_N = 100001
var G [MAX_N][]edge
var par[MAX_N][21]int
var depth[MAX_N] int
var visited [MAX_N]bool
var weight[MAX_N]int
var N int
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

func predfs(here,d int){

	visited[here] = true
	depth[here] =d
	weight[here] = 1

	for _,next := range G[here]{
		if visited[next.next]{
			continue
		}
		par[next.next][0] = here
		predfs(next.next,d+1)
		weight[here] += weight[next.next]

	}
}
var color = 1
var p = 0
var colored [MAX_N]int
var haednode[MAX_N]int
var nodetopos[MAX_N]int
var pos[MAX_N]int
func dfs(here int,T *RMQ){
	visited[here] = true
	heavy := -1
	heavynode := 0
	pos[p] = here
	nodetopos[here]=p
	p++
	for _,next := range G[here]{
		if visited[next.next]{
			continue
		}
		if heavy < weight[next.next]{
			heavy = weight[next.next]
			heavynode = next.next
		}
	}
	for _,next := range G[here]{
		if visited[next.next]{
			continue
		}
		if heavynode != next.next{
			color ++
		}
		colored[next.next] = color
		T.update()
		dfs(next.next,T)
	}
}
func preprocessing(){
	for j:=1; j < 21; j++{
		for i := 1; i <= N; i++{
			par[i][j] = par[par[i][j-1]][j-1]
		}
	}
}
func lca(x,y int)int{
	if depth[x] > depth[y]{
		y, x = x,y

	}


	for i := 20; i>=0; i--{

		if depth[y] - depth[x] >= int(1<<uint(i)){
			y = par[y][i]
		}
	}
	if x == y {
		return x
	}
	for i:=20; i>=0; i--{
		if par[x][i] != par[y][i]{r
			x = par[x][i]
			y = par[y][i]
		}

	}
	return par[x][0]
}
type edge struct{
	next int
	weight int
}
func main() {
	reader := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	reader.Scan()
	N, _ = strconv.Atoi(reader.Text())
	tree := NewRMQ(N)
	for i:= 0; i < N-1; i++{
		reader.Scan()
		line := reader.Text()
		s := strings.Split(line," ")
		u,_ := strconv.Atoi(s[0])
		v,_ := strconv.Atoi(s[1])
		w,_ := strconv.Atoi(s[2])
		G[u] = append(G[u],edge{v,w})
		G[v] = append(G[v],edge{u,w})
	}
	predfs(1,0)
	for i:=0;i <=N; i++{
		visited[i] = false
	}
	preprocessing()
	dfs(1,tree)




}
