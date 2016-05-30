
package main  


import (

	"fmt"
	"time"
	// "net/http"
	// "os"
	// "io/ioutil"
)

type World struct{}

func (w *World) String() string {

	return "Corinde Stensland"
}

func main() {

	start := time.Now()

	// .........

	fmt.Printf("Hello, %s\n", "Corinde")


	fmt.Printf("Hello, %s\n", new(World))

	day := time.Now().Weekday()

	fmt.Printf("Hello, %s (%d)\n", day, day)

	// .............
/*
	response, err := http.Get("http://google.com/")

	if err != nil {

		fmt.Printf("%s", err)

		os.Exit(1)

	} else {

		defer response.Body.Close()

		contents, err := ioutil.ReadAll(response.Body)

		if err != nil {

			fmt.Printf("%s", err)

			os.Exit(2)
		}

		// fmt.Printf("%s\n", string(contents))

		_ = contents
	}
*/

	// .............

	fmt.Println(time.Hour + time.Since(start))
	fmt.Printf("%d\n", time.Hour + time.Since(start))
}