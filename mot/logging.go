package mot

import (
	"github.com/motclub/common/logging"
	"github.com/sirupsen/logrus"
)

func init() {
	logging.AddHandler(func(entry *logrus.Entry) error {
		return nil
	})
}

func PRINT(v ...interface{}) {
	logging.PRINT(v...)
}

func DEBUG(v ...interface{}) {
	logging.DEBUG(v...)
}

func WARN(v ...interface{}) {
	logging.WARN(v...)
}

func INFO(v ...interface{}) {
	logging.INFO(v...)
}

func ERROR(v ...interface{}) {
	logging.ERROR(v...)
}

func FATAL(v ...interface{}) {
	logging.FATAL(v...)
}

func PANIC(v ...interface{}) {
	logging.PANIC(v...)
}
