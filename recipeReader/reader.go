package recipeReader

import (
  "log"
  "time"
  "math/rand"
)

type RecipeReader struct {}

func New() RecipeReader {
  return RecipeReader{}
}

func (r RecipeReader) Do() error {
  millis := rand.Intn(400) + 100

  log.Printf("Reader doing %d milliseconds worth of work", millis)
  time.Sleep(time.Duration(int(time.Millisecond) * millis))

  return nil
}
