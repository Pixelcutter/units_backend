package model

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TestItem struct {
	Item Item
	R1   string `json:"r1"`
	R2   string `json:"r2"`
}
