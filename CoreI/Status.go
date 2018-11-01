package CoreI

import (
	"v2ray.com/core"
	"os"
	"strconv"
)

type Status struct {
	IsRunning       bool
	VpnSupportnodup bool
	PackageName     string

	Vpoint core.Server
}

func (v *Status) GetDataDir() string {
	return v.getDataDir()
}

func (v *Status) getDataDir() string {
	var datadir = "/data/data/com.v2ray.ang/"
	if v.PackageName != "" {
		datadir = "/data/user/" + strconv.Itoa(os.Getuid()/100000) + "/" + v.PackageName + "/"
	}
	return datadir
}