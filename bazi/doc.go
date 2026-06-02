// Package bazi provides BaZi (Eight Characters) related types and calculations,
// including Heavenly Stems (TianGan), Earthly Branches (DiZhi), Stem-Branch
// combinations (GanZhi), the Four Pillars, Hidden Stems, and EightChar interface.
//
// # Planned exports (Phase 3 migration)
//
// Types:     TianGan (uint32), DiZhi (uint32), GanZhi (uint32)
//            EightChar (interface), Pillar, FourPillarsStruct
//            EightCharIndex (int)
// Constants: TianGanJia..TianGanMax, DiZhiZi..DiZhiMax, GanZhiJiaZi..GanZhiMax
//            — from ../gan_zhi_enum.go
// Variables: DiZhiHiddenStems — from ../wuxing_data.go
//            tianGanChinese, diZhiChinese, ganZhiChinese — from ../gan_zhi.go
// Functions: ShiZhu, RiZhu, YueZhu, NianZhu (Four Pillars)
//            parseGanZhi, getTianGan, getDiZhi, splitGanZhi
//            FourPillarsFromArr
// Type aliases: StemBranch = GanZhi, PillarHour/Day/Month/Year
//
// Dependencies: none on other chronos sub-packages
//               requires github.com/6tail/lunar-go/calendar (via EightChar impl)
package bazi
