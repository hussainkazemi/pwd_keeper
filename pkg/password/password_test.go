package password

import "testing"

func TestEncrypt(t *testing.T) {
	tCases := []string{
		"mohsen",
		"ali",
		"kazem",
		"ehsan",
	}

	for _, v := range tCases {
		encrypt, err1 := Encrypt([]byte(v))
		if err1 != nil {
			t.Fatalf("can not encrypt because %s", err1.Error())
		}
		decrypt, err2 := Decrypt(encrypt)
		if err2 != nil {
			t.Fatalf("can not decrypt because %s", err2.Error())
		}
		if string(decrypt) != v {
			t.Errorf("for %s wanted %s but get %s", v, v, string(decrypt))
		}
	}
}
