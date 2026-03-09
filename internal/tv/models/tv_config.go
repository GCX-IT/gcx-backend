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
	AutoNext           bool           `json:"autoNext" gorm:"default:true"`
	Loop               bool           `json:"loop" gorm:"default:true"`
	Playlist           datatypes.JSON `json:"playlist" gorm:"type:json;default:'[]'"`
	Images             datatypes.JSON `json:"images" gorm:"type:json;default:'[]'"`
	VideoDuration      int            `json:"videoDuration" gorm:"default:60"`
	MarketDataDuration int            `json:"marketDataDuration" gorm:"default:10"`
	ImageDuration      int            `json:"imageDuration" gorm:"default:120"`
	EnableRotation     bool           `json:"enableRotation" gorm:"default:false"`
	CreatedAt          time.Time      `json:"-"`
	UpdatedAt          time.Time      `json:"-"`
	DeletedAt          gorm.DeletedAt `json:"-" gorm:"index"`
}

func (TVConfig) TableName() string {
	return "tv_config"
}
