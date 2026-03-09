package tv_models

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// TVConfig is a singleton row that holds the GCX TV display configuration.
// Playlist and Images are stored as JSON columns in MySQL.
type TVConfig struct {
	ID                 uint           `json:"-" gorm:"primaryKey"`
	NowPlaying         *string        `json:"nowPlaying"`
	NowPlayingId       *string        `json:"nowPlayingId"`
	AutoNext           bool           `json:"autoNext"`
	Loop               bool           `json:"loop"`
	Playlist           datatypes.JSON `json:"playlist" gorm:"type:json"`
	Images             datatypes.JSON `json:"images" gorm:"type:json"`
	VideoDuration      int            `json:"videoDuration"`
	MarketDataDuration int            `json:"marketDataDuration"`
	ImageDuration      int            `json:"imageDuration"`
	EnableRotation     bool           `json:"enableRotation"`
	CreatedAt          time.Time      `json:"-"`
	UpdatedAt          time.Time      `json:"-"`
	DeletedAt          gorm.DeletedAt `json:"-" gorm:"index"`
}

func (TVConfig) TableName() string {
	return "tv_config"
}
