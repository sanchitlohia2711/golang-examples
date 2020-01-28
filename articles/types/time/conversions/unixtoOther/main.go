package main

import (
	"fmt"
	"time"
)

func main() {
	tUnix := time.Now().Unix()
	fmt.Printf("timeUnix %d\n", tUnix)

	//Conversion to time.Time
	timeT := time.Unix(tUnix, 0)
	fmt.Printf("time.Time: %s\n", timeT)

	//Conversion to Time Unix Millisecond
	tUnixMili := tUnix * int64(time.Microsecond)
	fmt.Printf("timeUnixMilli: %d\n", tUnixMili)

	//Conversion to Time Unix Microsecond
	tUnixMicro := tUnix * int64(time.Millisecond)
	fmt.Printf("timeUnixMicro: %d\n", tUnixMicro)

	//Conversion to Time Unix Nanosecond
	tUnixNano := tUnix * int64(time.Second)
	fmt.Printf("timeUnixNano: %d\n", tUnixNano)

}
