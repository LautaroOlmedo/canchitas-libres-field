package web

import (
	"canchitas-libres-field/internal/pkg/domain"
	"canchitas-libres-field/internal/pkg/infrastructure/web/dto"
	"canchitas-libres-field/internal/pkg/infrastructure/web/mapper"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type Service interface {
	GetAll() ([]domain.Field, error)
	GetByID(id string) (domain.Field, error)
	Add(field domain.Field) error
	Delete(id string) error
	Update(id string, field domain.Field) error
	GetByType(typeF string) ([]domain.Field, error)
}

type Handler struct {
	Service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		Service: service,
	}
}

var (
	getAllRe    = regexp.MustCompile(`^\/field\/?$`)
	getOneRe    = regexp.MustCompile(`^\/field\/([a-fA-F0-9-]{36})$`)
	createRe    = regexp.MustCompile(`^\/field\/?$`)
	getByTypeRe = regexp.MustCompile(`^\/field\/type\/([^\/]+)$`)
)

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && getAllRe.MatchString(r.URL.Path):
		w.Header().Set("Content-Type", "application/json")
		h.GetAllFields(w, r)
		return
	case r.Method == http.MethodGet && getOneRe.MatchString(r.URL.Path):
		w.Header().Set("Content-Type", "application/json")
		h.GetFieldByID(w, r)
		return
	case r.Method == http.MethodPost && createRe.MatchString(r.URL.Path):
		w.Header().Set("Content-Type", "application/json")
		h.CreateField(w, r)
		return
	case r.Method == http.MethodDelete:
		h.DeleteField(w, r)
		return
	case r.Method == http.MethodPut:
		h.UpdateField(w, r)
		return
	case r.Method == http.MethodGet && getByTypeRe.MatchString(r.URL.Path):
		w.Header().Set("Content-Type", "application/json")
		h.GetByType(w, r)
		return

	default:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Method not allowed"))
		return
	}
}

func (handler *Handler) GetAllFields(w http.ResponseWriter, r *http.Request) {
	fields, err := handler.Service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fieldsJSON, jsonErr := json.Marshal(fields) //lo transforma de codigo json a uno legible para el lenguaje
	if jsonErr != nil {
		return
	}

	// Configures the response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Send the JSON response to the client
	_, _ = w.Write(fieldsJSON)

}

func (handler *Handler) GetFieldByID(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	parts := strings.Split(path, "/")
	id := parts[len(parts)-1]

	err := dto.ValidateInputId(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	field, err := handler.Service.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fieldJSON, err := json.Marshal(field)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(fieldJSON)
}

func (handler *Handler) CreateField(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fieldDto := dto.FieldDto{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &fieldDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = dto.ValidateFieldCreateDto(fieldDto.Name, fieldDto.Type, fieldDto.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fieldDomain, err := mapper.ToDomainField(fieldDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.Service.Add(fieldDomain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("field was created"))
}

func (handler *Handler) UpdateField(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	parts := strings.Split(path, "/")
	id := parts[len(parts)-1]

	err := dto.ValidateInputId(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fieldDto := dto.FieldDto{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &fieldDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = dto.ValidateFieldUpdateDto(fieldDto.Name, fieldDto.Type, fieldDto.Price, fieldDto.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fieldDomain, err := mapper.ToDomainField(fieldDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.Service.Update(id, fieldDomain)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("field was updated"))
}

func (handler *Handler) DeleteField(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	parts := strings.Split(path, "/")
	id := parts[len(parts)-1]

	err := dto.ValidateInputId(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.Service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("field was eliminated"))
}

func (handler *Handler) GetByType(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		http.Error(w, "Missing type parameter", http.StatusBadRequest)
		return
	}
	typeValue := parts[len(parts)-1]
	typeF, err := url.PathUnescape(typeValue)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(typeF)
	if typeF == "" {
		http.Error(w, "Invalid field type", http.StatusBadRequest)
		return
	}

	fields, err := handler.Service.GetByType(typeF)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fieldsJSON, err := json.Marshal(fields)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(fieldsJSON)
}
