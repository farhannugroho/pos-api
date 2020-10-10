package model

type User struct {
	Model
	Code        string   `json:"code"`
	Name        string   `json:"name"`
	ImageUrl    string   `json:"image_url"`
	Username    string   `json:"username"`
	Password    string   `json:"password"`
	Pin         string   `json:"pin"`
	Address     string   `json:"address"`
	Phone       string   `json:"phone"`
	Email       string   `gorm:"unique" json:"email"`
	CompanyId   int      `json:"company_id"`
	UserRolesId int      `json:"user_roles_id"`
	UserRole    UserRole `gorm:"foreignKey: UserRolesId" json:"user_role"`
	IsActive    bool     `json:"is_active"`
	IsSuperUser bool     `json:"is_super_user"`
}
