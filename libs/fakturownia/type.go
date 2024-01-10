package facturownia

type DocKind string

const (
	KindPZ DocKind = "pz"
	KindWZ DocKind = "wz"
	KindRW DocKind = "rw"
	KindPW DocKind = "pw"
	KindMM DocKind = "mm"
)

type WarehouseAction struct {
    ProductName        string  `json:"product_name"`
    ProductId		   string  `json:"product_id"`
	PurchaseTax        int     `json:"purchase_tax"`
    PurchasePriceNet   float64 `json:"purchase_price_net"`
    Quantity           int     `json:"quantity"`

}

type WarehouseDocument struct {
    Kind            DocKind            `json:"kind"`
    Number          *int               `json:"number,omitempty"` // używamy wskaźnika, by umożliwić wartość nil
    WarehouseID     string             `json:"warehouse_id"`
    IssueDate       string             `json:"issue_date"`
    DepartmentName  string             `json:"department_name"`
    ClientName      string             `json:"client_name"`
    WarehouseActions []WarehouseAction `json:"warehouse_actions"`
}

type Payload struct {
    APIToken          string             `json:"api_token"`
    WarehouseDocument WarehouseDocument  `json:"warehouse_document"`
}