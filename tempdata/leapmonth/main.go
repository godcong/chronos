package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/godcong/chronos/v2/runes"
)

var path = "E:\\workspace\\project\\golang\\chronos\\tempdata\\leapmonth\\runyue.txt"
var number = runes.Runes(`一二三四五六七八九十十一十二`)

func main() {
	file, err := ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Println("reading", file)
	sta := 1900
	//stop := 3000
	//idx := 0
	date := make([]byte, (3000-1900+1)*2)
	//for i := sta; i <= stop; i++ {
	for _, s := range file {
		//s, ok := file[i]
		//if ok {
		year, month, sb, err := decodeStr(s)
		if err != nil {
			panic(err)
		}
		idx := (year - sta) * 2
		date[idx] = byte(month)
		date[idx+1] = byte(sb)
		//	if year == i {
		//		date[i-1900] = byte(month)
		//	}
		//	idx++
		//}
	}

	//}
	fmt.Printf("leap year:%x", date)
	err = WriteByteToFile("DataLeapMonth", date)
	if err != nil {
		panic(err)
	}

}

func ReadFile(path string) (map[int]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	r := bufio.NewReader(file)
	ret := make(map[int]string)
	idx := 0
	for {
		line, _, err := r.ReadLine()
		//line, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		//fmt.Println("reading", string(line))
		ret[idx] = string(line)
		idx++
	}
	return ret, nil
}

func decodeStr(str string) (year int, month int, sb int, err error) {
	y := str[0:4]
	fmt.Println("decode year", y)
	year, err = strconv.Atoi(y)
	if err != nil {
		return 0, 0, 0, err
	}
	r := []rune(str[4:])
	month = chineseMonthToInt(r[1])
	//if err != nil {
	//	return 0, 0, false, err
	//}
	sb = 2
	if r[3] == rune('小') {
		sb = 1
	}
	return
}

func chineseMonthToInt(r rune) int {
	if n := number.Index(r); n != -1 {
		return n + 1
	}
	return -1
}

func WriteByteToFile(name string, b []byte) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return err
	}
	_, err = file.Write(b)
	return err
}
