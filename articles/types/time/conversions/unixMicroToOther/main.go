package main

import (
	"fmt"
	"time"
)

func main() {
	tUnixMicro := int64(time.Nanosecond) * time.Now().UnixNano() / int64(time.Microsecond)
	fmt.Printf("tUnixMicro %d\n", tUnixMicro)

	//Conversion to time.Time
	tUnix := tUnixMicro / int64(time.Millisecond)
	tUnixNanoRemainder := (tUnixMicro % int64(time.Millisecond)) * int64(time.Microsecond)
	timeT := time.Unix(tUnix, tUnixNanoRemainder)
	fmt.Printf("time.Time: %s\n", timeT)

	//Converstion to Time Unix
	tUnix = tUnixMicro / int64(time.Millisecond)
	fmt.Printf("timeUnix: %d\n", tUnix)

	//Converstion to Time Unix Milli
	tUnixMilli := tUnixMicro / int64(time.Microsecond)
	fmt.Printf("timeUnixMill: %d\n", tUnixMilli)

	//Converstion to Time Unix Nano
	tUnixNano := tUnixMicro * int64(time.Microsecond)
	fmt.Printf("timeUnixNano: %d\n", tUnixNano)
}
