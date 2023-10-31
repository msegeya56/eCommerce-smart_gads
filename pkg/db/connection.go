package db

import (
	"errors"
	"fmt"
	"log"

	"github.com/msegeya56/eCommerce-smart_gads/pkg/config"
	"github.com/msegeya56/eCommerce-smart_gads/pkg/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB variables
var (
	DB  *gorm.DB
	err error
)

// To connect database
func ConnToDB(cfg config.Config) (*gorm.DB, error) {

	dsn := cfg.DATABASE

	if DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Failed to Connect Database")
		return nil, errors.New("failed to connect database")
	}
	fmt.Println("Successfully Connected to database")
	// defer DB.Close()
	// Migrate models
	err := DB.AutoMigrate(
		// Users
		domain.Users{},
		domain.Admin{},
		domain.Address{},

		// Product tables
		domain.Category{},
		domain.Product{},
		domain.ProductItem{},
		domain.ProductConfig{},
		domain.Variation{},
		domain.VariationOption{},
		domain.ProductImage{},

		// cart tables
		domain.Cart{},
		domain.CartItem{},

		// wishlist
		domain.Wishlist{},

		// payment tables
		domain.PaymentDetails{},
		domain.PaymentMethod{},
		domain.PaymentStatus{},

		// Order tables
		domain.ShopOrder{},
		domain.OrderStatus{},
		domain.OrderLine{},
		domain.ShopOrder{},
		domain.DeliveryStatus{},

		// Coupon
		domain.Coupon{},

		// Return
		domain.Return{},

		// wallet
		domain.Wallet{},
	)
	if err != nil {
		log.Fatal("Migration failed")
		return nil, nil
	}
	fmt.Println("DB migration success")
	return DB, nil

}
