
//This program uses the gopacket library to capture tcp packets that are
//destined to port 80 and displays the contents.


package main

import (
	"fmt"
	"log"
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
	"github.com/google/gopacket/layers"
	"os"

)

var (
	iface = "en0"
	snaplen int32 = 1024
	promisc = false
	timeout = pcap.BlockForever
	filter = "tcp and port 80"
	devFound = false
	nPackets = 100
)

func main(){

	f, _ := os.Create("output.pcap")
	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(uint32(snaplen), layers.LinkTypeEthernet)
	defer f.Close()

	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	for _, device := range devices{
		if device.Name == iface{
			devFound = true
		}
	}
	if !devFound{
		log.Panicf("Device %s does not exist.\n", iface)
	}

	handle, err := pcap.OpenLive(iface,snaplen,promisc,timeout)
	if err != nil{
		log.Panicln(err)
	}

	source := gopacket.NewPacketSource(handle,handle.LinkType())
	count := 0
	for packet := range source.Packets(){
	//	printPacketInfo(packet)
	//	w.WritePacket(packet.Metadata().CaptureInfo, packet.Data())
	fmt.Println(packet)
		count++
		if count >= nPackets{
			break
		}
	}

}

func printPacketInfo(packet gopacket.Packet){
	etherNetLayer := packet.Layer(layers.LayerTypeEthernet)
	if etherNetLayer != nil {
		fmt.Println("Ethernet layer detected.")
		etherNetPacket, _ := etherNetLayer.(*layers.Ethernet)
		fmt.Println("Source MAC: ", etherNetPacket.SrcMAC)
		fmt.Println("Destination MAC: ", etherNetPacket.DstMAC)
		fmt.Println("Ethernet type: ", etherNetPacket.EthernetType)
		fmt.Println()
	}

	tcpLayer := packet.Layer(layers.LayerTypeTCP)
	if tcpLayer != nil{
		tcpPacket, _ := tcpLayer.(*layers.TCP)
		fmt.Println("TCP packet detected.")
		fmt.Printf("\nSrcPrt: %d, DstPort: %d\n", tcpPacket.SrcPort, tcpPacket.DstPort)
		fmt.Println("Sequence number: ", tcpPacket.Seq)
		fmt.Println()
	}
}
