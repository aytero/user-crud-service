package service

import (
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func ComparePasswords(hashedPwd string, plainPwd []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Error().Msgf("%v\n", err)
		return false
	}
	return true
}

func HashPassword(password string) (string, error) {
	hpass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hpass), nil
}

func CreateFile(filename string, data string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}

func DeleteFile(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return err
	}
	return nil
}
