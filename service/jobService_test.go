package service_test

import (
	"goginmvc/database"
	"goginmvc/service"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("JobService", func() {
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

			jobSvc := service.NewJobService()
			jobs, err := jobSvc.GetJobs()
			Expect(err).NotTo(HaveOccurred())
			Expect(jobs).To(Equal([]*service.Job{{ID: 1}}))
		})
	})
})
