package main

import (
	"flag"
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"net"
	"os"
	"time"
)

func usage() {
	fmt.Printf("Usage: %s -file <pcap-file> [-port <listening-port>] [-proto tcp|udp] [-bpf <bpf-filter>] [-delay <packet-delay-ms>] [-wait <wait-flag-toggle>]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	port := flag.String("port", "8484", "Layer 4 listening port")
	file := flag.String("file", "", "PCAP file to replay")
	proto := flag.String("proto", "tcp", "Layer 4 protocol to use")
	bpf := flag.String("bpf", "", "BPF filter")
	delay := flag.Int("delay", 100, "Delay between sent packets in milliseconds")
	wait := flag.Bool("wait", true, "Wait for first packet before starting replay")

	flag.Parse()

	if *file == "" || (*proto != "tcp" && *proto != "udp") {
		usage()
	}

	ln, err := net.Listen(*proto, fmt.Sprintf(":%s", *port))
	if err != nil {
		panic(err)
	}

	packets := loadPcap(file, bpf)

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go sendPackets(conn, packets, delay, wait)
	}
}

// Pcap files can be filtered with bpf to localize packets
// going to or from a certain direction, host, or port.
// See http://biot.com/capstats/bpf.html for the syntax details
func loadPcap(file *string, bpf *string) []gopacket.Packet {
	hpcap, err := pcap.OpenOffline(*file)
	if err != nil {
		panic(err)
	}

	if *bpf != "" {
		err = hpcap.SetBPFFilter(*bpf)
		if err != nil {
			panic(err)
		}
	}

	packets := make([]gopacket.Packet, 0)
	psource := gopacket.NewPacketSource(hpcap, hpcap.LinkType())
	for packet := range psource.Packets() {
		packets = append(packets, packet)
	}

	return packets
}

// This goroutine passes the packets once to the connection and intentionally
// leaves the connection hanging to ensure that the client receives all the data
func sendPackets(conn net.Conn, packets []gopacket.Packet, delay *int, wait *bool) {
	if *wait {
		tmp := make([]byte, 1)
		conn.Read(tmp)
	}

	for _, packet := range packets {
		time.Sleep(time.Duration(*delay) * time.Millisecond)
		conn.Write(packet.ApplicationLayer().Payload())
	}
}
