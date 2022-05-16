package controller

import (
	"encoding/json"
	"net/http"

	"github.com/brackbk/purchase-service/dto"
	"github.com/brackbk/purchase-service/dto/purchaseItem"
	"github.com/brackbk/purchase-service/logger"
	"github.com/brackbk/purchase-service/service"
)

type PurchaseItemController struct {
	Service         service.PurchaseItemService
	PurchaseService service.PurchaseService
}

func (p PurchaseItemController) Update(w http.ResponseWriter, r *http.Request) {
	logger.Info("Update")
	var request purchaseItem.UpdateRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.Error(err.Error())
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	purchase, errP := p.PurchaseService.GetById(request.PurchaseId)
	if errP != nil {
		logger.Error(err.Error())
		writeResponse(w, http.StatusBadRequest, errP)
		return
	}

	e := p.Service.Update(request)
	if e != nil {
		logger.Error(e.AsMessage())
		writeResponse(w, e.Code, e.AsResponse())
		return
	}
	errC := p.PurchaseService.SaveNewTotal(purchase)
	if errC != nil {
		logger.Error(err.Error())
		writeResponse(w, http.StatusBadRequest, errC)
		return
	}
	writeResponse(w, http.StatusOK, nil)
}

func (c PurchaseItemController) Delete(w http.ResponseWriter, r *http.Request) {
	logger.Info("Delete Item")
	var request purchaseItem.DeleteRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		logger.Error(err.Error())
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	errV := dto.Validate(request)
	if errV != nil {
		writeResponse(w, errV.Code, errV)
		return
	}
	purchase, errP := c.PurchaseService.GetById(request.PurchaseId)
	if errP != nil {
		writeResponse(w, http.StatusBadRequest, errP)
		return
	}
	errV = c.PurchaseService.VerifyItemQty(*purchase)
	if errV != nil {
		writeResponse(w, errV.Code, errV)
		return
	}
	e := c.Service.Delete(request)
	if e != nil {
		logger.Error(e.AsMessage())
		writeResponse(w, e.Code, e.AsResponse())
		return
	}
	errC := c.PurchaseService.SaveNewTotal(purchase)
	if errC != nil {
		logger.Error(err.Error())
		writeResponse(w, http.StatusBadRequest, errC)
		return
	}
	writeResponse(w, http.StatusOK, nil)
}
