# Calendar 日期对象接口
```
 chronos.New() //创建当前时间
 chronos.New(time.Now()) //该参数的类型为time.Time
 chronos.New("2017/11/14 08:17") //参数是指定格式的字符串#年/月/日 时:分#
```

# Lunar 农历日期显示
```
chronos.New().Lunar() //获取月历
```
# Solar 公历日期显示
```
chornos.New().Solar() //获取日历
```


