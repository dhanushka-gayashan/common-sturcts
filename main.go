package structs

type Item struct {
	ItemId      string  `json:"itemId"`
	SupplierId  string  `json:"supplierId"`
	ItemTypeId  string  `json:"itemTypeId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	Stock       int16   `json:"stock"`
	AvgStock    int16   `json:"avgStock"`
	MinStock    int16   `json:"minStock"`
	ReOrder     bool    `json:"reOrder"`
	ProfileURL  string  `json:"profileURL"`
}
