package main


import (
  "log"
	"net"
	"strconv"
	"strings"
	"regexp"
)

const (
	message       = "pcmeasure.com1.1"
	StopCharacter = "\r\n\r\n"
)

func ReadData(ip string, port int) {
	var rgx = regexp.MustCompile(`value=(.*?);`)

	addr := strings.Join([]string{ip, strconv.Itoa(port)}, ":")
	conn, err := net.Dial("tcp", addr)

	defer conn.Close()

	if err != nil {
		log.Printf("MessPC scheint offline zu sein.")
		log.Fatalln(err)
	}

	conn.Write([]byte(message))
	conn.Write([]byte(StopCharacter))
	log.Printf("Sende: %s", message)

	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	log.Printf("Antwort von MessPC")
	rs := rgx.FindAllStringSubmatch(string(buff[:n]),-1)
	log.Printf("%q", rs[0][1]);
}

func main() {
	var (
		ip   = "10.28.12.4"
		port = 4000
	)

	ReadData(ip, port)
}
