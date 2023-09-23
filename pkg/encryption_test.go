package pkg

import (
	"encoding/base64"
	"testing"
)

func BenchmarkEncryption(b *testing.B) {
	key := GenerateRandomKey()
	encryption := &Encryption{Key: key}
	plaintext := []byte("Hello, this is a test plaintext!")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ciphertext, err := encryption.Encrypt(plaintext)
		if err != nil {
			b.Fatalf("Error encrypting: %v", err)
		}

		decrypted, err := encryption.Decrypt(ciphertext)
		if err != nil {
			b.Fatalf("Error decrypting: %v", err)
		}

		if string(decrypted) != string(plaintext) {
			b.Fatal("Decrypted text does not match original plaintext")
		}
	}
}

func TestEncryption_Encrypt(t *testing.T) {
	type fields struct {
		Key []byte
	}
	type args struct {
		plaintext []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "Encrypts plaintext",
			fields:  fields{Key: GenerateRandomKey()},
			args:    args{plaintext: []byte("Hello, this is a test plaintext!")},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encryption{
				Key: tt.fields.Key,
			}
			got, err := e.Encrypt(tt.args.plaintext)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// get base64 string
			base64String := base64.StdEncoding.EncodeToString(got)
			t.Log("Plaintext: ", string(tt.args.plaintext))
			t.Log("Encrypted string: ", base64String)
		})
	}
}
