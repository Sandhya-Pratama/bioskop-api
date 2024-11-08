package handler

import (
	"net/http"
	"strconv"

	"github.com/Sandhya-Pratama/bioskop-api/entity"
	"github.com/Sandhya-Pratama/bioskop-api/internal/service"
	"github.com/labstack/echo/v4"
)

type JadwalHandler struct {
	jadwalService service.JadwalTayangUsecase
}

func NewJadwalHandler(jadwalService service.JadwalTayangUsecase) *JadwalHandler {
	return &JadwalHandler{jadwalService}
}

func (h *JadwalHandler) CreateJadwal(ctx echo.Context) error {
	// Define the expected input structure for creating a new schedule
	var input struct {
		FilmID        int    `json:"film_id" validate:"required"`
		BioskopID     int    `json:"bioskop_id" validate:"required"`
		TanggalTayang string `json:"tanggal_tayang" validate:"required"` // Menggunakan string dulu
		WaktuTayang   string `json:"waktu_tayang" validate:"required"`
	}

	// Error checking with binding and validation
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	// Create new JadwalTayang entity using the input values as a pointer
	jadwal := entity.NewJadwal(input.FilmID, input.BioskopID, input.TanggalTayang, input.WaktuTayang)
	// Call the service to create the jadwal in the database
	err := h.jadwalService.CreateJadwal(ctx.Request().Context(), jadwal)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	// Return success message indicating jadwal creation
	return ctx.JSON(http.StatusCreated, "Jadwal created successfully")
}

func (h *JadwalHandler) GetJadwalByID(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid id"})
	}

	jadwal, err := h.jadwalService.GetJadwalByID(ctx.Request().Context(), id)

	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Jadwal not found"})
	}
	// Return success message get jadwal
	return ctx.JSON(http.StatusOK, jadwal)
}

func (h *JadwalHandler) UpdateJadwal(ctx echo.Context) error {
	// Define the expected input structure for updating a jadwal
	var input struct {
		FilmID        int    `json:"film_id" validate:"required"`
		BioskopID     int    `json:"bioskop_id" validate:"required"`
		TanggalTayang string `json:"tanggal_tayang" validate:"required"`
		WaktuTayang   string `json:"waktu_tayang" validate:"required"`
	}

	// Error checking with binding and validation
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	// Create new JadwalTayang entity using the input values as a pointer
	jadwal := entity.UpdateJadwal(input.FilmID, input.BioskopID, input.TanggalTayang, input.WaktuTayang)
	// Call the service to update the jadwal in the database
	err := h.jadwalService.UpdateJadwal(ctx.Request().Context(), jadwal)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	// Return success message indicating jadwal update
	return ctx.JSON(http.StatusOK, "Jadwal updated successfully")
}

func (h *JadwalHandler) DeleteJadwal(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid id"})
	}

	err = h.jadwalService.DeleteJadwal(ctx.Request().Context(), id)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": "Jadwal not found"})
	}
	// Return success message indicating jadwal deletion
	return ctx.JSON(http.StatusOK, "Jadwal deleted successfully")
}
