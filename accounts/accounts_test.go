package accounts

import (
	"crypto/ecdsa"
	"encoding/hex"
	"reflect"
	"testing"
)

func privateKeyToHex(privateKey *ecdsa.PrivateKey) string {
	return hex.EncodeToString(privateKey.D.Bytes())
}

func Test_getPrivateKeyFromKeystore(t *testing.T) {
	password := "Razor@123"

	type args struct {
		keystoreFilePath string
		password         string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test 1: When keystore file is present and getPrivateKeyFromKeystore function executes successfully",
			args: args{
				keystoreFilePath: "test_accounts/UTC--2024-03-20T07-03-56.358521000Z--911654feb423363fb771e04e18d1e7325ae10a91",
				password:         password,
			},
			want:    "b110b1f06b7b64323a6fb768ceab966abe9f65f4e6ab3c39382bd446122f7b01",
			wantErr: false,
		},
		{
			name: "Test 2: When there is no keystore file present at the desired path",
			args: args{
				keystoreFilePath: "test_accounts/UTC--2024-03-20T07-03-56.358521000Z--211654feb423363fb771e04e18d1e7325ae10a91",
				password:         password,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test 3: When password is incorrect for the desired keystore file",
			args: args{
				keystoreFilePath: "test_accounts/UTC--2024-03-20T07-03-56.358521000Z--911654feb423363fb771e04e18d1e7325ae10a91",
				password:         "Razor@456",
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPrivateKey, err := getPrivateKeyFromKeystore(tt.args.keystoreFilePath, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPrivateKeyFromKeystore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// If there's no error and a private key is expected, compare the keys
			if !tt.wantErr && tt.want != "" {
				gotPrivateKeyHex := privateKeyToHex(gotPrivateKey)
				if gotPrivateKeyHex != tt.want {
					t.Errorf("GetPrivateKey() got private key = %v, want %v", gotPrivateKeyHex, tt.want)
				}
			}
		})
	}
}

func TestGetPrivateKey(t *testing.T) {
	password := "Razor@123"
	keystoreDirPath := "test_accounts"

	type args struct {
		address  string
		password string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test 1: When input address with correct password is present in keystore directory",
			args: args{
				address:  "0x911654feb423363fb771e04e18d1e7325ae10a91",
				password: password,
			},
			want:    "b110b1f06b7b64323a6fb768ceab966abe9f65f4e6ab3c39382bd446122f7b01",
			wantErr: false,
		},
		{
			name: "Test 2: When input upper case address with correct password is present in keystore directory",
			args: args{
				address:  "0x2F5F59615689B706B6AD13FD03343DCA28784989",
				password: password,
			},
			want:    "726223b8b95628edef6cf2774ddde39fb3ea482949c8847fabf74cd994219b50",
			wantErr: false,
		},
		{
			name: "Test 3: When provided address is not present in keystore directory",
			args: args{
				address: "0x911654feb423363fb771e04e18d1e7325ae10a91_not_present",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Test 4: When input address with incorrect password is present in keystore directory",
			args: args{
				address:  "0x911654feb423363fb771e04e18d1e7325ae10a91",
				password: "incorrect password",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			am := NewAccountManager(keystoreDirPath)
			gotPrivateKey, err := am.GetPrivateKey(tt.args.address, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPrivateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// If there's no error and a private key is expected, compare the keys
			if !tt.wantErr && tt.want != "" {
				gotPrivateKeyHex := privateKeyToHex(gotPrivateKey)
				if gotPrivateKeyHex != tt.want {
					t.Errorf("GetPrivateKey() got private key = %v, want %v", gotPrivateKeyHex, tt.want)
				}
			}
		})
	}
}

func TestSignData(t *testing.T) {
	password := "Razor@123"

	hexHash := "a3b8d42c7015c1e9354f8b9c2161d9b2e1ad89e6b6c7a9610e029fd7afec27ae"
	hashBytes, err := hex.DecodeString(hexHash)
	if err != nil {
		log.Fatal("Failed to decode hex string")
	}

	type args struct {
		address  string
		password string
		hash     []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test 1: When Sign function returns no error",
			args: args{
				address:  "0x911654feb423363fb771e04e18d1e7325ae10a91",
				password: password,
				hash:     hashBytes,
			},
			want:    "f14cf1b8c9486777e4280b4da6cb7c314d5c19b7e30d32a46f83767a1946e35a39f6941df71375d7ffceaddac81e2454e9a129896803d02f633eb78ab7883ff200",
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting private key",
			args: args{
				address:  "0x_invalid_address",
				password: password,
				hash:     hashBytes,
			},
			want:    hex.EncodeToString(nil),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			am := NewAccountManager("test_accounts")
			got, err := am.SignData(tt.args.hash, tt.args.address, tt.args.password)
			if !reflect.DeepEqual(hex.EncodeToString(got), tt.want) {
				t.Errorf("Sign() got = %v, want %v", hex.EncodeToString(got), tt.want)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
