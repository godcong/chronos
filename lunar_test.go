package chronos_test

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/godcong/chronos"
)

// TestGetSolarTerm ...
func TestGetSolarTerm(t *testing.T) {

}

// TestYearDays ...
func TestYearDays(t *testing.T) {

}

// TestGetDayString ...
func TestGetDayString(t *testing.T) {
	//log.Println(GetDayString(20))
	//reader := transform.NewReader(strings.NewReader(, simplifiedchinese.GB18030.NewDecoder())
	//all, _ := ioutil.ReadAll(reader)
	fmt.Println(strconv.Itoa(0x97783), strconv.Itoa(0x97bd0), strconv.Itoa(0x97c36), strconv.Itoa(0xb0b6f), strconv.Itoa(0xc9274), strconv.Itoa(0xc91aa))
}

// TestGetTerm ...
func TestGetTerm(t *testing.T) {
	for i := 1; i <= 24; i++ {
		i := chronos.GetTermInfo(2018, i)
		log.Println(i)
	}

}

// TestGetZodiac ...
func TestGetZodiac(t *testing.T) {
	t.Log(chronos.GetZodiac(chronos.New("2020/01/24 18:40").Lunar()) == "猪")
	t.Log(chronos.GetZodiac(chronos.New("2020/01/25 18:40").Lunar()) == "鼠")
	t.Log(chronos.GetZodiac(chronos.New("2021/02/11 18:40").Lunar()) == "鼠")
	t.Log(chronos.GetZodiac(chronos.New("2021/02/12 18:40").Lunar()) == "牛")
}

// TestStemBranchYear ...
func TestStemBranchYear(t *testing.T) {
	log.Println(chronos.StemBranchYear(2017))

}

// TestStemBranchMonth ...
func TestStemBranchMonth(t *testing.T) {
	log.Println(chronos.StemBranchMonth(2017, 11, 14))
}

// TestStemBranchDay ...
func TestStemBranchDay(t *testing.T) {
	log.Println(chronos.StemBranchDay(2017, 11, 14))

}

// TestStemBranchHour ...
func TestStemBranchHour(t *testing.T) {
	log.Print(12, chronos.StemBranchDay(2018, 1, 12), "日")
	for i := 0; i <= 23; i++ {

		log.Println(i, chronos.StemBranchHour(2019, 11, 2, i))

	}
	log.Print(13, chronos.StemBranchHour(2018, 1, 13, 8), "日")
	log.Print(14, chronos.StemBranchDay(2017, 11, 14), "日")
	log.Println(8, chronos.StemBranchHour(2017, 11, 14, 8))
}

// TestNewLunar ...
func TestNewLunar(t *testing.T) {
	log.Print(chronos.New().Lunar().Date())
}

// TestCalculateLunar ...
func TestCalculateLunar(t *testing.T) {
	//log.Print("now: ", chronos.Solar2Lunar(time.Parse()))
	//检查日柱是否正确
	log.Print(chronos.New("2019/10/31 23:13").Lunar().EightCharacter())
	log.Print(chronos.New("2019/11/01 0:13").Lunar().EightCharacter())
	log.Print(chronos.New("2019/11/01 1:13").Lunar().EightCharacter())

	//检查立春是否有效
	log.Print(chronos.New("2020/02/03 23:13").Lunar().EightCharacter())
	log.Print(chronos.New("2020/02/04 1:13").Lunar().EightCharacter())
	log.Print(chronos.New("2020/02/04 18:13").Lunar().EightCharacter())

	log.Print(chronos.New("2020/01/24 0:40").LunarDate())

	log.Print(chronos.New("2020/01/25 0:40").LunarDate())
	log.Print(chronos.New(time.Now()).LunarDate())

	log.Print(chronos.New("1989/01/07 18:40").Lunar().EightCharacter())

	//output:戊辰年十一月三十日
	log.Print(chronos.New("1989/01/07 18:40").LunarDate())
	log.Print(chronos.New("1989/01/07 0:40").LunarDate())

	log.Print(chronos.New("2019/06/01 0:40").LunarDate())

}

func TestLunar_EightCharacter(t *testing.T) {
	//output 己 亥 甲 戌 庚 子 丁 丑
	log.Println(chronos.New("2019/10/30 01:30").Lunar().EightCharacter())
	//output 己 亥 甲 戌 庚 子 丙 子
	log.Println(chronos.New("2019/10/30 23:30").Lunar().EightCharacter())

}
