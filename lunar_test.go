package lunar_test

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/godcong/lunar"
)

func TestGetSolarTerm(t *testing.T) {
	//GetSolarTerm(time.Now(), 0)
}

func TestYearDays(t *testing.T) {
	//log.Println(YearDays(1900))
}

func TestGetDayString(t *testing.T) {
	//log.Println(GetDayString(20))
	//reader := transform.NewReader(strings.NewReader(, simplifiedchinese.GB18030.NewDecoder())
	//all, _ := ioutil.ReadAll(reader)
	fmt.Println(strconv.Itoa(0x97783), strconv.Itoa(0x97bd0), strconv.Itoa(0x97c36), strconv.Itoa(0xb0b6f), strconv.Itoa(0xc9274), strconv.Itoa(0xc91aa))
}

func TestGetTerm(t *testing.T) {
	for i := 1; i <= 24; i++ {
		i := lunar.GetTermInfo(2018, i)
		log.Println(i)
	}

}

func TestGetZodiac(t *testing.T) {
	log.Println(lunar.GetZodiac(time.Now()))
}

func TestStemBranchYear(t *testing.T) {
	log.Println(lunar.StemBranchYear(2017))

}

func TestStemBranchMonth(t *testing.T) {
	log.Println(lunar.StemBranchMonth(2017, 11, 14))
}

func TestStemBranchDay(t *testing.T) {
	log.Println(lunar.StemBranchDay(2017, 11, 14))

}

func TestStemBranchHour(t *testing.T) {
	log.Print(12, lunar.StemBranchDay(2018, 1, 12), "日")
	for i := 0; i <= 23; i++ {

		log.Println(i, lunar.StemBranchHour(2018, 1, 12, i))

	}
	log.Print(13, lunar.StemBranchHour(2018, 1, 13, 8), "日")
	log.Print(14, lunar.StemBranchDay(2017, 11, 14), "日")
	log.Println(8, lunar.StemBranchHour(2017, 11, 14, 8))
}
func TestNewLunar(t *testing.T) {
	log.Print(lunar.NewLunar(nil).Date())
}

func TestCalculateLunar(t *testing.T) {
	log.Print("now: ", lunar.Solar2Lunar(time.Now()))
	log.Print(lunar.NewLunar(nil).Date())

}
