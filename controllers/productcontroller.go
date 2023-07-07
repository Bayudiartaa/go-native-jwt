package controllers

import (
	"encoding/json"
	"go-native-jwt/configs"
	"go-native-jwt/helpers"
	"go-native-jwt/models"
	"net/http"

	"github.com/gorilla/mux"
)


func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product	

	if err := configs.DB.Find(&products).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "List Products", products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product

	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	defer r.Body.Close()

    if product.Name == "" {
		helpers.Response(w, 400, "Name cannot be empty", nil)
		return
	}

	if product.Price == 0 {
		helpers.Response(w, 400, "Price must be provided", nil)
		return
	}

	if product.Quantity == 0 {
		helpers.Response(w, 400, "Quantity must be provided", nil)
		return
	}

	if product.Description == "" {
		helpers.Response(w, 400, "Description cannot be empty", nil)
		return
	}

	if err := configs.DB.Create(&product).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 201, "Product created successfully", product)
}


func GetProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product models.Product
	if err := configs.DB.First(&product, params["id"]).Error; err != nil {
		helpers.Response(w, 404, "Product not found", nil)
		return
	}

	helpers.Response(w, 200, "Product found", product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product models.Product
	if err := configs.DB.First(&product, params["id"]).Error; err != nil {
		helpers.Response(w, 404, "Product not found", nil)
		return
	}

	var updatedProduct models.Product
	if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	// Check if Price is provided and valid
	if updatedProduct.Price <= 0 {
		helpers.Response(w, 400, "Price must be greater than 0", nil)
		return
	}

	configs.DB.Model(&product).Updates(updatedProduct)

	helpers.Response(w, 200, "Product updated successfully", product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var product models.Product
	if err := configs.DB.First(&product, params["id"]).Error; err != nil {
		helpers.Response(w, 404, "Product not found", nil)
		return
	}

	configs.DB.Delete(&product)

	helpers.Response(w, 200, "Product deleted successfully", nil)
}
