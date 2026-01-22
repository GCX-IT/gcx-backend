package services

import (
	"gcx-cms/internal/models"

	"gorm.io/gorm"
)

type PartnerService struct {
	db *gorm.DB
}

func NewPartnerService(db *gorm.DB) *PartnerService {
	return &PartnerService{db: db}
}

// CreatePartner creates a new partner
func (s *PartnerService) CreatePartner(partner *models.Partner) error {
	return s.db.Create(partner).Error
}

// GetPartnerByID retrieves a partner by ID
func (s *PartnerService) GetPartnerByID(id uint) (*models.Partner, error) {
	var partner models.Partner
	err := s.db.First(&partner, id).Error
	if err != nil {
		return nil, err
	}
	return &partner, nil
}

// GetAllPartners retrieves all partners with pagination and filtering
func (s *PartnerService) GetAllPartners(page, limit int, search, category, status string) ([]models.Partner, int64, error) {
	var partners []models.Partner
	var total int64

	query := s.db.Model(&models.Partner{})

	// Apply search filter
	if search != "" {
		query = query.Where("name ILIKE ? OR description ILIKE ? OR email ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	// Apply category filter
	if category != "" {
		query = query.Where("category = ?", category)
	}

	// Apply status filter
	if status != "" {
		query = query.Where("status = ?", status)
	}

	// Get total count
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Apply pagination
	offset := (page - 1) * limit
	err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&partners).Error

	return partners, total, err
}

// UpdatePartner updates a partner
func (s *PartnerService) UpdatePartner(id uint, partner *models.Partner) error {
	return s.db.Model(&models.Partner{}).Where("id = ?", id).Updates(partner).Error
}

// DeletePartner soft deletes a partner
func (s *PartnerService) DeletePartner(id uint) error {
	return s.db.Delete(&models.Partner{}, id).Error
}

// GetPartnersByCategory retrieves partners by category
func (s *PartnerService) GetPartnersByCategory(category string) ([]models.Partner, error) {
	var partners []models.Partner
	err := s.db.Where("category = ? AND status = ?", category, "active").Order("name ASC").Find(&partners).Error
	return partners, err
}

// GetActivePartners retrieves all active partners
func (s *PartnerService) GetActivePartners() ([]models.Partner, error) {
	var partners []models.Partner
	err := s.db.Where("status = ?", "active").Order("category ASC, name ASC").Find(&partners).Error
	return partners, err
}
