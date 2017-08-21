package main

import (
    "github.com/jlr52/exer-3-concurrent-hash-map/concurrent_hashmap"
    "sync"
    "runtime"
    "strconv"
    "math/rand"
    "fmt"
)


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
    concurrent_hashmap_instance := concurrent_hashmap.NewConcurrentHashMapStruct(uint32(4))
    concurrent_hashmap_instance.Put("hi", "Hi")

	var wg sync.WaitGroup


    for n := 0; n < 10; n++ {
        wg.Add(1)
    	go insert(concurrent_hashmap_instance, &wg)
    
    }

    wg.Wait()

    fmt.Println(concurrent_hashmap_instance)
}



func insert(concurrent_hashmap_instance *concurrent_hashmap.ConcurrentHashMapStruct, wg *sync.WaitGroup) {
	var rand_int int
    for element_num := 0; element_num < 10; element_num++ {
    	rand_int = rand.Intn(100)
    	concurrent_hashmap_instance.Put(strconv.Itoa(rand_int), rand_int)
    }
    defer wg.Done()
}
