package recipeManager

import (
  "log"
  "errors"
)

type RecipeManager struct {
  messages chan string
  running bool
  workers []*RecipeReader
  workerMessageBoxes []chan string
}

func New(numWorkers int) (manager RecipeManager, messages chan string) {
  messages = make(chan string)
  manager.messages = messages

  manager.workers = make([]*RecipeReader, numWorkers)
  manager.workerMessageBoxes = make([]chan string, numWorkers)

  for i := range manager.workers {
    worker, workerMessageBox := NewReader()

    manager.workers[i] = &worker
    manager.workerMessageBoxes[i] = workerMessageBox
  }

  return
}

func (m *RecipeManager) MustStart() {
  err := m.Start()
  if err != nil {
    panic(err.Error())
  }
}

func (m *RecipeManager) Start() (err error) {
  if m.running {
    err = errors.New("Recipe Manager is already running")
    return
  }

  m.running = true
  log.Println("Starting the recipe manager")

  go func() {
    for _, worker := range m.workers {
      err = worker.StartReader()
      if err != nil {
        return
      }
    }

    m.watchRunners()
  }()

  return
}

func (m RecipeManager) Running() bool {
  return m.running
}

func (m *RecipeManager) watchRunners() {
  for _, workerMessageBox := range m.workerMessageBoxes {
    switch <-workerMessageBox {
    case "Done":
      log.Println("A worker is done")
    }
  }

  log.Println("All workers are complete")
  m.messages <- "Done"
  m.running = false
}
