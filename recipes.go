package main

import (
  "log"
  "github.com/jasocox/recipes/recipeFinder"
  "github.com/jasocox/recipes/recipeReader"
  "github.com/jasocox/goworker"
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
