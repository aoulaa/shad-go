package main

import (
    "fmt"
    "os"
	"io"
	"net/http"
	"time"
	"bytes"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func ReadSoruce(url string, ch chan []string) {
	start := time.Now()
	resp, err := http.Get(url)
		 if err != nil {
			s := []string{
				"The requested URL was not found.",
			}
			ch <- s
		 } else {

		defer resp.Body.Close()
		var responseBody bytes.Buffer
		if resp.Body != nil {
			_, err := io.Copy(&responseBody, resp.Body)
			check(err)
		}

		s := []string{
				fmt.Sprintf("%fs", time.Since(start).Seconds()),
				fmt.Sprintf("%d", responseBody.Len()),
				url,
			}
		ch <- s
		 }
}		


func main() {
   	start := time.Now()
	ch := make(chan []string)

	urls := os.Args[1:]

	for _,v := range urls {
		
		go ReadSoruce(v, ch)
	
	}
	for range urls { 
		fmt.Println(<-ch)
	}

	overallTime := fmt.Sprintf("%fs elapsed", time.Since(start).Seconds())
	
	fmt.Println(overallTime)

}
