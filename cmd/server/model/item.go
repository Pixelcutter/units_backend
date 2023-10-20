package model

import "time"

type Item struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	SKU         string    `json:"sku"`
	CategoryID  int       `json:"category_id"`
	ImgPath     string    `json:"img_path"`
	Quantity    float32   `json:"quantity"`
	Price       float32   `json:"price"`
	ForSale     bool      `json:"for_sale"`
	CreatedBy   int       `json:"created_by"`
	UpdatedBy   int       `json:"updated_by"`
	Cost        float32   `json:"cost"`
	Unit        string    `json:"unit"`
}

type Component struct {
	ChildId  int     `json:"child_id"`
	Quantity float32 `json:"quantity"`
}

type NewItem struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	SKU         string      `json:"sku"`
	CategoryID  int         `json:"category_id"`
	ImgPath     string      `json:"img_path"`
	Quantity    float32     `json:"quantity"`
	Price       float32     `json:"price"`
	Cost        float32     `json:"cost"`
	ForSale     bool        `json:"for_sale"`
	CreatedBy   int         `json:"created_by"`
	UpdatedBy   int         `json:"updated_by"`
	Unit        string      `json:"unit"`
	Components  []Component `json:"components"`
}
