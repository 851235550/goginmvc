package database_test

import (
	"goginmvc/database"
	"goginmvc/globalconst"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DB", func() {
	Context("LoadDBConfig", func() {
		It("should work at DEV mode", func() {
			cfg, err := database.LoadDBConfig(globalconst.DEV)
			Expect(err).NotTo(HaveOccurred())
			Expect(cfg).ShouldNot(BeNil())
		})
		It("should work at PROD mode", func() {
			os.Setenv("DB_USERNAME", "dbusername")
			os.Setenv("DB_PASSWORD", "dbpassword")
			os.Setenv("DB_HOST", "dbhost")
			os.Setenv("DB_PORT", "dbport")
			os.Setenv("DB_NAME", "dbname")

			cfg, err := database.LoadDBConfig(globalconst.PROD)
			Expect(err).NotTo(HaveOccurred())
			Expect(cfg).ShouldNot(BeNil())

			os.Unsetenv("DB_USERNAME")
			os.Unsetenv("DB_PASSWORD")
			os.Unsetenv("DB_HOST")
			os.Unsetenv("DB_PORT")
			os.Unsetenv("DB_NAME")
		})
		It("should error when env value is not set", func() {
			os.Setenv("DB_USERNAME", "")
			os.Setenv("DB_PASSWORD", "")
			os.Setenv("DB_HOST", "")

			cfg, err := database.LoadDBConfig(globalconst.PROD)
			Expect(err).Should(HaveOccurred())
			Expect(cfg).Should(BeNil())

			os.Unsetenv("DB_USERNAME")
			os.Unsetenv("DB_PASSWORD")
			os.Unsetenv("DB_HOST")
		})
	})
})
