package ipLocationService

import (
	"bytes"
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"
)

func AddIpRange(ipBlocks []IpBlock)  {
	var temp IpRange
	for i:=0; i<len(ipBlocks); i++{
		temp = GetIpRange(ipBlocks[i].Cidr)
		ipBlocks[i].FirstHost = temp.First
		ipBlocks[i].LastHost = temp.Last
	}
}

func Sort(ipBlock []IpBlock){
	var temp IpBlock
	var i int
	var j int
	length := len(ipBlock)
	for i =0; i<length; i++{
		for j = 0; j< length-1 ; j++{
			if whichBlockBig(ipBlock[j], ipBlock[j+1]) == "first"{
				temp = ipBlock[j]
				ipBlock[j] = ipBlock[j+1]
				ipBlock[j+1] = temp
			}
		}
	}
}

func whichBlockBig(a IpBlock, b IpBlock) string {
	if bytes.Compare(a.LastHost, b.FirstHost) == 1{
		return "first"
	}else {
		return "second"
	}
}

func GetCountry(ip net.IP) string {
	blocks := GetSortedIpBlocks("ipblocks")
	size := len(blocks)

	index := binarySearch(blocks, 0, int64(size-1), ip)

	if index == -1{
		return ""
	}else {
		return blocks[index].Country
	}
}

func binarySearch(blocks []IpBlock, start int64, end int64, ip net.IP) int64 {
	mid := int64((end + start)/2)

	if start> mid{
		return -1
	}

	if IsIpBetween(ip, blocks[mid].Cidr) == true {
		return  mid
	}else if bytes.Compare(ip, blocks[mid].LastHost) == 1{
		return binarySearch(blocks, mid+1, end, ip)
	}else{
		return 	binarySearch(blocks, start, mid, ip)
	}
}

func dupIP(ip net.IP) net.IP {
	dup := make(net.IP, len(ip))
	copy(dup, ip)
	return dup
}

func GetIpRange(cidr string) IpRange{
	var ipRange IpRange

	cidrParts := strings.Split(cidr, "/")

	ip := net.ParseIP(cidrParts[0])
	ipRange.First = dupIP(ip)

	prefix, _ := strconv.Atoi(cidrParts[1])

	numberOfBlocksToManipulate := math.Ceil(float64(32 - prefix)/float64(8))

	numberOfBit := (32 - prefix) % 8
	condition := 15 - int(numberOfBlocksToManipulate) + 1

	for i :=  0; i <int(numberOfBlocksToManipulate); i++ {
		temp := 15 - i
		if temp == condition{
			ip[temp] = byte(GetHighestRange(ip[temp], numberOfBit))
		}else{
			ip[temp] = 255
		}
	}
	ipRange.Last = ip
	//fmt.Println(ipRange.first)
	return ipRange
}

func IsIpBetween(current net.IP, cidr string) bool {
	var ipRange IpRange

	ipRange = GetIpRange(cidr)

	if current.To4() == nil {
		fmt.Printf("%v is not an IPv4 address\n", current)
		return false
	}
	if bytes.Compare(current, ipRange.First) >= 0 && bytes.Compare(current, ipRange.Last) <= 0 {
		fmt.Printf("%v is between %v and %v\n", current, ipRange.First, ipRange.Last)
		return true
	}
	fmt.Printf("%v is NOT between %v and %v\n", current, ipRange.First, ipRange.Last)
	return false
}

func AdjustLength(bin string) string{
	count := len(bin)
	temp := ""

	for i:= 0; i< (8-count); i++{
		temp += "0"
	}
	return temp + bin
}

func GetHighestRange(block uint8, index int) uint64 {
	bin := strconv.FormatUint(uint64(block), 2)
	bin = AdjustLength(bin)
	if(len(bin)< 8){

	}
	var modifiedBin string

	// fmt.Println("bin: ",bin)
	for i := 0; i<8; i++{
		if i >= 8 - index {
			modifiedBin += "1"
		}else{
			modifiedBin += string([]rune(bin)[i])
		}
	}

	// fmt.Println("Modified Bin: ", modifiedBin)

	out, _ := strconv.ParseUint(modifiedBin, 2, 8)

	return out
}
