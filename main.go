// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// func main() {
// 	c := make(chan string)
// 	links := []string{
// 		"http://www.amazon.com",
// 		"http://www.google.com",
// 		"http://www.apple.com",
// 		"http://www.microsoft.com",
// 		"http://www.facebook.com",
// 		"http://www.twitter.com",
// 		"http://www.netflix.com",
// 		"http://www.walmart.com",
// 		"http://www.stackoverflow.com",
// 	}
// 	for _, url := range links {
// 		go checkActiveLink(url, c)
// 	}
// 	for l := range c {
// 		go func(link string) {
// 			time.Sleep(5 * time.Second)
// 			checkActiveLink(link, c)
// 		}(l)
// 	}
// }

// func checkActiveLink(s string, c chan string) {

// 	_, err := http.Get(s)
// 	if err != nil {
// 		fmt.Println(s, "server is down")
// 		c <- s
// 		return
// 	}
// 	fmt.Println(s, "server is up")
// 	c <- s

// }

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{"https://amazon.com", "https://stackoverflow.com", "https://google.com", "https://facebook.com"}

	//create channer with string type

	channel := make(chan string)

	//loop through the urls
	for _, url := range urls {
		// using go key word we initialize new routine
		go makeApiCall(url, channel)
	}

	//we loop the channel response
	for channelResp := range channel {

		// we are calling the make api call infinely with 5 sec time gap
		go func(url string) {

			//make 5 sec sleep
			time.Sleep(5 * time.Second)
			makeApiCall(url, channel)
		}(channelResp)
	}
}

func makeApiCall(url string, channel chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("It seems server down")
		return
	}

	fmt.Println("Yes it's up...", url)
	channel <- url

}
