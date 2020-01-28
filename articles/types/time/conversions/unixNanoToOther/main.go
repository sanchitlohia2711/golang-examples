package main

import (
	"fmt"
	"time"
)

func main() {
	tUnixNano := time.Now().UnixNano()
	fmt.Printf("tUnixNano %d\n", tUnixNano)

	tUnix := tUnixNano / int64(time.Second)
	tUnixNanoRemainder := (tUnixNano % int64(time.Second))
	timeT := time.Unix(tUnix, tUnixNanoRemainder)
	fmt.Printf("time.Time: %s\n", timeT)

	//Conversion to Time Unix
	tUnix = tUnixNano / int64(time.Second)
	fmt.Printf("timeUnix: %d\n", tUnix)

	//Conversion to Time Unix Milli
	tUnixMilli := tUnixNano / int64(time.Millisecond)
	fmt.Printf("timeUnixMilli: %d\n", tUnixMilli)

	//Conversion to Time Unix Micro
	tUnixMicro := tUnixNano / int64(time.Microsecond)
	fmt.Printf("timeUnixMicro: %d\n", tUnixMicro)
}
