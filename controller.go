package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const adamEndpoint string = "http://localhost:80/"

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter the digital pin number: ")

	pinNum, _ := reader.ReadString('\n')

	fmt.Print("Enter 0 for OFF and 1 for ON: ")

	signal, _ := reader.ReadString('\n')

	setDigitalOutput(strings.TrimSpace(pinNum), strings.TrimSpace(signal))
	main()
}

func setDigitalOutput(pinNumber string, signal string) {
	data := url.Values{}
	data.Add(fmt.Sprintf("DO%s", pinNumber), signal)
	fmt.Println(data)
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, adamEndpoint+"digitaloutput/all/value", strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	req.Header.Add("Authorization", "Basic cm9vdDowMDAwMDAwMA==")
	r, err := client.Do(req)

	if err != nil {
		fmt.Println("Error sending req")
		fmt.Print(err.Error())
	} else {
		defer r.Body.Close()
		fmt.Println("Request sent succesfully")
	}

}
