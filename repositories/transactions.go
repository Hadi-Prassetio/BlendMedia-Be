package repositories

import (
	"ibox/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions(ID int) ([]models.Transaction, error)
	FindIncomes() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	GetOneTransaction(ID string) (models.Transaction, error)
	CreateTransaction(transactions models.Transaction) (models.Transaction, error)
	UpdateTransaction(status string, ID string) error
	GetCartTransaction(CartId int, Status string) (models.Cart, error)
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransactions(ID int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Cart").Preload("Buyer").Find(&transactions, "buyer_id = ?", ID).Error

	return transactions, err
}

func (r *repository) FindIncomes() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("Cart").Preload("Cart.Order.Product").Preload("Buyer").Find(&transactions).Error

	return transactions, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transactions models.Transaction
	err := r.db.Preload("Cart").Preload("Cart.Order").Preload("Buyer").Find(&transactions, "id = ?", ID).Error

	return transactions, err
}

// Create GetOneTransaction method here ...
func (r *repository) GetOneTransaction(ID string) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("Cart").Preload("Cart.Order").Preload("Buyer").First(&transaction, "id = ?", ID).Error

	return transaction, err
}

func (r *repository) CreateTransaction(transactions models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transactions).Error

	return transactions, err
}

// Create UpdateTransaction method here ...
func (r *repository) UpdateTransaction(status string, ID string) error {
	var transaction models.Transaction
	r.db.Preload("Cart").Preload("Cart.Order").First(&transaction, ID)

	// If is different & Status is "success" decrement product quantity
	if status != transaction.Status && status == "success" {
		transaction.Status = status
	}

	err := r.db.Save(&transaction).Error

	return err
}

func (r *repository) GetCartTransaction(CartId int, Status string) (models.Cart, error) {
	var Cart models.Cart
	err := r.db.Preload("User").Preload("Order").Preload("Order.Product").Where("user_id = ? AND status = ?", CartId, Status).First(&Cart).Error

	return Cart, err
}
