package main

import (
	"fmt"
	"sync"
	"time"
)

// Task definition
type Task interface {
	Process()
}

// Email task definition
type EmailTask struct {
	Email, Subject, Message string
}

// Way to process the Email task
func (t *EmailTask) Process() {
	fmt.Printf("Sending the email to %s\n", t.Email)
	// Simulating a time-consuming process
	time.Sleep(2 * time.Second)
}

// Image task definition
type ImageTask struct {
	ImageURL string
}

func (t *ImageTask) Process() {
	fmt.Printf("Processing the image: %s\n", t.ImageURL)
	// Simulating a time-consuming process
	time.Sleep(5 * time.Second)
}

// Worker pool definition
type WorkerPool struct {
	Tasks       []Task
	concurrency int
	taskChan    chan Task
	wg          sync.WaitGroup
}

// Functions to execute the worker pool
func (wp *WorkerPool) worker() {
	for task := range wp.taskChan {
		task.Process()
		wp.wg.Done()
	}
}

// Run method initialize the channel, sets the concurrency, creates the goroutines and
// sends tasks over the channel
func (wp *WorkerPool) Run() {
	// Initialize the tasks channel with the capacity equals to the number of tasks
	wp.taskChan = make(chan Task, len(wp.Tasks))

	// Start workers
	for i := 0; i < wp.concurrency; i++ {
		go wp.worker()
	}

	// Send tasks to the channel
	wp.wg.Add(len(wp.Tasks))
	for _, t := range wp.Tasks {
		wp.taskChan <- t
	}
	close(wp.taskChan)

	// Wait for all tasks
	wp.wg.Wait()
}
