package recipeFinder

import (
  "log"
  "time"
  "goworker"
)

type RecipeFinder struct {
  goworker.Worker
}

func New() RecipeFinder {
  name := "Recipe Finder"
  log.Println("Creating " + name)

  finder := RecipeFinder{goworker.NewWorker(name, func() {
    log.Println(name + " is starting")

    time.Sleep(200 * time.Millisecond)

    log.Println(name + " is done")
  })}

  return finder
}
