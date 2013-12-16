package main

import (
  "log"
  "recipes/recipeFinder"
)

func main() {
  log.Println("Starting Recipes App")

  finder, finderMessages := recipeFinder.New()
  finder.Start()

  select {
    case message := <-finderMessages:
      switch message {
        case "Done":
          log.Println("Finder is done")
      }
  }

  log.Println("Stopping Recipes App")
}
