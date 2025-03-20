package captcha_service_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCaptchaService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CaptchaService Suite")
}
