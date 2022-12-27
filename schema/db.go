package schema

import (
	"gorm.io/datatypes"
)

type StatsSnapshot struct {
	Date  datatypes.Date `gorm:"uniqueIndex:idx_name,sort:desc"`
	Stats string         `gorm:"type:text"`
}
