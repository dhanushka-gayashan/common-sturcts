package structs

import (
	"time"
)

const (
	Add          string = "Add"
	Deduct              = "Deduct"
	PlaceOrder          = "PlaceOrder"
	ReceiveOrder        = "ReceiveOrder"
)

type Card struct {
	Id         string    `json:"id"`
	CardNumber string    `json:"cardNumber"`
	CVV        int16     `json:"cvv"`
	CardHolder string    `json:"cardHolder"`
	ExpireDate time.Time `json:"expireDate"`
}

type Customer struct {
	CustomerId    string `json:"customerId"`
	Email         string `json:"email"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Phone         string `json:"phone"`
	Address       string `json:"address"`
	MasterCard    Card   `json:"masterCard"`
	SecondaryCard Card   `json:"secondaryCard"`
}

type Supplier struct {
	SupplierId string `json:"supplierId"`
	Email      string `json:"email"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	ProfileURL string `json:"profileURL"`
}

type Employee struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

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

type ItemType struct {
	ItemTypeId string `json:"itemTypeId"`
	Name       string `json:"name"`
}

type ListItem struct {
	ItemId   string  `json:"itemId"`
	Price    float32 `json:"price"`
	Quantity int16   `json:"quantity"`
}

type CustomerOrder struct {
	CustomerOrderId string     `json:"customerOrderId"`
	CustomerId      string     `json:"customerId"`
	Date            time.Time  `json:"date"`
	Progress        string     `json:"progress"`
	Items           []ListItem `json:"items"`
}

type CustomerReturn struct {
	CustomerReturnId string     `json:"customerReturnId"`
	CustomerId       string     `json:"customerId"`
	Date             time.Time  `json:"date"`
	Progress         string     `json:"progress"`
	Items            []ListItem `json:"items"`
}

type SupplierGoodReceive struct {
	SupplierGoodReceiveId string     `json:"supplierGoodReceiveId"`
	SupplierId            string     `json:"supplierId"`
	Date                  time.Time  `json:"date"`
	Progress              string     `json:"progress"`
	Items                 []ListItem `json:"items"`
}

type SupplierGoodReturn struct {
	SupplierGoodReturnId string     `json:"supplierGoodReturnId"`
	SupplierId           string     `json:"supplierId"`
	Date                 time.Time  `json:"date"`
	Progress             string     `json:"progress"`
	Items                []ListItem `json:"items"`
}

type SupplierPurchaseOrder struct {
	SupplierPurchaseOrderId string     `json:"supplierPurchaseOrderId"`
	SupplierId              string     `json:"supplierId"`
	Date                    time.Time  `json:"date"`
	Progress                string     `json:"progress"`
	Items                   []ListItem `json:"items"`
}

type OperationEvent struct {
	UniqueId       string `json:"uniqueId"`
	Table          string `json:"table"`
	PartitionKey   string `json:"partitionKey"`
	DocId          string `json:"docId"`
	Progress       string `json:"progress"`
	NotifyOnly     bool   `json:"notifyOnly"`
	ItemUpdateOnly bool   `json:"itemUpdateOnly"`
	Operation      string `json:"operation"`
	Email          bool   `json:"email"`
	SMS            bool   `json:"sms"`
}

type OperationUpdateItem struct {
	ItemId   string `json:"itemId"`
	Quantity int16  `json:"quantity"`
	OpType   string `json:"opType"`
}

type Response struct {
	StatusCode int16  `json:"statusCode"`
	Message    string `json:"message"`
}
