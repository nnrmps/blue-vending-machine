package service_test

import (
	"github.com/nnrmps/blue-vending-machine/be/internal/app/service"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Health Service", func() {
	var testTarget service.HealthService

	BeforeEach(func() {
		testTarget = service.NewHealthService()
	})

	When("Health Check", func() {
		It("should return error is nil", func() {

			err := testTarget.HealthCheck()
			Expect(err).To(BeNil())

		})
	})
})
