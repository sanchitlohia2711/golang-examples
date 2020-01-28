package main

import (
	"fmt"
	"time"
)

func main() {
	timeT := time.Now() //It will return time.Time object with current timestamp
	fmt.Printf("time.Time %s\n", timeT)

	//Converstion to Time Unix also known as epoch time
	tUnix := timeT.Unix()
	fmt.Printf("timeUnix: %d\n", tUnix)

	//Conversion to Time Unix Millisecond
	tUnixMilli := int64(time.Nanosecond) * timeT.UnixNano() / int64(time.Millisecond)
	fmt.Printf("timeUnixMilli: %d\n", tUnixMilli)

	//Conversion to Time Unix Microsecond
	tUnixMicro := int64(time.Nanosecond) * timeT.UnixNano() / int64(time.Microsecond)
	fmt.Printf("timeUnixMicro: %d\n", tUnixMicro)

	//Conversion to Time Unix Nanosecond
	tUnixNano := timeT.UnixNano()
	fmt.Printf("timeUnixNano: %d\n", tUnixNano)

}
