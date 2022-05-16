package repository

import (
	"github.com/brackbk/purchase-service/domain"
	"github.com/brackbk/purchase-service/dto"
	"github.com/brackbk/purchase-service/errs"
	"github.com/brackbk/purchase-service/logger"
	"gorm.io/gorm"
)

type PurchaseRepositoryAdapter struct {
	client *gorm.DB
}

type PurchaseRepository interface {
	Create(*domain.Purchase) (*domain.Purchase, *errs.Error)
	DeleteById(id *int) *errs.Error
	List(u *dto.ListRequest) (*[]domain.Purchase, *errs.Error)
	Update(u *domain.Purchase) *errs.Error
	UpdateWithItems(u *domain.Purchase, items []domain.PurchaseItem) *errs.Error
	GetById(Id *int) (*domain.Purchase, *errs.Error)
	VerifyExist(Name string) *errs.Error
	Count() (int, *errs.Error)
}

func NewPurchaseRepositoryAdapter(client *gorm.DB) PurchaseRepositoryAdapter {
	return PurchaseRepositoryAdapter{client}
}

func (pr PurchaseRepositoryAdapter) Create(p *domain.Purchase) (*domain.Purchase, *errs.Error) {
	logger.Info("Register")
	err := pr.client.Create(&p).Error
	if err != nil {
		return p, errs.ResponseError("Error to insert Purchase", 500)
	}
	return nil, nil
}

func (p PurchaseRepositoryAdapter) DeleteById(Id *int) *errs.Error {
	purchase, err := p.GetById(Id)
	if err != nil {
		return errs.UnprocessableEntityError(err.AsMessage())
	}

	tx := p.client.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return errs.ResponseError("Delete Error", 500)
	}

	if err := tx.Delete(&purchase.PurchaseItem).Error; err != nil {
		tx.Rollback()
		return errs.ResponseError("Update Error", 500)
	}

	if err := tx.Delete(&purchase).Error; err != nil {
		tx.Rollback()
		return errs.ResponseError("Update Error", 500)
	}

	saveError := tx.Commit().Error
	if saveError != nil {
		return errs.ResponseError("Update Error", 500)
	}
	return nil
}

func (pr PurchaseRepositoryAdapter) UpdateWithItems(u *domain.Purchase, items []domain.PurchaseItem) *errs.Error {
	logger.Info("Updated")

	tx := pr.client.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return errs.ResponseError("Update Error", 500)
	}

	if err := tx.Delete(&u.PurchaseItem).Error; err != nil {
		tx.Rollback()
		return errs.ResponseError("Update Error", 500)
	}

	u.PurchaseItem = items
	if err := tx.Save(&u).Error; err != nil {
		tx.Rollback()
		return errs.ResponseError("Update Error", 500)
	}

	saveError := tx.Commit().Error
	if saveError != nil {
		return errs.ResponseError("Update Error", 500)
	}
	return nil
}

func (p PurchaseRepositoryAdapter) Update(u *domain.Purchase) *errs.Error {
	logger.Info("Updated")
	if e := p.client.Save(&u).Error; e != nil {
		return errs.ResponseError("Update Error", 500)
	}
	return nil
}
func (pr PurchaseRepositoryAdapter) List(p *dto.ListRequest) (*[]domain.Purchase, *errs.Error) {
	logger.Info("Listed")
	purchase := []domain.Purchase{}
	offset := (p.Page - 1) * p.Limit
	err := pr.client.Offset(offset).
		Preload("PurchaseItem").
		Limit(p.Limit).
		Find(&purchase).Error
	if err != nil {
		return nil, errs.ResponseError("List not found", 500)
	}
	return &purchase, nil
}

func (p PurchaseRepositoryAdapter) GetById(Id *int) (*domain.Purchase, *errs.Error) {
	logger.Info("Get Purchase by Id")
	purchase := domain.Purchase{}
	err := p.client.Where("id = ?", Id).
		Preload("PurchaseItem").
		Find(&purchase).Error
	if err != nil {
		return nil, errs.ResponseError("Server error", 500)
	}
	if purchase.ID == 0 {
		return nil, errs.ResponseError("Purchase not found", 404)
	}
	return &purchase, nil
}

func (p PurchaseRepositoryAdapter) VerifyExist(Name string) *errs.Error {
	purchase := domain.Purchase{}
	err := p.client.Where("UPPER(name) = UPPER(?)", Name).Find(&purchase).Error
	if err != nil {
		return errs.ResponseError("Server error", 500)
	}
	if purchase.ID != 0 {
		return errs.ResponseError("Can not create a new Purchase, that name is already taken", 422)
	}
	return nil
}
func (p PurchaseRepositoryAdapter) Count() (int, *errs.Error) {
	var list []domain.Purchase
	result := p.client.Find(&list)
	if result.Error != nil {
		return 0, errs.ResponseError("Something went wrong. Please retry late", 500)
	}
	return int(result.RowsAffected), nil
}
