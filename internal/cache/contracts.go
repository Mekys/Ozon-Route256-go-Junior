package cache

import (
	"context"
	"homework-3/internal/models"
	"time"
)

// NewOrderContacts wrapper
func NewOrderContacts(ttl time.Duration) *OrderContacts {
	return &OrderContacts{
		cli: NewTTLClient[string, []models.Order](ttl),
	}
}

type OrderContacts struct {
	cli *TTLClient[string, []models.Order]
}

func (p *OrderContacts) Get(ctx context.Context, key string) ([]models.Order, bool) {
	return p.cli.Get(key)
}

func (p *OrderContacts) Set(ctx context.Context, key string, contacts []models.Order, now time.Time) error {
	p.cli.Set(key, contacts, now)
	return nil
}

func (p *OrderContacts) Clear(ctx context.Context) error {
	p.cli.Clear()
	return nil
}
