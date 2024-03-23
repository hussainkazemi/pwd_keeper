package password

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	random "math/rand/v2"
	"pwsd_keeper/model"
)

const (
	key = "a very very very very secret key"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func Encrypt(text []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	b := base64.StdEncoding.EncodeToString(text)
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))
	return ciphertext, nil
}

func Decrypt(text []byte) ([]byte, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	if len(text) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := text[:aes.BlockSize]
	text = text[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(text, text)
	data, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GenerateRandomPassword(rPasswordModel model.RandomPassword) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyz")
	var capital = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	var numbers = []rune("0123456789")
	var signs = []rune("!@#$%^&*()_+")
	b := make([]rune, rPasswordModel.Length)
	j := 0
	for i := uint8(0); i < rPasswordModel.N_letter; i++ {
		b[j] = letter[random.IntN(len(letter))]
		j++
	}
	for i := uint8(0); i < rPasswordModel.N_capita; i++ {
		b[j] = capital[random.IntN(len(capital))]
		j++
	}
	for i := uint8(0); i < rPasswordModel.N_number; i++ {
		b[j] = numbers[random.IntN(len(numbers))]
		j++
	}
	for i := uint8(0); i < rPasswordModel.N_signs; i++ {
		b[j] = signs[random.IntN(len(signs))]
		j++
	}

	random.Shuffle(int(rPasswordModel.Length), func(i, j int) { b[i], b[j] = b[j], b[i] })

	return string(b)
}
