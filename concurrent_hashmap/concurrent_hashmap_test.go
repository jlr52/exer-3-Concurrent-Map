package concurrent_hashmap

import (
    "testing"
    "runtime"
    "math/rand"
    "strconv"
    "sync"
)


func BenchmarkConcurrentMapPut(b *testing.B) {
	var wg sync.WaitGroup

    runtime.GOMAXPROCS(runtime.NumCPU())

    concurrent_hashmap_instance := NewConcurrentHashMapStruct(uint32(4))
    for n := 0; n < b.N; n++ {
    	wg.Add(1)
    	go insert(concurrent_hashmap_instance, &wg)
    
    }
    wg.Wait()
}

func insert(concurrent_hashmap_instance *ConcurrentHashMapStruct, wg *sync.WaitGroup) {
	var rand_int int
    for element_num := 0; element_num < 10; element_num++ {
    	rand_int = rand.Intn(100)
    	concurrent_hashmap_instance.Put(strconv.Itoa(rand_int), rand_int)
    	//concurrent_hashmap_instance.Get(strconv.Itoa(rand_int))
    }
    defer wg.Done()
}

func BenchmarkMapSynchronizePut(b *testing.B) {
    my_map := make(map[string]interface{})
    var rand_int int
    for n := 0; n < b.N; n++ {
    	rand_int = rand.Intn(100)
        my_map[strconv.Itoa(rand_int)] = rand_int
    }
}