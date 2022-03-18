// Package bootstrap
package bootstrap

import (
	"github.com/orn-id/wedding-api/src/consts"
	"github.com/orn-id/wedding-api/src/pkg/logger"
	"github.com/orn-id/wedding-api/src/pkg/msg"
)

func RegistryMessage()  {
	err := msg.Setup("msg.yaml", consts.CONFIG_PATH)
	if err != nil {
		logger.Fatal(logger.SetMessageFormat("file message multi language load error %s", err.Error()))
	}

}
