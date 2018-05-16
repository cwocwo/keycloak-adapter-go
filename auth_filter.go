package keycloak


import (
	"github.com/emicklei/go-restful"
 	"log"
	//"encoding/json"

	//"github.com/SermoDigital/jose/jwt"
	"github.com/euforia/keycloak-client"

	//"io/ioutil"

)

const (
	UserVar = "_username_"
    RoleVar = "_roles_"
)

// basic auth Filter
func BasicAuthFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	log.Printf("[global-filter (logger)] %s,%s\n", req.Request.Method, req.Request.URL)
	chain.ProcessFilter(req, resp)

}


// basic auth Filter
func TokenAuthFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	log.Printf("[token-auth-filter (logger)] %s,%s\n", req.Request.Method, req.Request.URL)
	if Config != nil {
		var kc, err = keycloak.NewHttpKeycloakAuthenticatorFromConfig(Config)

		if err == nil {
			keycloakJWT, roles, err := kc.AuthenticateRequest(req.Request)
			if err == nil {
				req.SetAttribute(UserVar, keycloakJWT.GetUsername())
				req.SetAttribute(RoleVar, roles)


				log.Printf("[token-auth-filter (logger)] %s,%s\n", keycloakJWT.GetUsername(), roles)
			} else {
				log.Printf("[token-auth-filter (logger)] %s\n", err)
			}
		}
	}

	chain.ProcessFilter(req, resp)
}

func loadConfig() {
	//config, err := ioutil.ReadFile("./keycloak-config.json")
	//if err != nil {
	//	log.Fatal(err)
	//}
}
