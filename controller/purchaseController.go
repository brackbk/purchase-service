package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/brackbk/purchase-service/dto"
	"github.com/brackbk/purchase-service/dto/purchase"
	"github.com/brackbk/purchase-service/errs"
	"github.com/brackbk/purchase-service/logger"
	"github.com/brackbk/purchase-service/service"
)

type PurchaseController struct {
	Service service.PurchaseService
}

func (p PurchaseController) Create(w http.ResponseWriter, r *http.Request) {
	logger.Info("Register")
	var request purchase.RegisterRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.Error(err.Error())
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	purchaseRegisterResponse, e := p.Service.Create(request)
	if e != nil {
		logger.Error(e.AsMessage())
		writeResponse(w, e.Code, e)
		return
	}

	response := purchase.MountResponse(*purchaseRegisterResponse)
	writeResponse(w, http.StatusCreated, response)
}

func (p PurchaseController) Delete(w http.ResponseWriter, r *http.Request) {
	logger.Info("Delete")
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		logger.Error(err.Error())
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	e := p.Service.Delete(&id)
	if e != nil {
		logger.Error(e.AsMessage())
		writeResponse(w, e.Code, e.AsResponse())
		return
	}
	writeResponse(w, http.StatusOK, " Deleted successfuly ")

}

func (p PurchaseController) Update(w http.ResponseWriter, r *http.Request) {
	logger.Info("Update")
	var request purchase.UpdateRequest
	if e := json.NewDecoder(r.Body).Decode(&request); e != nil {
		writeResponse(w, http.StatusBadRequest, e.Error())
		return
	}

	if e := dto.Validate(request); e != nil {
		writeResponse(w, http.StatusBadRequest, e)
		return
	}

	if e := p.Service.Update(request); e != nil {
		writeResponse(w, e.Code, e)
		return

	}
	writeResponse(w, http.StatusOK, nil)
}

func (p PurchaseController) List(w http.ResponseWriter, r *http.Request) {
	var request dto.ListRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	list, errList := p.Service.List(request)
	if errList != nil {
		writeResponse(w, errList.Code, errList)
		return
	}
	var rows []purchase.ResponseDTO
	for _, item := range *list {
		rows = append(rows, purchase.MountResponse(item))
	}
	count, e := p.Service.Count()
	if e != nil {
		writeResponse(w, errList.Code, errList)
		return
	}
	response := purchase.MountListResponse(count, rows)
	writeResponse(w, http.StatusOK, response)
}

func (p PurchaseController) Get(w http.ResponseWriter, r *http.Request) {
	logger.Info("Current Warehouse")
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		logger.Error(err.Error())
		writeResponse(w, http.StatusBadRequest, errs.ResponseError("No id was passed", 400))
		return
	}
	domain, e := p.Service.GetById(id)
	if e != nil {
		logger.Error(e.AsMessage())
		writeResponse(w, e.Code, e)
		return
	}
	response := purchase.MountResponse(*domain)
	writeResponse(w, http.StatusOK, response)
}
