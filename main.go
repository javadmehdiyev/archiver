package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const banner = `
    █████╗ ██████╗  ██████╗██╗  ██╗██╗██╗   ██╗███████╗██████╗ 
   ██╔══██╗██╔══██╗██╔════╝██║  ██║██║██║   ██║██╔════╝██╔══██╗
   ███████║██████╔╝██║     ███████║██║██║   ██║█████╗  ██████╔╝
   ██╔══██║██╔══██╗██║     ██╔══██║██║╚██╗ ██╔╝██╔══╝  ██╔══██╗
   ██║  ██║██║  ██║╚██████╗██║  ██║██║ ╚████╔╝ ███████╗██║  ██║
   ╚═╝  ╚═╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝╚═╝  ╚═══╝  ╚══════╝╚═╝  ╚═╝
                                                                 
   Web Archive Scanner v1.0
   Author: MIRI
`

func main() {
	workers := flag.Int("w", 10, "Number of concurrent workers")
	pattern := flag.String("p", "*.example.com/*", "URL pattern to search for")
	output := flag.String("o", "", "Output file (optional)")
	flag.Parse()

	fmt.Print(banner)
	fmt.Printf("[*] Starting scan with %d workers\n", *workers)
	fmt.Printf("[*] Pattern: %s\n", *pattern)
	fmt.Printf("[*] Time: %s\n\n", time.Now().Format("2006-01-02 15:04:05"))

	var out *os.File
	var err error
	if *output != "" {
		out, err = os.Create(*output)
		if err != nil {
			log.Fatal("Error creating output file:", err)
		}
		defer out.Close()
	}

	response, err := http.Get(fmt.Sprintf("https://web.archive.org/cdx/search/cdx?url=%s&output=txt&fl=original&collapse=urlkey", *pattern))
	if err != nil {
		log.Fatal("Error fetching from web archive:", err)
	}
	body, err := io.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Fatal("Error reading response:", err)
	}

	urls := strings.Split(string(body), "\n")
	fmt.Printf("[*] Found %d URLs to check\n\n", len(urls))

	checkIfAlive(urls, *workers, out)
}

func checkIfAlive(urls []string, workers int, out *os.File) {
	var wg sync.WaitGroup
	urlChan := make(chan string)
	results := make(chan string)

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range urlChan {
				response, err := http.Get(url)
				if err != nil {
					continue
				}
				if response.StatusCode != 404 && response.StatusCode != 500 && response.StatusCode != 403 && response.StatusCode != 400 {
					results <- fmt.Sprintf("[+] %s --> Alive (Status: %d)", url, response.StatusCode)
				}
				response.Body.Close()
			}
		}()
	}

	go func() {
		for result := range results {
			fmt.Println(result)
			if out != nil {
				fmt.Fprintln(out, result)
			}
		}
	}()

	for _, url := range urls {
		if url != "" {
			urlChan <- url
		}
	}
	close(urlChan)
	wg.Wait()
	close(results)
}
