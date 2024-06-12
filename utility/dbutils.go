package utility

import (
	"gorm.io/gorm"
)

func Paginate(dataParams map[string]interface{}) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// q := r.URL.Query()
		page := dataParams["page"].(int)
		if page <= 0 {
			page = 1
		}

		pageSize := dataParams["per_page"].(int)
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
