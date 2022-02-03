package utils

import (
	"github.com/avast/retry-go"
	"github.com/stretchr/testify/mock"
	"razor/utils/mocks"
	"reflect"
	"testing"
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
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "TODO API",
			args:    args{url: "https://jsonplaceholder.typicode.com/todos/1"},
			want:    getAPIByteArray(0),
			wantErr: false,
		},
		{
			name:    "Comments API",
			args:    args{url: "https://jsonplaceholder.typicode.com/comments/1"},
			want:    getAPIByteArray(1),
			wantErr: false,
		},
		{
			name:    "When API is invalid",
			args:    args{url: "https:api.gemini.com/v1/pubticker"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "When API is not responding",
			args:    args{url: "https://api.gemini.com/v1/pubticker/TEST"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optionsMock := new(mocks.OptionUtils)
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				Options:        optionsMock,
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			optionsMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetDataFromAPI(tt.args.url)
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
			optionsMock := new(mocks.OptionUtils)
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				Options:        optionsMock,
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			got, err := utils.GetDataFromJSON(tt.args.jsonObject, tt.args.selector)
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
		url      string
		selector string
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
				url:      "https://coinmarketcap.com/all/views/all/",
				selector: `div h1`,
			},
			want:    "All Cryptocurrencies",
			wantErr: false,
		},
		{
			name: "Test 2: Test for invalid website",
			args: args{
				url:      "http://razor-go.com/",
				selector: `table tbody tr td a[href="/en/coins/bitcoin"]`,
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			optionsMock := new(mocks.OptionUtils)
			utilsMock := new(mocks.Utils)

			optionsPackageStruct := OptionsPackageStruct{
				Options:        optionsMock,
				UtilsInterface: utilsMock,
			}
			utils := StartRazor(optionsPackageStruct)

			got, err := utils.GetDataFromHTML(tt.args.url, tt.args.selector)
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
