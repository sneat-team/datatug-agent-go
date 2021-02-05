package server

import (
	"fmt"
	"github.com/datatug/datatug/packages/server/endpoints"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strings"
)

var router *httprouter.Router

func initRouter() {
	router = httprouter.New()
	handlerFunc := func(method, path string, handler http.HandlerFunc) {
		wrappedHandler := func(w http.ResponseWriter, r *http.Request) {
			log.Println(method, r.RequestURI)
			handler(w, r)
		}
		router.HandlerFunc(method, path, wrappedHandler)
	}

	router.GlobalOPTIONS = http.HandlerFunc(globalOptionsHandler)
	handlerFunc(http.MethodGet, "/ping", endpoints.Ping)

	handlerFunc(http.MethodGet, "/projects", endpoints.GetProjects)
	handlerFunc(http.MethodGet, "/project-summary", endpoints.GetProjectSummary)
	handlerFunc(http.MethodGet, "/", root)

	handlerFunc(http.MethodPost, "/dbserver-add", endpoints.AddDbServer)
	handlerFunc(http.MethodDelete, "/dbserver-delete", endpoints.DeleteDbServer)
	handlerFunc(http.MethodGet, "/dbserver-summary", endpoints.GetDbServerSummary)
	handlerFunc(http.MethodGet, "/dbserver-databases", endpoints.GetServerDatabases)
	handlerFunc(http.MethodGet, "/project-full", endpoints.GetProjectFull)
	handlerFunc(http.MethodGet, "/environment-summary", endpoints.GetEnvironmentSummary)
	handlerFunc(http.MethodPost, "/execute", endpoints.ExecuteCommandsHandler)
	handlerFunc(http.MethodGet, "/select", endpoints.ExecuteSelectHandler)

	handlerFunc(http.MethodGet, "/entities/all_entities", endpoints.GetEntities)
	handlerFunc(http.MethodGet, "/entities/entity", endpoints.GetEntity)
	handlerFunc(http.MethodPost, "/entities/create_entity", endpoints.SaveEntity)
	handlerFunc(http.MethodPut, "/entities/save_entity", endpoints.SaveEntity)
	handlerFunc(http.MethodDelete, "/entities/delete_entity", endpoints.DeleteEntity)

	handlerFunc(http.MethodGet, "/boards/board", endpoints.GetBoard)
	handlerFunc(http.MethodPost, "/boards/create_board", endpoints.CreateBoard)
	handlerFunc(http.MethodPut, "/boards/save_board", endpoints.SaveBoard)
	handlerFunc(http.MethodDelete, "/boards/delete_board", endpoints.DeleteBoard)

	handlerFunc(http.MethodGet, "/data/recordsets", endpoints.GetRecordsetsSummary)
	handlerFunc(http.MethodGet, "/data/recordset_definition", endpoints.GetRecordsetDefinition)
	handlerFunc(http.MethodGet, "/data/recordset_data", endpoints.GetRecordsetData)
	handlerFunc(http.MethodPost, "/data/recordset_add_rows", endpoints.AddRowsToRecordset)
	handlerFunc(http.MethodDelete, "/data/recordset_delete_rows", endpoints.DeleteRowsFromRecordset)
	handlerFunc(http.MethodPut, "/data/recordset_update_rows", endpoints.UpdateRowsInRecordset)
}

// globalOptionsHandler handles OPTIONS requests
func globalOptionsHandler(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	accessControlRequestMethod := r.Header.Get("Access-Control-Request-Method")
	if origin == "" || accessControlRequestMethod == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintln(w, "origin: ", origin)
		_, _ = fmt.Fprintln(w, "accessControlRequestMethod: ", accessControlRequestMethod)
		return
	}
	if !IsSupportedOrigin(origin) {
		w.WriteHeader(http.StatusForbidden)
		_, _ = fmt.Fprintf(w, "Unsupported origin: %v", origin)
		return
	}
	// Set CORS headers BEFORE calling w.WriteHeader() or w.Write()
	responseHeader := w.Header()
	responseHeader.Set("Access-Control-Allow-Origin", origin)
	responseHeader.Set("Access-Control-Allow-Methods", accessControlRequestMethod)
	accessControlRequestHeaders := r.Header.Get("Access-Control-Request-Headers")
	if accessControlRequestHeaders != "" {
		responseHeader.Set("Access-Control-Allow-Headers", accessControlRequestHeaders)
	}
	w.WriteHeader(http.StatusNoContent) // Set response status code to 204
}

// IsSupportedOrigin check provided origin is allowed
func IsSupportedOrigin(origin string) bool {
	switch origin {
	case "http://localhost:8100", "https://datatug.app":
		return true
	default:
		return strings.HasPrefix(origin, "https://") && strings.HasSuffix(origin, ".datatug.app")
	}
}
