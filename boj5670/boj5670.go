package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{}) error{
	_,err:=fmt.Fscanf(reader, f, a...)
	return err
}
const MAXN = 100001
type trie struct{
	children [27]*trie
	terminal bool
	cnt int
}
func initTrie() *trie{
	ret := new(trie)
	ret.terminal = false
	ret.cnt = 0
	for i := 0; i < 26; i++{
		ret.children[i] = nil
	}
	return ret
}
func (t *trie)insert(key *string, idx int){


	if len(*key) == idx{
		t.terminal = true
		return
	}else{

		next := int((*key)[idx]) - int('a')
		if t.children[next] == nil{
			t.children[next] = initTrie()
			t.cnt++
		}
		t.children[next].insert(key,idx+1)

	}

}
func (t *trie)typekey(key *string,idx int)int{

	if len(*key) == idx{
		return 0
	}

	next := int((*key)[idx]) - int('a')
	if(!t.terminal && t.cnt ==1) {
		return t.children[next].typekey(key,idx+1)
	}
	return t.children[next].typekey(key,idx+1) +1
}
func main(){

	defer writer.Flush()
	var n int
	keys := make([]string,MAXN)
	for{
		if scanf("%d\n",&n) == io.EOF{
			return
		}
		root := initTrie()



		for i := 0; i < n; i++{
			scanf("%s\n",&keys[i])
			root.insert(&keys[i],0)
		}
		root.cnt = 2
		s := 0
		for i := 0; i < n; i++{
			s += root.typekey(&keys[i],0)
		}
		printf("%.2f\n",(float32(s)/float32(n)))
		//fmt.Printf("%.2f\n",(float32(s)/float32(n)))
	}

	
}
