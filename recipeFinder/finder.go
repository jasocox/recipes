package recipeFinder

import (
  "log"
  "time"
)

type RecipeFinder struct {}

func New() RecipeFinder {
  return RecipeFinder{}
}

func (f RecipeFinder) Do() error {
  log.Println("Finder working for 200 milliseconds")

  time.Sleep(2000 * time.Millisecond)

  return nil
}
