package main

/* Workers are a very powerful concept in Go. They are used to perform tasks concurrently.
 * In Go, we can create multiple workers to perform tasks concurrently. This is very useful when we have to perform multiple tasks that are independent of each other.
 * For example, if we have to download multiple files from the internet, we can create multiple workers to download each file concurrently.
 */

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

// Worker is a struct that represents a worker. A worker is a goroutine that performs a task.
type Worker struct {
	ID int
}

type Task func()

func exampleTask() {
	// This can be long running task. For example, downloading a file from the internet.
	// For the sake of simplicity, we are just printing a message here. and sleeping for 1 second.
	log.Println("Starting task ...")
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	log.Println("Task completed")
}

// Start is a method on Worker that starts the worker.
func (w *Worker) Start(task Task) {
	log.Printf("Worker %d started", w.ID)
	go func() {
		task()
	}()
}

// NewWorker is a function that creates a new worker.
func NewWorker(id int) *Worker {
	return &Worker{
		ID: id,
	}
}

// WorkerPool is a struct that represents a pool of workers.
type WorkerPool struct {
	Workers []*Worker
}

// NewWorkerPool is a function that creates a new worker pool.
func NewWorkerPool(numWorkers int) *WorkerPool {
	workers := make([]*Worker, numWorkers)
	for i := range numWorkers {
		workers[i] = NewWorker(i)
	}
	return &WorkerPool{
		Workers: workers,
	}
}

// Start is a method on WorkerPool that starts all the workers in the pool.
func (wp *WorkerPool) Start(task Task) {
	var wg sync.WaitGroup
	for _, worker := range wp.Workers {
		wg.Add(1)
		go func(w *Worker) {
			defer func() {
				log.Printf("Worker %d stopped", worker.ID)
				wg.Done()
			}()
			w.Start(task)
		}(worker)
	}
	wg.Wait()
}

// RunWorkerPool is a function that demonstrates how to use the WorkerPool.
func RunWorkerPool() {
	// Create a new worker pool with 5 workers
	wp := NewWorkerPool(10)

	// Start the worker pool
	log.Println("Starting worker pool ...")
	wp.Start(exampleTask)
}

func main() {
	RunWorkerPool()
}
