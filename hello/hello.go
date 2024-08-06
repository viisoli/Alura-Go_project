package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
	"strconv"
)

const monitoring = 2
const delay = 2

func main() {
	
	showIntro()

	for {
	showOptions()

	command := readCommand()

	switch command {
	case 1:
		initMonitoring()
	case 2:
		fmt.Println("Showing logs...")
		printLogs()
	case 0:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid command.")
		os.Exit(-1)
	}
}
}

func showIntro() {
	name := "Name"
	version := 1.1
	fmt.Println("Hello,", name,"!")
	fmt.Println("This program is on", version, "version.")
}

func showOptions() {
	fmt.Println("--- OPTIONS ---")
	fmt.Println("1- Start Monitoring")
	fmt.Println("2- Show logs")
	fmt.Println("0- Exit")
}

func readCommand() int {
	var command int

	fmt.Scan(&command)
	fmt.Println("The selected command is", command)
	fmt.Println("")

	return command
}

func initMonitoring() {
	fmt.Println("Monitoring...")
	sites := readFile()

	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("Position", i, "Site:", site)

			monitoringSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")
}

func monitoringSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("An error occurred:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "was loaded successfully")
		createLog(site, true)
	} else {
		fmt.Println("Site:", site, "is not working. Status Code:", resp.StatusCode)
		createLog(site, false)
	}
}

func readFile() []string{
	var sites []string

	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("An error occurred:", err)
		os.Exit(0)
	}

	reader := bufio.NewReader(file)
	for{
	line, err := reader.ReadString('\n')
	line = strings.TrimSpace(line)

	sites = append(sites, line)

	if err ==  io.EOF {
		break
	}	
	}

	file.Close()

	return sites
}

func createLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("An error occurred:", err)
	}

	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()
}

func printLogs() {
	file, err := os.ReadFile("log.txt")

	if err != nil {
		fmt.Println("An error occurred:", err)
	}

	fmt.Println(string(file))
}