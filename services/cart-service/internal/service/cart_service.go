package service

import (
	"errors"

	"github.com/gocart-v2/cart-service/internal/repository"
	"github.com/gocart-v2/shared/model"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrCartNotFound    = errors.New("cart not found")
	ErrInvalidCart     = errors.New("invalid cart data")
	ErrEmptyCart       = errors.New("cart is empty")
)

type CartService struct {
	cartRepo *repository.CartRepository
}

func NewCartService(cartRepo *repository.CartRepository) *CartService {
	return &CartService{
		cartRepo: cartRepo,
	}
}

// CreateCart creates a new cart
func (s *CartService) CreateCart(customerID int) (*model.Cart, error) {
	if customerID < 1 {
		return nil, ErrInvalidCart
	}

	return s.cartRepo.Create(customerID)
}

// AddItemToCart adds an item to a cart
func (s *CartService) AddItemToCart(cartID int, productID int, quantity int) error {
	if cartID < 1 || productID < 1 || quantity < 1 {
		return ErrInvalidCart
	}

	// Verify cart exists
	_, err := s.cartRepo.GetByID(cartID)
	if err == repository.ErrCartNotFound {
		return ErrCartNotFound
	}
	if err != nil {
		return err
	}

	// Add item to cart
	item := model.CartItem{
		ProductID: productID,
		Quantity:  quantity,
	}

	return s.cartRepo.AddItem(cartID, item)
}

// CheckoutCart processes checkout for a cart
func (s *CartService) CheckoutCart(cartID int) (int, error) {
	if cartID < 1 {
		return 0, ErrInvalidCart
	}

	// Get cart
	cart, err := s.cartRepo.GetByID(cartID)
	if err == repository.ErrCartNotFound {
		return 0, ErrCartNotFound
	}
	if err != nil {
		return 0, err
	}

	// Validate cart has items
	if len(cart.Items) == 0 {
		return 0, ErrEmptyCart
	}

	// In a real system, this would:
	// 1. Reserve inventory
	// 2. Process payment
	// 3. Create order
	// 4. Delete cart

	// For now, just generate an order ID and delete the cart
	orderID := cartID * 1000 // Simple order ID generation

	err = s.cartRepo.Delete(cartID)
	if err != nil {
		return 0, err
	}

	return orderID, nil
}

// GetCart retrieves a cart
func (s *CartService) GetCart(cartID int) (*model.Cart, error) {
	if cartID < 1 {
		return nil, ErrInvalidCart
	}

	cart, err := s.cartRepo.GetByID(cartID)
	if err == repository.ErrCartNotFound {
		return nil, ErrCartNotFound
	}
	return cart, err
}
