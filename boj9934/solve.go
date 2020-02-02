package main

import (
	"bufio"
	"fmt"
	"os"
)

var builtTree [10][]int
var K int

func travelNode(data *[]int,left,right, depth int) {

	if right < left{
		return
	}
	if left == right{
		builtTree[depth] = append(builtTree[depth],(*data)[left])
		return
	}
	mid := int((left+right) /2)

	travelNode(data,left,mid-1,depth+1)
	builtTree[depth] = append(builtTree[depth],(*data)[mid])
	travelNode(data,mid+1,right,depth+1)
}

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{}) { fmt.Fscanf(reader, f, a...) }
func main() {

	defer writer.Flush()
	scanf("%d\n",&K)
	nodes := int(1<<uint(K)) -1
	data := make([]int,nodes)
	for i :=0; i < nodes; i++{
		scanf("%d ",&data[i])
	}

	travelNode(&data,0,nodes-1,0)
	for h :=0; h < K; h++{
		for i := 0; i < len(builtTree[h]); i++{
			printf("%d ",builtTree[h][i])
		}
		printf("\n")
	}

}
