package service

import (
	"time"

	"github.com/brackbk/purchase-service/domain"
	"github.com/brackbk/purchase-service/dto"
	"github.com/brackbk/purchase-service/dto/purchase"
	"github.com/brackbk/purchase-service/errs"
	"github.com/brackbk/purchase-service/repository"
)

type PurchaseService interface {
	Create(p purchase.RegisterRequest) (*domain.Purchase, *errs.Error)
	Delete(id *int) *errs.Error
	Update(p purchase.UpdateRequest) *errs.Error
	List(p dto.ListRequest) (*[]domain.Purchase, *errs.Error)
	GetById(Id int) (*domain.Purchase, *errs.Error)
	SaveNewTotal(p *domain.Purchase) *errs.Error
	VerifyItemQty(p domain.Purchase) *errs.Error
	Count() (int, *errs.Error)
}
type DefaultPurchaseService struct {
	PurchaseRepo repository.PurchaseRepository
}

func NewDefaultPurchaseService(PurchaseRepo repository.PurchaseRepository) DefaultPurchaseService {
	return DefaultPurchaseService{PurchaseRepo}
}
func (d DefaultPurchaseService) Create(p purchase.RegisterRequest) (*domain.Purchase, *errs.Error) {
	err := dto.Validate(p)
	if err != nil {
		return nil, err
	}
	newPurchase := domain.Purchase{
		PaymentStatus: p.PaymentStatus,
		PaymentType:   p.PaymentType,
		CompanyId:     p.CompanyId,
		ProviderId:    p.ProviderId,
		LocationId:    p.LocationId,
		CurrencyId:    p.CurrencyId,
	}
	Total := 0.0
	var purchaseItems []domain.PurchaseItem
	for _, item := range p.PurchaseItem {
		total := (float64(item.Quantity) * item.UnitPrice)
		Total += total
		purchaseItems = append(purchaseItems, domain.PurchaseItem{
			PurchaseId: int(newPurchase.ID),
			ProductId:  item.ProductId,
			Quantity:   item.Quantity,
			UnitPrice:  item.UnitPrice,
			TotalPrice: total,
		})
	}
	newPurchase.PurchaseItem = purchaseItems
	newPurchase.Total = &Total
	if p.PurchaseDate != nil {
		newPurchase.PurchaseDate = p.PurchaseDate
	} else {
		newPurchase.PurchaseDate = &time.Time{}
	}
	if err != nil {
		return nil, err
	}
	_, err = d.PurchaseRepo.Create(&newPurchase)
	if err != nil {
		return nil, err
	}
	return &newPurchase, nil
}

func (d DefaultPurchaseService) GetById(Id int) (*domain.Purchase, *errs.Error) {
	purchase, err := d.PurchaseRepo.GetById(&Id)
	if err != nil {
		return nil, err
	}
	return purchase, nil
}

func (d DefaultPurchaseService) Delete(Id *int) *errs.Error {

	err := d.PurchaseRepo.DeleteById(Id)
	if err != nil {
		return err
	}
	return err
}

func (d DefaultPurchaseService) Update(p purchase.UpdateRequest) *errs.Error {
	err := dto.Validate(p)
	if err != nil {
		return err
	}
	purchase, err := d.PurchaseRepo.GetById(&p.Id)
	if err != nil {
		return err
	}
	if p.PaymentType != "" {
		purchase.PaymentType = p.PaymentType
	}

	if p.ProviderId != nil {
		purchase.ProviderId = p.ProviderId
	}

	if p.PurchaseDate != nil {
		purchase.PurchaseDate = p.PurchaseDate
	}

	Total := 0.0
	var purchaseItems []domain.PurchaseItem
	for _, item := range p.PurchaseItem {
		total := (float64(item.Quantity) * item.UnitPrice)
		Total += total

		purchaseItem := domain.PurchaseItem{
			PurchaseId: int(purchase.ID),
			ProductId:  item.ProductId,
			Quantity:   item.Quantity,
			UnitPrice:  item.UnitPrice,
			TotalPrice: total,
		}

		purchaseItems = append(purchaseItems, purchaseItem)

	}

	purchase.Total = &Total

	updateError := d.PurchaseRepo.UpdateWithItems(purchase, purchaseItems)
	if updateError != nil {
		return updateError
	}
	return nil
}

func (d DefaultPurchaseService) List(p dto.ListRequest) (*[]domain.Purchase, *errs.Error) {
	errList := dto.Validate(p)
	if errList != nil {
		return nil, errList
	}

	purchase, err := d.PurchaseRepo.List(&p)
	if err != nil {
		return nil, err
	}
	return purchase, nil
}

func (d DefaultPurchaseService) SaveNewTotal(p *domain.Purchase) *errs.Error {
	newTotal := 0.0
	for _, pi := range p.PurchaseItem {
		newTotal += pi.TotalPrice
	}
	p.Total = &newTotal
	err := d.PurchaseRepo.Update(p)
	if err != nil {
		return err
	}
	return nil
}

func (d DefaultPurchaseService) VerifyItemQty(p domain.Purchase) *errs.Error {
	if len(p.PurchaseItem) > 1 {
		return nil
	}
	return errs.UnprocessableEntityError("This item cannot be deleted, please verify your purchase items")
}
func (d DefaultPurchaseService) Count() (int, *errs.Error) {
	count, e := d.PurchaseRepo.Count()
	if e != nil {
		return 0, e
	}
	return count, e
}
