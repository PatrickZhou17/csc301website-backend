package orm

import "gorm.io/gorm"

func Paginate(page, pageSize int, isOffset bool) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 && !isOffset {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		if isOffset {
			offset = page
		}

		return db.Offset(offset).Limit(pageSize)
	}
}

func ScopeTime(createAt, createAtEnd, updateAt, updateAtEnd string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if createAt != "" {
			db.Where("created_at >= ?", createAt)
		}
		if createAtEnd != "" {
			db.Where("created_at <= ?", createAtEnd)
		}
		if updateAt != "" {
			db.Where("updated_at >= ?", updateAt)
		}
		if updateAtEnd != "" {
			db.Where("updated_at <= ?", updateAtEnd)
		}
		return db
	}
}
