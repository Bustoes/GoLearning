package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()	//获取时间对象 Time
	fmt.Printf("%v\n", now)
	fmt.Printf("%#v\n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	fmt.Println(year, month, day, hour, minute, second)

	//// 时间戳 1970.1.1 CST 8:00:00 到现在的毫秒数
	//unix := now.Unix()
	//unixMicro := now.UnixMicro()
	//unixNano := now.UnixNano()
	//fmt.Println(unix, unixMicro, unixNano)
	//
	////将时间戳转换为具体的时间格式
	//t := time.Unix(unix, 0)
	//fmt.Println(t)
	//
	////时间间隔	一个数据类型，int64
	//n := 5	//int	强制类型转换
	//time.Sleep(time.Duration(n) * time.Second)
	//fmt.Println("over")

	//时间方法
	t1 := now.Add(time.Hour)
	fmt.Printf("%v\n", t1)
	t2 := t1.Sub(now)
	fmt.Printf("%v\n", t2)

	////定时器
	//tick := time.Tick(time.Second)	//返回值类型是一个通道 channel
	//for tmp := range tick {
	//	fmt.Printf("%v\n", tmp)
	//}

	//时间格式化
	timeFormat := now.Format("2006.01.02 15:04:05 PM  时区：-07")
	fmt.Printf("%v\n", timeFormat)

	//解析字符串类型的时间
	timeStr := "2019/08/07 15:23:48"
	//1.拿到时区
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	//2.根据时区解析一个字符串格式的时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", timeStr, location)
	if err != nil {
		fmt.Printf("parse timeStr failed, err:%v\n", err)
		return
	}
	fmt.Println(timeObj)

}
