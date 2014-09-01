package main

import (
	"sort"
	"fmt"
	"strings"
	"bufio"
	"log"
	"os"
	"strconv"
)

func GrabLines(args string) []string {
	var lines []string
	file, err := os.Open(args)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func ToIntSlice(va []string) []int {
	var intslice []int
	for _, s := range va {
		newint, _ := strconv.Atoi(s)
		intslice = append(intslice, newint) 
	}
	return intslice
}

func ScalarProduct(va []int, vb []int, i int, numints int) {
	total := 0
	sort.Sort(sort.Reverse(sort.IntSlice(va)))
	sort.Sort(sort.IntSlice(vb))
	for i, e := range va { 
		total += e * vb[i]
	}
	fmt.Printf("Case #%d: ", i)
	fmt.Println(total)
}

func main() {
	argsWithoutProgram := os.Args[1]
	lines := GrabLines(argsWithoutProgram)
	cases, _ := strconv.Atoi(lines[0])
	index := 1
	for i := 1; i <= cases; i ++ {
		numints, _ := strconv.Atoi(lines[index])
		va := strings.Split(lines[index+1], " ")
		vb := strings.Split(lines[index+2], " ")
		newva := ToIntSlice(va)
		newvb := ToIntSlice(vb)
		ScalarProduct(newva, newvb, i, numints)
		index += 3
	}
}
