package goshopify

import (
	"fmt"
	"net/http"
	"time"
)

const inventoryLevelBasePath = "inventory_levels"

type InventoryLevelService interface {
	List(interface{}) ([]InventoryLevel, error)
	ListWithPagination(interface{}) ([]InventoryLevel, *Pagination, error)
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

func (s *InventoryLevelServiceOp) ListWithPagination(options interface{}) ([]InventoryLevel, *Pagination, error) {
	path := fmt.Sprintf("%s.json", inventoryLevelBasePath)
	resource := new(InventoryLevelsResource)
	headers := http.Header{}

	headers, err := s.client.createAndDoGetHeaders("GET", path, nil, options, resource)
	if err != nil {
		return nil, nil, err
	}
	linkHeader := headers.Get("Link")
	pagination, err := extractPagination(linkHeader)
	if err != nil {
		return nil, nil, err
	}
	return resource.InventoryItems, pagination, nil
}
func (s *InventoryLevelServiceOp) List(options interface{}) ([]InventoryLevel, error) {
	path := fmt.Sprintf("%s.json", inventoryLevelBasePath)
	resource := new(InventoryLevelsResource)
	err := s.client.Get(path, resource, options)
	return resource.InventoryItems, err
}
