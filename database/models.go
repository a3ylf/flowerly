package database


type Plant struct {
    ID             int     `json:"id"`
    Name           string  `json:"name"`
    ScientificName string  `json:"scientific_name"`
    Description    string  `json:"description"`
    Category       string  `json:"category"`
    Price          float64 `json:"price"`
    StockQuantity  int     `json:"stock_quantity"`
    ImageURL       string  `json:"image_url"`
    OriginLocation string  `json:"origin_location"`
}


type Vendor struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
    CPF      string `json:"cpf"`
}
type Client struct {
    Vendor
    Rua      string `json:"rua"`
    Num      int16  `json:"num"`
}
type Cart struct {
	ID       int `json:"id"`         // ID do carrinho
	ClientID int `json:"client_id"`  // ID do cliente associado
}
type CartItem struct {
	CartID    int       `json:"cart_id"`    // ID do carrinho
	ProductID int       `json:"product_id"` // ID do produto
	Quantity  int       `json:"quantity"`   // Quantidade de itens
}
type Product struct {
    ProductID   int    `json:"product_id"`
    ProductName string `json:"product_name"`
    Quantity    int    `json:"quantity"`
}

type Purchase struct {
    PurchaseID     int       `json:"purchase_id"`
    PurchaseDate   string    `json:"purchase_date"`
    TotalAmount    float64   `json:"total_amount"`
    PaymentStatus  string    `json:"payment_status"`
    PaymentMethod  string    `json:"payment_method"`
    Products       []Product `json:"products"`
}

type ClientPurchases struct {
    ClientID   int       `json:"client_id"`
    ClientName string    `json:"client_name"`
    Purchases  []Purchase `json:"purchases"`
}