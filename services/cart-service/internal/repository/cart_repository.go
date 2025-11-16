package repository

import (
	"errors"
	"sync"

	"github.com/gocart-v2/shared/model"
)

var (
	ErrCartNotFound = errors.New("cart not found")
)

type CartRepository struct {
	carts      map[int]*model.Cart
	mu         sync.RWMutex
	nextCartID int
}

func NewCartRepository() *CartRepository {
	return &CartRepository{
		carts:      make(map[int]*model.Cart),
		nextCartID: 1,
	}
}

// Create creates a new cart
func (r *CartRepository) Create(customerID int) (*model.Cart, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	cart := &model.Cart{
		CartID:     r.nextCartID,
		CustomerID: customerID,
		Items:      []model.CartItem{},
	}

	r.carts[r.nextCartID] = cart
	r.nextCartID++

	return cart, nil
}

// GetByID retrieves a cart by its ID
func (r *CartRepository) GetByID(cartID int) (*model.Cart, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	cart, exists := r.carts[cartID]
	if !exists {
		return nil, ErrCartNotFound
	}

	// Return a copy
	cartCopy := *cart
	cartCopy.Items = make([]model.CartItem, len(cart.Items))
	copy(cartCopy.Items, cart.Items)

	return &cartCopy, nil
}

// AddItem adds an item to a cart
func (r *CartRepository) AddItem(cartID int, item model.CartItem) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	cart, exists := r.carts[cartID]
	if !exists {
		return ErrCartNotFound
	}

	// Check if product already exists in cart, if so update quantity
	found := false
	for i, existingItem := range cart.Items {
		if existingItem.ProductID == item.ProductID {
			cart.Items[i].Quantity += item.Quantity
			found = true
			break
		}
	}

	if !found {
		cart.Items = append(cart.Items, item)
	}

	return nil
}

// Delete removes a cart (used after checkout)
func (r *CartRepository) Delete(cartID int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.carts[cartID]; !exists {
		return ErrCartNotFound
	}

	delete(r.carts, cartID)
	return nil
}
