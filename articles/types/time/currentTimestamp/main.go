package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now() //It will return time.Time object with current timestamp
	fmt.Printf("time.Time %s\n", t)

	tUnix := t.Unix()
	fmt.Printf("timeUnix: %d\n", tUnix)

	tUnixNano := t.UnixNano()
	fmt.Printf("timeUnixNano: %d\n", tUnixNano)

	tUnixMilli := int64(time.Nanosecond) * t.UnixNano() / int64(time.Millisecond)
	fmt.Printf("timeUnixMilli: %d\n", tUnixMilli)

	tUnixMicro := int64(time.Nanosecond) * t.UnixNano() / int64(time.Microsecond)
	fmt.Printf("timeUnixMicro: %d\n", tUnixMicro)

}
