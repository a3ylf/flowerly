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
