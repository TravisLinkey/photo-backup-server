package utils

import (
  "log"
  "fmt"
  "os"
)

func ExitErrorf(msg string, args ...interface{}) {
  log.Printf("Should show in logs")
  log.Printf(msg)
  fmt.Printf(msg)
  fmt.Fprintf(os.Stderr, msg + "\n", args...)
  os.Exit(1)
}
