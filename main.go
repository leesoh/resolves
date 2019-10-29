package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/miekg/dns"
)

func main() {
	domain := flag.String("d", "www.google.com", "Domain to look up")
	maxrtt := flag.Int("t", 100, "Max RTT in ms")
	concurrency := flag.Int("c", 20, "Number of concurrent requests")
	flag.Parse()
	maxtime := time.Duration(*maxrtt) * time.Millisecond
	var wg sync.WaitGroup
	resolvers := make(chan string)
	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for r := range resolvers {
				if LookupHost(*domain, r, maxtime) {
					fmt.Println(r)
				}
			}
		}()
	}
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		res := strings.ToLower(strings.TrimSpace(sc.Text()))
		if res == "" {
			continue
		}
		resolvers <- res
	}
	close(resolvers)
	wg.Wait()
}

func LookupHost(domain, resolver string, maxtime time.Duration) bool {
	c := new(dns.Client)
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), dns.TypeA)
	m.RecursionDesired = true
	result, rtt, err := c.Exchange(m, resolver+":53")
	if err != nil {
		return false
	}
	if len(result.Answer) == 0 {
		return false
	}
	for _, answer := range result.Answer {
		if _, ok := answer.(*dns.A); ok {
			if rtt <= maxtime {
				return true
			}
		}
	}
	return false
}
