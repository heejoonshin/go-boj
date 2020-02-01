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
func (t *trie)insert(key []byte)bool{
	if t.terminal{
		return false
	}
	if len(key) == 0{
		t.terminal = true
		return true
	}else{
		k := "0"
		next := int(key[0])-int(k[0])
		if t.children[next] == nil{
			t.children[next] = initTrie()

		}
		return t.children[next].insert(key[1:])
	}

}
func (t *trie)find(key []byte) bool{

	if t.terminal && len(key) == 0{
		return true
	}
	if t.terminal{
		return false
	}
	k := "0"
	next := int(key[0])-int(k[0])
	return t.children[next].find(key[1:])

}
func main() {
	defer writer.Flush()
	var t int


	scanf("%d\n",&t)

	for T := 0; T < t; T++{
		trie := initTrie()

		scanf("%d\n",&N)
		str := make([]string,N)
		for i := 0; i < N; i++{
			scanf("%s\n",&str[i])
			bstr := []byte(str[i])
			trie.insert(bstr)

		}
		ans := true
		for i :=0; i < N; i++{
			bstr := []byte(str[i])
			if !trie.find(bstr){
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
