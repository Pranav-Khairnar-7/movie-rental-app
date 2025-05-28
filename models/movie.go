package model

type Movie struct {
	Base
	Title       string      `json:"title" gorm:"type:VARCHAR(200)"`
	Description string      `json:"description" gorm:"type:VARCHAR(200)"`
	ReleaseDate string      `json:"releaseDate" gorm:"type:DATE"`
	Genre       string      `json:"genre" gorm:"type:VARCHAR(200)"`
	MovieID     int         `json:"movieID" gorm:"type:INT"`
	ImdbCode    int         `json:"imdbCode" gorm:"type:INT"`
	PricingType PricingType `json:"pricingType" gorm:"type:VARCHAR(20)"`
}

func (Movie) TableName() string {
	return "movie"
}

type PricingType string

const (
	PRICING_REGULAR PricingType = "Regular"
	PRICING_NEW     PricingType = "New Release"
	PRICING_CLASSIC PricingType = "Classic"
)

type PriceStore struct {
	Base
	PricingType PricingType `json:"pricingType" gorm:"type:VARCHAR(20)"`
	Price       float64     `json:"price" gorm:"type:DECIMAL(4,2)"`
}
