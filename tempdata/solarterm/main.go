package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/godcong/chronos/v2/utils"
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

//
//func main() {
//	file, err := ReadFile("1900.json")
//	if err != nil {
//		return
//	}
//	for _, m := range file {
//		for i := 0; i < 24; i++ {
//			istr := strconv.Itoa(i)
//			j, ok := m[istr]
//			if ok {
//				fmt.Printf("%v:\"%v\",\n", istr, j.SanHou)
//			}
//		}
//	}
//
//	fmt.Println(">>>")
//	fmt.Println(">>>")
//	fmt.Println(">>>")
//
//	for _, m := range file {
//		for i := 0; i < 24; i++ {
//			istr := strconv.Itoa(i)
//			j, ok := m[istr]
//			if ok {
//				fmt.Printf("%v:\"%v\",\n", istr, j.Explanation)
//			}
//		}
//	}
//	for _, m := range file {
//		for i := 0; i < 24; i++ {
//			istr := strconv.Itoa(i)
//			j, ok := m[istr]
//			if ok {
//				fmt.Printf("%v:\"%v\",\n", istr, j.Year)
//			}
//		}
//	}
//	fmt.Println("start time:", time.Time{}.UTC().String())
//
//	for _, m := range file {
//		jieqi := make([]byte, 24)
//		for i := 0; i < 24; i++ {
//			istr := strconv.Itoa(i)
//			j, ok := m[istr]
//			if ok {
//				//fmt.Printf("%v:\"%v\",\n", istr, j.Year)
//				parse, err := time.ParseInLocation("2006-01-02 15:04:05", j.Year, time.Local)
//				if err != nil {
//					fmt.Println("error parsing", "time", j.Year, "err", err)
//					continue
//				}
//				s := makeSolarTerTable(jieqi, parse.Month(), parse.Day())
//				//fmt.Printf("%v:\"%v\",\n", istr, parse.Format("2006-01-02 15:04:05"))
//				//fmt.Printf("%v:\"%v\",\n", istr, uint64(parse.Unix()))
//				//fi := strconv.FormatUint(uint64(parse.Unix()), 16)
//				//fmt.Printf("0x%X,\n", uint64(parse.Unix()))
//				fmt.Println("generated:", s)
//				//fmt.Printf("%v:\"%v\",\n", istr, time.Unix(int64(uint32(parse.Unix())), 0).UTC().Format("2006-01-02 15:04:05"))
//			}
//		}
//		fmt.Printf("generated: 0x%x", jieqi)
//	}
//}

var filepath = "jieqi"

func main() {
	fileSta := 1900
	fileEnd := 3000
	for ; fileSta <= fileEnd; fileSta++ {
		bytes, err := ReadBytes(fmt.Sprintf(filepath+"\\%v.json", fileSta))
		if err != nil {
			panic(err)
		}
		fmt.Printf("reading %v, hex:%x\n", fileSta, bytes)
		err = WriteByteToFile("data/DataSolarTerm", yearOffset(fileSta), bytes)
		if err != nil {
			panic(err)
		}
	}
}

func yearOffset(year int) int64 {
	return int64((year - 1900) * 24 * 8)
}

func ReadBytes(name string) ([]byte, error) {
	file, err := ReadFile(name)
	if err != nil {
		return nil, err
	}
	for _, m := range file {
		jieqi := make([]byte, 24*8)
		for i := 0; i < 24; i++ {
			istr := strconv.Itoa(i)
			j, ok := m[istr]
			if ok {
				parse, err := time.ParseInLocation("2006-01-02 15:04:05", j.Year, time.Local)
				if err != nil {
					fmt.Println("error parsing", "time", j.Year, "err", err)
					continue
				}
				jieqiIdx := (i + 2) % 24
				makeSolarTerTable(jieqi, jieqiIdx, parse.Unix())
			}
		}
		return jieqi, nil
	}
	return nil, errors.New("not found")
}

func makeSolarTerTable(b []byte, i int, unix int64) []byte {
	idx := int((i) * 8)
	//if day > 10 {
	//	idx += 8
	//}
	//runes := []rune(strconv.Itoa(day))
	//if len(runes) > 0 {
	//b[idx] = byte(day)
	//idx++
	//t := big.NewInt(unix).Bytes()
	t := utils.Int64ToBytes(unix)
	for i := range t {
		//fmt.Println("write index:", int(idx), len(t), unix)
		b[idx] = t[i]
		idx++
	}
	//}
	//idx++
	//if len(runes) > 1 {
	//	b[idx] = runes[1]
	//}
	return b
}

func WriteByteToFile(name string, offset int64, b []byte) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	_, err = file.Seek(offset, 0)
	if err != nil {
		return err
	}
	if len(b) != 24*8 {
		return errors.New("wrong byte size")
	}
	_, err = file.Write(b)
	return err
}
