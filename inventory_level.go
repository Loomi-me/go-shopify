package goshopify

import (
	"fmt"
	"time"
)

const inventoryLevelBasePath = "inventory_levels"

type InventoryLevelService interface {
	List(interface{}) ([]InventoryLevel, error)
}

type InventoryLevelServiceOp struct {
	client *Client
}

type InventoryLevel struct {
	InventoryItemId   int64      `json:"inventory_item_id,omitempty"`
	LocationId        int64      `json:"location_id,omitempty"`
	Available         int64      `json:"available,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
	AdminGraphqlAPIID string     `json:"admin_graphql_api_id,omitempty"`
}

type InventoryLevelsResource struct {
	InventoryItems []InventoryLevel `json:"inventory_levels"`
}

func (s *InventoryLevelServiceOp) List(options interface{}) ([]InventoryLevel, error) {
	path := fmt.Sprintf("%s.json", inventoryLevelBasePath)
	resource := new(InventoryLevelsResource)
	err := s.client.Get(path, resource, options)
	return resource.InventoryItems, err
}
