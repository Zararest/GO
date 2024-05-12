package handler

import (
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

//go:generate go run github.com/golang/mock/mockgen -source=handler.go -destination=internal/mock_file_service/file_service.go --package=mock_file_service

// fileService interface.
type fileService interface {
	Get(string) ([]byte, error)
	Upload(string, []byte) error
	Delete(string) error
}

// handler using for create new mux router. handler contains service for operations.
type handler struct {
	service fileService
}

// NewHandler return mux.Router with HandlerFunc mounts.
func NewHandler(r *mux.Router, service fileService) *mux.Router {
	h := handler{
		service: service,
	}

	r.HandleFunc("/ping", h.ping).Methods(http.MethodOptions, http.MethodGet)
	r.HandleFunc("/get/{name}", h.get).Methods(http.MethodOptions, http.MethodGet)
	r.HandleFunc("/upload", h.upload).Methods(http.MethodOptions, http.MethodPost)
	r.HandleFunc("/delete", h.delete).Methods(http.MethodOptions, http.MethodDelete)

	return r
}

// ping returns pong (health).
//
// @Summary assign ping returns pong (health).
// @Router /ping [get]
// @Tags healt
// @Success 200
func (h *handler) ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

// get returns file by name from mux.Vars.
//
// @Summary eturns file by name from mux.Vars.
// @Router /get/{name} [get]
// @Tags read
// @Param name path string true "file name"
// @Success 200
// @Failure 400 {object} string
// @Failure 408 {object} string
func (h *handler) get(w http.ResponseWriter, r *http.Request) {
	// Parse mux vars.
	vars := mux.Vars(r)
	filename := vars["name"]

	// Get file from system using service.
	data, err := h.service.Get(filename)
	if err != nil {
		// Wrap error.
		code := getMappedStatusCode(errorsGetMap, err)
		http.Error(w, err.Error(), code)
		return
	}
	w.Write(data)
}

// upload add file into base dir.
func (h *handler) upload(w http.ResponseWriter, r *http.Request) {
	// Read form file from request.
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, fmt.Errorf("error while get form file: %w", err).Error(), http.StatusBadRequest)
		return
	}

	// Get name of file and read file data.
	name := filepath.Base(header.Filename)
	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, fmt.Errorf("error while read form file: %w", err).Error(), http.StatusBadRequest)
		return
	}

	// Upload file in base dir using service.
	if err := h.service.Upload(name, data); err != nil {
		// Wrap error.
		code := getMappedStatusCode(errorsUploadMap, err)
		http.Error(w, err.Error(), code)
		return
	}
}

// delete removes file from base dir by url query param.
//
// @Summary removes file from base dir by url query param.
// @Router /delete [delete]
// @Tags write
// @Param name query string true "file name"
// @Success 200
// @Failure 404 {object} string
func (h *handler) delete(w http.ResponseWriter, r *http.Request) {
	// Get name from url query.
	name := r.URL.Query().Get("name")

	// Delete file from system using service.
	err := h.service.Delete(name)
	if err != nil {
		// Wrap error.
		code := getMappedStatusCode(errorsDeleteMap, err)
		http.Error(w, err.Error(), code)
		return
	}
}
