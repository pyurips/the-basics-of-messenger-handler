package configs

import "os"

func GetAPIEndpoint() string {
	emulator := os.Getenv("EMULATOR")
	accessToken := os.Getenv("ACCESS_TOKEN")
	if emulator == "true" {
		return "http://localhost:8081/?access_token=" + accessToken
	}
	return "https://graph.facebook.com/v2.6/me/messages?access_token=" + accessToken
}
