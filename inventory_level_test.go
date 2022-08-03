package goshopify

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"testing"
)

func TestInventoryLevel(t *testing.T) {
	setup()
	defer teardown()

	route := fmt.Sprintf("https://fooshop.myshopify.com/%s/inventory_levels.json?inventory_item_ids=%s&limit=250", client.pathPrefix, "1%2C2%2C3%2C4%2C5%2C6%2C7%2C8%2C9%2C10%2C11%2C12%2C13%2C14%2C15%2C16%2C17%2C18%2C19%2C20%2C21%2C22%2C23%2C24%2C25%2C26%2C27")
	httpmock.RegisterResponder("GET", route,
		httpmock.NewBytesResponder(200, loadFixture("inventory_levels.json")))

	items, err := client.InventoryLevel.List(InventoryLevelListOptions{
		IDs:   []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27},
		Limit: 250,
	})
	if err != nil {
		t.Errorf("InventoryLevels.List returned error: %v", err)
	}

	if len(items) == 0 {
		t.Fail()
	}
}
