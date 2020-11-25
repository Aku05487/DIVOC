package pkg

import (
	"encoding/json"
	"github.com/divoc/portal-api/config"
	"github.com/divoc/portal-api/swagger_gen/restapi/operations"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type GenericResponse struct {
	statusCode int
}

func (o *GenericResponse) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {
	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses
	rw.WriteHeader(o.statusCode)
}

func NewGenericStatusOk() middleware.Responder {
	return &GenericResponse{statusCode: 200}
}

func NewGenericServerError() middleware.Responder {
	return &GenericResponse{statusCode: 500}
}
func SetupHandlers(api *operations.DivocPortalAPIAPI) {
	api.CreateMedicineHandler = operations.CreateMedicineHandlerFunc(createMedicineHandler)
	api.CreateProgramHandler = operations.CreateProgramHandlerFunc(createProgramHandler)
	api.PostFacilitiesHandler = operations.PostFacilitiesHandlerFunc(postFacilitiesHandler)
	api.PostVaccinatorsHandler = operations.PostVaccinatorsHandlerFunc(postVaccinatorsHandler)
	api.GetFacilitiesHandler = operations.GetFacilitiesHandlerFunc(getFacilitiesHandler)
	api.GetVaccinatorsHandler = operations.GetVaccinatorsHandlerFunc(getVaccinatorsHandler)
	api.GetMedicinesHandler = operations.GetMedicinesHandlerFunc(getMedicinesHandler)
	api.GetProgramsHandler = operations.GetProgramsHandlerFunc(getProgramsHandler)
}

func getProgramsHandler(params operations.GetProgramsParams, principal interface{}) middleware.Responder {
	return NewGenericStatusOk()
}
func getMedicinesHandler(params operations.GetMedicinesParams, principal interface{}) middleware.Responder {
	return NewGenericStatusOk()
}

func getVaccinatorsHandler(params operations.GetVaccinatorsParams, principal interface{}) middleware.Responder {
	return NewGenericStatusOk()
}

func getFacilitiesHandler(params operations.GetFacilitiesParams, principal interface{}) middleware.Responder {
	return NewGenericStatusOk()
}

func createMedicineHandler(params operations.CreateMedicineParams, principal interface{}) middleware.Responder {
	log.Infof("Create medicine %+v", params.Body)
	objectId:="Medicine"
	requestBody, err := json.Marshal(params.Body)
	if err != nil {
		return operations.NewCreateMedicineBadRequest()
	}
	requestMap := make(map[string]interface{})
	err = json.Unmarshal(requestBody, &requestMap)
	if err != nil {
		log.Info(err)
		return NewGenericServerError()
	}
	return makeRegistryCreateRequest(requestMap, objectId)
}

func createProgramHandler(params operations.CreateProgramParams, principal interface{}) middleware.Responder {
	log.Infof("Create Program %+v", params.Body)
	objectId:="Program"
	requestBody, err := json.Marshal(params.Body)
	if err != nil {
		return operations.NewCreateProgramBadRequest()
	}
	requestMap := make(map[string]interface{})
	err = json.Unmarshal(requestBody, &requestMap)
	if err != nil {
		log.Info(err)
		return NewGenericServerError()
	}
	return makeRegistryCreateRequest(requestMap, objectId)
}
func postFacilitiesHandler(params operations.PostFacilitiesParams, principal interface{}) middleware.Responder {
	data := NewScanner(params.File)
	defer params.File.Close()
	for data.Scan() {
		createFacility(&data)
		log.Info(data.Text("serialNum"), data.Text("facilityName"))
	}
	return operations.NewPostFacilitiesOK()
}

func postVaccinatorsHandler(params operations.PostVaccinatorsParams, principal interface{}) middleware.Responder {
	data := NewScanner(params.File)
	defer params.File.Close()
	for data.Scan() {
		createVaccinator(&data)
		log.Info("Created ", data.Text("serialNum"), data.Text("facilityName"))
	}
	return operations.NewPostFacilitiesOK()
}

func registryUrl(operationId string) string {
	url := config.Config.Registry.Url + "/" + operationId
	return url
}