package service

import (
	"github.com/brackbk/purchase-service/domain"
	"github.com/brackbk/purchase-service/dto"
	"github.com/brackbk/purchase-service/dto/purchaseItem"
	"github.com/brackbk/purchase-service/errs"
	"github.com/brackbk/purchase-service/repository"
)

type PurchaseItemService interface {
	Delete(p purchaseItem.DeleteRequest) *errs.Error
	Update(p purchaseItem.UpdateRequest) *errs.Error
	Create(p purchaseItem.RegisterRequest) *errs.Error
}

type DefaultPurchaseItemService struct {
	PurchaseItemRepo repository.PurchaseItemRepository
}

func NewDefaultPurchaseItemService(PurchaseItemRepo repository.PurchaseItemRepository) DefaultPurchaseItemService {
	return DefaultPurchaseItemService{PurchaseItemRepo}
}
func (d DefaultPurchaseItemService) Delete(p purchaseItem.DeleteRequest) *errs.Error {
	err := dto.Validate(p)
	if err != nil {
		return err
	}
	purchaseItem, err := d.PurchaseItemRepo.GetById(p.Id, p.PurchaseId)
	if err != nil {
		return err
	}
	errD := d.PurchaseItemRepo.Delete(*purchaseItem)
	if errD != nil {
		return errD
	}
	return nil

}

func (d DefaultPurchaseItemService) Update(p purchaseItem.UpdateRequest) *errs.Error {
	err := dto.Validate(p)
	if err != nil {
		return err
	}
	PurchaseItem, err := d.PurchaseItemRepo.GetById(p.Id, p.PurchaseId)
	if err != nil {
		return err
	}
	if p.Quantity != 0 {
		PurchaseItem.Quantity = p.Quantity
	}
	if p.UnitPrice != 0 {
		PurchaseItem.UnitPrice = p.UnitPrice
	}

	PurchaseItem.TotalPrice = (float64(PurchaseItem.Quantity) * PurchaseItem.UnitPrice)

	updateError := d.PurchaseItemRepo.Update(*PurchaseItem) //2
	if updateError != nil {
		return updateError
	}

	return nil
}
func (d DefaultPurchaseItemService) Create(p purchaseItem.RegisterRequest) *errs.Error {
	err := dto.Validate(p)
	if err != nil {
		return err
	}
	total := (float64(p.Quantity) * p.UnitPrice)
	purchaseItem := domain.PurchaseItem{
		PurchaseId: p.PurchaseId,
		ProductId:  p.ProductId,
		Quantity:   p.Quantity,
		UnitPrice:  p.UnitPrice,
		TotalPrice: total,
	}
	err = d.PurchaseItemRepo.Create(&purchaseItem)
	if err != nil {
		return err
	}
	return nil
}
