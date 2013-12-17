package recipeFinder

import (
  "log"
  "time"
  "errors"
)

type RecipeFinder struct {
  messages chan string
  running bool
}

func New() (finder RecipeFinder, messages chan string) {
  log.Println("Creating Recipe Finder")

  messages = make(chan string)
  finder.messages = messages

  return
}

func (f *RecipeFinder) MustStart() {
  err := f.Start()
  if err != nil {
    panic(err.Error())
  }
}

func (f *RecipeFinder) Start() (err error) {
  if f.running {
    err = errors.New("Recipe Finder is already running")
    return
  }

  f.running = true
  log.Println("Started the recipe finder")

  go func() {
    time.Sleep(2000 * time.Millisecond)
    f.running = false
    f.messages <- "Done"
  }()

  return
}

func (f RecipeFinder) Running() bool {
  return f.running
}
