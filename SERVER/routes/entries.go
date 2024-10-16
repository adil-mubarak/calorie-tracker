package routes

import (
	"fmt"
	"net/http"

	"calorie-tracker/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB 

func InitDB(database *gorm.DB){
    db = database
}

func AddEntry(c *gin.Context) {
    var entry models.Entry

    if err := c.BindJSON(&entry); err != nil{
        c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
        fmt.Println(err)
        return
    }

    if err := db.Create(&entry).Error; err != nil{
        c.JSON(http.StatusInternalServerError,gin.H{"error":"Could not create entry"})
        fmt.Println(err)
        return
    }
    c.JSON(http.StatusOK,entry)
}

func GetEntries(c *gin.Context) {
    var entries []models.Entry

    if err := db.Find(&entries).Error; err != nil{
        c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
        fmt.Println(err)
        return
    }

    c.JSON(http.StatusOK,entries)
}

func GetEntriesByIngredient(c *gin.Context) {

}

func GetEntryById(c *gin.Context) {

}

func UpdateIngredient(c *gin.Context) {

}

func UpdateEntry(c *gin.Context) {

}

func DeleteEntry(c *gin.Context) {
	entryID := c.Params.ByName("id")

	if err := db.Delete(&models.Entry{}, entryID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, "Entry deleted successfully")
}
