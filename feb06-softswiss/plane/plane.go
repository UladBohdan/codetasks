package plane

import (
  "sort"
  "strconv"
  "strings"
)

const seatRow = "ABCDEFGIJK"

var possibleArrangements = []string{"BCDE", "DEFG", "FGIJ"}

// ArrangeFamilies ...
func ArrangeFamilies(nRows int, occupied string) int {
  sum := 0
  seatsOccupied := processInput(occupied)
  for row := 1; row <= nRows; row++ {
    rowOccupied := seatsOccupied[strconv.Itoa(row)]
    tempRow := seatRow
    for _, s := range rowOccupied {
      pos := strings.IndexRune(tempRow, s)
      if pos == -1 {
        panic("probably, a duplicate in input")
      }
      tempRow = tempRow[0:pos] + tempRow[pos+1:len(tempRow)]
    }
    for _, arr := range possibleArrangements {
      if strings.Index(tempRow, arr) != -1 {
        sum++
      }
    }
  }
  return sum
}

// processInput refactores input string into a map of rowNumber:sortedSeats
func processInput(occupied string) map[string]string {
  seats := make(map[string]string)
  occupiedSorted := strings.Fields(occupied)
  sort.Strings(occupiedSorted)
  for _, seat := range occupiedSorted {
    row, place := seat[0:len(seat)-1], seat[len(seat)-1:]
    seats[row] += string(place)
  }
  return seats
}
