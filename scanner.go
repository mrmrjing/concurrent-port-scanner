package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

func scanPort(in chan string){
	for{
		hostport :=<- in

		conn, err := net.DialTimeout("tcp", hostport, 1*time.Second)

		if err != nil {
			fmt.Println(hostport + " is closed")
		} else {
			fmt.Println(hostport + " is open")
			conn.Close()
		}
		}
}

var (
	hosts string
	ports string

	// Number of routines to be used, by default it is 2500 from empirical results. 
	// If resource exceeds, the concurrent port scanner works too for smaller numbers of routines but
	// time it takes would increase as seen from the graph in the report.
	number_of_routines int = 2500
)

// https://stackoverflow.com/questions/67788289/iterating-over-a-multiline-variable-to-return-all-ips
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

// Parse the arguments passed to the program
// go run scanner.go -p ports hosts
func parse_arguments() error {
	// Set up flags
	flag.StringVar(&ports, "p", "", "help message")
	flag.Parse() // Scans the arg list and sets up flags

	// Check if hosts were provided
	if len(flag.Args()) < 1 {
		return errors.New("No hosts provided")
	}

	hosts = flag.Args()[0] // get the first arg

	// Check if ports were provided
	if ports == "" {
		return errors.New("No ports provided")
	}

	// Print the value of arguments passed (hosts and ports)
	fmt.Println("Hosts input: ", hosts)
	fmt.Println("Ports input: ", ports)

	// Parse ports of the form 22-25,80,8080 to array
	var stringPortArray = strings.Split(ports, ",") // Split on comma

	var portArray []int
	// Loop through the ports
	for _, port := range stringPortArray {
		// Check if port is on the form 22-25
		if strings.Contains(port, "-") {
			// Split on dash
			var rangeArray = strings.Split(port, "-")
			// Convert to int
			start, err := strconv.Atoi(rangeArray[0])
			if err != nil {
				return err
			}

			end, err := strconv.Atoi(rangeArray[1])
			if err != nil {
				return err
			}

			// Loop through the range
			for i := start; i <= end; i++ {
				portArray = append(portArray, i)
			}
		} else {
			// Convert to int
			realPort, err := strconv.Atoi(port)
			if err != nil {
				return err
			}

			portArray = append(portArray, realPort)
		}
	}

	// Print parsed results
	fmt.Println("Parsed ports: ", portArray)

	// Create channel to pass the addresses
	out := make(chan string)

	var rangeArray = strings.Split(hosts, "/")
	var ip = net.ParseIP(rangeArray[0])
	_, subnet, err := net.ParseCIDR(hosts)
	if err != nil {
		return err
	}

	var count int = 0;
	for ip := ip.Mask(subnet.Mask); subnet.Contains(ip); inc(ip) {
		for range portArray {
			count++
		}
	}


	// Start the GoRoutines
	if count < number_of_routines {
		number_of_routines = count
	}

	for i := 0; i < number_of_routines; i++ {
		go scanPort(out)
	}


	// https://stackoverflow.com/questions/67788289/iterating-over-a-multiline-variable-to-return-all-ips
	for ip := ip.Mask(subnet.Mask); subnet.Contains(ip); inc(ip) {
		var host = ip.String()
		for _, port := range portArray {
			out <- host + ":" + fmt.Sprint(port)
		}
	}

	return nil
}

func main() {
	err := parse_arguments()
	if err != nil {
		log.Fatal(err)
	}
}
