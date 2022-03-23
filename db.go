package ipLocationService

import (
	"encoding/json"
	"fmt"
)

var ipBlocks IpBlocks

func LoadData(key string) {
	ipBlocksString, _ := RedisClient.Get(key).Bytes()
	json.Unmarshal(ipBlocksString, &ipBlocks)
	fmt.Println("Data Loaded")
}

func GetSortedIpBlocks(key string) []IpBlock {
	if ipBlocks.Values == nil {
		LoadData(key)
	}
	return ipBlocks.Values
}
