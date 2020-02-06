package plane

import (
  "testing"
)

// TestArrangeFamilies ...
func TestArrangeFamilies(t *testing.T) {
  data := []struct {
    descr    string
    numRows  int
    occupied string
    expected int
  }{
    {
      descr:    "nothing occupied",
      numRows:  4,
      occupied: "",
      expected: 12,
    },
    {
      descr:    "all occupied",
      numRows:  1,
      occupied: "1K 1J 1I 1G 1F 1E 1D 1A 1B 1C",
      expected: 0,
    },
    {
      descr:    "some occupied",
      numRows:  2,
      occupied: "1D 1G 2D",
      expected: 1,
    },
    {
      descr:    "there are double-digits",
      numRows:  100,
      occupied: "15D",
      expected: 3*99 + 1,
    },
  }
  for _, tc := range data {
    t.Run(tc.descr, func(t *testing.T) {
      result := ArrangeFamilies(tc.numRows, tc.occupied)
      if result != tc.expected {
        t.Errorf("got %v, expected %v", result, tc.expected)
      }
    })
  }
}
