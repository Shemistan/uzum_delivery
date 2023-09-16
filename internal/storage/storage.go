package storage

import "github.com/Shemistan/uzum_delivery/internal/models"

type IStorage interface {
	GetOrder(id int64) *models.Order
}
