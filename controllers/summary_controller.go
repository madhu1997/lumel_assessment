package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/madhu1997/lumel_assessment.git/db/initializers"
	"gorm.io/gorm"
)

type RevenueByProduct struct {
	ProductName string
	Revenue     float64
}

type RevenueByCategory struct {
	Category string
	Revenue  float64
}

type RevenueByRegion struct {
	RegionName string
	Revenue    float64
}

type TotalRevenueResult struct {
	Revenue float64
}

func RevenueCalculation(c *gin.Context) {
	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")
	filter_type := c.Query("filter_type")
	var revenueByCategory []RevenueByCategory
	var revenueByProduct []RevenueByProduct
	var revenueByRegion []RevenueByRegion
	var totalRevenue TotalRevenueResult
	layout := "2006-01-02"
	startDate, _ := time.Parse(layout, startDateStr)
	endDate, _ := time.Parse(layout, endDateStr)

	if startDate.After(endDate) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Start date should be less than end date",
		})
		return
	}

	query := getTotalRevenue(startDateStr, endDateStr, filter_type)
	if err := query.Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	switch filter_type {
	case "by_product":
		query.Scan(&revenueByProduct)
		c.JSON(http.StatusOK, gin.H{"result": revenueByProduct})
		return
	case "by_category":
		query.Scan(&revenueByCategory)
		c.JSON(http.StatusOK, gin.H{"result": revenueByCategory})
		return
	case "by_region":
		query.Scan(&revenueByRegion)
		c.JSON(http.StatusOK, gin.H{"result": revenueByRegion})
		return
	default:
		query.Scan(&totalRevenue)
		c.JSON(http.StatusOK, gin.H{"result": totalRevenue})
		return
	}
}

func getTotalRevenue(startDate string, endDate string, filter_type string) (query_result *gorm.DB) {
	if filter_type == "by_product" {
		query_result = initializers.DB.Raw(`SELECT products.product_name, SUM(od.quantity_sold * od.unit_price - od.discount) AS revenue
			FROM orders o
			JOIN order_details od ON o.order_id = od.order_id
			JOIN products on products.product_id = od.product_id
			WHERE o.date_of_sale BETWEEN ? AND ?
			GROUP BY products.product_name`, startDate, endDate)

	} else if filter_type == "by_category" {
		query_result = initializers.DB.Raw(`SELECT products.category, SUM(od.quantity_sold * od.unit_price - od.discount) AS revenue
			FROM orders o
			JOIN order_details od ON o.order_id = od.order_id
			JOIN products on products.product_id = od.product_id
			WHERE o.date_of_sale BETWEEN ? AND ?
			GROUP BY products.category`, startDate, endDate)

	} else if filter_type == "by_region" {
		query_result = initializers.DB.Raw(`SELECT regions.region_name, SUM(od.quantity_sold * od.unit_price - od.discount) AS revenue
			FROM orders o
			JOIN regions on regions.id = o.region_id
			JOIN order_details od ON o.order_id = od.order_id
			WHERE o.date_of_sale BETWEEN ? AND ?
			GROUP BY regions.region_name`, startDate, endDate)

	} else {
		query_result = initializers.DB.Raw(`
			SELECT SUM(od.quantity_sold * od.unit_price - od.discount) AS revenue
			FROM orders o
			JOIN order_details od ON o.order_id = od.order_id
			WHERE o.date_of_sale BETWEEN ? AND ?
		`, startDate, endDate)
	}

	return query_result

}
