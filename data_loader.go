package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/madhu1997/lumel_assessment.git/db/initializers"
	"github.com/madhu1997/lumel_assessment.git/models"
)

const (
	workerCount = 10 // Number of concurrent workers
)

func load_csv_data() {
	file, err := os.Open("sales_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	jobs := make(chan []string, len(records))
	var wg sync.WaitGroup

	//Start worker goroutines
	fmt.Printf("started data loading..... total length %v\n", len(records))
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go worker(jobs, &wg)
	}

	// Distribute records to workers
	for _, record := range records[1:] {
		//fmt.Print(record[0], record[1])
		jobs <- record
	}
	close(jobs)
	wg.Wait()
}

func worker(jobs <-chan []string, wg *sync.WaitGroup) {
	defer wg.Done()

	for record := range jobs {
		if err := validateRecord(record); err != nil {
			log.Printf("Skipping invalid record: %v\n", err)
			continue
		}
		orderID, _ := strconv.Atoi(record[0])
		productID := record[1]
		customerID := record[2]
		regionName := record[5]
		dateOfSale, _ := time.Parse("2006-01-02", record[6])
		quantitySold, _ := strconv.Atoi(record[7])
		unitPrice, _ := strconv.ParseFloat(record[8], 64)
		discount, _ := strconv.ParseFloat(record[9], 64)
		shippingCost, _ := strconv.ParseFloat(record[10], 64)

		var product models.Product
		var customer models.Customer
		var region models.Region

		db := initializers.DB

		db.FirstOrCreate(&product, models.Product{ProductID: productID, ProductName: record[3], Category: record[4]})
		db.FirstOrCreate(&customer, models.Customer{CustomerID: customerID, CustomerName: record[12], Email: record[13], Address: record[14]})
		db.FirstOrCreate(&region, models.Region{RegionName: regionName})

		var regionID uint
		db.Model(&region).Where("region_name = ?", regionName).Select("id").Scan(&regionID)

		order := models.Order{
			OrderID:       orderID,
			CustomerID:    customerID,
			RegionID:      int(regionID),
			DateOfSale:    dateOfSale,
			PaymentMethod: record[11],
			ShippingCost:  shippingCost,
		}
		orderDetail := models.OrderDetail{
			OrderID:      orderID,
			ProductID:    productID,
			QuantitySold: quantitySold,
			UnitPrice:    unitPrice,
			Discount:     discount,
		}

		db.Create(&order)
		db.Create(&orderDetail)
	}
}

func validateRecord(record []string) error {
	/// Check if required fields are valid
	if record[1] == "" || record[2] == "" || record[6] == "" || record[7] == "" {
		return fmt.Errorf("missing required fields")
	}

	// Validate date format
	if _, err := time.Parse("2006-01-02", record[6]); err != nil {
		return fmt.Errorf("invalid date format in record: %v", record[6])
	}

	// Validate numeric fields
	if _, err := strconv.Atoi(record[7]); err != nil {
		return fmt.Errorf("invalid quantitySold: %v", record[7])
	}
	if _, err := strconv.ParseFloat(record[8], 64); err != nil {
		return fmt.Errorf("invalid unitPrice: %v", record[8])
	}
	if _, err := strconv.ParseFloat(record[9], 64); err != nil {
		return fmt.Errorf("invalid discount: %v", record[9])
	}
	if _, err := strconv.ParseFloat(record[10], 64); err != nil {
		return fmt.Errorf("invalid shippingCost: %v", record[10])
	}

	return nil

}
