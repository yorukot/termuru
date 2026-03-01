package variable

import (
	"path/filepath"

	"github.com/adrg/xdg"
)

var (
	HomeDir           = xdg.Home
	MainDir  = filepath.Join(xdg.ConfigHome, AppName)
	FileCacheDir = filepath.Join(xdg.CacheHome, AppName)
	FileDataDir  = filepath.Join(xdg.DataHome, AppName)
	FileStateDir = filepath.Join(xdg.StateHome, AppName)

	SSHConfigFile = filepath.Join(FileDataDir, "ssh_config.toml")
)
