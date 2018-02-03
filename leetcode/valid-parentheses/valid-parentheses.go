package main

import "fmt"
import "container/list"


var CloseP map[byte]byte = map[byte]byte{')':'(' , ']':'[', '}':'{'}
var OpenP map[byte]byte = map[byte]byte{'(':')' , '[':']', '{':'}'}

func isValid2(s string) bool {
    buf := make([]byte,0, len(s))
    for i:=0;i<len(s);i++{
        c := s[i]
        if _,ok := OpenP[c]; ok{
            buf =append(buf,c)
            continue
        }
        openOne:=CloseP[c]
        bufLen := len(buf)
        if len(buf)<=0 || buf[bufLen-1] != openOne{
            return false
        }
        buf = buf[:bufLen-1]
    }
    return len(buf)==0
}

func isValid(s string) bool {
    l := list.New()
    pairs := map[rune]rune{
        ')':'(',
        '}':'{',
        ']':'[',        
    }
    for _, c := range(s) {
        if c == '(' ||
            c == '{' ||
            c == '['  {
            l.PushBack(c)
        } else if c == ')' ||
            c == '}' ||
            c == ']'  {
            e := l.Back()
            if e != nil && e.Value == pairs[c] {
                l.Remove(e)
            } else {
                return false
            }                        
        }
    }
    return l.Len() == 0
}

func main() {
    // fmt.Printf("%v\n", parse(342))
    // fmt.Printf("%v\n", parse(465))
    fmt.Printf("%v\n", isValid("["))
    fmt.Printf("%v\n", isValid("]"))
    fmt.Printf("%v\n", isValid("[{()}][]"))
    fmt.Printf("%v\n", isValid("[{(}}][]"))
}
