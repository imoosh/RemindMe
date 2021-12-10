package response

import "RemindMe/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
