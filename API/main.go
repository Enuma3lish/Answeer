package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// 定義匯率資料結構體
type ExchangeRates struct {
	Rates map[string]map[string]float64
}

// 介面
type CurrencyExchangeService interface {
	GetRate(source, target string) (float64, error)
	Convert(source, target string, amount float64) (float64, error)
}

// 貨幣兌換服務
type currencyExchangeService struct {
	rates ExchangeRates
}

// 透過依賴注入匯率資料
func NewCurrencyExchangeService(rates ExchangeRates) CurrencyExchangeService {
	return &currencyExchangeService{rates: rates}
}

// 實現 GetRate 方法
func (s *currencyExchangeService) GetRate(source, target string) (float64, error) {
	sourceRates, exists := s.rates.Rates[source]
	if !exists {
		return 0, fmt.Errorf("rate for source currency %s not found", source)
	}
	rate, exists := sourceRates[target]
	if !exists {
		return 0, fmt.Errorf("rate for target currency %s not found", target)
	}
	return rate, nil
}

func (s *currencyExchangeService) Convert(source, target string, amount float64) (float64, error) {
	rate, err := s.GetRate(source, target)
	if err != nil {
		return 0, err
	}
	converted := amount * rate
	return round(converted, 2), nil
}

// 四捨五入到指定的小數位
func round(val float64, precision int) float64 {
	format := fmt.Sprintf("%%.%df", precision)
	str := fmt.Sprintf(format, val)
	result, _ := strconv.ParseFloat(str, 64)
	return result
}

// 格式化數字為千分位表示
func formatWithComma(val float64) string {
	parts := strings.Split(fmt.Sprintf("%.2f", val), ".")
	intPart := parts[0]
	decPart := parts[1]

	var result strings.Builder
	for i, digit := range intPart {
		if i > 0 && (len(intPart)-i)%3 == 0 {
			result.WriteString(",")
		}
		result.WriteRune(digit)
	}
	return result.String() + "." + decPart
}

// 定義 API 的響應結構體
type RateResponse struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Amount string `json:"amount"`
	Msg    string `json:"msg,omitempty"`
	Error  string `json:"error,omitempty"`
}

// API 處理函數
func rateHandler(service CurrencyExchangeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		source := r.URL.Query().Get("source")
		target := r.URL.Query().Get("target")
		amountStr := r.URL.Query().Get("amount")

		amount, err := strconv.ParseFloat(strings.ReplaceAll(amountStr, ",", ""), 64)
		if err != nil {
			http.Error(w, "Invalid amount format", http.StatusBadRequest)
			return
		}

		converted, err := service.Convert(source, target, amount)
		var response RateResponse
		if err != nil {
			response = RateResponse{
				Source: source,
				Target: target,
				Amount: amountStr,
				Error:  err.Error(),
			}
			w.WriteHeader(http.StatusNotFound)
		} else {
			response = RateResponse{
				Source: source,
				Target: target,
				Amount: formatWithComma(converted),
				Msg:    "success",
			}
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}

func main() {
	// 靜態匯率資料
	rates := ExchangeRates{
		Rates: map[string]map[string]float64{
			"TWD": {
				"TWD": 1,
				"JPY": 3.669,
				"USD": 0.03281,
			},
			"JPY": {
				"TWD": 0.26956,
				"JPY": 1,
				"USD": 0.00885,
			},
			"USD": {
				"TWD": 30.444,
				"JPY": 111.801,
				"USD": 1,
			},
		},
	}

	// 使用依賴注入創建 CurrencyExchangeService 實例
	exchangeService := NewCurrencyExchangeService(rates)

	// 設置路由和處理器
	http.HandleFunc("/rate", rateHandler(exchangeService))

	// 啟動 HTTP 服務器
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed to start:", err)
	}
}
