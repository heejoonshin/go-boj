package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
const Max_N = 100001
var G [Max_N][]int
var par[Max_N][21]int
var depth[Max_N]int
var visited [Max_N]bool
var N int

func dfs(here,d int){

	visited[here] = true
	depth[here] = d
	for _,next := range G[here]{
		if visited[next]{
			continue
		}
		par[next][0] = here
		dfs(next,d+1)
	}
}

func preprocessing(){
	for j:=1; j < 21; j++{
		for i:=1; i <=N; i++{
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
		if par[x][i] != par[y][i]{
			x = par[x][i]
			y = par[y][i]
		}

	}
	return par[x][0]
}
func main() {

	reader := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	reader.Scan()
	N, _ = strconv.Atoi(reader.Text())
	for i:= 0; i < N-1; i++{
		reader.Scan()
		line := reader.Text()
		s := strings.Split(line," ")
		u,_ := strconv.Atoi(s[0])
		v,_ := strconv.Atoi(s[1])
		G[u] = append(G[u],v)
		G[v] = append(G[v],u)
	}
	dfs(1,0)
	preprocessing()
	reader.Scan()
	Q, _ := strconv.Atoi(reader.Text())
	for ; Q>0; Q--{
		reader.Scan()
		line := reader.Text()
		s := strings.Split(line," ")
		x,_ := strconv.Atoi(s[0])
		y,_ := strconv.Atoi(s[1])
		fmt.Fprintln(w,lca(x,y))

	}
	w.Flush()
}
