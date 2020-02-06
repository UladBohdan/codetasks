package main

import (
  "fmt"

  "github.com/uladbohdan/codetasks/feb06/plane"
)

func main() {
  numFamilies := plane.ArrangeFamilies(4, "1A 1B 3C")

  fmt.Printf("Number of families to arrange: %d\n", numFamilies)
}
