package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {

	if strings.Contains(err, "username") {
		return errors.New("Username sudah digunakan")
	}

	if strings.Contains(err, "email") {
		return errors.New("Email sudah digunakan")
	}

	if strings.Contains(err, "record not found") {
		return errors.New("Data tidak ditemukan")
	}

	if strings.Contains(err, "gagal verifikasi") {
		return errors.New("Gagal melakukan verifikasi, waktu sudah habis")
	}

	if strings.Contains(err, "title") {
		return errors.New("Title Already Taken")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("Password salah...")
	}
	return errors.New("not found")
}
