package meta

import "github.com/wonderivan/logger"

var Version = "1.0.0"
var ValidConfigVersion = []string{"1.0.0", ""}

func init() {
	logger.Info("PasteMe Go Account Version \"%s\"", Version)
}
