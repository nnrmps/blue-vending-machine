package service_test

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/persistence"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/service"
	repository_mock "github.com/nnrmps/blue-vending-machine/be/internal/mock/repository"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ = Describe("ReservedMoneyService", func() {

	var (
		ctx                         context.Context
		gormDB                      *gorm.DB
		mockReservedMoneyRepository *repository_mock.MockReservedMoneyRepository

		testTarget service.ReserveMoneyService
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

		mockReservedMoneyRepository = repository_mock.NewMockReservedMoneyRepository(ctrl)
		testTarget = service.NewReservedMoneyService(gormDB, mockReservedMoneyRepository)

	})

	Describe("Get Reserved Money", func() {

		reservedMoneyData := map[string]int64{
			"Coins1":   1,
			"Coins5":   1,
			"Coins10":  1,
			"Bank20":   1,
			"Bank50":   1,
			"Bank100":  1,
			"Bank500":  1,
			"Bank1000": 1,
		}

		When("Reserved money done not exit", func() {
			It("should return error", func() {
				mockReservedMoneyRepository.EXPECT().GetReservedMoney(gomock.Eq(ctx), gomock.Eq(gormDB)).Return(map[int64]int64{}, errors.New("reserved money done not exit"))
				_, err := testTarget.GetReservedMoney(ctx)

				Expect(err).To(HaveOccurred())
			})
		})

		When("Reserved money done is exit", func() {
			It("should return reserved money data", func() {
				mockReservedMoneyRepository.EXPECT().GetReservedMoney(gomock.Eq(ctx), gomock.Eq(gormDB)).Return(map[int64]int64{
					1:    reservedMoneyData["Coins1"],
					5:    reservedMoneyData["Coins5"],
					10:   reservedMoneyData["Coins10"],
					20:   reservedMoneyData["Bank20"],
					50:   reservedMoneyData["Bank50"],
					100:  reservedMoneyData["Bank100"],
					500:  reservedMoneyData["Bank500"],
					1000: reservedMoneyData["Bank1000"],
				}, nil)

				res, _ := testTarget.GetReservedMoney(ctx)

				Expect(res["coins1"]).To(Equal(reservedMoneyData["Coins1"]))
				Expect(res["coins5"]).To(Equal(reservedMoneyData["Coins5"]))
				Expect(res["coins10"]).To(Equal(reservedMoneyData["Coins10"]))
				Expect(res["bank20"]).To(Equal(reservedMoneyData["Bank20"]))
				Expect(res["bank50"]).To(Equal(reservedMoneyData["Bank50"]))
				Expect(res["bank100"]).To(Equal(reservedMoneyData["Bank100"]))
				Expect(res["bank500"]).To(Equal(reservedMoneyData["Bank500"]))
				Expect(res["bank1000"]).To(Equal(reservedMoneyData["Bank1000"]))
			})
		})
	})
	Describe("Update Reserved Money", func() {

		reservedMoneyData := persistence.ReservedMoney{
			Coins1:   100,
			Coins5:   100,
			Coins10:  100,
			Bank20:   100,
			Bank50:   100,
			Bank100:  100,
			Bank500:  100,
			Bank1000: 100,
		}

		When("update Reserved money fail", func() {
			It("should return error", func() {
				mockReservedMoneyRepository.EXPECT().UpdateReservedMoney(gomock.Eq(ctx), gomock.Eq(gormDB), reservedMoneyData).Return(errors.New("update reserved money fail"))
				err := testTarget.UpdateReservedMoney(ctx, reservedMoneyData)

				Expect(err).To(HaveOccurred())
			})
		})

		When("Update Reserved money success", func() {
			It("should return err is nil", func() {
				mockReservedMoneyRepository.EXPECT().UpdateReservedMoney(gomock.Eq(ctx), gomock.Eq(gormDB), reservedMoneyData).Return(nil)

				err := testTarget.UpdateReservedMoney(ctx, reservedMoneyData)

				Expect(err).To(BeNil())

			})
		})
	})

})
