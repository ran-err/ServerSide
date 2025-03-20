package captcha_service

type CaptchaService interface {
	GenerateCaptchaRaw(width, height int) (string, []byte, error)
	GenerateCaptchaBase64(width, height int) (string, string, error)
	VerifyCaptcha(captchaID, solution string) bool
}
