package woocommerce

type ProductType string

const (
	ProductTypeSimple   ProductType = "simple"
	ProductTypeGrouped  ProductType = "grouped"
	ProductTypeExternal ProductType = "external"
	ProductTypeVariable ProductType = "variable"
)

type ProductStatus string

const (
	ProductStatusDraft   ProductStatus = "draft"
	ProductStatusPending ProductStatus = "pending"
	ProductStatusPrivate ProductStatus = "private"
	ProductStatusPublish ProductStatus = "publish"
)

// ProductCommon contains the common fields of product and product variation.
type ProductCommon struct {
	ID           int           `json:"id,omitempty"`
	DateCreated  string        `json:"date_created,omitempty"`
	DateModified string        `json:"date_modified,omitempty"`
	Status       ProductStatus `json:"status,omitempty"`
	Description  string        `json:"description,omitempty"`
	SKU          string        `json:"sku,omitempty"`
	Price        NullFloat     `json:"price,omitempty"`
	RegularPrice NullFloat     `json:"regular_price,omitempty"`
	SalePrice    NullFloat     `json:"sale_price,omitempty"`
	MetaData     []MetaData    `json:"meta_data,omitempty"`
}

type Product struct {
	ProductCommon

	Name             string      `json:"name,omitempty"`
	Slug             string      `json:"slug,omitempty"`
	Type             ProductType `json:"type,omitempty"`
	Featured         bool        `json:"featured,omitempty"`
	ShortDescription string      `json:"short_description,omitempty"`
	ParentID         int         `json:"parent_id,omitempty"`
	Variations       []int       `json:"variations,omitempty"`
}

type ProductVariation struct {
	ProductCommon
}
