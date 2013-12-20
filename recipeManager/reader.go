package recipeManager

import (
  "log"
  "time"
  "math/rand"

  "goworker"
)

type RecipeReader struct {
  goworker.Worker
}

func NewReader() RecipeReader {
  name := "Recipe Reader"
  log.Println("Creating " + name)

  reader := RecipeReader{goworker.NewWorker(name, func() {
    log.Println(name + " is starting")

    millis := rand.Intn(400) + 100

    log.Printf("Reader doing %d milliseconds worth of work", millis)
    time.Sleep(time.Duration(int(time.Millisecond) * millis))

    log.Println(name + " is done")
  })}

  return reader
}
