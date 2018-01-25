package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hashwing/sxito/sxito-uapi/api/platform"
	"github.com/hashwing/sxito/sxito-uapi/api/app"
)

func NewRouter(root *mux.Router) {
	apiRoute := root.PathPrefix("/api").Subrouter()
	apiRoute.HandleFunc("/login", platform.Login)
	apiRoute.Handle("/user", apiHandler(platform.GetCurrentUser))
	brandRoute:=apiRoute.PathPrefix("/device/brand").Subrouter()
	brandRoute.HandleFunc("/find",  apiHandler(platform.FindBrands))
	brandRoute.HandleFunc("/get",  apiHandler(platform.GetBrand))
	brandRoute.HandleFunc("/create",apiHandler(platform.CreateBrand))
	brandRoute.HandleFunc("/update", apiHandler(platform.UpdateBrand))
	brandRoute.HandleFunc("/delete", apiHandler(platform.DelBrand))
	gatewayRoute:=apiRoute.PathPrefix("/device/gateway").Subrouter()
	gatewayRoute.Handle("/find",  apiHandler(platform.FindGateways))
	gatewayRoute.Handle("/get",  apiHandler(platform.GetGateway))
	gatewayRoute.Handle("/create",  apiHandler(platform.CreateGateway))
	gatewayRoute.Handle("/update",  apiHandler(platform.UpdateGateway))
	gatewayRoute.Handle("/delete",  apiHandler(platform.DelGateway))
	deviceRoute:=apiRoute.PathPrefix("/device").Subrouter()
	deviceRoute.Handle("/find",  apiHandler(platform.FindDevcies))
	deviceRoute.Handle("/get",  apiHandler(platform.GetDevcie))
	deviceRoute.Handle("/create",  apiHandler(platform.CreateDevice))
	deviceRoute.Handle("/update",  apiHandler(platform.UpdateDevice))
	deviceRoute.Handle("/delete",  apiHandler(platform.DelDevice))
	//
	appRoute:=apiRoute.PathPrefix("/app").Subrouter()
	appRoute.HandleFunc("/login", app.Login)
	appRoute.HandleFunc("/reg", app.AddUser)
	appDeviceRoute:=appRoute.PathPrefix("/device").Subrouter()
	appDeviceRoute.Handle("/find",  appHandler(app.FindDevices))
	appDeviceRoute.Handle("/add",  appHandler(app.CreateDevice))
}

type apiHandler func(http.ResponseWriter, *http.Request)

func (fn apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Access-Control-Allow-Origin","*")
	r.Header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
	r.Header.Set("Access-Control-Allow-Headers", "Auth");
	err := platform.Auth(w, r)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	fn(w, r)
}

type appHandler func(http.ResponseWriter, *http.Request)

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Access-Control-Allow-Origin","*")
	r.Header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
	r.Header.Set("Access-Control-Allow-Headers", "Auth");
	err := app.Auth(w, r)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	fn(w, r)
}