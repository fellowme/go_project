package remote_web

/*
	SendPhoneCode 没有企业验证无法发送验证码  待验证
	https://next.api.aliyun.com/api/Dysmsapi/2017-05-25/SendBatchSms?params={}
*/
func SendPhoneCode(code string, mobile string) error {
	//paramByte, _ := json.Marshal(map[string]string{"code": code})
	//sendSmsRequest := dysmsapi20170525.SendSmsRequest{
	//	SignName:      tea.String("阿里云短信测试"),
	//	TemplateCode:  tea.String("SMS_154950909"),
	//	TemplateParam: tea.String(string(paramByte)),
	//	PhoneNumbers:  tea.String(mobile),
	//}
	//return send_phone.AliYunSendSms(sendSmsRequest)
	return nil
}
