package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var (
	durationFlag = flag.String("d", "1s", "Duration of time between printing line rates, i.e. inverted frequency")
	windowFlag   = flag.Int("w", 1, "Number of durations to window together")
)

func main() {
	flag.Parse()

	duration, err := time.ParseDuration(*durationFlag)
	if err != nil {
		fmt.Printf("Frequency %q invalid: %v\n", *durationFlag, err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	counterChan := make(chan struct{}, 100)
	go func() {
		wg.Add(1)
		defer wg.Done()

		ticker := time.NewTicker(duration)
		defer ticker.Stop()
		windowCount := 0
		lines := 0

		for {
			select {
			case _, ok := <-counterChan:
				if !ok {
					return
				}
				lines++
			case <-ticker.C:
				fmt.Printf("%d line%s/%s\n", lines, plural(lines), duration*time.Duration(*windowFlag))
				windowCount++
				if windowCount >= *windowFlag {
					lines = 0
					windowCount = 0
				}
			}
		}
	}()

	stdin := os.Stdin
	buf := make([]byte, 1024)
	err = nil
	n := 0
	for err == nil {
		for i := 0; i < n; i++ {
			if buf[i] == '\n' {
				counterChan <- struct{}{}
			}
		}
		n, err = stdin.Read(buf)
	}

	close(counterChan)
	wg.Wait()

	if !errors.Is(err, io.EOF) {
		fmt.Println(err)
		os.Exit(1)
	}
}

func plural(n int) string {
	if n == 1 {
		return ""
	}
	return "s"
}
