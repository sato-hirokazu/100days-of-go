package animals

import (
 "testing"
)

func TestElephantFeed(t *testing.T) {
  expect := "Grass"
  actual := ElepahntFeed()

  if expect != actual {
     t.Errorf("%s != %s", expect, actual)
  }
}


