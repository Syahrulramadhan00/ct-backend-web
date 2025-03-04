package Repository

import (
	"ct-backend/Model"
	"ct-backend/Model/Dto"
	"ct-backend/Utils"
	"time"
	"gorm.io/gorm"
)

type (
	IAnalyticRepository interface {
		GetRevenueStream() (Dto.ChartData, error)
		GetStockMonitoring() (Dto.ChartData, error)
		GetHighestSales() (Dto.ChartData, error)
	}

	AnalyticRepository struct {
		DB *gorm.DB
	}
)

func AnalyticRepositoryProvider(DB *gorm.DB) *AnalyticRepository {
	return &AnalyticRepository{
		DB: DB,
	}
}

// GetRevenueStream - Fetches total revenue stream for the last six months
func (r *AnalyticRepository) GetRevenueStream() (Dto.ChartData, error) {
	var revenues []Model.Revenue
	var chartData Dto.ChartData

	startDate := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2025, 8, 1, 0, 0, 0, 0, time.UTC)

	err := r.DB.Model(&Model.Sale{}).
		Select("TO_CHAR(created_at, 'YYYY-MM') as month, SUM(quantity * price) as total_revenue").
		Where("created_at >= ? AND created_at < ?", startDate, endDate).
		Group("month").
		Order("month ASC").
		Scan(&revenues).Error

	if err != nil {
		return chartData, err
	}

	// Extract labels
	labels := make([]string, len(revenues))
	for i, revenue := range revenues {
		labels[i] = revenue.Month
	}

	// Convert labels to Indonesian month names
	indonesianLabels, err := Utils.ConvertLabelsToIndonesian(labels)
	if err != nil {
		return chartData, err
	}

	chartData.Labels = indonesianLabels
	chartData.Datasets = []Dto.Dataset{{Label: "Revenue Stream", Data: make([]float64, len(revenues))}}

	for i, revenue := range revenues {
		chartData.Datasets[0].Data[i] = revenue.Total
	}

	return chartData, nil
}

// GetStockMonitoring - Fetches the top 10 products with the highest stock
func (r *AnalyticRepository) GetStockMonitoring() (Dto.ChartData, error) {
	var stocks []Model.Stock
	var chartData Dto.ChartData

	err := r.DB.Model(&Model.Product{}).
		Select("name, stock").
		Order("stock DESC").
		Limit(5).
		Scan(&stocks).Error

	if err != nil {
		return chartData, err
	}

	chartData.Labels = make([]string, len(stocks))
	chartData.Datasets = []Dto.Dataset{{Label: "Stock Monitoring", Data: make([]float64, len(stocks))}}

	for i, stock := range stocks {
		chartData.Labels[i] = stock.Name
		chartData.Datasets[0].Data[i] = float64(stock.Stock)
	}

	return chartData, nil
}

// GetHighestSales - Fetches the top 5 highest-selling products in February 2025
func (r *AnalyticRepository) GetHighestSales() (Dto.ChartData, error) {
	var sales []Model.HighestSales
	var chartData Dto.ChartData

	startDate := time.Date(2025, 2, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2025, 3, 1, 0, 0, 0, 0, time.UTC)

	err := r.DB.Model(&Model.Sale{}).
	Select("p.name as product_name, SUM(sales.quantity) as total").
	Joins("JOIN products p ON sales.product_id = p.id").
	Where("sales.created_at >= ? AND sales.created_at < ?", startDate, endDate).
	Group("p.id, p.name").
	Order("total DESC").
	Limit(4).
	Scan(&sales).Error


	if err != nil {
		return chartData, err
	}

	chartData.Labels = make([]string, len(sales))
	chartData.Datasets = []Dto.Dataset{{Label: "Highest Sales", Data: make([]float64, len(sales))}}

	for i, sales := range sales {
		chartData.Labels[i] = sales.ProductName
		chartData.Datasets[0].Data[i] = float64(sales.Total)
	}

	return chartData, nil
}