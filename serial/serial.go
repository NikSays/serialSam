// Utils for working with serial files
package serial

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/tarm/serial"
)

// Infinite loop reading upto 128 bytes from serial file
// and sending them to the channel.
// Designed to run as a goroutine
func ReadLoop (port *serial.Port, dataChan chan<- []byte){
	for{
		buffer := make([]byte, 128)
		n, _ := port.Read(buffer)
		dataChan <- buffer[:n]
	}
}


// Infinite loop reading from STDIN
// and sending them to the serial file.
// Designed to run as a goroutine
func WriteLoop (port *serial.Port){
var data string;
	for {
		fmt.Scanf("%s", &data)
		port.Write([]byte(data))
	}
}

// Returns opened serial file
func Open(device string, baud int) (*serial.Port, error){
	return serial.OpenPort(&serial.Config{Name: device, Baud: baud})
}

// Returns list of probable serial files on linux
func FindFiles() ([]string, error) {
	var serialFiles []string = make([]string, 0);
	serialRegex, _ := regexp.Compile(`(?i)^tty(acm|s|usb)\d+$`);

	files, err := ioutil.ReadDir("/dev");
	if err != nil {
		return nil, errors.New("/dev not accessible")
	}
	for _, file := range files {
		if (serialRegex.MatchString(file.Name())){
			serialFiles = append(serialFiles, fmt.Sprintf("/dev/%s",file.Name()))
		}
	}

	files, err = ioutil.ReadDir("/dev/pts");
	if err != nil {
		return nil, errors.New("/dev/pts not accessible")
	}
	for _, file := range files {
		serialFiles = append(serialFiles, fmt.Sprintf("/dev/pts/%s",file.Name()))
	}

	return serialFiles, nil
}