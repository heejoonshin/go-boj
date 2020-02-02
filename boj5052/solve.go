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
var N int

type trie struct{
	children [10]*trie
	terminal bool
}
func initTrie() *trie{
	ret := new(trie)
	ret.terminal = false
	for i := 0; i < 10; i++{
		ret.children[i] = nil
	}
	return ret
}
func (t *trie)insert(key *[]byte,idx int)bool{
	if t.terminal{
		return false
	}
	if len(*key) == idx{
		t.terminal = true
		return true
	}else{

		next := int((*key)[idx])-int('0')
		if t.children[next] == nil{
			t.children[next] = initTrie()

		}
		return t.children[next].insert(key,idx+1)
	}

}
func (t *trie)find(key *[]byte,idx int) bool{

	if t.terminal && len(*key) == idx{
		return true
	}
	if t.terminal{
		return false
	}

	next := int((*key)[idx])-int('0')
	return t.children[next].find(key,idx+1)

}
func main() {
	defer writer.Flush()
	var t int


	scanf("%d\n",&t)

	for T := 0; T < t; T++{
		trie := initTrie()

		scanf("%d\n",&N)
		str := make([][]byte,N)
		for i := 0; i < N; i++{
			scanf("%s\n",&str[i])

			trie.insert(&str[i],0)

		}
		ans := true
		for i :=0; i < N; i++{

			if !trie.find(&str[i],0){
				ans = false
			}

		}
		if ans{
			printf("YES\n")
		}else{
			printf("NO\n")
		}


	}


	
}
