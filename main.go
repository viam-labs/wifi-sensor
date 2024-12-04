// package main is a module with a linux wifi sensor component
package main

import (
	"github.com/viam-labs/wifi-sensor/linuxwifi"
	"go.viam.com/rdk/components/sensor"
	"go.viam.com/rdk/module"
	"go.viam.com/rdk/resource"
)

func main() {
	module.ModularMain(resource.APIModel{API: sensor.API, Model: linuxwifi.Model})
}
