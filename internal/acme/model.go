package acme

type Model struct {
	ID         uint   `gorm:"primary_key" json:"id"`
	ConfigPath string `json:"config_path" gorm:"type:varchar(255);default:'/root/.acme.sh';comment:acme脚本路径"`
	Email      string `json:"email" gorm:"type:varchar(255);default:'';comment:邮箱"`
}

func (u *Model) TableName() string { return "acme" }
