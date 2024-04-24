package request

import "motionserver/app/database/schema"

type LoginRequest struct {
	Email    string `form:"email" json:"email" validate:"required,email"`
	Password string `form:"password" json:"password" validate:"required,min=8,max=255"`
}

type RegisterRequest struct {
	Email    string       `form:"email" json:"email" validate:"required,email"`
	Name     string       `form:"name" json:"name" validate:"required,min=3,max=255"`
	Password string       `form:"password" json:"password" validate:"required,min=8,max=255"`
	Role     *schema.Role `form:"role" json:"role" validate:"omitempty,oneof=admin superadmin user"`
}
