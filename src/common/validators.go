package common

import (
	"github.com/MrRezoo/CarApp/config"
	"github.com/MrRezoo/CarApp/pkg/logging"
	"regexp"
)

const MobilePattern = `^(\+98|0)?9\d{9}$`

var logger = logging.NewLogger(config.GetConfig())

func ValidateMobile(mobileNumber string) bool {
	res, err := regexp.MatchString(MobilePattern, mobileNumber)
	if err != nil {
		logger.Error(logging.Validation, logging.MobileValidation, err.Error(), nil)
	}
	return res
}
