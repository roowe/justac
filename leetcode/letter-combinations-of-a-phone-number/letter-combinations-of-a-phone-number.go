package main

import "fmt"
func letterCombinations(digits string) []string {
	mapNumber := make(map[string][]string)
	mapNumber["0"] = []string{""}
	mapNumber["1"] = []string{""}
	mapNumber["2"] = []string{"a", "b", "c"}
	mapNumber["3"] = []string{"d", "e", "f"}
	mapNumber["4"] = []string{"g", "h", "i"}
	mapNumber["5"] = []string{"j", "k", "l"}
	mapNumber["6"] = []string{"m", "n", "o"}
	mapNumber["7"] = []string{"p", "q", "r", "s"}
	mapNumber["8"] = []string{"t", "u", "v"}
	mapNumber["9"] = []string{"w", "x", "y", "z"}

	if len(digits) == 0 {
		return []string{}
	}

	result := []string{""}

	for i := 0; i < len(digits); i++ {
		tempResult := []string{}

		char := mapNumber[string(digits[i])]
		for j := 0; j < len(char); j++ {
			for k := 0; k < len(result); k++ {
				tempResult = append(tempResult, result[k]+char[j])
			}
		}
		result = tempResult
	}
	return result
}
func letterCombinations2(digits string) []string {
    n := len(digits)
    if n == 0 {
        return []string{}
    }
    for i:=0; i<n; i++ {
        var begin rune
        var end rune
        if (digits[i] >= '2' && digits[i] <= '6') {
            diff := rune(3*(digits[i] - '2'))
            begin = diff + 'a'
            end = diff + 'c'
        } else if digits[i] == '7'{
            begin = 'p'
            end = 's'
        } else if digits[i] == '8'{
            begin = 't'
            end = 'v'
        } else if digits[i] == '9'{
            begin = 'w'
            end = 'z'
        }
        ret := []string{}
        if i+1 == n {

            for j:=begin; j<=end; j++ {
                ret = append(ret, string(j))
            }
            return ret
        }        
        for j:=begin; j<=end; j++ {            
            ret0 := letterCombinations(digits[(i+1):])
            for _, s := range ret0 {
                ret = append(ret, string(j)+s)
            }            
        }
        return ret
    }
    return []string{}
}
func main() {
	fmt.Printf("%v\n", letterCombinations("23"))
    fmt.Printf("%v\n", letterCombinations("7"))
}
