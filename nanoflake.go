package nanoflakes

import (
	"strconv"
	"time"
)

type Nanoflake struct {
	epoch int64
	value int64
}

// returns the epoch value (in milliseconds) of the nanoflake
func (n Nanoflake) Epoch() int64 {
	return n.epoch
}

// returns the value of the nanoflake
func (n Nanoflake) Value() int64 {
	return n.value
}

// returns the value of the nanoflake as a string
func (n Nanoflake) ToString() string {
	return strconv.FormatInt(n.value, 10)
}

// returns the nanoflake's creation time  in milliseconds
func (n Nanoflake) CreationTimeMillis() int64 {
	return n.epoch + TimeStampValue(n.value)
}

// returns the nanoflake's creation time in UTC
func (n Nanoflake) CreationTime() time.Time {
	return time.UnixMilli(n.CreationTimeMillis())
}
