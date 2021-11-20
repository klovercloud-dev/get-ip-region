package ipLocationService

import (
	"encoding/json"
	"fmt"
)


var ipBlocks IpBlocks

func LoadData(key string)  {
	ipBlocksString, _ := RedisClient.Get(key).Bytes()

	json.Unmarshal(ipBlocksString, &ipBlocks)

	fmt.Println("Data Loaded")

	//Sort(ipBlocks.Values)
	//
	//fmt.Println("Data Sorted")

}

func GetSortedIpBlocks(key string)[]IpBlock  {
	if ipBlocks.Values == nil{
		LoadData(key)
	}
	return ipBlocks.Values
}
//
//
//func GetIpBlocks(client *redis.Client, key string) []byte {
//	blocks, _ := client.Get(key).Bytes()
//	return blocks
//}
//
//func SetIpBlocks(client *redis.Client, key string, value string)  {
//	client.Set(key, value, 0)
//}
//
//func GetSortedIpBlocks(client *redis.Client, key string)[]IpBlock  {
//	var ipBlocks IpBlocks
//
//	ipBlocksString, _ := client.Get(key).Bytes()
//
//	json.Unmarshal(ipBlocksString, &ipBlocks)
//
//	return ipBlocks.Values
//}
//
//func InitData(client *redis.Client, key string)  {
//	var ipBlocks IpBlocks
//
//	ipBlocksString, _ := client.Get(key).Bytes()
//
//	json.Unmarshal(ipBlocksString, &ipBlocks)
//
//	AddIpRange(ipBlocks.Values)
//
//	Sort(ipBlocks.Values)
//
//	value, _ := json.Marshal(ipBlocks)
//
//	SetIpBlocks(client, key, string(value))
//}