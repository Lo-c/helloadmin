package models

import "gorm.io/gorm"

type Menu struct {
	Model
	ParentId  int    `json:"parent_id"`
	Sort      int    `json:"sort"`
	Title     string `json:"title"`
	Icon      string `json:"icon"`
	Uri       string `json:"uri"`
	Extension string `json:"extension"`
	IsShow    int8   `json:"is_show"`
}

func (Menu) TableName() string {
	return "hi_menus"
}

type MenuTree struct {
	Menu
	Children []Menu
}

func (m Menu) List(db *gorm.DB, offset, size int) ([]*Menu, error) {
	var menu []*Menu
	var err error
	if offset >= 0 && size >= 0 {
		db = db.Offset(offset).Limit(size)
	}
	if m.Title != "" {
		db = db.Where("title = ?", m.Title)
	}
	if err = db.Order("sort DESC").Find(&menu).Error; err != nil {
		return nil, err
	}
	return menu, nil
}

func (m Menu) Options(db *gorm.DB) ([]map[string]interface{}, error) {
	var menu []*Menu
	var err error
	if err = db.Order("sort DESC").Find(&menu).Error; err != nil {
		return nil, err
	}
	var options = make([]map[string]interface{}, 0, len(menu))
	for i := 0; i < len(menu); i++ {
		ops := map[string]interface{}{"key": menu[i].ID, "value": menu[i].Title}
		options = append(options, ops)
	}
	return options, nil
}
