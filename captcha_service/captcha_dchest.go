package captcha_service

import (
	"bytes"
	"encoding/base64"
	"github.com/dchest/captcha"
)

type CaptchaServiceDchest struct{}

func NewCaptchaServiceDchest() *CaptchaServiceDchest {
	return &CaptchaServiceDchest{}
}

func (c *CaptchaServiceDchest) GenerateCaptchaRaw(width, height int) (string, []byte, error) {
	captchaID := captcha.New()
	var image bytes.Buffer
	if err := captcha.WriteImage(&image, captchaID, width, height); err != nil {
		return "", nil, err
	}
	return captchaID, image.Bytes(), nil
}

func (c *CaptchaServiceDchest) GenerateCaptchaBase64(width, height int) (string, string, error) {
	captchaID, imageBytes, err := c.GenerateCaptchaRaw(width, height)
	if err != nil {
		return "", "", err
	}
	encodedImage := base64.StdEncoding.EncodeToString(imageBytes)
	return captchaID, encodedImage, nil
}

func (c *CaptchaServiceDchest) VerifyCaptcha(captchaID, solution string) bool {
	return captcha.VerifyString(captchaID, solution)
}
