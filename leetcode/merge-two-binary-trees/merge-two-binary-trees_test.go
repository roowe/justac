package ac

import "fmt"
import "testing"


func FUN(x int) int {

}

func TestM(t *testing.T) {
    var tests = []struct{
        x int
        want int
    }{
        
    }
    for _, test := range tests {
        desc := fmt.Sprintf("FUN(%d)", test.x)
        got := reverse(test.x)
        if got != test.want {
            t.Errorf("%s = %d, want %d", desc, got, test.want)
        }
    }   
}
