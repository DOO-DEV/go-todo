package password

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytePassword), err
}

func ComparePassword(hashedPasswd, originalPasswd string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPasswd), []byte(originalPasswd))
}
