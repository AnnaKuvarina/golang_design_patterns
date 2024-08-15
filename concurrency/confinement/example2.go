package main

import "fmt"

// CounterCommand represents a command that can be sent to the counter goroutine
type CounterCommand struct {
	Action   string
	Value    int
	Response chan int
}

// StartCounter creates a goroutine that owns a counter and processes commands
func StartCounter() chan<- CounterCommand {
	commands := make(chan CounterCommand)

	go func() {
		counter := 0

		for cmd := range commands {
			switch cmd.Action {
			case "increment":
				counter += cmd.Value
				cmd.Response <- counter
			case "decrement":
				counter -= cmd.Value
				cmd.Response <- counter
			case "get":
				cmd.Response <- counter
			default:
				fmt.Println("Unknown command")
				close(cmd.Response)
			}
		}
	}()

	return commands
}

func main() {
	worker := StartCounter()

	incrementResp := make(chan int)
	worker <- CounterCommand{Action: "increment", Value: 10, Response: incrementResp}
	fmt.Println("Counter after incrementing by 10:", <-incrementResp) // Prints: 10

	decrementResp := make(chan int)
	worker <- CounterCommand{Action: "decrement", Value: 3, Response: decrementResp}
	fmt.Println("Counter after decrementing by 10:", <-decrementResp) // Prints: 7

	getResp := make(chan int)
	worker <- CounterCommand{Action: "get", Response: getResp}
	fmt.Println("Current counter value:", <-getResp) // Prints: 7

	unknownResp := make(chan int)
	worker <- CounterCommand{Action: "close", Response: unknownResp}
	fmt.Println("Current counter value:", <-unknownResp)

	close(worker)
}
