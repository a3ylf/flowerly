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
