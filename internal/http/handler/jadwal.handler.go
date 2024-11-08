package handler

import (
	"net/http"
	"time"

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
		FilmID        int       `json:"film_id" validate:"required"`
		BioskopID     int       `json:"bioskop_id" validate:"required"`
		TanggalTayang time.Time `json:"tanggal_tayang" validate:"required"`
		WaktuTayang   time.Time `json:"waktu_tayang" validate:"required"`
	}

	// Error checking with binding and validation
	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	// Create new JadwalTayang entity using the input values as a pointer
	jadwal := &entity.JadwalTayang{
		FilmID:        input.FilmID,
		BioskopID:     input.BioskopID,
		TanggalTayang: input.TanggalTayang,
		WaktuTayang:   input.WaktuTayang,
	}

	// Call the service to create the jadwal in the database
	err := h.jadwalService.CreateJadwal(ctx.Request().Context(), jadwal)
	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, echo.Map{"error": err.Error()})
	}

	// Return success message indicating jadwal creation
	return ctx.JSON(http.StatusCreated, echo.Map{"message": "Jadwal Tayang created successfully"})
}
