package workpool

import (
  "sync"
)

type Pool struct {
  Tasks   []Task

  concurrency   int
  input         chan Task
  output        chan Location
  errors        chan error
  wg            sync.WaitGroup
}

func New(tasks []Task, concurrency int) Pool {
  return Pool{
    Tasks:       tasks,
    concurrency: concurrency,
    input:       make(chan Task),
    output:      make(chan Location),
    errors:      make(chan error),
  }
}

func (p Pool) Output() <-chan Location {
  return p.output
}

func (p Pool) Errors() <-chan error {
  return p.errors
}

func (p Pool) Run() {
  for i := 1; i <= p.concurrency; i++ {
    worker := NewWorker(p.input, p.output, p.errors)
    worker.Start(&p.wg)
  }

  for i := range p.Tasks {
    p.input <- p.Tasks[i]
  }
  close(p.input)

  p.wg.Wait()
}

func(p Pool) Stop() {
  close(p.output)
  close(p.errors)
}

