package main

import (
    "github.com/davecgh/go-spew/spew"
    "strconv"
    "strings"
)

func subdomainVisits(ss []string) []string {
    m := map[string]int{}
    for _, v := range ss {
        vs := strings.Split(v, " ")
        count, _ := strconv.Atoi(vs[0])
        m[vs[1]] += count
        for i, ch := range vs[1] {
            if ch == '.' {
                m[vs[1][i+1:]] += count
            }
        }
    }
    res := make([]string, 0, len(m))
    for s, c := range m {
        res = append(res, strconv.Itoa(c) + " " + s)
    }
    return res
}

func main() {
    spew.Dump(subdomainVisits([]string{"9001 discuss.leetcode.com"}))
    spew.Dump(subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}