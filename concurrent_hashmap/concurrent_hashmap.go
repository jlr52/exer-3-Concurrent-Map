package concurrent_hashmap

import (
	"hash/fnv"
    "sync"
)

type ConcurrentHashMapStruct struct {
	hashMapBuckets []ThreadSafeHashMapStruct
	bucketCount uint32
}

type ThreadSafeHashMapStruct struct {
	mapLock *sync.RWMutex
	hashMap map[string]interface{}
}

// Constructor
func NewConcurrentHashMapStruct(bucketSize uint32) *ConcurrentHashMapStruct {
	var threadSafeHashMapStructInstance ThreadSafeHashMapStruct
	var hashMapBucketsInstance []ThreadSafeHashMapStruct

    for i := 0; i <= int(bucketSize); i++ {
    	threadSafeHashMapStructInstance = ThreadSafeHashMapStruct{&sync.RWMutex{}, make(map[string]interface{})}
    	hashMapBucketsInstance = append(hashMapBucketsInstance, threadSafeHashMapStructInstance)
    }
	return &ConcurrentHashMapStruct{hashMapBucketsInstance, bucketSize}
}


// Public API
func (cMap *ConcurrentHashMapStruct) Put(key string, val interface{}) {
	bucketIndex := hash(key) % cMap.bucketCount
	bucket := cMap.hashMapBuckets[bucketIndex]
	bucket.mapLock.Lock()
	bucket.hashMap[key] = val
	defer bucket.mapLock.Unlock()
}

func (cMap *ConcurrentHashMapStruct) Get(key string) interface{} {
	bucketIndex := hash(key) % cMap.bucketCount
	bucket := cMap.hashMapBuckets[bucketIndex]
	bucket.mapLock.RLock()
	returnVal := bucket.hashMap[key]
	bucket.mapLock.RUnlock()
	return returnVal
}

// Helper
func hash(s string) uint32 {
        h := fnv.New32a()
        h.Write([]byte(s))
        return h.Sum32()
}
