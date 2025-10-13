package api

import (
	"air-trail-backend/utils/env"
)

var (
	BASE_URL      = env.GetEnv("BASE_URL", "")
	BASE_URL_BDG  = env.GetEnv("BASE_URL_BDG", "")
	BASE_URL_NODE = env.GetEnv("BASE_URL_NODE", "")
	JETPHOTOS_URL = env.GetEnv("JETPHOTOS_URL", "")
)
