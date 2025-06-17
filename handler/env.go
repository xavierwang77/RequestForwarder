package handler

import (
	"OffMetaCore/cmn/log"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func Init() {
	logger = log.GetLogger()
}
