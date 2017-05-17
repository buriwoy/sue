package sue

import (
  "math"
  "time"
)


// numbers and letters
var set62 string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"


// returns unique id based on microseconds
func New () string {
  return intToBase(microEpoch(), "")
}


// returns unique id based on nanoseconds
func New2 () string {
  return intToBase(nanoEpoch(), "")
}


// converts integer to base62
func intToBase (decimal int64, acc string) string {
  base := int64(62)
  quotient := math.Floor(float64(decimal) / float64(base))
  remainder := decimal % base
  acc = string(set62[remainder]) + acc
  if quotient != 0 {
    acc = intToBase(int64(quotient), acc)
  }
  return acc
}


// creates epoch from 16-digit microseconds epoch
func microEpoch () int64 {
  // 01/01/2017 00:00:00
  start := int64(1483228800000000)
  return (time.Now().UnixNano() / 1000) - start
}


// creates epoch from 19-digit nanoseconds epoch
func nanoEpoch () int64 {
  // 01/01/2017 00:00:00
  start := int64(1483228800000000000)
  return time.Now().UnixNano() - start
}
