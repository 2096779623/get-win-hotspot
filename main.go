package main

import (
	"fmt"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func main() {
	var reg, err = registry.OpenKey(registry.LOCAL_MACHINE, `SYSTEM\CurrentControlSet\Services\icssvc\Settings`, registry.QUERY_VALUE|registry.BINARY)
	val, _, err := reg.GetBinaryValue("PrivateConnectionSettings")
	if err != nil {
		fmt.Println(err)
	}
	vals := string(val)
	rem := strings.Replace(vals, "", "", -1)
	fmt.Print(rem)
}
