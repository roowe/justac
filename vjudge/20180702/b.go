package main

import (
    //"math"
    "io"
    "os"
    "bufio"
    // "io/ioutil"
    // "fmt"
    // "bytes"
)
// func main() {
//     bs, _ := ioutil.ReadAll(os.Stdin)
//     if i := bytes.IndexByte(bs, '0'); i >= 0 {
//         fmt.Printf("%s", append(bs[:i], bs[i+1:]...))
//     } else {
//         fmt.Printf("%s", bs[:len(bs)-3])
//     }   
// }
// func main() {
//     var s string
//     fmt.Scan(&s)
//     var bs = []byte(s)
//     var i0 = len(s)
//     for i := range bs {
//         if s[i] == '0' {
//             i0 = i
//             break
//         }
//     }
//     if i0 == len(s) {
//         fmt.Printf("%s", bs[:i0-1])
//     } else {
//         fmt.Printf("%s", append(bs[:i0], bs[i0+1:]...))
//     }
// }
func main() {    
    r := bufio.NewReader(os.Stdin)
    for {
        prefix, err := r.ReadSlice('0')
        //fmt.Println(err, prefix)
        n := len(prefix)
        if err == nil {
            os.Stdout.Write(prefix[:n-1])
            io.Copy(os.Stdout, r)
            break
        } else if err == io.EOF {
            os.Stdout.Write(prefix[:n-3])
            break
        } else {
            os.Stdout.Write(prefix)            
        }
    }    
}
