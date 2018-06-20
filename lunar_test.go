package chronos_test

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/godcong/chronos"
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
		i := chronos.GetTermInfo(2018, i)
		log.Println(i)
	}

}

func TestGetZodiac(t *testing.T) {
	log.Println(chronos.GetZodiac(time.Now()))
}

func TestStemBranchYear(t *testing.T) {
	log.Println(chronos.StemBranchYear(2017))

}

func TestStemBranchMonth(t *testing.T) {
	log.Println(chronos.StemBranchMonth(2017, 11, 14))
}

func TestStemBranchDay(t *testing.T) {
	log.Println(chronos.StemBranchDay(2017, 11, 14))

}

func TestStemBranchHour(t *testing.T) {
	log.Print(12, chronos.StemBranchDay(2018, 1, 12), "日")
	for i := 0; i <= 23; i++ {

		log.Println(i, chronos.StemBranchHour(2018, 1, 12, i))

	}
	log.Print(13, chronos.StemBranchHour(2018, 1, 13, 8), "日")
	log.Print(14, chronos.StemBranchDay(2017, 11, 14), "日")
	log.Println(8, chronos.StemBranchHour(2017, 11, 14, 8))
}
func TestNewLunar(t *testing.T) {
	log.Print(chronos.New().Lunar().Date())
}

func TestCalculateLunar(t *testing.T) {
	log.Print("now: ", chronos.Solar2Lunar(time.Now()))
	log.Print(chronos.New().Lunar().EightCharacter())
}
