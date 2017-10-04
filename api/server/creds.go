package server

import (
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/libopenstorage/openstorage/api"
	"github.com/libopenstorage/openstorage/creds"
)

type credAPI struct {
	restBase
}

func newCredAPI(name string) restServer {
	return &credAPI{restBase{version: creds.APIVersion, name: "Credential API"}}
}

func (c *credAPI) String() string {
	return c.name
}

func (c *credAPI) create(w http.ResponseWriter, r *http.Request) {
	method := "create"
	var credential creds.CredentialEntry

	if err := json.NewDecoder(r.Body).Decode(&credential); err != nil {
		c.sendError(c.name, method, w, err.Error(), http.StatusBadRequest)
		return
	}

	cm, err := creds.Inst()
	if err != nil {
		c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = cm.Create(credential)
	if err != nil {
		c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}

	credResponse := &api.CredentialResponse{Error: err.Error()}
	json.NewEncoder(w).Encode(credResponse)
}

func (c *credAPI) list(w http.ResponseWriter, r *http.Request) {
	method := "list"

	cm, err := creds.Inst()
	if err != nil {
		c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}

	credentials, err := cm.List()

	if err != nil {
		c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(credentials)
}

func (c *credAPI) update(w http.ResponseWriter, r *http.Request) {
	method := "update"
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	var credential creds.CredentialEntry

	if err := json.NewDecoder(r.Body).Decode(&credential); err != nil {
		c.sendError(c.name, method, w, err.Error(), http.StatusBadRequest)
		return
	}

	cm, err := creds.Inst()

	if err != nil {
		c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = cm.Update(uuid, credential)

	if err != nil {
		c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}

	credResponse := &api.CredentialResponse{Error: err.Error()}
	json.NewEncoder(w).Encode(credResponse)
}

func (c *credAPI) delete(w http.ResponseWriter, r *http.Request) {
	method := "delete"
	vars := mux.Vars(r)
	uuid := vars["uuid"]

	cm, err := creds.Inst()
	if err != nil {
		c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = cm.Delete(uuid)
	if err != nil {
		c.sendError(c.name, method, w, err.Error(), http.StatusInternalServerError)
		return
	}

	credResponse := &api.CredentialResponse{Error: err.Error()}
	json.NewEncoder(w).Encode(credResponse)
}

func credVersion(route, version string) string {
	return "/" + version + "/" + route
}

func credPath(route, version string) string {
	return credVersion("creds"+route, version)
}

func (c *credAPI) Routes() []*Route {
	return []*Route{
		{verb: "GET", path: credPath("/listcreds", creds.APIVersion), fn: c.list},
		{verb: "PUT", path: credPath("/addcreds", creds.APIVersion), fn: c.create},
		{verb: "DELETE", path: credPath("/deletecreds", creds.APIVersion), fn: c.delete},
		{verb: "POST", path: credPath("/updatecreds", creds.APIVersion), fn: c.update},
	}
}