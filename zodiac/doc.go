// Package zodiac provides the Chinese Zodiac (ShengXiao) type and calculations,
// including LiChun boundary correction for accurate year-based Zodiac determination.
//
// # Planned exports (Phase 3 migration)
//
// Type:      Zodiac (uint32) — from ../zodiac.go
// Constants: ZodiacZi..ZodiacMax — from ../zodiac_enum.go
// Functions: YearZodiac, YearZodiacDay, YearZodiacNoFix — from ../zodiac.go
// Interface: ChineseSupport (via root package re-export)
//
// This package has no internal dependencies on other chronos sub-packages.
package zodiac
