package utils

import "github.com/nbittich/factsfood/types"

func GetAlertClassKey(t types.MessageType) string {
	switch t {
	case types.INFO:
		return "alert-info"
	case types.SUCCESS:
		return "alert-success"
	case types.ERROR:
		return "alert-danger"
	case types.WARNING:
		return "alert-warning"
	}
	return ""
}
