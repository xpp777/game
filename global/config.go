package global

import (
	"github.com/xiaomingping/game/config"
	"github.com/xiaomingping/game/iface"
)

var (
	Config *config.WebSocketConfig
	Server iface.Server
	Logo   string
)

func init() {
	Logo = `
.------..------..------..------..------..------..------.
|G.--. ||A.--. ||M.--. ||E.--. ||-.--. ||W.--. ||S.--. |
| :/\: || (\/) || (\/) || (\/) || (\/) || :/\: || :/\: |
| :\/: || :\/: || :\/: || :\/: || :\/: || :\/: || :\/: |
| '--'G|| '--'A|| '--'M|| '--'E|| '--'-|| '--'W|| '--'S|
.------..------..------..------..------..------..------.`
}
