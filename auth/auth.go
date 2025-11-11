package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

type Service struct {
	appKey    string
	appSecret string
}

func NewService(appKey, appSecret string) *Service {
	return &Service{
		appKey:    appKey,
		appSecret: appSecret,
	}
}

func (s *Service) GenerateSignature(socketID, channel string, channelData *string) string {
	var stringToSign string
	if channelData != nil {
		stringToSign = fmt.Sprintf("%s:%s:%s", socketID, channel, *channelData)
	} else {
		stringToSign = fmt.Sprintf("%s:%s", socketID, channel)
	}

	h := hmac.New(sha256.New, []byte(s.appSecret))
	h.Write([]byte(stringToSign))
	signature := hex.EncodeToString(h.Sum(nil))

	return signature
}

// return format: app_key:signature
func (s *Service) GenerateAuthString(socketID, channel string, channelData *string) string {
	signature := s.GenerateSignature(socketID, channel, channelData)
	return fmt.Sprintf("%s:%s", s.appKey, signature)
}

func (s *Service) ValidateAuth(authString, socketID, channel string, channelData *string) bool {
	expectedAuth := s.GenerateAuthString(socketID, channel, channelData)
	return authString == expectedAuth
}

func ParseAuthString(authString string) (appKey, signature string, err error) {
	parts := strings.SplitN(authString, ":", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("invalid auth string format")
	}
	return parts[0], parts[1], nil
}
