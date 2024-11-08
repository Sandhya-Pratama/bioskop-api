package entity

type Jadwal_Tayang struct {
	JadwalID      int    `gorm:"primaryKey" json:"jadwal_id"`
	FilmID        int    `json:"film_id"`
	BioskopID     int    `json:"bioskop_id"`
	TanggalTayang string `json:"tanggal_tayang"`
	WaktuTayang   string `json:"waktu_tayang"`
}

func NewJadwal(FilmID int, BioskopID int, TanggalTayang, WaktuTayang string) *Jadwal_Tayang {
	return &Jadwal_Tayang{
		FilmID:        FilmID,
		BioskopID:     BioskopID,
		TanggalTayang: TanggalTayang,
		WaktuTayang:   WaktuTayang,
	}
}

func UpdateJadwal(FilmID int, BioskopID int, TanggalTayang, WaktuTayang string) *Jadwal_Tayang {
	return &Jadwal_Tayang{
		FilmID:        FilmID,
		BioskopID:     BioskopID,
		TanggalTayang: TanggalTayang,
		WaktuTayang:   WaktuTayang,
	}
}

func (Jadwal_Tayang) TableName() string {
	return "Jadwal_Tayang"
}
