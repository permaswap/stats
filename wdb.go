package stats

import (
	"encoding/json"
	"time"

	"github.com/permaswap/stats/schema"
	"gorm.io/datatypes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type WDB struct {
	db *gorm.DB
}

func NewWDB(dsn string) *WDB {
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	return &WDB{db}
}

func (w *WDB) Migrate() {
	w.db.AutoMigrate(&schema.StatsSnapshot{})
}

func (w *WDB) LoadStats() (statSnapshot *schema.StatsSnapshot, err error) {
	err = w.db.Order("date desc").First(&statSnapshot).Error
	return
}

func (w *WDB) LoadAllStats() (statSnapshots []*schema.StatsSnapshot, err error) {
	err = w.db.Order("date desc").Find(&statSnapshots).Error
	return
}

func (w *WDB) SaveStatsSnapshot(stats *schema.Stats, tx *gorm.DB) (err error) {
	if tx == nil {
		tx = w.db
	}
	snapshot, err := json.Marshal(stats)
	if err != nil {
		return
	}
	err = tx.Where("date = ?", datatypes.Date(stats.Date)).First(&schema.StatsSnapshot{}).Error
	if err == nil {
		return tx.Model(&schema.StatsSnapshot{}).Where("date = ?", datatypes.Date(stats.Date)).
			Update("Stats", string(snapshot)).Error
	} else if err == gorm.ErrRecordNotFound {
		return tx.Create(&schema.StatsSnapshot{
			Date:  datatypes.Date(stats.Date),
			Stats: string(snapshot),
		}).Error
	} else {
		return
	}
}

func (w *WDB) FindStatsSnapshot(date time.Time) (statSnapshot *schema.StatsSnapshot, err error) {
	err = w.db.Model(&schema.StatsSnapshot{}).Where("date = ?", datatypes.Date(date)).First(&statSnapshot).Error
	return
}
