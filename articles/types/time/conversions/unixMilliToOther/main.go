package main

import (
	"fmt"
	"time"
)

func main() {
	tUnixMilli := int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Millisecond)
	fmt.Printf("timeMilli %d\n", tUnixMilli)

	time.Parse

	//Conversion to time.Time
	tUnix := tUnixMilli / int64(time.Microsecond)
	tUnixNanoRemainder := (tUnixMilli % int64(time.Microsecond)) * int64(time.Millisecond)
	timeT := time.Unix(tUnix, tUnixNanoRemainder)
	fmt.Printf("time.Time: %s\n", timeT)

	//Conversion to Time Unix
	tUnix = tUnixMilli / int64(time.Microsecond)
	fmt.Printf("timeUnix: %d\n", tUnix)

	//Conversion to Time Unix Microsecond
	tUnixMicro := tUnixMilli * int64(time.Microsecond)
	fmt.Printf("timeUnixMicro: %d\n", tUnixMicro)

	//Conversion to Time Unix Nanosecond
	tUnixNano := tUnixMilli * int64(time.Millisecond)
	fmt.Printf("timeUnixNano: %d\n", tUnixNano)

}
