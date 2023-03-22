package netw

import "github.com/xpp777/game/iface"

var (
	config *iface.Config
)

func SetConfig(c *iface.Config) {
	config = c
}
