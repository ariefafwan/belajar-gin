package controllers

import (
	"belajar-gin/internal/helpers"
	"belajar-gin/internal/models"
	"belajar-gin/internal/repositories"
	"belajar-gin/internal/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BioskopController struct {
	Repo repositories.BioskopRepository
}

func NewBioskopController(repo repositories.BioskopRepository) *BioskopController {
	return &BioskopController{Repo: repo}
}

func (bc *BioskopController) GetAll(c *gin.Context) {
	bioskops, err := bc.Repo.GetAll()
	if err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Failed to get bioskops")
		return
	}
	helpers.Success(c, http.StatusOK, bioskops, "List bioskop berhasil diambil")
}

func (bc *BioskopController) GetByID(c *gin.Context) {
	id := c.Param("id")
	bioskop, err := bc.Repo.GetByID(id)
	if err != nil {
		helpers.Error(c, http.StatusNotFound, err.Error(), "Bioskop tidak ditemukan")
		return
	}
	helpers.Success(c, http.StatusOK, bioskop, "Bioskop ditemukan")
}

func (bc *BioskopController) Create(c *gin.Context) {
	var req request.CreateBioskopRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Format request tidak valid")
		return
	}

	if errors := request.ValidateStruct(req); errors != nil {
        helpers.Error(c, http.StatusUnprocessableEntity, errors, "Validasi gagal")
        return
    }

	bioskop := models.Bioskop{
		Nama:   req.Nama,
		Lokasi: req.Lokasi,
		Rating: req.Rating,
	}

	if err := bc.Repo.Create(&bioskop); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Gagal menyimpan bioskop")
		return
	}

	helpers.Success(c, http.StatusCreated, bioskop, "Bioskop berhasil dibuat")
}

func (bc *BioskopController) Update(c *gin.Context) {
	id := c.Param("id")
	existing, err := bc.Repo.GetByID(id)
	if err != nil {
		helpers.Error(c, http.StatusNotFound, err.Error(), "Bioskop tidak ditemukan")
		return
	}

	var req request.UpdateBioskopRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, err.Error(), "Format request tidak valid")
		return
	}

	if errors := request.ValidateStruct(req); errors != nil {
        helpers.Error(c, http.StatusUnprocessableEntity, errors, "Validasi gagal")
        return
    }

	existing.Nama = req.Nama
	existing.Lokasi = req.Lokasi
	existing.Rating = req.Rating

	if err := bc.Repo.Update(existing); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Gagal memperbarui bioskop")
		return
	}

	helpers.Success(c, http.StatusOK, existing, "Bioskop berhasil diperbarui")
}

func (bc *BioskopController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := bc.Repo.Delete(id); err != nil {
		helpers.Error(c, http.StatusInternalServerError, err.Error(), "Gagal menghapus bioskop")
		return
	}
	helpers.Success(c, http.StatusOK, nil, "Bioskop berhasil dihapus")
}
