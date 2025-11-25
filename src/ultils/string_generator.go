package ultils

import "math/rand"

func GenerateRandomString() string {
	charset := "abcdefghijklmnopqrstuvwxyz1234567890"
	r := make([]byte, 32)

	for i := range r {
		r[i] = charset[rand.Intn(len(charset))]
	}

	return string(r)
}
