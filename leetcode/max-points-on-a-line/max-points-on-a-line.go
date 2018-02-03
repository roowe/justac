package main

import "fmt"

type Point struct {
    X int
    Y int
}

type K struct {
    X int
    Y int
    Kx int
    Ky int
}
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
func maxPoints(p []Point) int {
    n := len(p)
    if (n == 0) {
        return 0
    }
    max := 1
    for i:=0; i<n; i++ {
        h := make(map[K]int)
        dumpSize := 0
        tmpMax := 1
        for j:=i+1; j<n; j++ {
            if p[i] == p[j] {
                dumpSize++
            } else {
                kx := p[i].X-p[j].X
                ky := p[i].Y-p[j].Y
                c := gcd(kx, ky)
                kx /= c
                ky /= c
                k := K{p[i].X, p[i].Y, kx, ky}                
                //fmt.Printf("%v\n", k)
                s := 2
                if oldSize, ok := h[k]; ok {
                    s = oldSize + 1                                      
                }
                h[k] = s
                if (s > tmpMax) {
                    tmpMax = s
                }
            }
        }
        //fmt.Println(h)
        if (tmpMax+dumpSize > max) {
            max = tmpMax+dumpSize
        }
    }
    return max
}

func main() {
	data := []Point{
        Point{1, 2},
        Point{3, 6},
        Point{0, 0},
        Point{1, 3},
    }
	//ret := maxPoints(data)
    //fmt.Printf("%v\n", ret)

    data = []Point{
        Point{0, 0},
        Point{94911151, 94911150},
        Point{94911152, 94911151},
    }
	ret := maxPoints(data)    
	fmt.Printf("%v\n", ret)
}
