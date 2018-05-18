package keycloak

import (
	"github.com/emicklei/go-restful"
	"log"
	"github.com/euforia/crud-rbac"
)

func RbacAuthFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	log.Printf("[rbac-auth-filter] %s,%s\n", req.Request.Method, req.Request.URL)
	username := req.Attribute(UserVar)

	//TODO anonymous support

    if username == nil {
		log.Println("[rbac-auth-filter] user not login!")
		// TODO render error
	}

	rolesInterface := req.Attribute(RoleVar)
	if rolesInterface == nil {
		log.Println("[rbac-auth-filter] user' roles not exists!")
		// TODO render error
	} else {
		roles := rolesInterface.([]crudrbac.Role)
		for _, role := range roles {
			log.Printf("[rbac-auth-filter] %s\n", role.Name)
		}
	}



	uri := req.Request.RequestURI



	chain.ProcessFilter(req, resp)
}