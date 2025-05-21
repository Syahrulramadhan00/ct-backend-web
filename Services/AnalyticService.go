package Services

import (
	"ct-backend/Model/Dto"
	"ct-backend/Repository"
	"time"
)

type (
	IAnalyticService interface {
		GetRevenueStream(startDate time.Time, endDate time.Time) (Dto.ChartData, error)
		GetStockMonitoring(yearMonth string) (Dto.ChartData, error)
		GetHighestSales(startDate, endDate time.Time) (Dto.ChartData, error)
		GetExpenses(startDate time.Time, endDate time.Time) (Dto.ChartData, error)
		GetTopSpenders(yearMonth string) (Dto.ChartData, error)
		GetAvailableMonths(table string) ([]string, []string, error)
	}

	AnalyticService struct {
		AnalyticRepository Repository.IAnalyticRepository
	}
)

func AnalyticServiceProvider(analyticRepository Repository.IAnalyticRepository) *AnalyticService {
	return &AnalyticService{
		AnalyticRepository: analyticRepository,
	}
}

func (s *AnalyticService) GetRevenueStream(startDate time.Time, endDate time.Time) (Dto.ChartData, error) {
	return s.AnalyticRepository.GetRevenueStream(startDate, endDate)
}

func (s *AnalyticService) GetStockMonitoring(yearMonth string) (Dto.ChartData, error) {
	return s.AnalyticRepository.GetStockMonitoring(yearMonth)
}


func (s *AnalyticService) GetHighestSales(startDate, endDate time.Time) (Dto.ChartData, error) {
	return s.AnalyticRepository.GetHighestSales(startDate, endDate)
}


func (s *AnalyticService) GetExpenses(startDate, endDate time.Time) (Dto.ChartData, error) {
    return s.AnalyticRepository.GetExpenses(startDate, endDate)
}


func (s *AnalyticService) GetTopSpenders(yearMonth string) (Dto.ChartData, error) {
	return s.AnalyticRepository.GetTopSpenders()
}

func (s *AnalyticService) GetAvailableMonths(tableName string) ([]string, []string, error) {
    return s.AnalyticRepository.GetAvailableMonths(tableName)
}