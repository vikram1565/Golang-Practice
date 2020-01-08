package main

import (
	"fmt"

	"github.com/tidwall/gjson"
)

func main() {
	jtext := `[
		{"ALCCode":"37210079","ALCMobileNo":"9730462910","ALCName":"ERA Computer Center","AllocatedSMS":1010,"UtilizedSMS":1003,"BalanceSMS":7,"Am":""},
		{"ALCCode":"13210041","ALCMobileNo":"9922402753","ALCName":"Shashi Computer Academy","AllocatedSMS":51,"UtilizedSMS":51,"BalanceSMS":0,"Am":""},
		{"ALCCode":"13-000-000-864","ALCMobileNo":null,"ALCName":" Goverment polytechnic","AllocatedSMS":500,"UtilizedSMS":0,"BalanceSMS":0,"Am":""},
		{"ALCCode":"3-44-000-131","ALCMobileNo":null,"ALCName":"ZILS NAC","AllocatedSMS":500,"UtilizedSMS":500,"BalanceSMS":0,"Am":""},
		{"ALCCode":"7868","ALCMobileNo":null,"ALCName":"Balaji Computer Education","AllocatedSMS":500,"UtilizedSMS":500,"BalanceSMS":0,"Am":""},
		{"ALCCode":"5675","ALCMobileNo":null,"ALCName":"Navdeep Infoworld","AllocatedSMS":500,"UtilizedSMS":500,"BalanceSMS":0,"Am":""},
		{"ALCCode":"9879","ALCMobileNo":null,"ALCName":"CHITNIS COMPUTER SERVICES","AllocatedSMS":500,"UtilizedSMS":500,"BalanceSMS":0,"Am":""}
		]`
	rs := gjson.Parse(jtext)

	fmt.Print(rs.Value())
}
