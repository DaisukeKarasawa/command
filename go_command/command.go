package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	maxAndMinArr := []int{1, 2, 3, 4, 5}
	min := 10000
	max := 0

	for _, v := range maxAndMinArr {
		if min >= v {
			min = v
		}
		if max <= v {
			max = v
		}
	}
	fmt.Println(max, min)

	result := math.Max(10, 20)
	fmt.Printf("%T >> %v\n", result, result)

	tr := -1.5
	fmt.Println(math.Trunc(tr), math.Floor(tr))

	nd := 100
	nf := 3.1415
	s := "golang"
	arr1 := []int{1, 2, 3}
	arr2 := map[int]string{1: "Go", 2: "Ruby"}
	fmt.Printf("%d\n", nd)
	fmt.Printf("%f : %.2f\n", nf, nf)
	fmt.Printf("%s\n", s)
	fmt.Printf("%x >> %X\n", s, s)
	fmt.Printf("%v\n", arr1)
	fmt.Printf("%v\n", arr2)

	str := "golang"
	str2 := "go! golang"
	str3 := "gRPC"
	str4 := "I study golang"
	fmt.Println(strings.Index(str, "go"))
	fmt.Println(strings.Index(str, "Go"))
	fmt.Println(strings.LastIndex(str2, "go"))
	fmt.Println(strings.LastIndex(str2, "Go"))
	fmt.Println(strings.IndexAny(str2, "go"))
	fmt.Println(strings.IndexAny(str2, "Go"))
	fmt.Println(strings.LastIndexAny(str2, "go"))
	fmt.Println(strings.LastIndexAny(str2, "Go"))
	fmt.Println(strings.Contains(str2, "go"))
	fmt.Println(strings.Contains(str2, "Go"))
	fmt.Println(strings.ContainsAny(str2, "go"))
	fmt.Println(strings.ContainsAny(str2, "Go"))
	fmt.Println(strings.HasPrefix(str2, "go"))
	fmt.Println(strings.HasPrefix(str2, "Go"))
	fmt.Println(strings.HasSuffix(str2, "go"))
	fmt.Println(strings.HasSuffix(str2, "Go"))
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str3))
	fmt.Println(strings.ReplaceAll(str4, " ", ""))

	// fmt.Println(int(str))
}
