package captcha_service_test

import (
	"captcha_service"
	"encoding/base64"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// This is a group of characterization tests.
var _ = Describe("CaptchaDchest", func() {
	var service captcha_service.CaptchaService
	const width, height int = 100, 50

	BeforeEach(func() {
		service = captcha_service.NewCaptchaServiceDchest()
	})

	Describe("GenerateCaptchaRaw", func() {
		It("should generate a valid captcha ID and image bytes", func() {
			captchaID, imageBytes, err := service.GenerateCaptchaRaw(width, height)
			Expect(err).ToNot(HaveOccurred())
			Expect(captchaID).ToNot(BeEmpty())
			Expect(imageBytes).ToNot(BeEmpty())
		})
	})

	Describe("GenerateCaptchaBase64", func() {
		It("should generate a valid base64-encoded image", func() {
			captchaID, encodedImage, err := service.GenerateCaptchaBase64(width, height)
			Expect(err).ToNot(HaveOccurred())
			Expect(captchaID).ToNot(BeEmpty())
			Expect(encodedImage).ToNot(BeEmpty())

			decodedBytes, err := base64.StdEncoding.DecodeString(encodedImage)
			Expect(err).ToNot(HaveOccurred())
			Expect(decodedBytes).ToNot(BeEmpty())
		})
	})
})
