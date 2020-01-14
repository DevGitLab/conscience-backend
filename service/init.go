package service

import (
	"github.com/snlansky/glibs/logging"
)

var logger = logging.MustGetLogger("service")

var DefaultFaceIDService *FaceIDService

func Init() {
	DefaultFaceIDService = NewFaceIDService()
}
