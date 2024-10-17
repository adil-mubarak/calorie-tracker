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
    ingredient := c.Params.ByName("ingredient")
    var entries []models.Entry

    if err := db.Where("ingredients LIKE ?","%"+ingredient+"%").Find(&entries).Error; err != nil{
        c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
        fmt.Println(err)
        return
    }

    c.JSON(http.StatusOK,entries)
}

func GetEntryById(c *gin.Context) {
    entryID := c.Params.ByName("id")
    var entry models.Entry

    if err := db.First(&entry,entryID).Error; err != nil{
        c.JSON(http.StatusInternalServerError,gin.H{"error":"Entry not found"})
        fmt.Println(err)
        return
    }

    c.JSON(http.StatusOK,entry)
}

func UpdateIngredient(c *gin.Context) {
    entryID := c.Params.ByName("id")
    var ingredient struct{
        Ingredient *string `json:"ingredients"`
    }

    if err := c.BindJSON(&ingredient); err != nil{
        c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
        fmt.Println(err)
        return
    }

    if err := db.Model(&models.Entry{}).Where("id = ?",entryID).Update("ingredients",ingredient.Ingredient).Error; err != nil{
        c.JSON(http.StatusInternalServerError,gin.H{"error":"Could not update ingredients"})
        fmt.Println(err)
        return
    }

    c.JSON(http.StatusOK,gin.H{"message":"Ingredients updated successfully"})

}

func UpdateEntry(c *gin.Context) {
    entryID := c.Params.ByName("id")
    var entry models.Entry

    if err := c.BindJSON(&entry); err != nil{
        c.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
        fmt.Println(err)
        return
    }

    if err := db.Model(&entry).Where("id = ?",entryID).Updates(entry).Error; err != nil{
        c.JSON(http.StatusInternalServerError,gin.H{"error":"Could not update the entrie"})
        fmt.Println(err)
        return
    }

    c.JSON(http.StatusOK,entry)
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
