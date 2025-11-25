package ultils

import (
	"crypto/sha256"
	"fmt"
)

func HashingCodeChallenge(code string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(code)))
}
