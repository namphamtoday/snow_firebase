package authen

type FireBaseUser struct {
	Id            string `json:"id" form:"id"`
	Email         string `json:"email" form:"email"`
	EmailVerified bool   `json:"email_verified" form:"email_verified"`
	PhoneNumber   string `json:"phone_number" form:"phone_number"`
	Password      string `json:"password" form:"password"`
	DisplayName   string `json:"display_name" form:"display_name"`
	PhotoURL      string `json:"photo_url" form:"photo_url"`
	Disabled      bool   `json:"disabled" form:"disabled"`
}
