package main

import (
  "log"
  "recipes/recipeFinder"
  "recipes/recipeManager"
)

func main() {
  log.Println("Starting Recipes App")

  finder := recipeFinder.New()
  log.Println("Starting the recipe finder")
  finder.MustStart()

  manager := recipeManager.New(10)
  log.Println("Starting the recipe manager")
  manager.MustStart()

  <-finder.Messages()
  log.Println("Finder is done")
  <-manager.Messages()
  log.Println("Manager is done")

  log.Println("Stopping Recipes App")
}
