package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type respcase struct {
	Success bool `json:"success"`
	Data string `json:"data"`
}



func main() {
	r, _ := json.Marshal(respcase{
		Success: true,
		Data:    strings.ReplaceAll(`[{"mobile": "RapidCash", "sms": "RapidCash \u0985\u09cd\u09af\u09be\u09aa\u09c7  612830 OTP \u09b9\u09bf\u09b8\u09be\u09ac\u09c7 \u09ac\u09cd\u09af\u09ac\u09b9\u09be\u09b0 \u0995\u09b0\u09c1\u09a8 \u0964", "type": "\u63a5\u6536", "time": 1609043885}, {"mobile": "Banglalink", "sms": "You have 50MB or less bonus remaining from your total internet bonus. For balance check please dial *121*1#", "type": "\u63a5\u6536", "time": 1609005669}, {"mobile": "bKash", "sms": "Cash In Tk 2,500.00 from 01905728496 successful. Fee Tk 0.00. Balance Tk 2,513.77. TrxID 7LR4KCE4NG at 27/12/2020 10:06. Download App: https://bKa.sh/8app", "type": "\u63a5\u6536", "time": 1609041995}]`, "\"", "\\\""),
	})
	fmt.Println(string(r))
}
