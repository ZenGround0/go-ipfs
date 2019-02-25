package plugin

import (
	coreiface "gx/ipfs/QmNk3oZXh51B2Es72TTbLuHorvUL3487DGDZ64YjjubvGz/interface-go-ipfs-core"
)

// PluginDaemon is an interface for daemon plugins. These plugins will be run on
// the daemon and will be given access to an implementation of the CoreAPI.
type PluginDaemon interface {
	Plugin

	Start(coreiface.CoreAPI) error
	Close() error
}
