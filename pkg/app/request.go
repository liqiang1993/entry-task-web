package app

import (
	"github.com/astaxie/beego/validation"

	"github.com/lucky-cheerful-man/phoenix_gateway/pkg/log"
)

// MarkErrors logs error logs
func MarkErrors(requestID string, errors []*validation.Error) {
	for _, err := range errors {
		log.Infof("request:%s, err.key:%s, err.message:%s ", requestID, err.Key, err.Message)
	}
}
