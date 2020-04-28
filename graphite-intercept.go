package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	var listenAddr = flag.String("listen", "localhost:0", "TCP/UDP on which to listen")
	var targetAddr = flag.String("target", "localhost:2003", "Target address and port")
	var fileName = flag.String("file", "metrics.log", "File in which to write the metrics")

	flag.Parse()

	metricChannel := make(chan string, 10)

	listener, err := net.Listen("tcp", *listenAddr)
	if err != nil {
		log.Fatalf("Unable to open listening port : %v", err)
	}
	defer listener.Close()
	log.Infof("Listening on %s", listener.Addr().String())

	go dumpMetrics(metricChannel, *fileName)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting new connection : %v", err)
		}
		go handleConnection(conn, *targetAddr, metricChannel)
	}
}

func handleConnection(conn net.Conn, targetAddr string, c chan<- string) {
	logPrefix := fmt.Sprintf("%s %s - ", conn.RemoteAddr().Network(), conn.RemoteAddr().String())
	log.Infof("%s New connection", logPrefix)

	targetConn, err := net.Dial("tcp", targetAddr)
	if err != nil {
		log.Errorf("Unable to open connection to target : %v", err)
	} else {
		defer targetConn.Close()
	}

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		log.Infof("%s New data: %s", logPrefix, line)
		if targetConn != nil {
			fmt.Fprintf(targetConn, "%s\n", line)
		}
		c <- line
	}

	if err := conn.Close(); err != nil {
		log.Warnf("%s Problem closing connection : %v", logPrefix, err)
	}
	log.Infof("%s Connection closed", logPrefix)
}

func dumpMetrics(c <-chan string, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Panicf("Unable to open file : %v", err)
		panic(err)
	}
	defer file.Close()

	for line := range c {
		fmt.Fprintln(file, line)
	}
}
