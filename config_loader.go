package keycloak

import (
	"github.com/euforia/keycloak-client"
)

var Config, _ = keycloak.LoadClientConfig("./keycloak-config.json")
