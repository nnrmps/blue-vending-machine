package service_test

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/persistence"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/service"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/setting"
	repository_mock "github.com/nnrmps/blue-vending-machine/be/internal/mock/repository"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ = Describe("UserService", func() {

	var (
		ctx                context.Context
		gormDB             *gorm.DB
		mockUserRepository *repository_mock.MockUserRepository

		testTarget service.UserService
	)

	BeforeEach(func() {
		ctx = context.Background()
		db, _, _ := sqlmock.New()

		gormDB, _ = gorm.Open(postgres.New(postgres.Config{
			DSN:        "sql_mock_db_0",
			DriverName: "postgres",
			Conn:       db,
		}), &gorm.Config{})

		ctrl := gomock.NewController(GinkgoT(), gomock.WithOverridableExpectations())

		mockUserRepository = repository_mock.NewMockUserRepository(ctrl)
		testTarget = service.NewUserService(gormDB, mockUserRepository)

		setting.AppConfig = setting.Configuration{
			SecretKey: "t_t",
		}
	})

	Describe("Checkout Product", func() {
		username := "admin"
		password := "12345abcde"
		When("User Login done not exits in DB", func() {
			It("should return error", func() {
				mockUserRepository.EXPECT().Login(gomock.Eq(ctx), gomock.Eq(gormDB), username, password).Return(persistence.User{}, errors.New("user done not exit"))
				_, err := testTarget.Login(ctx, username, password)

				Expect(err).To(HaveOccurred())
			})
		})

		When("User Login Success", func() {
			It("should return token", func() {
				mockUserRepository.EXPECT().Login(gomock.Eq(ctx), gomock.Eq(gormDB), username, password).Return(persistence.User{}, nil)
				token, err := testTarget.Login(ctx, username, password)

				Expect(err).To(BeNil())
				Expect(token).ToNot(BeEmpty())
			})
		})

	})

})
