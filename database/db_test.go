package database_test

import (
	"goginmvc/database"
	"goginmvc/globalconst"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DB", func() {
	Context("LoadDBConfig", func() {
		It("should work", func() {
			cfg, err := database.LoadDBConfig(globalconst.DEV)
			Expect(err).NotTo(HaveOccurred())
			Expect(cfg).ShouldNot(BeNil())
		})
	})
})
