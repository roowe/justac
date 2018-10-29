package main

import (
    "fmt"
)

type Solution struct {
    beginWord string
    endWord   string
    wordList  []string

    wordMap  map[string]struct{}
    prevPath map[string][]string
    result   [][]string
    dfsPath  []string
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
    p := &Solution{
        beginWord: beginWord,
        endWord:   endWord,
        wordList:  wordList,
    }

    p.wordMap = make(map[string]struct{}, len(wordList))
    for _, w := range wordList {
        p.wordMap[w] = struct{}{}
    }

    if !p.bfs() {
        return nil
    }
    p.dfsPath = make([]string, 0, len(wordList))
    p.dfs(p.endWord)
    return p.result
}

func (p *Solution) dfs(w string) {

    p.dfsPath = append(p.dfsPath, w)
    if len(p.prevPath[w]) == 0 {
        path := make([]string, len(p.dfsPath))
        for i:=0; i<len(path); i++ {
            path[i] = p.dfsPath[len(path)-1-i]
        }
        p.result = append(p.result, path)
    } else {
        for _, nextW := range p.prevPath[w] {
            p.dfs(nextW)
        }
    }
    p.dfsPath = p.dfsPath[:len(p.dfsPath)-1]
}
func (p *Solution) bfs() bool {
    n := len(p.wordList)
    p.prevPath = make(map[string][]string, n)
    for i := range p.prevPath {
        p.prevPath[i] = make([]string, n)
    }
    var rows [2]map[string]struct{}
    for i := range rows {
        rows[i] = make(map[string]struct{}, n)
    }

    cur, prev := 0, 1
    rows[cur][p.beginWord] = struct{}{}
    for {
        cur = cur ^ 1
        prev = prev ^ 1
        for w := range rows[prev] {
            delete(p.wordMap, w)
        }
        rows[cur] = make(map[string]struct{}, n)
        for w := range rows[prev] {
            bs := []byte(w)
            for i := range bs {
                for c := 'a'; c <= 'z'; c++ {
                    if w[i] == byte(c) {
                        continue
                    }
                    bs[i] = byte(c)
                    nw := string(bs)
                    _, ok := p.wordMap[nw]
                    if ok {
                        p.prevPath[nw] = append(p.prevPath[nw], w)
                        rows[cur][nw] = struct{}{}
                    }
                }
                bs[i] = w[i]
            }
        }
        if len(rows[cur]) == 0 {
            return false
        }
        _, ok := rows[cur][p.endWord]
        if ok {
            return true
        }
    }
    return false
}
func main() {
    fmt.Println(findLadders("hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}))
}
