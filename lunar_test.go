package lunar

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"
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
	i := GetTermInfo(2018, 3)
	log.Println(i)
}

func TestGetZodiac(t *testing.T) {
	log.Println(GetZodiac(time.Now()))
}

func TestStemBranchYear(t *testing.T) {
	log.Println(StemBranchYear(2017))

}

func TestStemBranchMonth(t *testing.T) {
	log.Println(StemBranchMonth(2017, 11, 14))
}

func TestStemBranchDay(t *testing.T) {
	log.Println(StemBranchDay(2017, 11, 14))

}

func TestStemBranchHour(t *testing.T) {
	log.Print(12, StemBranchDay(2018, 1, 12), "日")
	for i := 0; i <= 23; i++ {

		log.Println(i, StemBranchHour(2018, 1, 12, i))

	}
	log.Print(13, StemBranchHour(2018, 1, 13, 8), "日")
	log.Print(14, StemBranchDay(2017, 11, 14), "日")
	log.Println(8, StemBranchHour(2017, 11, 14, 8))
}
func TestNewLunar(t *testing.T) {
	log.Print(NewLunar(nil).Date())
}
