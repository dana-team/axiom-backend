package e2e_tests

import (
	"github.com/joho/godotenv"
	"log"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "e2e test Suite")
}

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("No .env file found")
	}
}
