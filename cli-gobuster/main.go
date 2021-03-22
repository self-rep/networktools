/*
Credits: self-rep
While Extremely Terrible Code it is the first time i have attempted Any Networking / Pentesting Tools
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

var (
	timeout  time.Duration
	debug    = true
	link_ext []string
)

func main() {
	// Get Flags
	wordlist_path := flag.String("wordlist", "none", "Extensions Wordlist")
	addr := flag.String("host", "none", "URL/IP")
	addr_port := flag.Int("port", 80, "Port")
	show_output := flag.Bool("output", false, "Show Output With Status Codes!")
	timeout := flag.String("timeout", "1000ms", "Timeout In Milliseconds! E.G 1000ms or 1000s")
	flag.Parse()
	// Load Word List
	wordlist_file, err := os.Open(*wordlist_path)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer wordlist_file.Close()
	scanner := bufio.NewScanner(wordlist_file)
	for scanner.Scan() {
		if *addr_port == 80 {
			if *show_output == true {
				url := fmt.Sprintf("%s/%s", *addr, scanner.Text())
				req, err := http.Get(url)
				if err != nil {
					fmt.Println(err.Error())
				}
				defer req.Body.Close()
				code := req.StatusCode
				response := fmt.Sprintf("%s/%s : %d", *addr, scanner.Text(), code)
				fmt.Println(response)
				dur, err := time.ParseDuration(*timeout)
				if err != nil {
					fmt.Println(err)
					return
				}
				time.Sleep(dur)
			}

		} else {
			if *show_output == true {
				url := fmt.Sprintf("%s:%d/%s", *addr, *addr_port, scanner.Text())
				req, err := http.Get(url)
				if err != nil {
					fmt.Println(err.Error())
				}
				defer req.Body.Close()
				code := req.StatusCode
				response := fmt.Sprintf("%s:%d/%s : %d", *addr, *addr_port, scanner.Text(), code)
				fmt.Println(response)
				dur, err := time.ParseDuration(*timeout)
				if err != nil {
					fmt.Println(err)
					return
				}
				time.Sleep(dur)
			}
		}

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

}
