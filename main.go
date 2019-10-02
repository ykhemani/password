package main

import (
  "fmt"
  "log"
  "flag"
  "github.com/sethvargo/go-password/password"
)

func main() {
  // Parse command line options with default values defined below.
  // Generate a password that is lengthPtr (64) characters long
  // with digitsPtr (10) digits, symbolPtr (10) symbols,
  // allowing upper and lower case letters unless
  // no_upperPtr is set to true, disallowing repeat characters
  // unless allow_repeatPtr is set to true.
  lengthPtr  := flag.Int("length", 32, "length of password")
  digitsPtr  := flag.Int("digits",  6, "number of digits")
  symbolsPtr := flag.Int("symbols", 8, "number of symbols")
  no_upperPtr := flag.Bool("no_upper", false, "No uppercase characters (default: false)")
  allow_repeatPtr := flag.Bool("allow_repeat", false, "Allow repeating characters (default: false)")

  flag.Parse()

  //res, err := password.Generate(64, 10, 10, false, false)
  res, err := password.Generate(*lengthPtr, *digitsPtr, *symbolsPtr, *no_upperPtr, *allow_repeatPtr)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(res)
}
