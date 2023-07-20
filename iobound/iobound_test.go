package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
	"testing"
)

func concurrent(){
	db := map[string] [][]string{
		"CrudeOil.csv": nil,
		"GM_players_statistics.csv": nil,
		"Movies_data.csv": nil,
		"Spotify_Song_Attributes.csv": nil,
	}
	var wg sync.WaitGroup
	wg.Add(len(db))
	for file := range db{
		go func(file string){
			defer wg.Done()
			db[file] = ReadCsv(file)
		}(file)
	}
	wg.Wait()
}
func serial(){
	db := map[string] [][]string{
		"CrudeOil.csv": nil,
		"GM_players_statistics.csv": nil,
		"Movies_data.csv": nil,
		"Spotify_Song_Attributes.csv": nil,
	}
    
	for file := range db {
        db[file] = ReadCsv(file)
    }
}
func ReadCsv(filepath string) [][]string {
    f, err := os.Open("../dataset/" +filepath)
    if err != nil {
        fmt.Println(err)
    }
    defer f.Close()    
	csvr := csv.NewReader(f)
    rows, err := csvr.ReadAll()
    if err != nil {
        fmt.Println(err)
    }    
	return rows
}
func BenchmarkReadFileConcurrent(b *testing.B) {
    for i := 0; i < b.N; i++ {
        concurrent()
    }
}
func BenchmarkReadFileSerial(b *testing.B) {
    for i := 0; i < b.N; i++ {
        serial()
    }
}
