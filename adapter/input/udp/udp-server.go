package udp

import (
	"encoding/json"
	"fmt"
	"github.com/frederico-neres/my-own-log-management--litelogs/application/domain"
	"github.com/frederico-neres/my-own-log-management--litelogs/application/port"
	"log"
	"net"
	"strings"
)

type LogReceived struct {
	Level   int    `json:"level"`
	Service string `json:"service"`
	Message string `json:"message"`
}

func (l *LogReceived) ToDomain() *domain.LogDomain {
	return domain.NewLogDomain(l.Level, l.Service, l.Message)
}

type UdpServer struct {
	SaveLogServicePort port.SaveLogServicePort
}

func NewUdpServer(saveLogServicePort port.SaveLogServicePort) *UdpServer {
	return &UdpServer{
		SaveLogServicePort: saveLogServicePort,
	}
}

func (u *UdpServer) Listen() {
	udpServer, err := net.ListenPacket("udp", ":1053")
	if err != nil {
		log.Fatal(err)
	}
	defer udpServer.Close()

	for {
		buf := make([]byte, 1024)
		_, addr, err := udpServer.ReadFrom(buf)
		if err != nil {
			continue
		}
		u.received(udpServer, addr, buf)
	}

}

func (u *UdpServer) received(udpServer net.PacketConn, addr net.Addr, buf []byte) {
	receivedStr := strings.Replace(string(buf), "\x00", "", -1)
	var logReceived LogReceived
	err := json.Unmarshal([]byte(receivedStr), &logReceived)
	if err != nil {
		fmt.Println(err)
		udpServer.WriteTo([]byte(`{"status":0}`), addr)
	}

	udpServer.WriteTo([]byte(receivedStr), addr)
	err = u.SaveLogServicePort.Exec(logReceived.ToDomain())
	if err != nil {
		fmt.Println(err)
	}
}
