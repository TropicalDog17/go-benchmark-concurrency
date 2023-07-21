package cpubound

import (
	"math/rand"
	"sync"
	"testing"
)

var maxNumber = 100000000

func FindSum(list []int) int {
	sum := 0
	for _, number := range list {
		sum += number
	}
	return sum

}
func FindSumConc(list []int) int {
	sum := []int{0, 0, 0, 0, 0, 0, 0, 0}
	total := 0
	var wg sync.WaitGroup
	noOfGoRoutines := 8
	chunkSize := len(list) / noOfGoRoutines
	for i := 0; i < noOfGoRoutines; i++ {
		wg.Add(1)
		go func(i int) {
			localSum := 0
			defer wg.Done()
			start := i * chunkSize
			end := start + chunkSize
			for _, number := range list[start:end] {
				localSum += number
			}
			sum[i] = localSum
		}(i)
	}
	wg.Wait()
	for _, number := range sum {
		total += number
	}
	return total
}

func BenchmarkFindSumSerial(b *testing.B) {
	list := make([]int, 0)
	for i := 0; i < maxNumber; i++ {
		list = append(list, rand.Intn(maxNumber))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindSum(list)
	}
}
func BenchmarkFindSumConcurrent(b *testing.B) {
	list := make([]int, 0)
	for i := 0; i < maxNumber; i++ {
		list = append(list, rand.Intn(maxNumber))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FindSumConc(list)
	}
}

func TestSerialAndConcurrent(t *testing.T) {
	list := make([]int, 0)
	for i := 0; i < maxNumber; i++ {
		list = append(list, rand.Intn(maxNumber))
	}
	serialSum := FindSum(list)
	concurrentSum := FindSumConc(list)
	if serialSum != concurrentSum {
		t.Error("Sum not equal")
	}
}
