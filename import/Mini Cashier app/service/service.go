package service

import (
	"a21hc3NpZ25tZW50/database"
	"a21hc3NpZ25tZW50/entity"
	"errors"
)

// Service is package for any logic needed in this program

type ServiceInterface interface {
	AddCart(productName string, quantity int) error
	RemoveCart(productName string) error
	ShowCart() ([]entity.CartItem, error)
	ResetCart() error
	GetAllProduct() ([]entity.Product, error)
	Paid(money int) (entity.PaymentInformation, error)
}

type Service struct {
	database database.DatabaseInterface
}

func NewService(database database.DatabaseInterface) *Service {
	return &Service{
		database: database,
	}
}

func (s *Service) AddCart(productName string, quantity int) error {
	if quantity <= 0 {
		return errors.New("invalid quantity")
	}
	product, err := s.database.GetProductByname(productName)
	if err != nil {
		return err
	}
	data, err := s.database.Load()
	if err != nil {
		return err
	}
	exists := false
	for _, datum := range data {
		if datum.ProductName == productName {
			datum.Quantity += quantity
			exists = true
			break
		}
	}

	if !exists {
		data = append(data, entity.CartItem{productName, product.Price, quantity})
	}
	err = s.database.Save(data)

	if err != nil {
		return nil
	}

	return nil // TODO: replace this
}

func (s *Service) RemoveCart(productName string) error {
	_, err := s.database.GetProductByname(productName)
	if err != nil {
		return err
	}
	data, err := s.database.Load() //product
	if err != nil {
		return err
	}

	//jika tidak ada yg bisa dijadikan kondisi, maka buat kondisi sendiri
	exists := false
	for a := 0; a < len(data); a++ {
		if data[a].ProductName == productName {
			data = append(data[:a], data[a+1:]...)
			exists = true
		}
	}
	if !exists {
		return errors.New("product not found")
	}
	err = s.database.Save(data)
	if err != nil {
		return err
	}
	return nil // TODO: replace this
}

func (s *Service) ShowCart() ([]entity.CartItem, error) {
	carts, err := s.database.Load()
	if err != nil {
		return nil, err
	}

	return carts, nil
}

func (s *Service) ResetCart() error {

	err := s.database.Save([]entity.CartItem{})

	if err != nil {
		return err
	}

	return nil // TODO: replace this
}

func (s *Service) GetAllProduct() ([]entity.Product, error) {

	return s.database.GetProductData(), nil // TODO: replace this
}

func (s *Service) Paid(money int) (entity.PaymentInformation, error) {

	data, err := s.database.Load()
	if err != nil {
		return entity.PaymentInformation{}, err
	}
	totalPrice := 0
	for _, check := range data {
		totalPrice += check.Price * check.Quantity
	}

	if money < totalPrice {
		return entity.PaymentInformation{}, errors.New("money is not enough")
	}
	err = s.ResetCart()
	if err != nil {
		return entity.PaymentInformation{}, err
	}

	return entity.PaymentInformation{
		ListProduct: data,
		TotalPrice:  totalPrice,
		MoneyPaid:   money,
		Change:      money - totalPrice,
	}, nil // TODO: replace this
}
