package cpubound

import (
	"math/rand"
	"sync"
	"testing"
)

func FindSum(list []int) int {
    sum := 0
    for _, number := range list {
        sum += number
    }
    return sum

}
func FindSumConc(list []int) int {
    sum := 0
    var rwm sync.RWMutex    
	var wg sync.WaitGroup
    noOfGoRoutines := 8
    chunkSize := len(list) / noOfGoRoutines
    for i := 0; i < noOfGoRoutines; i++ {
        wg.Add(1)
        go func(i int) {
            defer wg.Done()
            start := i * chunkSize
            end := start + chunkSize
            rwm.Lock()
            for _, number := range list[start:end] {
                sum += number
            }
            rwm.Unlock()
        }(i)
    }
    wg.Wait()    
	return sum
}

func BenchmarkFindSumSerial(b *testing.B){
	list := make([]int, 0)
    for i := 0; i < 10000000; i++ {
        list = append(list, rand.Intn(10000000))
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        FindSum(list)
    }
}
func BenchmarkFindSumConcurrent(b *testing.B){
	list := make([]int, 0)
	for i:= 0; i < 10000000; i++{
		list = append(list, rand.Intn(10000000))
	}
	b.ResetTimer()
	for i:= 0; i < b.N; i++{
		FindSumConc(list)
	}
}
