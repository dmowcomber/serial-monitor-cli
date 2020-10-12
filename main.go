package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/tarm/serial"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("expected device argument: go run serial_monitor.go /dev/the-device")
	}
	device := os.Args[1]

	serialConfig := &serial.Config{Name: device, Baud: 9600}
	serialPort, err := serial.OpenPort(serialConfig)
	if err != nil {
		log.Printf("err: %s", err.Error())
		os.Exit(1)
	}

	br := bufio.NewReader(serialPort)

	for {
		line, _, err := br.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(line))
	}
}
