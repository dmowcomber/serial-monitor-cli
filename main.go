package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"

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
	retryReadSerial(serialConfig)
}

func retryReadSerial(serialConfig *serial.Config) {
	var lastErr error
	for {
		err := readSerial(serialConfig)
		if lastErr == nil || err.Error() != lastErr.Error() {
			log.Println(err.Error())
			log.Println("attemtping again...")
		}
		lastErr = err
	}
}

func readSerial(serialConfig *serial.Config) error {
	serialPort, err := serial.OpenPort(serialConfig)
	if err != nil {
		return fmt.Errorf("failed to open serial port: %s", err.Error())
	}

	br := bufio.NewReader(serialPort)
	for {
		line, _, err := br.ReadLine()
		if err != nil {
			return fmt.Errorf("failed to read: %s", err.Error())
		}
		log.Println(string(line))
	}
}
