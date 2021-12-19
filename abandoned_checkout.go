package goshopify

import (
	"github.com/shopspring/decimal"
	"time"
)

type AbandonedCartResource struct {
	Checkouts []AbandonedCart `json:"checkouts,omitempty"`
}

type AbandonedCart struct {
	ID                    int                     `json:"id,omitempty"`
	Token                 string                  `json:"token,omitempty"`
	CartToken             string                  `json:"cart_token,omitempty"`
	Email                 string                  `json:"email,omitempty"`
	Gateway               *string                 `json:"gateway,omitempty"`
	BuyerAcceptsMarketing bool                    `json:"buyer_accepts_marketing,omitempty"`
	CreatedAt             *time.Time              `json:"created_at,omitempty"`
	UpdatedAt             *time.Time              `json:"updated_at,omitempty"`
	LandingSite           *string                 `json:"landing_site,omitempty"`
	Note                  *string                 `json:"note,omitempty"`
	NoteAttributes        []NoteAttribute         `json:"note_attributes,omitempty"`
	ReferringSite         string                  `json:"referring_site,omitempty"`
	ShippingLines         []CheckoutShippingLines `json:"shipping_lines,omitempty"`

	TaxesIncluded            bool                `json:"taxes_included,omitempty"`
	TotalWeight              int                 `json:"total_weight,omitempty"`
	Currency                 string              `json:"currency,omitempty"`
	CompletedAt              *time.Time          `json:"completed_at,omitempty"`
	ClosedAt                 *time.Time          `json:"closed_at,omitempty"`
	UserID                   *int64              `json:"user_id,omitempty"`
	LocationID               *int64              `json:"location_id,omitempty"`
	SourceIdentifier         string              `json:"source_identifier,omitempty"`
	SourceURL                string              `json:"source_url,omitempty"`
	DeviceID                 *int64              `json:"device_id,omitempty"`
	Phone                    string              `json:"phone,omitempty"`
	CustomerLocale           string              `json:"customer_locale,omitempty"`
	LineItems                []CheckoutLineItems `json:"line_items,omitempty"`
	Name                     string              `json:"name,omitempty"`
	AbandonedCheckoutURL     string              `json:"abandoned_checkout_url,omitempty"`
	DiscountCodes            []DiscountCode      `json:"discount_codes,omitempty"`
	TaxLines                 []TaxLine           `json:"tax_lines,omitempty"`
	SourceName               string              `json:"source_name,omitempty"`
	PresentmentCurrency      string              `json:"presentment_currency,omitempty"`
	BuyerAcceptsSmsMarketing bool                `json:"buyer_accepts_sms_marketing,omitempty"`
	SmsMarketingPhone        *string             `json:"sms_marketing_phone,omitempty"`
	TotalDiscounts           *decimal.Decimal    `json:"total_discounts,omitempty"`
	TotalLineItemsPrice      *decimal.Decimal    `json:"total_line_items_price,omitempty"`
	TotalPrice               *decimal.Decimal    `json:"total_price,omitempty"`
	TotalTax                 *decimal.Decimal    `json:"total_tax,omitempty"`
	SubtotalPrice            *decimal.Decimal    `json:"subtotal_price,omitempty"`
	TotalDuties              *decimal.Decimal    `json:"total_duties,omitempty"`
	BillingAddress           *Address            `json:"billing_address,omitempty"`
	ShippingAddress          *Address            `json:"shipping_address,omitempty"`
	Customer                 *Customer           `json:"customer,omitempty"`
}

type CheckoutLineItems struct {
	AppliedDiscount         *AppliedDiscount      `json:"applied_discount,omitempty"`
	DiscountAllocations     []DiscountAllocations `json:"discount_allocations"`
	Key                     string                `json:"key"`
	GiftCard                bool                  `json:"gift_card"`
	Grams                   int                   `json:"grams"`
	PresentmentTitle        string                `json:"presentment_title"`
	PresentmentVariantTitle string                `json:"presentment_variant_title"`
	ProductID               int                   `json:"product_id"`
	Properties              []NoteAttribute       `json:"properties"`
	Quantity                int                   `json:"quantity"`
	RequiresShipping        bool                  `json:"requires_shipping"`
	Sku                     string                `json:"sku"`
	TaxLines                []TaxLine             `json:"tax_lines"`
	Taxable                 bool                  `json:"taxable"`
	Title                   string                `json:"title"`
	VariantID               int                   `json:"variant_id"`
	VariantTitle            string                `json:"variant_title"`
	VariantPrice            *decimal.Decimal      `json:"variant_price"`
	Vendor                  string                `json:"vendor"`
	UserID                  int64                 `json:"user_id"`
	LinePrice               *decimal.Decimal      `json:"line_price"`
	Price                   string                `json:"price"`
}

type CheckoutShippingLines struct {
	Title            string            `json:"title"`
	Price            *decimal.Decimal            `json:"price"`
	Code             string            `json:"code"`
	Source           string            `json:"source"`
	AppliedDiscounts []AppliedDiscount `json:"applied_discounts"`
	ID               string            `json:"id"`
}
