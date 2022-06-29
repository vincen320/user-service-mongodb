package web

type UserCreateRequest struct {
	Username string `validate:"required,min=6,max=15" bson:"username,omitempty" json:"username,omitempty"`
	Password string `validate:"required,min=6,max=20" bson:"password,omitempty" json:"password,omitempty"`
}
