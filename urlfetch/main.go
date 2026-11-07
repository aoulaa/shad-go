package main

import (
    "fmt"
    "os"
	"io"
	"net/http"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
   

	for _,v := range os.Args[1:] {
		
		 resp, err := http.Get(v)
		 if err != nil {
			fmt.Println("The requested URL was not found.")
			os.Exit(1)
		 }

		 defer resp.Body.Close()
		 body, err := io.ReadAll(resp.Body)
		 check(err)
		 fmt.Println(string(body))

	}
	

}
