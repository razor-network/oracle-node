package utils

import (
	"encoding/hex"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"reflect"
	"testing"
)

func TestEcRecover(t *testing.T) {
	hash, _ := hex.DecodeString("a4077fc17f659439ff7d7ce1d0735af97fbc1dcdca28dbe50667830fbffa5f85")
	signedData, _ := hex.DecodeString("b0150b4852635a700a8207121ccdeab8db2c5be683ee3f786b2e2b24cea826963be63651d8f492414edd0c9cdb8e12117acb74b9fbd3483ca6e0cbc7732e915401")

	modifiedSignedData, _ := hex.DecodeString("b0150b4852635a700a8207121ccdeab8db2c5be683ee3f786b2e2b24cea826963be63651d8f492414edd0c9cdb8e12117acb74b9fbd3483ca6e0cbc7732e915401")
	modifiedSignedData[64] = 27

	type args struct {
		data hexutil.Bytes
		sig  hexutil.Bytes
	}
	tests := []struct {
		name    string
		args    args
		want    common.Address
		wantErr bool
	}{
		{
			name: "Test 1: Recovers address successfully",
			args: args{
				data: hash,
				sig:  signedData,
			},
			want:    common.HexToAddress("0x57Baf83BAD5bee0F7F44d84669A50C35c57E3576"),
			wantErr: false,
		},
		{
			name: "Test 2: Signature length is not equal to 65",
			args: args{
				data: nil,
				sig:  nil,
			},
			want:    common.HexToAddress(""),
			wantErr: true,
		},
		{
			name: "Test 3: Cannot recover address due to error in getting public key from signature",
			args: args{
				data: hash,
				sig:  modifiedSignedData,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EcRecover(tt.args.data, tt.args.sig)
			if (err != nil) != tt.wantErr {
				t.Errorf("EcRecover() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EcRecover() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateCacheKey(t *testing.T) {
	type args struct {
		url  string
		body map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test when the request is of type GET with body as nil",
			args: args{
				url:  "https:api.gemini.com/v1/pubticker",
				body: nil,
			},
			want:    "0x9e9d7684772b773287cfd315aea97e574b7059bf63599ad1e008b9f695ab46e2",
			wantErr: false,
		},
		{
			name: "Test when the request is of type POST with body",
			args: args{
				url:  "https://staging-v3.skalenodes.com/v1/staging-aware-chief-gianfar",
				body: map[string]interface{}{"jsonrpc": "2.0", "method": "eth_chainId", "params": nil, "id": 0},
			},
			want:    "0xecd57ac05ee0584932ac9b969f6e8a851b7a5f1bbd46ae756f48ad9e2747a0ff",
			wantErr: false,
		},
		{
			name: "Test when url and body is nil",
			args: args{
				url:  "",
				body: nil,
			},
			want:    "0xc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470",
			wantErr: false,
		},
		{
			name: "Test when marshalling fails",
			args: args{
				url: "http://example.com",
				body: map[string]interface{}{
					"key": func() {}, // functions cannot be marshaled and will cause an error
				},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateCacheKey(tt.args.url, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateCacheKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("generateCacheKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}
