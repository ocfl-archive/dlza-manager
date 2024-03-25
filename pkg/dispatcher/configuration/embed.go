package configuration

import (
	"embed"
)

//go:embed config.toml
var ConfigFS embed.FS
