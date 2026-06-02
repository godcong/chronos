// Package term provides Solar Term (JieQi) related types and calculations,
// including the 24 Solar Terms, their dates, and SanHou (three periods) data.
//
// # Planned exports (Phase 3 migration)
//
// Types:     SolarTerm (uint32), SolarTermDetail (struct)
// Constants: SolarTermXiaoHan..SolarTermMax — from ../solar_term_enum.go
// Functions: GetSolarTerm, getSolarTermDay, getSolarTermTime
//            YearSolarTermDate, YearSolarTermDay, YearSolarTermMonth
//            YearSolarTermDetail, IsSolarTermDetailDay
//            yearLiChunDay
// Variables: solarTermChinese, sanHouData, solarTermData — from ../solar_term_data.go
//
// Dependencies: none on other chronos sub-packages
//               requires github.com/6tail/lunar-go/calendar (for date computation)
package term
