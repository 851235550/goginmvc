package database_test

import (
	"goginmvc/database"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("JobManager", func() {
	var mockDB sqlmock.Sqlmock
	BeforeEach(func() {
		db, mock, err := sqlmock.New()
		Expect(err).NotTo(HaveOccurred())
		Expect(db).ShouldNot(BeNil())
		Expect(mock).ShouldNot(BeNil())

		mockDB = mock
		database.MockDB(db, "postgres")
	})

	Context("GetJobs", func() {
		It("should work", func() {
			mockDB.ExpectQuery(`SELECT \* FROM jobs`).WillReturnRows(
				sqlmock.NewRows(
					[]string{"id"}).AddRow("1"))

			jobManager := database.NewJobManager()
			jobs, err := jobManager.GetJobs()
			Expect(err).NotTo(HaveOccurred())
			Expect(jobs).To(Equal([]*database.Job{{ID: 1}}))
		})
	})
})
