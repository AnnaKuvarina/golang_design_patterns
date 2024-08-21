package main

import "fmt"

func main() {
	// create 20 tasks
	tasks := []Task{
		&EmailTask{
			Email:   "test1@gmai.com",
			Subject: "test",
			Message: "test",
		},
		&EmailTask{
			Email:   "test2@gmai.com",
			Subject: "test",
			Message: "test",
		},
		&ImageTask{ImageURL: "test.jpeg"},
		&ImageTask{ImageURL: "test2.jpeg"},
		&EmailTask{
			Email:   "test3@gmai.com",
			Subject: "test",
			Message: "test",
		},
		&ImageTask{ImageURL: "test3.jpeg"},
		&ImageTask{ImageURL: "test4.jpeg"},
		&EmailTask{
			Email:   "test4@gmai.com",
			Subject: "test",
			Message: "test",
		},
		&EmailTask{
			Email:   "test5@gmai.com",
			Subject: "test",
			Message: "test",
		},
	}

	// Create a worker pool
	wp := WorkerPool{
		Tasks:       tasks,
		concurrency: 5, // Number of the workers that can run at a time
	}

	// Run the pool
	wp.Run()

	fmt.Println("All tasks have been processed")
}
