package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
)

// var (
// 	KEY string
// )

// func loadEnv() error {
// 	// Get the current working directory
// 	cwd, err := os.Getwd()
// 	if err != nil {
// 		return fmt.Errorf("error getting current working directory: %v", err)
// 	}

// 	// Get the directory of the executable
// 	ex, err := os.Executable()
// 	if err != nil {
// 		return fmt.Errorf("error getting executable path: %v", err)
// 	}
// 	exPath := filepath.Dir(ex)

// 	// List of possible locations for .env file
// 	envLocations := []string{
// 		filepath.Join(cwd, ".env"),
// 		filepath.Join(exPath, ".env"),
// 		"/var/www/ilits2025-backend/.env", // Add the expected location when run as a service
// 	}

// 	// Try to load .env from each location
// 	for _, loc := range envLocations {
// 		err := godotenv.Load(loc)
// 		if err == nil {
// 			fmt.Printf("Loaded .env from: %s\n", loc)
// 			return nil
// 		}
// 	}

// 	return fmt.Errorf("no .env file found in any of the expected locations")
// }

// func Init() error {
// 	if err := loadEnv(); err != nil {
// 		return err
// 	}

// 	// Get the AES key from environment variable after loading .env
// 	KEY = os.Getenv("AES_KEY")
// 	if KEY == "" {
// 		return errors.New("AES_KEY not set in environment")
// 	}

// 	return nil
// }

// // https://www.melvinvivas.com/how-to-encrypt-and-decrypt-data-using-aes

// )

func AESEncrypt(stringToEncrypt string) (encryptedString string, err error) {
	//Since the key is in string, we need to convert decode it to bytes
	// if err := Init(); err != nil {
	// 	return "", err
	// }
	KEY := os.Getenv("AES_KEY")
	key, _ := hex.DecodeString(KEY)
	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext), nil
}

func AESDecrypt(encryptedString string) (decryptedString string, err error) {
	// if err := Init(); err != nil {
	// 	return "", err
	// }
	defer func() {
		if r := recover(); r != nil {
			decryptedString = ""
			err = errors.New("error in decrypting")
		}
	}()

	KEY := os.Getenv("AES_KEY")
	key, _ := hex.DecodeString(KEY)
	enc, _ := hex.DecodeString(encryptedString)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", nil
	}

	return string(plaintext), nil
}
