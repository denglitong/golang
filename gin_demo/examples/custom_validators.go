package examples

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

type Booking struct {
	// validation: today < checkIn < CheckOut
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn,bookabledate" time_format:"2006-01-02"`
	Cost     float64   `form:"cost" binding:"required,costable"`
}

var bookableDate validator.Func = func(fl validator.FieldLevel) bool {
	date, ok := fl.Field().Interface().(time.Time)
	if ok {
		today := time.Now()
		if today.After(date) {
			return false
		}
	}
	return true
}

var costable validator.Func = func(fl validator.FieldLevel) bool {
	cost, ok := fl.Field().Interface().(float64)
	if ok {
		return cost > 10
	}
	return false
}

func ShowCustomValidators() error {
	route := gin.Default()

	// register custom validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("bookabledate", bookableDate)
		if err != nil {
			return err
		}
		err = v.RegisterValidation("costable", costable)
		if err != nil {
			return err
		}
	}

	// curl "localhost:8091/bookable?check_in=2118-04-16&check_out=2118-04-17"
	// 	valid
	// curl "localhost:8091/bookable?check_in=2118-03-10&check_out=2118-03-09"
	// 	"error": "Key: 'Booking.CheckOut' Error:Field validation for 'CheckOut' failed on the 'gtfield' tag"
	// http://localhost:8091/bookable?check_in=2023-04-16&check_out=2118-04-17
	// "error": "Key: 'Booking.CheckIn' Error:Field validation for 'CheckIn' failed on the 'bookabledate' tag"
	route.GET("/bookable", getBookable)
	return route.Run(":8091")
}

func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "Booking dates are valid!",
			"book":    b,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}
