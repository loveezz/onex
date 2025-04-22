package controller

import (
	"net/http"
	"one/db"
	"one/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InitProducts(zxc *gin.Context){

	var products  []models.Product

	db.DB.Find(&products)

	zxc.JSON(http.StatusOK, products)


}

func GetProduct(zxc *gin.Context){
	var product models.Product

	id,_ :=strconv.Atoi(zxc.Param("id"))

	r := db.DB.First(&product, id)


	if r.Error != nil {

		zxc.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})

		return
	}
	zxc.JSON(http.StatusOK, product)

}

func CreateProduct(zxc *gin.Context){
	var product models.Product

	if err := zxc.ShouldBindJSON(&product); err != nil {
		zxc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	db.DB.Create(&product)

	zxc.JSON(http.StatusCreated, product)
}


func UpdateProduct(zxc *gin.Context) {

    id, err := strconv.Atoi(zxc.Param("id"))

    if err != nil {
        zxc.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID"})
        return
    }


    var existingProduct models.Product

    if err := db.DB.First(&existingProduct, id).Error; err != nil {
        zxc.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
        return
    }


    var updateData models.Product

    if err := zxc.ShouldBindJSON(&updateData); err != nil {
        zxc.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }


    existingProduct.Name = updateData.Name
    existingProduct.Price = updateData.Price
    existingProduct.Description = updateData.Description


    if err := db.DB.Save(&existingProduct).Error; err != nil {
        zxc.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при обновлении"})
        return
    }

    zxc.JSON(http.StatusOK, existingProduct)
}

func DeleteProduct(zxc* gin.Context){

	var product models.Product

	id,_ := strconv.Atoi(zxc.Param("id"))


	if err := db.DB.First(&product, id ).Error; err != nil {
		zxc.JSON(http.StatusNotFound, gin.H{"error": "Товар не найден"})
		return
	}


	db.DB.Delete(&product)

	zxc.JSON(http.StatusNoContent, nil)


}
