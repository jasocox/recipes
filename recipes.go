package main

import (
  "log"
  "recipes/recipeFinder"
  "recipes/recipeManager"
)

func main() {
  log.Println("Starting Recipes App")

  finder, finderMessages := recipeFinder.New()
  finder.MustStart()

  manager, managerMessages := recipeManager.New(10)
  manager.MustStart()

  for finder.Running() && manager.Running() {
    select {
    case message := <-managerMessages:
      switch message {
      case "Done":
        log.Println("Manager is done")
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
