package recipeManager

import (
  "log"
  "time"
  "errors"
  "math/rand"
)

type RecipeReader struct {
  messages chan string
  running bool
}

func NewReader() (reader RecipeReader, messages chan string) {
  log.Println("Creating Recipe Reader")

  messages = make(chan string)
  reader.messages = messages

  return
}

func (r *RecipeReader) MustStartReader() {
  err := r.StartReader()
  if err != nil {
    panic(err.Error())
  }
}

func (r *RecipeReader) StartReader() (err error) {
  if r.running {
    err = errors.New("Recipe Reader is already running")
    return
  }

  r.running = true
  log.Println("Starting the recipe reader")

  go func() {
    seconds := rand.Intn(4)

    log.Printf("Reader doing %d seconds worth of work", seconds)
    time.Sleep(time.Duration(int(time.Second) * seconds))

    r.running = false
    r.messages <- "Done"
  }()

  return
}

func (r RecipeReader) RunningReader() bool {
  return r.running
}
