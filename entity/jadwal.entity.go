package entity

import (
	"time"
)

type JadwalTayang struct {
	JadwalID      uint      `gorm:"primaryKey" json:"jadwal_id"`
	FilmID        int       `json:"film_id"`
	BioskopID     int       `json:"bioskop_id"`
	TanggalTayang time.Time `json:"tanggal_tayang"`
	WaktuTayang   time.Time `json:"waktu_tayang"`
}

func NewJadwal() {

}
