package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hashwing/sxiot/sxiot-uapi/api/platform"
	"github.com/hashwing/sxiot/sxiot-uapi/api/app"
)

func NewRouter(root *mux.Router) {
	apiRoute := root.PathPrefix("/api").Subrouter()
	apiRoute.HandleFunc("/login", platform.Login)
	userRoute:=apiRoute.PathPrefix("/user").Subrouter()
	userRoute.Handle("/get", apiHandler(platform.GetCurrentUser))
	userRoute.Handle("/find", apiHandler(platform.FindUsers))
	userRoute.Handle("/create", apiHandler(platform.CreateUser))
	userRoute.Handle("/del", apiHandler(platform.DelUser))
	brandRoute:=apiRoute.PathPrefix("/device/brand").Subrouter()
	brandRoute.Handle("/find",  apiHandler(platform.FindBrands))
	brandRoute.Handle("/get",  apiHandler(platform.GetBrand))
	brandRoute.Handle("/create",apiHandler(platform.CreateBrand))
	brandRoute.Handle("/update", apiHandler(platform.UpdateBrand))
	brandRoute.Handle("/delete", apiHandler(platform.DelBrand))
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
	appRoute.HandleFunc("/son", app.GetSonDevices)
	appRoute.HandleFunc("/reg", app.AddUser)
	appRoute.HandleFunc("/template", platform.FindBrands)
	appDeviceRoute:=appRoute.PathPrefix("/device").Subrouter()
	appDeviceRoute.Handle("/find",  appHandler(app.FindDevices))
	appDeviceRoute.Handle("/all",  appHandler(app.GetAllDevices))
	appDeviceRoute.Handle("/add",  appHandler(app.CreateDevice))
	appDeviceRoute.Handle("/del",  appHandler(app.DelDevice))
	appDeviceRoute.Handle("/update",  appHandler(app.UpdateDevice))
	appSonRoute:=appRoute.PathPrefix("/son").Subrouter()
	appSonRoute.Handle("/find",  appHandler(app.GetSonDevices))
	appSonRoute.Handle("/update",  appHandler(app.UpdateSonDevice))
}

type apiHandler func(http.ResponseWriter, *http.Request)

func (fn apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Access-Control-Allow-Origin","*")
	r.Header.Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
	r.Header.Set("Access-Control-Allow-Headers", "Auth");
	err := platform.Auth(w, r)
	if err != nil {
		fmt.Println(err)
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