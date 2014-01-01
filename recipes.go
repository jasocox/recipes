package main

import (
  "log"
  "recipes/recipeFinder"
  "recipes/recipeReader"
  "goworker"
)

func main() {
  name := "Recipes Task Manager"

  app := goworker.NewManager(name, 10)

  app.Exec(recipeFinder.New())
  for i:=0; i<99; i++ {
    app.Exec(recipeReader.New())
  }

  app.Finish()

  log.Println("Stopping Recipes App")
}
