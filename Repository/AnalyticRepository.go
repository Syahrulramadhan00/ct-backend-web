package Repository

import (
	"ct-backend/Model"
	"ct-backend/Model/Dto"
	"ct-backend/Utils"
	"time"
	"gorm.io/gorm"
	"fmt"
)

type (
	IAnalyticRepository interface {
		GetRevenueStream(startDate, endDate time.Time) (Dto.ChartData, error)
		GetStockMonitoring(yearMonth string) (Dto.ChartData, error)
		GetHighestSales(startDate time.Time, endDate time.Time) (Dto.ChartData, error)
		GetExpenses(startDate, endDate time.Time) (Dto.ChartData, error)
		GetTopSpenders() (Dto.ChartData, error)
		GetAvailableMonths(string) ([]string, []string, error)
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


func (r *AnalyticRepository) GetRevenueStream(startDate, endDate time.Time) (Dto.ChartData, error) {
	var revenues []Model.Revenue
	var chartData Dto.ChartData

	err := r.DB.Model(&Model.Sale{}).
		Select("TO_CHAR(created_at, 'YYYY-MM') as month, SUM(quantity * price) as total_revenue").
		Where("created_at >= ? AND created_at < ?", startDate, endDate).
		Group("month").
		Order("month ASC").
		Scan(&revenues).Error

	if err != nil {
		return chartData, err
	}

	chartData.Labels = make([]string, len(revenues))
	chartData.Datasets = []Dto.Dataset{{Label: "Revenue Stream", Data: make([]float64, len(revenues))}}

	for i, revenue := range revenues {
		chartData.Labels[i] = revenue.Month
		chartData.Datasets[0].Data[i] = revenue.Total
	}

	return chartData, nil
}


func (r *AnalyticRepository) GetExpenses(startDate, endDate time.Time) (Dto.ChartData, error) {
	var expenses []Model.Expenses
	var chartData Dto.ChartData

	err := r.DB.Raw(`
		SELECT 
			TO_CHAR(DATE_TRUNC('month', created_at), 'YYYY-MM-DD"T00:00:00+07:00"') AS month,
			SUM(count * price) AS total_expenses
		FROM purchases
		WHERE created_at >= ? AND created_at < ?
		GROUP BY month
		ORDER BY month ASC
	`, startDate, endDate).Scan(&expenses).Error

	if err != nil {
		return chartData, err
	}

	chartData.Labels = make([]string, len(expenses))
	chartData.Datasets = []Dto.Dataset{{Label: "Expenses", Data: make([]float64, len(expenses))}}

	for i, expense := range expenses {
		chartData.Labels[i] = expense.Month
		chartData.Datasets[0].Data[i] = expense.Total
	}

	return chartData, nil
}


func (r *AnalyticRepository) GetStockMonitoring(yearMonth string) (Dto.ChartData, error) {
	var stocks []Model.Stock
	var chartData Dto.ChartData

	err := r.DB.Model(&Model.Product{}).
		Select("name, stock").
		Where("TO_CHAR(created_at, 'YYYY-MM') = ?", yearMonth).
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


func (r *AnalyticRepository) GetHighestSales(startDate, endDate time.Time) (Dto.ChartData, error) {
	var sales []Model.HighestSales
	var chartData Dto.ChartData

	// Query the database with dynamic date range
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

	// Prepare the chart data
	chartData.Labels = make([]string, len(sales))
	chartData.Datasets = []Dto.Dataset{{Label: "Highest Sales", Data: make([]float64, len(sales))}}

	for i, sale := range sales {
		chartData.Labels[i] = sale.ProductName
		chartData.Datasets[0].Data[i] = float64(sale.Total)
	}

	return chartData, nil
}


func (r *AnalyticRepository) GetTopSpenders() (Dto.ChartData, error) {
	var topSpenders []Model.TopSpenders
	var chartData Dto.ChartData

	err := r.DB.Raw(`
		SELECT c.name AS name, SUM(i.total_price) AS total
		FROM invoices i
		JOIN clients c ON i.client_id = c.id
		WHERE i.invoice_status_id IN (SELECT id FROM invoice_statuses WHERE name = 'paid')
		GROUP BY c.id, c.name
		ORDER BY total DESC
		LIMIT 3
	`).Scan(&topSpenders).Error

	if err != nil {
		return chartData, err
	}

	chartData.Labels = make([]string, len(topSpenders))
	chartData.Datasets = []Dto.Dataset{{Label: "Top Spenders", Data: make([]float64, len(topSpenders))}}

	for i, spender := range topSpenders {
		chartData.Labels[i] = spender.Name
		chartData.Datasets[0].Data[i] = spender.Total
	}

	return chartData, nil
}

func (r *AnalyticRepository) GetAvailableMonths(tableName string) ([]string, []string, error) {
    var months []string
    var labels []string

    // Ensure tableName is valid (Prevent SQL Injection)
    validTables := map[string]string{
        "sales":     "sales",
        "purchases": "purchases",
        "products":  "products",
        "invoices":  "invoices",
    }

    table, exists := validTables[tableName]
    if !exists {
        return nil, nil, fmt.Errorf("invalid table name")
    }

    query := fmt.Sprintf(`
        SELECT DISTINCT TO_CHAR(created_at, 'YYYY-MM') AS month
        FROM %s
        ORDER BY month ASC
    `, table)

    err := r.DB.Raw(query).Scan(&months).Error
    if err != nil {
        return nil, nil, err
    }

    for _, month := range months {
		label, err := Utils.ConvertMonthToIndonesian(month)
        if err != nil {
            return nil, nil, err
        }
        labels = append(labels, label)
    }

    return months, labels, nil
}
