package api

import (
	"air-trail-backend/utils/env"
)

var (
	BASE_URL      = env.GetEnv("BASE_URL", "")
	JETPHOTOS_URL = env.GetEnv("JETPHOTOS_URL", "")
)
