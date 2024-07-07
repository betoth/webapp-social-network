package cookie

import (
	"encoding/hex"
	"log"
	"net/http"
	"os"
	"webapp-social-network/src/config"

	"github.com/gorilla/securecookie"
	"github.com/joho/godotenv"
)

var s *securecookie.SecureCookie = securecookie.New(config.HashKey, config.BlockKey)

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	hashKeyHex := os.Getenv("HASH_KEY")
	blockKeyHex := os.Getenv("BLOCK_KEY")

	// Ensure keys are valid hex strings
	if _, err := hex.DecodeString(hashKeyHex); err != nil {
		log.Fatalf("Invalid HASH_KEY: %v", err)
	}
	if _, err := hex.DecodeString(blockKeyHex); err != nil {
		log.Fatalf("Invalid BLOCK_KEY: %v", err)
	}

	// Decode hex strings to byte slices
	hashKey, err := hex.DecodeString(hashKeyHex)
	if err != nil {
		log.Fatalf("Failed to decode HASH_KEY: %v", err)
	}
	blockKey, err := hex.DecodeString(blockKeyHex)
	if err != nil {
		log.Fatalf("Failed to decode BLOCK_KEY: %v", err)
	}

	// Validate blockKey length
	if len(blockKey) != 16 && len(blockKey) != 24 && len(blockKey) != 32 {
		log.Fatalf("BLOCK_KEY has invalid size: %d", len(blockKey))
	}

	s = securecookie.New(hashKey, blockKey)
}

// Save saves a cookie in browser
func Save(w http.ResponseWriter, ID, token string) error {

	data := map[string]string{
		"id":    ID,
		"token": token,
	}

	EncryptedData, err := s.Encode("data", data)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "data",
		Value:    EncryptedData,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}

// Read return values of a cookie
func Read(r *http.Request) (map[string]string, error) {
	cookie, err := r.Cookie("data")
	if err != nil {
		return nil, err
	}

	cookieValue := make(map[string]string)

	err = s.Decode("data", cookie.Value, &cookieValue)
	if err != nil {
		return nil, err
	}

	return cookieValue, nil
}
