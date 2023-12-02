package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	for {
		response, err := http.Get("http://localhost:8080")

		if err != nil {
			fmt.Println("We cannot get the weather currently because the server is down.")
			fmt.Fprintln(os.Stderr, "Error sending request:", err)
			os.Exit(1)
		}

		defer response.Body.Close()

		if response.StatusCode == http.StatusOK {
			body, err := io.ReadAll(response.Body)

			if err != nil {
				fmt.Println("We cannot get the weather currently because the server returned an invalid response.")
				fmt.Fprintln(os.Stderr, "Error reading response:", err)
				os.Exit(2)
			}

			fmt.Println(string(body))
			os.Exit(0)
		} else if response.StatusCode == http.StatusTooManyRequests {
			if retryAfter := response.Header.Get("Retry-After"); retryAfter != "" {
				if retryAfterSeconds, err := strconv.Atoi(retryAfter); err == nil {
					delay(time.Duration(retryAfterSeconds) * time.Second)
				} else if retryAfterTime, err := http.ParseTime(retryAfter); err == nil {
					delay(retryAfterTime.Sub(time.Now().UTC()))
				} else {
					fmt.Fprintln(os.Stderr, "Error parsing Retry-After header ", retryAfter, ", we'll wait 200 ms.")
					delay(2 * time.Millisecond)
				}
			} else {
				fmt.Fprintln(os.Stderr, "No Retry-After header provided, we'll wait 200 ms.")
				delay(2 * time.Millisecond)
			}
			continue
		} else {
			fmt.Println("We cannot get the weather currently because the server returned unexpected status code.")
			fmt.Fprintln(os.Stderr, "Got unexpected status code:", response.StatusCode)
			os.Exit(3)
		}
	}
}

func delay(d time.Duration) {
	if d > 5*time.Second {
		fmt.Println("We cannot get the weather currently because the server is too busy.")
		fmt.Fprintln(os.Stderr, "Delay is ", d)
		os.Exit(4)
	} else if d > time.Second {
		fmt.Println("It may take a bit longer to get the weather. Thank you for your patience.")
	}

	time.Sleep(d)
}
