package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"
)

type jieqi struct {
	Explanation string `json:"Explanation"`
	SanHou      string `json:"SanHou"`
	Year        string `json:"Year"`
}

//var jieqi24 []map[string]jieqi

func ReadFile(path string) ([]map[string]jieqi, error) {
	filedata, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var jieqi24 []map[string]jieqi
	err = json.Unmarshal(filedata, &jieqi24)
	if err != nil {
		return nil, err
	}
	return jieqi24, nil
}

func main() {
	file, err := ReadFile("1900.json")
	if err != nil {
		return
	}
	for _, m := range file {
		for i := 0; i < 24; i++ {
			istr := strconv.Itoa(i)
			j, ok := m[istr]
			if ok {
				fmt.Printf("%v:\"%v\",\n", istr, j.SanHou)
			}
		}
	}

	fmt.Println(">>>")
	fmt.Println(">>>")
	fmt.Println(">>>")

	for _, m := range file {
		for i := 0; i < 24; i++ {
			istr := strconv.Itoa(i)
			j, ok := m[istr]
			if ok {
				fmt.Printf("%v:\"%v\",\n", istr, j.Explanation)
			}
		}
	}
	for _, m := range file {
		for i := 0; i < 24; i++ {
			istr := strconv.Itoa(i)
			j, ok := m[istr]
			if ok {
				fmt.Printf("%v:\"%v\",\n", istr, j.Year)
			}
		}
	}
	fmt.Println("start time:", time.Time{}.UTC().String())

	for _, m := range file {
		for i := 0; i < 24; i++ {
			istr := strconv.Itoa(i)
			j, ok := m[istr]
			if ok {

				fmt.Printf("%v:\"%v\",\n", istr, j.Year)
				parse, err := time.ParseInLocation("2006-01-02 15:04:05", j.Year, time.UTC)
				if err != nil {
					fmt.Println("error parsing", "time", j.Year, "err", err)
					continue
				}
				fmt.Printf("%v:\"%v\",\n", istr, parse.Format("2006-01-02 15:04:05"))
				fmt.Printf("%v:\"%v\",\n", istr, uint64(parse.Unix()))
				fmt.Printf("%v:\"%v\",\n", istr, time.Unix(int64(uint64(parse.Local().Unix())), 0).UTC().Format("2006-01-02 15:04:05"))
			}
		}
	}

}
