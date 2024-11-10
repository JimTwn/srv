package main

import "fmt"

// Some version constants.
const (
	AppVendor  = "Jim Teeuwen <jimteeuwen@proton.me>"
	AppName    = "srv"
	AppVersion = "v1.0.0"
)

// Version returns the version string.
func Version() string {
	return fmt.Sprintf("%s %s %s", AppVendor, AppName, AppVersion)
}
