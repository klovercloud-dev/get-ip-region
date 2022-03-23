package enums

type KEY string

const (
	REDISPORT       = KEY("REDIS_PORT")
	REDISCONNECTURL = KEY("REDIS_CONNECT_URL")
	IPBLOCKS        = KEY("ipBlocks")
)
