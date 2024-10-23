package utils

import "math/rand"

const LETTERBYTES = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const NUMBERBYTES = "0123456789"
const SPECIALBYTES = "!@#$%^&*()_+"

func GenerateRandomString(length int) string {
	res := make([]byte, length)
	for i := 0; i < length; i++ {
		res[i] = LETTERBYTES[rand.Intn(len(LETTERBYTES))]
	}
	return string(res)
}
