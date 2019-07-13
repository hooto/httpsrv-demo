package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/hooto/hlog4g/hlog"
	"github.com/lessos/lessgo/encoding/json"
)

var (
	Prefix  string
	Config  ConfigCommon
	AppName = "httpsrv-demo"
	Version = "0.0.1"
	Release = "1"
	err     error
)

type ConfigCommon struct {
	HttpPort     uint16 `json:"http_port"`
	HttpBasePath string `json:"http_base_path,omitempty"`
}

func Setup() error {

	Prefix, err = filepath.Abs(filepath.Dir(os.Args[0]) + "/..")
	hlog.Printf("info", "project prefix path %s", Prefix)

	if err := json.DecodeFile(Prefix+"/etc/config.json", &Config); err != nil &&
		!os.IsNotExist(err) {
		return err
	}

	if Config.HttpPort == 0 {
		Config.HttpPort = 8080
	}

	Config.HttpBasePath = strings.Trim(filepath.Clean(Config.HttpBasePath), ".")
	if !strings.HasSuffix(Config.HttpBasePath, "/") {
		Config.HttpBasePath += "/"
	}

	return Flush()
}

func Flush() error {
	return json.EncodeToFile(Config, Prefix+"/etc/config.json", "  ")
}
