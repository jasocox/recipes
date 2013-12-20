package main

import (
  "log"
  "recipes/recipeFinder"
  "recipes/recipeManager"
  "goworker"
)

func main() {
  name := "Recipes App"
  log.Println("Starting " + name)

  app := goworker.NewManagerFromTasks(name, []func()goworker.Worker{createRecipeFinder, createRecipeManager})
  app.MustStart()
  <-app.Messages()

  log.Println("Stopping Recipes App")
}

func createRecipeFinder() goworker.Worker {
  return recipeFinder.New()
}

func createRecipeManager() goworker.Worker {
  return recipeManager.New(10)
}
