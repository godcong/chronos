# chronos

[English](#english) | [中文](#中文)

---

## English

### Overview

chronos is a Go library for Chinese calendar calculations. It provides:

- Solar/Lunar date conversion (1900–3000)
- Heavenly Stems and Earthly Branches (天干地支) with type-safe enums
- Solar Terms (节气) with astronomical precision via [lunar-go](https://github.com/6tail/lunar-go)
- Chinese Zodiac (生肖) with LiChun boundary correction
- Western Constellation (星座) lookup
- Eight Characters (八字) analysis — Four Pillars, Five Elements, NaYin, Ten Gods, Hidden Stems, DaYun
- Five Element (五行) favorability analysis for name selection — Balance method and Pattern (格局) method

### Installation

```bash
go get github.com/godcong/chronos/v2
```

### Quick Start

```go
package main

import (
    "fmt"
    "time"

    "github.com/godcong/chronos/v2"
    "github.com/godcong/chronos/v2/fate"
)

func main() {
    // Parse a solar date
    cal := chronos.ParseSolarTime(time.Now())

    // Access lunar calendar data
    lunar := cal.Lunar()
    solar := cal.Solar()

    fmt.Println("Zodiac:", lunar.GetZodiac().Chinese())
    fmt.Println("Constellation:", solar.GetConstellation().Chinese())
    fmt.Println("Four Pillars:", lunar.GetEightChar().FourPillars())
    fmt.Println("Five Elements:", lunar.GetEightChar().FiveElements())
    fmt.Println("NaYin:", lunar.GetEightChar().NaYin())

    // Fate analysis for name selection
    data, err := fate.GetFateData(fate.FateInput{
        Calendar:     cal,
        Gender:       1,
        XiYongMethod: fate.XiYongMethodBalance,
    })
    if err != nil {
        panic(err)
    }
    fmt.Println("Xi (Favorable):", data.XiYongJiChou.Xi)
    fmt.Println("Yong (Useful):", data.XiYongJiChou.Yong)
    fmt.Println("Ji (Unfavorable):", data.XiYongJiChou.Ji)
    fmt.Println("Chou (Hostile):", data.XiYongJiChou.Chou)
}
```

### Parsing Dates

```go
// From time.Time
cal := chronos.ParseSolarTime(time.Now())

// From year/month/day/hour/minute/second
cal := chronos.ParseSolarDate(2024, 6, 15, 10, 30, 0)

// From string
cal := chronos.ParseSolarString("2024/06/15 10:30:00")

// Current time
cal := chronos.ParseSolarNow()

// From lunar date (with optional leap month flag)
cal := chronos.ParseLunarDate(2024, 5, 10, 10, 30, 0)
cal = chronos.ParseLunarDate(2024, 4, 10, 10, 30, 0, true) // leap month
```

### Package Structure

```
chronos/v2/
├── calendar.go          # Calendar constructors and date parsing
├── solar.go             # Solar (Gregorian) calendar
├── lunar.go             # Lunar calendar with BaZi support
├── interface.go         # Core interfaces (Calendar, Lunar, Solar, EightChar, etc.)
├── ganzhi.go            # Heavenly Stems, Earthly Branches, Stem-Branch enums
├── solarterm.go         # 24 Solar Terms with lunar-go astronomical engine
├── zodiac.go            # 12 Chinese Zodiac animals
├── constellation.go     # 12 Western constellations
├── eightchar.go         # Eight Characters (八字) implementation
├── fate/                # BaZi analysis and Five Element calculations
│   ├── fate.go          # FateInput → FateData entry point
│   ├── wuxing.go        # Five Element strength and favorability
│   ├── xiyong_balance.go # Balance method (平衡用神法)
│   ├── xiyong_geju.go   # Pattern method (格局用神法)
│   └── types.go         # Shared types (BaziInfo, XiYongJiChou, GeJuInfo, etc.)
└── utils/               # Internal utilities
```

### Key Types

| Type | Description |
|------|-------------|
| `Calendar` | Root interface for Solar + Lunar access |
| `Lunar` | Lunar calendar with 200+ traditional attributes |
| `Solar` | Gregorian calendar data |
| `TianGan` | Heavenly Stem enum (甲乙丙丁戊己庚辛壬癸) |
| `DiZhi` | Earthly Branch enum (子丑寅卯辰巳午未申酉戌亥) |
| `GanZhi` | Stem-Branch combination enum (60 甲子) |
| `SolarTerm` | Solar Term enum (24 节气) |
| `Zodiac` | Zodiac animal enum (12 生肖) |
| `Constellation` | Western constellation enum (12 星座) |
| `EightChar` | Eight Characters interface (Four Pillars, Five Elements, NaYin, etc.) |

### Chinese Conversion

All enum types implement the `ChineseSupport` interface:

```go
fmt.Println(chronos.TianGanJia.Chinese())          // "甲"
fmt.Println(chronos.DiZhiZi.Chinese())              // "子"
fmt.Println(chronos.GanZhiJiaZi.Chinese())          // "甲子"
fmt.Println(chronos.SolarTermLiChun.Chinese())      // "立春"
fmt.Println(chronos.ZodiacDragon.Chinese())         // "龙"
fmt.Println(chronos.ConstellationAries.Chinese())   // "白羊座"
```

### Fate Analysis

The `fate` sub-package provides BaZi analysis with two methods for determining Xi-Yong-Ji-Chou (喜用忌仇):

**Balance Method (平衡用神法)** — analyzes Five Element strength and balances the Day Master:

```go
data, _ := fate.GetFateData(fate.FateInput{
    Calendar:     cal,
    Gender:       1,
    XiYongMethod: fate.XiYongMethodBalance,
})
// data.XiYongJiChou → {Xi, Yong, Ji, Chou}
// data.WuxingStrength → {WuxingFen: map[木/火/土/金/水]float64, Total}
```

**Pattern Method (格局用神法)** — identifies the GeJu (格局) from the month branch and derives Xi-Yong-Ji-Chou:

```go
data, _ := fate.GetFateData(fate.FateInput{
    Calendar:     cal,
    Gender:       1,
    XiYongMethod: fate.XiYongMethodGeJu,
})
// data.GeJuInfo → {Type, Name, YongShen, XiShen, JiShen, ChouShen, Analysis}
// Supported patterns: 正官格, 七杀格, 正财格, 偏财格, 正印格, 偏印格, 食神格, 伤官格
```

### License

MIT License

---

## 中文

### 概述

chronos 是一个 Go 语言的中国农历计算库，提供：

- 阳历/阴历日期互转（支持 1900–3000 年）
- 天干地支类型安全的枚举
- 基于天文算法的 24 节气计算（通过 [lunar-go](https://github.com/6tail/lunar-go)）
- 带立春边界修正的生肖查询
- 西方星座查询
- 八字分析 — 四柱、五行、纳音、十神、藏干、大运
- 五行喜用忌仇分析（支持平衡用神法和格局用神法）

### 安装

```bash
go get github.com/godcong/chronos/v2
```

### 快速开始

```go
package main

import (
    "fmt"
    "time"

    "github.com/godcong/chronos/v2"
    "github.com/godcong/chronos/v2/fate"
)

func main() {
    // 解析阳历日期
    cal := chronos.ParseSolarTime(time.Now())

    // 获取阴历数据
    lunar := cal.Lunar()
    solar := cal.Solar()

    fmt.Println("生肖:", lunar.GetZodiac().Chinese())
    fmt.Println("星座:", solar.GetConstellation().Chinese())
    fmt.Println("四柱:", lunar.GetEightChar().FourPillars())
    fmt.Println("五行:", lunar.GetEightChar().FiveElements())
    fmt.Println("纳音:", lunar.GetEightChar().NaYin())

    // 命理分析（起名用）
    data, err := fate.GetFateData(fate.FateInput{
        Calendar:     cal,
        Gender:       1,
        XiYongMethod: fate.XiYongMethodBalance,
    })
    if err != nil {
        panic(err)
    }
    fmt.Println("喜神:", data.XiYongJiChou.Xi)
    fmt.Println("用神:", data.XiYongJiChou.Yong)
    fmt.Println("忌神:", data.XiYongJiChou.Ji)
    fmt.Println("仇神:", data.XiYongJiChou.Chou)
}
```

### 日期解析

```go
// 从 time.Time 解析
cal := chronos.ParseSolarTime(time.Now())

// 从 年/月/日/时/分/秒 解析
cal := chronos.ParseSolarDate(2024, 6, 15, 10, 30, 0)

// 从字符串解析
cal := chronos.ParseSolarString("2024/06/15 10:30:00")

// 当前时间
cal := chronos.ParseSolarNow()

// 从阴历日期解析（可选闰月标志）
cal := chronos.ParseLunarDate(2024, 5, 10, 10, 30, 0)
cal = chronos.ParseLunarDate(2024, 4, 10, 10, 30, 0, true) // 闰月
```

### 包结构

```
chronos/v2/
├── calendar.go          # 日历构造器和日期解析
├── solar.go             # 阳历（公历）
├── lunar.go             # 阴历（农历），支持八字
├── interface.go         # 核心接口（Calendar, Lunar, Solar, EightChar 等）
├── ganzhi.go            # 天干、地支、干支枚举
├── solarterm.go         # 24 节气（基于 lunar-go 天文算法）
├── zodiac.go            # 12 生肖
├── constellation.go     # 12 西方星座
├── eightchar.go         # 八字实现
├── fate/                # 八字分析与五行计算
│   ├── fate.go          # FateInput → FateData 入口
│   ├── wuxing.go        # 五行力量与喜忌
│   ├── xiyong_balance.go # 平衡用神法
│   ├── xiyong_geju.go   # 格局用神法
│   └── types.go         # 共享类型（BaziInfo, XiYongJiChou, GeJuInfo 等）
└── utils/               # 内部工具
```

### 核心类型

| 类型 | 说明 |
|------|------|
| `Calendar` | 根接口，提供 Solar + Lunar 访问 |
| `Lunar` | 阴历，包含 200+ 传统属性 |
| `Solar` | 阳历数据 |
| `TianGan` | 天干枚举（甲乙丙丁戊己庚辛壬癸） |
| `DiZhi` | 地支枚举（子丑寅卯辰巳午未申酉戌亥） |
| `GanZhi` | 干支组合枚举（60 甲子） |
| `SolarTerm` | 节气枚举（24 节气） |
| `Zodiac` | 生肖枚举（12 生肖） |
| `Constellation` | 西方星座枚举（12 星座） |
| `EightChar` | 八字接口（四柱、五行、纳音等） |

### 中文转换

所有枚举类型均实现 `ChineseSupport` 接口：

```go
fmt.Println(chronos.TianGanJia.Chinese())          // "甲"
fmt.Println(chronos.DiZhiZi.Chinese())              // "子"
fmt.Println(chronos.GanZhiJiaZi.Chinese())          // "甲子"
fmt.Println(chronos.SolarTermLiChun.Chinese())      // "立春"
fmt.Println(chronos.ZodiacDragon.Chinese())         // "龙"
fmt.Println(chronos.ConstellationAries.Chinese())   // "白羊座"
```

### 命理分析

`fate` 子包提供两种喜用忌仇分析方法：

**平衡用神法** — 分析五行力量，平衡日主：

```go
data, _ := fate.GetFateData(fate.FateInput{
    Calendar:     cal,
    Gender:       1,
    XiYongMethod: fate.XiYongMethodBalance,
})
// data.XiYongJiChou → {Xi, Yong, Ji, Chou}
// data.WuxingStrength → {WuxingFen: map[木/火/土/金/水]float64, Total}
```

**格局用神法** — 从月支识别格局，推导喜用忌仇：

```go
data, _ := fate.GetFateData(fate.FateInput{
    Calendar:     cal,
    Gender:       1,
    XiYongMethod: fate.XiYongMethodGeJu,
})
// data.GeJuInfo → {Type, Name, YongShen, XiShen, JiShen, ChouShen, Analysis}
// 支持格局：正官格、七杀格、正财格、偏财格、正印格、偏印格、食神格、伤官格
```

### 许可证

MIT License
