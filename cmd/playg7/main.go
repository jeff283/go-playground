package main

import (
	"fmt"
	"slices"
	"sync"
	"time"
)

// var mx = sync.Mutex{}
var mx = sync.RWMutex{}
var wg = sync.WaitGroup{}

var dbData = []string{"id1", "id2", "id3", "id4", "id5", "id6"}
var results = []string{}

func main() {
	t0 := time.Now()

	for i := range dbData {
		wg.Add(1)
		go db(i)
	}
	wg.Wait()

	fmt.Println()
	fmt.Println("Results: ", results)
	slices.Sort(results)
	fmt.Println("Sorted Results:", results)

	fmt.Println("Total Execution Time: ", time.Since(t0))
}

func db(i int) /*string*/ {
	var delay float32 = 2000 /* *rand.Float32() */
	time.Sleep(time.Duration(delay) * time.Millisecond)

	res := dbData[i]

	fmt.Println("The db results is: ", res)

	save(res)
	log()

	wg.Done()

	// return res
}

func save(res string) {
	mx.Lock()
	results = append(results, res)
	mx.Unlock()
}

func log() {
	mx.RLock()
	fmt.Println("The current results are: ", results)
	mx.RUnlock()
}
