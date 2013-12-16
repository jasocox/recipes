package recipeFinder

import (
  "log"
  "time"
)

type RecipeFinder struct {
  messages chan string
}

func New() (RecipeFinder, chan string) {
  log.Println("Creating Recipe Finder")

  var finder RecipeFinder
  messages := make(chan string)

  finder.messages = messages

  return finder, messages
}

func (f RecipeFinder) Start() {
  log.Println("Started the recipe finder")

  go func() {
    time.Sleep(2000 * time.Millisecond)
    f.messages <- "Done"
  }()
}
