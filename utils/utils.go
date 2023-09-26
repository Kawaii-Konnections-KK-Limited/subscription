package utils

import (
	"encoding/base64"
)

func SubscriptionBuilder(l *[]string) string {
	var subs string
	for _, v := range *l {
		subs += v + "\n"
	}
	// transform to base64
	uEnc := base64.StdEncoding.EncodeToString([]byte(subs))
	return uEnc

}

var GetUserProfiles = getUserProfiles
var VerifyToken = verifyToken
var SubUpdateSender = subscriptionUpdateSender

func getUserProfiles(token *string) *[]string {

	subs := []string{"vless://60d08a55-df01-4729-cb72-025a7bf9177b@m2rel.siasepid.sbs:80?mode=multi&security=reality&encryption=none&pbk=HgrpXJzQo2liQMY9YAPq1_PuiDXNNBLx8hRyVVfUZko&fp=firefox&spx=%2F&type=grpc&serviceName==tgju.org&sid=af41f983#🇩🇪 آلمان : @KralVPN", "vmess://eydhZGQnOiAnc2VydmVyMzMuaXJhbi1zZXJ2aWNlLm9ubGluZScsICdhaWQnOiAnMCcsICdob3N0JzogJ2dvb2dsZS5jb20nLCAnaWQnOiAnNDMxMmU0MDEtNzhhOS00MmViLWE2ZTUtNmE1NjI4NjZhM2MxJywgJ25ldCc6ICd3cycsICdwYXRoJzogJy8nLCAncG9ydCc6ICc4MDgwJywgJ3BzJzogJ/Cfh7Pwn4exINmH2YTZhtivIDogQEtyYWxWUE4nLCAnc2N5JzogJ2F1dG8nLCAnc25pJzogJycsICd0bHMnOiAnJywgJ3R5cGUnOiAnJywgJ3YnOiAnMid9", "vless://37a72384-f5cc-48bb-b577-fc50551491b8@wlf.ecodryfruit.com:443?security=reality&encryption=none&alpn=h2&pbk=GOWK9HVvqiOnATiX3e2sgT_GU5WBkbIglk7MAFWiLxk&headerType=none&fp=chrome&type=tcp&flow=xtls-rprx-vision&sni=www.whitepages.com&sid=ee739b59#🇳🇱 هلند : @KralVPN"}
	return &subs
}
func verifyToken(token *string) bool {
	return true
}
func subscriptionUpdateSender(token *string) {

}
