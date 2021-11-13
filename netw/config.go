package netw

import "github.com/xiaomingping/game/iface"

var (
	config  *iface.Config
)

func SetConfig(c *iface.Config)  {
	config = c
}

