package ac

import "fmt"
import "testing"


func judgeCircle(moves string) bool {
    c := make(map[rune]int, 4)
    for _, s := range(moves) {
        c[s]++
    }
    return c['R']-c['L']+c['U']-c['D'] == 0 
}

func TestM(t *testing.T) {
    var tests = []struct{
        x string
        want bool
    }{
        {"UD", true},
        {"LL", false},
    }
    for _, test := range tests {
        desc := fmt.Sprintf("judgeCircle(%s)", test.x)
        got := judgeCircle(test.x)
        if got != test.want {
            t.Errorf("%s = %s, want %s", desc, got, test.want)
        }
    }   
}
