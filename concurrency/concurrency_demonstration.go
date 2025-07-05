package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	args := os.Args
	if len(args) < 4 {
		fmt.Println("Usage: go run concurrency_demonstration.go <number_of_records> <number_of_workers> <strategy>")
		fmt.Println("Example: go run concurrency_demonstration.go 1000 4 naive")
		return
	}
	numberOfRecords, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Error parsing number of records:", err)
		return
	}
	numberOfWorkers, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Println("Error parsing number of workers:", err)
		return
	}
	fmt.Println("Number of CPUs:", numberOfWorkers)
	records := make([]int, 0, numberOfRecords)
	for i := range numberOfRecords {
		records = append(records, i)
	}

	strategy := args[3]
	switch strategy {
	case "naive":
		startTime := time.Now()
		fmt.Println(NaiveWork(records))
		fmt.Println("NaiveWork duration:", time.Since(startTime))
	case "mutex":
		fmt.Println(SplitWorkUsingMutex(records, numberOfWorkers))
	case "channels":
		fmt.Println(SplitWorkUsingChannels(records, numberOfWorkers))
	default:
		fmt.Println("Unknown strategy:", strategy)
	}
}

func NaiveWork(records []int) int {
	total := 0
	for _, record := range records {
		total += record
	}
	return total
}

func SplitWorkUsingMutex(records []int, cpus int) int {
	startTime := time.Now()

	total := 0
	var mutex sync.Mutex
	var wg sync.WaitGroup
	workLoad := len(records) / cpus
	for i := range cpus {
		wg.Add(1)
		work := func(start, end int) {
			defer wg.Done()
			localTotal := 0
			for j := start; j < end; j++ {
				localTotal += records[j]
			}
			mutex.Lock()
			total += localTotal
			mutex.Unlock()
		}
		if i == cpus-1 {
			work(i*workLoad, len(records))
		} else {
			work(i*workLoad, (i+1)*workLoad)
		}
	}
	wg.Wait()
	fmt.Println("SplitWorkUsingMutex duration:", time.Since(startTime))
	return total
}

func SplitWorkUsingChannels(records []int, cpus int) int {
	startTime := time.Now()

	total := 0
	workLoad := len(records) / cpus
	results := make(chan int, cpus)
	var wg sync.WaitGroup

	for i := range cpus {
		wg.Add(1)
		work := func(start, end int) {
			defer wg.Done()
			localTotal := 0
			for j := start; j < end; j++ {
				localTotal += records[j]
			}
			results <- localTotal
		}
		if i == cpus-1 {
			work(i*workLoad, len(records))
		} else {
			work(i*workLoad, (i+1)*workLoad)
		}
	}
	wg.Wait()
	close(results)

	for result := range results {
		total += result
	}

	fmt.Println("SplitWorkUsingChannels duration:", time.Since(startTime))
	return total
}
