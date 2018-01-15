package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hashwing/sxito/sxito-uapi/api/platform"
)

func NewRouter(root *mux.Router) {
	apiRoute := root.PathPrefix("/api").Subrouter()
	apiRoute.HandleFunc("/login", platform.Login)
	apiRoute.Handle("/user", apiHandler(platform.GetCurrentUser))
	brandRoute:=apiRoute.PathPrefix("/device/brand").Subrouter()
	brandRoute.HandleFunc("/find",  platform.FindBrands)
	brandRoute.HandleFunc("/get",  platform.GetBrand)
	brandRoute.HandleFunc("/create",  platform.CreateBrand)
	brandRoute.HandleFunc("/update",  platform.UpdateBrand)
	brandRoute.HandleFunc("/delete",  platform.DelBrand)
	gatewayRoute:=apiRoute.PathPrefix("/device/gateway").Subrouter()
	gatewayRoute.HandleFunc("/find",  apiHandler(platform.FindGateways))
	gatewayRoute.HandleFunc("/get",  apiHandler(platform.GetGateway))
	gatewayRoute.HandleFunc("/create",  apiHandler(platform.CreateGateway))
	gatewayRoute.HandleFunc("/update",  apiHandler(platform.UpdateGateway))
	gatewayRoute.HandleFunc("/delete",  apiHandler(platform.DelGateway))
	deviceRoute:=apiRoute.PathPrefix("/device").Subrouter()
	deviceRoute.HandleFunc("/find",  apiHandler(platform.FindDevcies))
	deviceRoute.HandleFunc("/get",  apiHandler(platform.GetDevcie))
	deviceRoute.HandleFunc("/create",  apiHandler(platform.CreateDevice))
	deviceRoute.HandleFunc("/update",  apiHandler(platform.UpdateDevice))
	deviceRoute.HandleFunc("/delete",  apiHandler(platform.DelDevice))
}

type apiHandler func(http.ResponseWriter, *http.Request)

func (fn apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := platform.Auth(w, r)
	if err == nil {
		fn(w, r)
	}
}
