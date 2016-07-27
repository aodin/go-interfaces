package main

import "fmt"
import "log"
import "time"

type Job interface {
	Log(...interface{})
	Pid() uint64
	Start(halt <-chan bool) error
}

type crunch struct{}

func (j crunch) Pid() uint64 { return 1 }

func (j crunch) Start(halt <-chan bool) error {
	for chunk := 0; chunk < 10; chunk += 1 {
		select {
		case <-halt:
			return fmt.Errorf("Halting during chunk %d", chunk)
		default:
			break
		}
		time.Sleep(time.Second) // Or any high difficulty task
		log.Printf("Finished chunk %d", chunk)
	}
	return nil
}

func main() {
	halt := make(chan bool)
	go func() {
		time.Sleep(time.Second * 3)
		halt <- true
	}()
	var job crunch
	if err := job.Start(halt); err != nil {
		log.Println(err)
	}
}
