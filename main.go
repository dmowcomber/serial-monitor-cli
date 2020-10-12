package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/tarm/serial"
)

func main() {
	var device string
	var baud int
	flag.StringVar(&device, "device", "", "the device to listen to. Example: /dev/tty.usbserial. REQUIRED")
	flag.IntVar(&baud, "baud", 9600, "the baud of the serial connection")
	flag.Parse()

	if device == "" {
		log.Fatal("expected device argument: go run main.go -device=/dev/the-device -baud=9600")
	}

	serialConfig := &serial.Config{Name: device, Baud: baud}
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
