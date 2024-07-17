package utils

import (
	"encoding/hex"
	"razor/cache"
	"razor/core/types"
	"reflect"
	"testing"
	"time"
)

func getAPIByteArray(index int) []byte {
	apiData := [][]byte{
		[]byte(`{
  "userId": 1,
  "id": 1,
  "title": "delectus aut autem",
  "completed": false
}`),
		[]byte(`{
  "postId": 1,
  "id": 1,
  "name": "id labore ex et quam laborum",
  "email": "Eliseo@gardner.biz",
  "body": "laudantium enim quasi est quidem magnam voluptate ipsam eos\ntempora quo necessitatibus\ndolor quam autem quasi\nreiciendis et nam sapiente accusantium"
}`),
	}
	return apiData[index]
}

func TestGetDataFromAPI(t *testing.T) {
	//postRequestInput := `{"type": "POST","url": "https://staging-v3.skalenodes.com/v1/staging-aware-chief-gianfar","body": {"jsonrpc": "2.0","method": "eth_chainId","params": [],"id": 0},"header": {"content-type": "application/json"}}`
	sampleChainId, _ := hex.DecodeString("7b226a736f6e727063223a22322e30222c22726573756c74223a223078616133376463222c226964223a307d0a")

	type args struct {
		urlStruct types.DataSourceURL
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "TODO API",
			args: args{
				urlStruct: types.DataSourceURL{
					Type:   "GET",
					URL:    "https://jsonplaceholder.typicode.com/todos/1",
					Body:   nil,
					Header: nil,
				},
			},
			want:    getAPIByteArray(0),
			wantErr: false,
		},
		{
			name: "Comments API",
			args: args{
				urlStruct: types.DataSourceURL{Type: "GET",
					URL:    "https://jsonplaceholder.typicode.com/comments/1",
					Body:   nil,
					Header: nil,
				},
			},
			want:    getAPIByteArray(1),
			wantErr: false,
		},
		{
			name: "When API is invalid",
			args: args{
				urlStruct: types.DataSourceURL{
					Type:   "GET",
					URL:    "https:api.gemini.com/v1/pubticker",
					Body:   nil,
					Header: nil,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "When URL is invalid and there is an error in http.NewRequest()",
			args: args{
				urlStruct: types.DataSourceURL{
					Type: "GET",
					URL:  "\x00",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "When API is not responding",
			args: args{
				urlStruct: types.DataSourceURL{
					Type:   "GET",
					URL:    "https://api.gemini.com/v1/pubticker/TEST",
					Body:   nil,
					Header: nil,
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Post request to fetch chainId",
			args: args{
				urlStruct: types.DataSourceURL{
					Type:   "POST",
					URL:    "https://sepolia.optimism.io",
					Body:   map[string]interface{}{"jsonrpc": "2.0", "method": "eth_chainId", "params": nil, "id": 0},
					Header: map[string]string{"content-type": "application/json"},
				},
			},
			want: sampleChainId,
		},
		{
			name: "Header requires API_KEY from environment but its not present",
			args: args{
				urlStruct: types.DataSourceURL{
					Type:   "POST",
					URL:    "https://sepolia.optimism.io",
					Body:   map[string]interface{}{"jsonrpc": "2.0", "method": "eth_chainId", "params": nil, "id": 0},
					Header: map[string]string{"auth": "${API_KEY}", "content-type": "application/json"},
				},
			},
			want: sampleChainId,
		},
		{
			name: "Body for POST request is incorrect",
			args: args{
				urlStruct: types.DataSourceURL{
					Type: "POST",
					URL:  "https://sepolia.optimism.io",
					Body: map[string]interface{}{"fail": func() {}, "jsonrpc": 1},
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid request type",
			args: args{
				urlStruct: types.DataSourceURL{
					Type: "",
					URL:  "https://sepolia.optimism.io",
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Gneerating cache key throws error",
			args: args{
				urlStruct: types.DataSourceURL{
					URL: "http://example.com",
					Body: map[string]interface{}{
						"key": func() {}, // functions cannot be marshaled and will cause an error
					},
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			localCache := cache.NewLocalCache(time.Second * 10)
			got, err := GetDataFromAPI(tt.args.urlStruct, localCache)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDataFromAPI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDataFromAPI() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDataFromJSON(t *testing.T) {
	type args struct {
		jsonObject map[string]interface{}
		selector   string
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name: "ETH-USD",
			args: args{
				jsonObject: map[string]interface{}{"bid": "2695.07", "ask": "2696.23", "volume": map[string]interface{}{"ETH": "46896.86087403", "USD": "121448883.148503684", "timestamp": 1622629500000}, "last": "2697.15"},
				selector:   "last",
			},
			want:    "2697.15",
			wantErr: false,
		},
		{
			name: "BTC-USD",
			args: args{
				jsonObject: map[string]interface{}{"bid": "37179.05", "ask": "37196.47", "volume": map[string]interface{}{"BTC": "2375.5393065136", "USD": "86722499.942276466292", "timestamp": 1622629800000}, "last": "37176.33"},
				selector:   "last",
			},
			want:    "37176.33",
			wantErr: false,
		},
		{
			name: "MSFT-USD",
			args: args{
				jsonObject: map[string]interface{}{
					"Global Quote": map[string]interface{}{
						"01. symbol":             "MSFT",
						"02. open":               "251.2300",
						"03. high":               "251.2900",
						"04. low":                "246.9600",
						"05. price":              "247.4000",
						"06. volume":             "23213310",
						"07. latest trading day": "2021-06-01",
						"08. previous close":     "249.6800",
						"09. change":             "-2.2800",
						"10. change percent":     "-0.9132%",
					},
				},
				selector: `["Global Quote"]["05. price"]`,
			},
			want:    "247.4000",
			wantErr: false,
		},
		{
			name: "nth nesting",
			args: args{
				jsonObject: map[string]interface{}{
					"id":       1,
					"name":     "Leanne Graham",
					"username": "Bret",
					"email":    "Sincere@april.biz",
					"address": map[string]interface{}{
						"street":  "Kulas Light",
						"suite":   "Apt. 556",
						"city":    "Gwenborough",
						"zipcode": "92998-3874",
						"geo": map[string]interface{}{
							"lat": "-37.3159",
							"lng": "81.1496",
						},
					},
					"phone":   "1-770-736-8031 x56442",
					"website": "hildegard.org",
					"company": map[string]interface{}{
						"name":        "Romaguera-Crona",
						"catchPhrase": "Multi-layered client-server neural-net",
						"bs":          "harness real-time e-markets",
					},
				},
				selector: `address["geo"]["lng"]`,
			},
			want:    "81.1496",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDataFromJSON(tt.args.jsonObject, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDataFromJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDataFromJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDataFromHTML(t *testing.T) {
	type args struct {
		urlStruct types.DataSourceURL
		selector  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test 1: Test data from coin market cap",
			args: args{
				urlStruct: types.DataSourceURL{URL: "https://coinmarketcap.com/all/views/all/"},
				selector:  `/html/body/div[1]/div[2]/div[2]/div/div[1]/h1`,
			},
			want:    "All Cryptocurrencies",
			wantErr: false,
		},
		{
			name: "Test 2: Test for invalid website",
			args: args{
				urlStruct: types.DataSourceURL{URL: "http://razor-go.com/"},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDataFromXHTML(tt.args.urlStruct, tt.args.selector)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDataFromHTML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetDataFromHTML() got = %v, want %v", got, tt.want)
			}
		})
	}
}
