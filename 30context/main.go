package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

const (
	minLatency = 10
	maxLatency = 5000
	timeout    = 3000
)

func main() {
	rootCtx := context.Background()
	ctxWithTimeout, cancel := context.WithTimeout(rootCtx, time.Duration(timeout)*time.Millisecond)
	defer cancel()
	res, err := Search(ctxWithTimeout, "gujrat", "Pune")
	if err != nil {
		fmt.Println(" \n Got error !!! ran out of time")
	}
	fmt.Println("Got results :", res)
}

func Search(ctx context.Context, from, to string) ([]string, error) {
	res := make(chan []string)
	go func() {
		res <- slowSearch(from, to)
		close(res)
	}()

	for {
		select {
		case dst := <-res:
			return dst, nil
		case <-ctx.Done():
			return []string{}, ctx.Err()
		}
	}

}

func slowSearch(from, to string) []string {
	// sleep for a random period between 10 and 5000 ms

	// rand.Seed(time.Now().Unix())
	latency := rand.Intn(maxLatency-minLatency+1) - minLatency
	fmt.Printf("started to search for %s-%s takes %dms...", from, to, latency)
	time.Sleep(time.Duration(latency) * time.Millisecond)

	return []string{from + "-" + to + "-jet-airways-11am", from + "-" + to + "-indigo-airlines-12am"}
}
