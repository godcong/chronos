# Calendar Date object interface
```
 chronos.New() //Create current time
 chronos.New(time.Now()) //The parameter is of type time.Time
 chronos.New("2017/11/14 08:17") //Parameter is a specified format string
```

# Lunar 农历日期显示
```
chronos.New().Lunar() //Get the lunar calendar
```
# Solar 公历日期显示
```
chornos.New().Solar() //Get the solar calendar
```


