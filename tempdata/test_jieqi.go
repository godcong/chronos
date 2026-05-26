package main

import (
	"fmt"
	"github.com/6tail/lunar-go/calendar"
)

func main() {
	solar := calendar.NewSolar(1900, 7, 1, 12, 0, 0)
	lunar := calendar.NewLunarFromSolar(solar)
	table := lunar.GetJieQiTable()
	for k, v := range table {
		fmt.Printf("%s -> %d/%02d/%02d %02d:%02d:%02d\n", k, v.GetYear(), v.GetMonth(), v.GetDay(), v.GetHour(), v.GetMinute(), v.GetSecond())
	}
	fmt.Println("---")
	solar2 := calendar.NewSolar(3000, 7, 1, 12, 0, 0)
	lunar2 := calendar.NewLunarFromSolar(solar2)
	table2 := lunar2.GetJieQiTable()
	for k, v := range table2 {
		fmt.Printf("%s -> %d/%02d/%02d %02d:%02d:%02d\n", k, v.GetYear(), v.GetMonth(), v.GetDay(), v.GetHour(), v.GetMinute(), v.GetSecond())
	}
}
