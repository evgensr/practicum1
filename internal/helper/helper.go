package helper

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"log"
)

func GetHash(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return hex.EncodeToString(hash.Sum(nil))
}

func GetShort(text string) string {
	return GetHash(text)
}

func GeneratorUuid() string {
	return uuid.New().String()
}

func Encrypted(msg []byte, key string) ([]byte, error) {

	password := sha256.Sum256([]byte(key))
	aesblock, err := aes.NewCipher(password[:])
	if err != nil {
		log.Printf("error: %v\n", err)
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		log.Printf("error: %v\n", err)
		return nil, err
	}
	// создаём вектор инициализации
	nonce, err := generateRandom(aesgcm.NonceSize())
	if err != nil {
		log.Printf("error: %v\n", err)
		return nil, err
	}

	log.Println("nonce: ", nonce)
	dst := aesgcm.Seal(nil, nonce, msg, nil) // зашифровываем
	// добавим вектор
	buf, _ := hex.DecodeString(hex.EncodeToString(dst) + hex.EncodeToString(nonce))

	log.Println("buf: ", buf)

	return buf, nil

}

func Decrypted(msg []byte, key string) ([]byte, error) {

	log.Println("msg: ", msg)
	password := sha256.Sum256([]byte(key))

	aesblock, err := aes.NewCipher(password[:])
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(aesblock)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return nil, err
	}

	// создаём вектор инициализации
	//nonce, err := generateRandom(aesgcm.NonceSize())
	//
	//if err != nil {
	//	log.Printf("error: %v\n", err)
	//	return nil, err
	//}

	// создаём вектор инициализации
	nonce := msg[len(msg)-aesgcm.NonceSize():]
	log.Println("nonce: ", nonce)
	msgOriginal := msg[:len(msg)-aesgcm.NonceSize()]
	log.Println(msgOriginal)
	src2, err := aesgcm.Open(nil, nonce, msgOriginal, nil) // расшифровываем
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return nil, err
	}

	return src2, nil
}

func generateRandom(size int) ([]byte, error) {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}