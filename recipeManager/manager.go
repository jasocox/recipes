package recipeManager

import (
  "log"
  "goworker"
)

type RecipeManager struct {
  goworker.Manager
}

func New(numWorkers int) RecipeManager {
  name := "Recipe Manager"
  log.Println("Creating " + name)

  manager := RecipeManager{
    goworker.NewManager(name, numWorkers, createRecipeReader),
  }

  return manager
}

func createRecipeReader() goworker.Worker {
  return NewReader()
}
