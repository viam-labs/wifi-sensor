package linuxwifi

import (
	"os/exec"
	"regexp"
	"strings"
)

var ssidRegex = regexp.MustCompile(`SSID (\d+) : (.+)$`)
var fieldRegex = regexp.MustCompile(`^\s+([^:]+)\s*:\s+(.+)\s*$`)

func platformReadings(_ string) (map[string]interface{}, error) {
	out, err := exec.Command("netsh", "wlan", "show", "networks").Output()
	if err != nil {
		return nil, err
	}
	ret := make(map[string]interface{})
	headerFound := false
	for _, line := range strings.Split(string(out), "\n") {
		if !headerFound {
			if found := ssidRegex.FindStringSubmatch(line); found != nil {
				headerFound = true
				ret["ssid"] = found[2]
			}
		} else {
			if found := fieldRegex.FindStringSubmatch(line); found != nil {
				ret[found[1]] = found[2]
			} else {
				break
			}
		}
	}
	return ret, nil
}
