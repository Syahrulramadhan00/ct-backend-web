package Services

import (
	"ct-backend/Model/Dto"
	"ct-backend/Repository"
)

type (
	IAnalyticService interface {
		GetRevenueStream() (Dto.ChartData, error)
		GetStockMonitoring() (Dto.ChartData, error)
		GetHighestSales() (Dto.ChartData, error)
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

func (s *AnalyticService) GetRevenueStream() (Dto.ChartData, error) {
	return s.AnalyticRepository.GetRevenueStream()
}

func (s *AnalyticService) GetStockMonitoring() (Dto.ChartData, error) {
	return s.AnalyticRepository.GetStockMonitoring()
}

func (s *AnalyticService) GetHighestSales() (Dto.ChartData, error) {
	return s.AnalyticRepository.GetHighestSales()
}
