// Package chronos provides Chinese calendar calculations including solar/lunar date
// conversion, Heavenly Stems and Earthly Branches (干支), Solar Terms (节气),
// Zodiac (生肖), and Constellation (星座) lookups.
//
// The package wraps github.com/6tail/lunar-go for astronomical calculations
// and provides type-safe Go enums for all Chinese calendar concepts.
//
// # Quick Start
//
// Parse a solar date and access calendar information:
//
//	cal := chronos.ParseSolarTime(time.Now())
//	fmt.Println(cal.Lunar().Zodiac().Chinese())
//	fmt.Println(cal.Lunar().EightChar().FourPillars())
//
// # Sub-packages
//
// The fate sub-package provides BaZi (八字) analysis and Five Element (五行)
// calculations for name selection:
//
//	import "github.com/godcong/chronos/v2/fate"
//
//	data, err := fate.GetFateData(&fate.FateInput{BirthDate: time.Now()})
package chronos
