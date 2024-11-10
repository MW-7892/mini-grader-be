package utils

import "strconv"

func StringToUint(value string) uint {
  value_int, err := strconv.ParseUint(value, 10, 0)
  if err != nil {
    panic(err)
  }
  return uint(value_int)
}

func UintToString(value uint) string {
  value_string := strconv.FormatUint(uint64(value), 10)
  return value_string
}
