// A gui/cli serial port communication software
package main

import (
	"fmt"
	"log"

	"github.com/niksays/serialSam/gui"
	"github.com/niksays/serialSam/serial"
) 

func main() {
	go gui.OpenWindow()
	files, err := serial.FindFiles()
	if err != nil {
		log.Fatal("Couldn't list files: ", err)
	}
	fmt.Println(files)

	var device string;
	var baud int;
	fmt.Scanf("%s", &device)
	if device == "" {
		device = "/dev/ttyACM0"
	}
	fmt.Scanf("%d", &baud)
	if baud == 0 {
		baud = 9600
	}

	sPort, err := serial.Open(device, baud)
	if err != nil {
		log.Fatal("Couldn't open device: ", err)
	}

	resChan := make(chan []byte)
	
	go serial.ReadLoop(sPort, resChan)
	go serial.WriteLoop(sPort)

	var rxData []byte = make([]byte, 0)
	for {
		rxData = <- resChan
		fmt.Print(string(rxData))
	}
}
