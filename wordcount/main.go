package main

import (
    "fmt"
    "os"
	"strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
	m := make(map[string]int)
	var s []string
   

	for _,v := range os.Args[1:] {
		
		 a, err := os.ReadFile(v)
		 
		 if len(a) == 0 {
			return
		 }
    	 check(err)
		
		 s = append(s, strings.Split(string(a), "\n")...)
	}
	
	for _ ,v:= range s {
		m[v]++
	}
	

	for line,count := range m {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, line)
		}
	}
}
