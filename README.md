# Concurrent Port Scanner

## Overview
The Concurrent Port Scanner is a high-performance tool written in Go that allows users to scan a range of IP addresses and ports to check for open ports. It utilizes goroutines to handle scanning in parallel, significantly speeding up the scanning process. This tool is designed to be flexible and efficient, using a user-defined number of concurrent goroutines based on available system resources.

## Features
- Concurrent scanning with configurable number of goroutines.
- Supports single ports, lists of ports, and ranges.
- Efficient IP address iteration for CIDR notations.
- Detailed output of open and closed port status.

## Prerequisites
Before you start using the Concurrent Port Scanner, you need to have Go installed on your machine. You can download and install Go from [here](https://golang.org/dl/).

## Installation
Clone the repository or download the source code to your local machine:
```bash
git clone https://github.com/mrmrjing/concurrent-port-scanner.git
```
To build the project, run: 
```bash
go build
```
This command will compile the source code and create an executable file in the same directory 

## Usage 
Run the scanner using the following command: 
```bash
./ConcurrentPortScanner -p [port_ranges] [hosts]
```
`[port_ranges]` should be specified in the format 22-25,80,8080 to represent ranges and specific ports.
`[hosts]` can be a single IP, multiple IPs separated by commas, or a CIDR notation.

### Example
To scan ports 22 to 25 and port 80 on all IPs in the subnet 192.168.0.1/24: 
```bash 
./ConcurrentPortScanner -p 22-25,80 192.168.0.1/24
```

## Output 
The program will output the status of each port on each IP address, indicating whether it is open or closed 
