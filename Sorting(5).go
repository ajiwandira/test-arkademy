package main

import(
	"fmt"
	"strings"
	"sort"
	"strconv"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
    return len(s)
}

func SortString(s string) string {
    r := []rune(s)
    sort.Sort(sortRunes(r))
    return string(r)
}

func main(){
	numbers := 43206523032
    
	strNum := strconv.Itoa(numbers)
	split := strings.Split(strNum, "0")
	split1 := SortString(split[0])
	split2 := SortString(split[1])
	split3 := SortString(split[2])
	
	resp := split1 + split2 + split3
	reps, _ := strconv.Atoi(resp)
	fmt.Println(reps)
}