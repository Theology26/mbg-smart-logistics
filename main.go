package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Pengguna struct {
	gorm.Model
	Nama               string `gorm:"type:varchar(100);not null" json:"nama"`
	Email              string `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	KataSandi          string `gorm:"type:varchar(255);not null" json:"kata_sandi"`
	Peran              string `gorm:"type:enum('admin', 'kurir', 'guru');not null" json:"peran"`
	TipeKendaraan      string `gorm:"type:enum('motor', 'mobil', 'tidak_ada');default:'tidak_ada'" json:"tipe_kendaraan"`
	KapasitasKendaraan *int   `gorm:"default:0" json:"kapasitas_kendaraan"`

	Rute []Rute `gorm:"foreignKey:KurirID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"rute"`
}

type Lokasi struct {
	gorm.Model
	NamaLokasi        string  `gorm:"type:varchar(150);not null" json:"nama_lokasi"`
	TipeLokasi        string  `gorm:"type:enum('DAPUR', 'SEKOLAH', 'ESTAFET');not null" json:"tipe_lokasi"`
	Latitude          float64 `gorm:"not null" json:"latitude"`
	Longitude         float64 `gorm:"not null" json:"longitude"`
	KebutuhanBoks     *int    `gorm:"default:0" json:"kebutuhan_boks"`
	JamBuka           string  `gorm:"type:varchar(10);not null" json:"jam_buka"`
	JamTutup          string  `gorm:"type:varchar(10);not null" json:"jam_tutup"`
	WaktuLayananMenit *int    `gorm:"default:0" json:"waktu_layanan_menit"`

	PemberhentianRute []PemberhentianRute `gorm:"foreignKey:LokasiID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"pemberhentian_rute"`
}

type Rute struct {
	gorm.Model
	KurirID      uint      `gorm:"not null" json:"kurir_id"`
	Tanggal      time.Time `gorm:"type:date;not null" json:"tanggal"`
	TotalJarakKm float64   `gorm:"default:0" json:"total_jarak_km"`
	StatusRute   string    `gorm:"type:enum('TUNDA', 'PROSES', 'SELESAI');default:'TUNDA'" json:"status_rute"`

	PemberhentianRute []PemberhentianRute `gorm:"foreignKey:RuteID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"pemberhentian_rute"`
}

type PemberhentianRute struct {
	gorm.Model
	RuteID         uint      `gorm:"not null" json:"rute_id"`
	LokasiID       uint      `gorm:"not null" json:"lokasi_id"`
	UrutanBerhenti int       `gorm:"not null" json:"urutan_berhenti"`
	EstimasiTiba   time.Time `gorm:"not null" json:"estimasi_tiba"`
	BoksTurun      *int      `gorm:"default:0" json:"boks_turun"`
	BoksNaik       *int      `gorm:"default:0" json:"boks_naik"`
	ApakahEstafet  bool      `gorm:"default:false" json:"apakah_estafet"`
}

var DB *gorm.DB

func main() {
	dsn := "root:@tcp(localhost:3306)/mbg_smart_logistics?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	DB = db

	DB.AutoMigrate(&Pengguna{}, &Lokasi{}, &Rute{}, &PemberhentianRute{})

	r := gin.Default()

	r.POST("/lokasi", tambahLokasi)
	r.GET("/lokasi", lihatLokasi)

	r.Run(":8080")
}

func tambahLokasi(c *gin.Context) {
	var input Lokasi
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func lihatLokasi(c *gin.Context) {
	var lokasi []Lokasi
	if err := DB.Find(&lokasi).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": lokasi})
}
