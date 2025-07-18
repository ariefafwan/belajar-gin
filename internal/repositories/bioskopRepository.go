// internal/repositories/bioskop_repository.go
package repositories

import (
	"belajar-gin/internal/models"

	"gorm.io/gorm"
)

type BioskopRepository interface {
    GetAll() ([]models.Bioskop, error)
    GetByID(id string) (*models.Bioskop, error)
    Create(bioskop *models.Bioskop) error
    Update(bioskop *models.Bioskop) error
    Delete(id string) error
}

type bioskopRepo struct {
    db *gorm.DB
}

func NewBioskopRepository(db *gorm.DB) BioskopRepository {
    return &bioskopRepo{db}
}

func (r *bioskopRepo) GetAll() ([]models.Bioskop, error) {
    var bioskops []models.Bioskop
    err := r.db.Find(&bioskops).Error
    return bioskops, err
}

func (r *bioskopRepo) GetByID(id string) (*models.Bioskop, error) {
    var bioskop models.Bioskop
    err := r.db.First(&bioskop, "id = ?", id).Error
    if err != nil {
        return nil, err
    }
    return &bioskop, nil
}

func (r *bioskopRepo) Create(bioskop *models.Bioskop) error {
    return r.db.Create(bioskop).Error
}

func (r *bioskopRepo) Update(bioskop *models.Bioskop) error {
    return r.db.Save(bioskop).Error
}

func (r *bioskopRepo) Delete(id string) error {
    return r.db.Delete(&models.Bioskop{}, "id = ?", id).Error
}
