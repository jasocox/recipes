package main

import (
  "log"
  "recipes/recipeFinder"
  "recipes/recipeReader"
)

func main() {
  log.Println("Starting Recipes App")

  finder, finderMessages := recipeFinder.New()
  finder.MustStart()

  reader, readerMessages := recipeReader.New()
  reader.MustStart()

  for finder.Running() && reader.Running() {
    select {
    case message := <-readerMessages:
      switch message {
      case "Done":
        log.Println("Reader is done")
      }
    case message := <-finderMessages:
      switch message {
          case "Done":
            log.Println("Finder is done")
      }
    }
  }

  log.Println("Stopping Recipes App")
}
