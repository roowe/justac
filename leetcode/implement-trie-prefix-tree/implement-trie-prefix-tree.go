package main

import "fmt"

type Trie struct {
    end bool
    next [26]*Trie
}


/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    next := this
    for _, c := range word {
        i := c - 'a'
        if next.next[i] == nil {
            next.next[i] = &Trie{}
        }
        next = next.next[i]
    }
    next.end = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    next := this
    for _, c := range word {
        i := c - 'a'
        //fmt.Println(next)
        if next.next[i] == nil {
            return false
        }
        next = next.next[i]
    }
    return next.end
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(word string) bool {
    next := this
    for _, c := range word {
        i := c - 'a'
        //fmt.Println(next)
        if next.next[i] == nil {
            return false
        }
        next = next.next[i]
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
func main() {
    obj := Constructor();
    obj.Insert("word");
	fmt.Printf("%v\n", obj.Search("word"))
    fmt.Printf("%v\n", obj.Search("wordx"))
    fmt.Printf("%v\n", obj.Search("wor"))
    fmt.Printf("%v\n", obj.Search("xwor"))

    fmt.Printf("%v\n", obj.StartsWith("word"))
    fmt.Printf("%v\n", obj.StartsWith("wordx"))
    fmt.Printf("%v\n", obj.StartsWith("wor"))
    fmt.Printf("%v\n", obj.StartsWith("xwor"))

}
