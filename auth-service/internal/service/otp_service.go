package service

import (
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type OtpService interface {
	GenerateOtp(string) (string, error)
	VerifyOtp(string, string) bool
}

type otpService struct{}

func NewOtpService() OtpService {
	return &otpService{}
}

func GenerateSecret(account string) (*otp.Key, error) {
	return totp.Generate(totp.GenerateOpts{Issuer: "orto", AccountName: account})
}

func (o *otpService) GenerateOtp(secret string) (string, error) {
	return totp.GenerateCode(secret, time.Now())
}

func (o *otpService) VerifyOtp(otp string, secret string) bool {
	return totp.Validate(otp, secret)
}
