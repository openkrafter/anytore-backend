package auth

var PassHasher PasswordHasher

func InitPassHasher() {
	PassHasher = &BcryptHasher{}
}
