package controllers

import (
	"belajar-gin/config"
	"belajar-gin/internal/helpers"
	"belajar-gin/internal/models"
	"belajar-gin/internal/repositories"
	"belajar-gin/internal/request"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		helpers.Error(c, http.StatusInternalServerError, "Failed to get bioskops", err.Error())
		return
	}
	helpers.Success(c, http.StatusOK, bioskops, "List bioskop berhasil diambil")
}

func (bc *BioskopController) GetByID(c *gin.Context) {
	id := c.Param("id")
	bioskop, err := bc.Repo.GetByID(id)
	if err != nil {
		helpers.Error(c, http.StatusNotFound, "Bioskop tidak ditemukan", err.Error())
		return
	}
	helpers.Success(c, http.StatusOK, bioskop, "Bioskop ditemukan")
}

func (bc *BioskopController) Create(c *gin.Context) {
	var req request.CreateBioskopRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, "Format request tidak valid", err.Error())
		return
	}

	if err := config.Validate.Struct(req); err != nil {
		errors := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			errors[e.Field()] = e.Error()
		}
		helpers.Error(c, http.StatusBadRequest, "Validasi gagal", errors)
		return
	}

	bioskop := models.Bioskop{
		Nama:   req.Nama,
		Lokasi: req.Lokasi,
		Rating: req.Rating,
	}

	if err := bc.Repo.Create(&bioskop); err != nil {
		helpers.Error(c, http.StatusInternalServerError, "Gagal menyimpan bioskop", err.Error())
		return
	}

	helpers.Success(c, http.StatusCreated, bioskop, "Bioskop berhasil dibuat")
}

func (bc *BioskopController) Update(c *gin.Context) {
	id := c.Param("id")
	existing, err := bc.Repo.GetByID(id)
	if err != nil {
		helpers.Error(c, http.StatusNotFound, "Bioskop tidak ditemukan", err.Error())
		return
	}

	var req request.UpdateBioskopRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		helpers.Error(c, http.StatusBadRequest, "Format request tidak valid", err.Error())
		return
	}

	if err := config.Validate.Struct(req); err != nil {
		errors := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			errors[e.Field()] = e.Error()
		}
		helpers.Error(c, http.StatusBadRequest, "Validasi gagal", errors)
		return
	}

	existing.Nama = req.Nama
	existing.Lokasi = req.Lokasi
	existing.Rating = req.Rating

	if err := bc.Repo.Update(existing); err != nil {
		helpers.Error(c, http.StatusInternalServerError, "Gagal memperbarui bioskop", err.Error())
		return
	}

	helpers.Success(c, http.StatusOK, existing, "Bioskop berhasil diperbarui")
}

func (bc *BioskopController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := bc.Repo.Delete(id); err != nil {
		helpers.Error(c, http.StatusInternalServerError, "Gagal menghapus bioskop", err.Error())
		return
	}
	helpers.Success(c, http.StatusOK, nil, "Bioskop berhasil dihapus")
}
