package service

func GenerateToken(userId int) (string, error) {
	return generateToken(userId)
}
