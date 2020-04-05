package main

import "github.com/grill-tamasi/wscgo/plugins"

func GetAddons() []plugins.Addon {
	return []plugins.Addon{
		&mcp23017addon{},
		&pca9685addon{},
	}
}
