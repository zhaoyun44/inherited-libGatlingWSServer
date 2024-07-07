package modDataPackage

import "net/http"

type cWJWSRouter struct {
	handler          map[string]http.HandlerFunc
	homeHandler      http.HandlerFunc
	notFoundHandler  http.HandlerFunc
	upgradeRouterKey string
	server           *CGatlingWSServer
}

func newWJWSRouter(serverInst *CGatlingWSServer) *cWJWSRouter {
	return &cWJWSRouter{handler: make(map[string]http.HandlerFunc),
		homeHandler:      pageEmpty,
		notFoundHandler:  pageEmpty,
		upgradeRouterKey: "/ws",
		server:           serverInst,
	}
}

func (pInst *cWJWSRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		pInst.homeHandler(w, r)
		return
	}
	if r.URL.Path == pInst.upgradeRouterKey {
		pInst.server.Upgrade(w, r)
	}
	http.NotFound(w, r)
}

func (pInst *cWJWSRouter) HandlerFunc(pattern string, fn http.HandlerFunc) {
	newMap := make(map[string]http.HandlerFunc)
	for key, value := range pInst.handler {
		newMap[key] = value
	}
	newMap[pattern] = fn
	pInst.handler = newMap
}

/////////////////////////////// default function for empty call /////////////////////////////

func pageEmpty(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("not found"))
}