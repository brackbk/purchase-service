package repository

import (
	"github.com/brackbk/purchase-service/domain"
	"github.com/brackbk/purchase-service/errs"
	"github.com/brackbk/purchase-service/logger"
	"gorm.io/gorm"
)

type PurchaseItemRepositoryAdapter struct {
	client *gorm.DB
}

type PurchaseItemRepository interface {
	GetById(id int, purchaseId int) (*domain.PurchaseItem, *errs.Error)
	Delete(p domain.PurchaseItem) *errs.Error
	Update(p domain.PurchaseItem) *errs.Error
	Create(p *domain.PurchaseItem) *errs.Error
}

func NewPurchaseItemRepositoryAdapter(client *gorm.DB) PurchaseItemRepositoryAdapter {
	return PurchaseItemRepositoryAdapter{client}
}

func (p PurchaseItemRepositoryAdapter) GetById(id int, purchaseId int) (*domain.PurchaseItem, *errs.Error) {
	PurchaseItem := domain.PurchaseItem{}
	getPurchaseItemErr := p.client.Where("id = ? AND purchase_id = ?", id, purchaseId).Find(&PurchaseItem).Error
	if getPurchaseItemErr != nil {
		return nil, errs.ResponseError("Purchase Item not found", 422)
	}
	if PurchaseItem.ID == 0 {
		return nil, errs.UnprocessableEntityError("Purchase Item can not be found!")
	}
	return &PurchaseItem, nil
}

func (pr PurchaseItemRepositoryAdapter) Delete(p domain.PurchaseItem) *errs.Error {
	logger.Info("Deleted")
	errDelete := pr.client.Delete(&p).Error
	if errDelete != nil {
		return errs.ResponseError("Delete Error", 500)
	}
	return nil
}

func (pr PurchaseItemRepositoryAdapter) Update(p domain.PurchaseItem) *errs.Error {
	logger.Info("Updated")
	errUpdate := pr.client.Save(p).Error
	if errUpdate != nil {
		return errs.ResponseError("Update Error", 500)
	}
	return nil
}

func (pr PurchaseItemRepositoryAdapter) Create(p *domain.PurchaseItem) *errs.Error {
	logger.Info("Register")
	err := pr.client.Create(&p).Error
	if err != nil {
		return errs.ResponseError("Error to insert Purchase", 500)
	}
	return nil
}
