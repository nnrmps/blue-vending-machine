package service_test

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/persistence"
	"github.com/nnrmps/blue-vending-machine/be/internal/app/service"
	repository_mock "github.com/nnrmps/blue-vending-machine/be/internal/mock/repository"
	"github.com/nnrmps/blue-vending-machine/be/pkg/request_model"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"go.uber.org/mock/gomock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ = Describe("CheckoutService", func() {

	var (
		ctx                         context.Context
		gormDB                      *gorm.DB
		mockProductRepository       *repository_mock.MockProductRepository
		mockReservedMoneyRepository *repository_mock.MockReservedMoneyRepository
		productID                   string
		productDetail               persistence.Product
		total                       request_model.Money

		testTarget service.CheckoutService
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

		mockProductRepository = repository_mock.NewMockProductRepository(ctrl)
		mockReservedMoneyRepository = repository_mock.NewMockReservedMoneyRepository(ctrl)
		testTarget = service.NewCheckoutService(gormDB, mockReservedMoneyRepository, mockProductRepository)

		productID = "49030f6a-14a7-48db-9e00-39efd0f5e4f7"
		productDetail = persistence.Product{
			ProductID: uuid.MustParse(productID),
			Name:      "Product test ja",
			Image:     "http://image.png",
			Stock:     10,
			Price:     100,
		}

		total = request_model.Money{
			Coins1:   1,
			Coins5:   1,
			Coins10:  1,
			Bank20:   1,
			Bank50:   1,
			Bank100:  2,
			Bank500:  1,
			Bank1000: 1,
		}
	})

	Describe("Checkout Product", func() {
		When("Product does not exist", func() {
			It("should return error", func() {
				mockProductRepository.EXPECT().GetProductByID(gomock.Eq(ctx), gomock.Eq(gormDB), gomock.Eq(productID)).Return(persistence.Product{}, errors.New("product done not exit"))
				_, err := testTarget.CheckoutProduct(ctx, productID, total)

				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("product done not exit"))
			})
		})
		When("Product stock equal 0", func() {

			It("should return error", func() {
				productDetail.Stock = 0
				mockProductRepository.EXPECT().GetProductByID(gomock.Eq(ctx), gomock.Eq(gormDB), gomock.Eq(productID)).Return(persistence.Product{}, nil)
				_, err := testTarget.CheckoutProduct(ctx, productID, total)

				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("Sorry, product out of stock"))
			})
		})

		When("Total Deposit less than product price", func() {
			It("should return error", func() {

				total = request_model.Money{
					Coins1:   1,
					Coins5:   0,
					Coins10:  0,
					Bank20:   0,
					Bank50:   0,
					Bank100:  0,
					Bank500:  0,
					Bank1000: 0,
				}
				mockProductRepository.EXPECT().GetProductByID(gomock.Eq(ctx), gomock.Eq(gormDB), gomock.Eq(productID)).Return(persistence.Product{
					ProductID: productDetail.ProductID,
					Name:      productDetail.Name,
					Image:     productDetail.Image,
					Stock:     productDetail.Stock,
					Price:     productDetail.Price,
				}, nil)
				_, err := testTarget.CheckoutProduct(ctx, productID, total)

				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("Sorry, not enough deposit"))
			})
		})

		When("Get reserved money is error", func() {
			It("should return error", func() {
				mockProductRepository.EXPECT().GetProductByID(gomock.Eq(ctx), gomock.Eq(gormDB), gomock.Eq(productID)).Return(persistence.Product{
					productDetail.ProductID,
					productDetail.Name,
					productDetail.Image,
					productDetail.Stock,
					productDetail.Price,
				}, nil)
				mockReservedMoneyRepository.EXPECT().GetReservedMoney(gomock.Eq(ctx), gomock.Eq(gormDB)).Return(map[int64]int64{}, errors.New("err"))
				_, err := testTarget.CheckoutProduct(ctx, productID, total)
				Expect(err).To(HaveOccurred())
			})
		})
		When("Don't have enough change", func() {
			It("should return error", func() {
				productDetail = persistence.Product{
					ProductID: uuid.MustParse(productID),
					Name:      "Product test ja",
					Image:     "http://image.png",
					Stock:     10,
					Price:     109,
				}

				total = request_model.Money{
					Coins1:   0,
					Coins5:   0,
					Coins10:  1,
					Bank20:   0,
					Bank50:   0,
					Bank100:  1,
					Bank500:  0,
					Bank1000: 0,
				}
				mockProductRepository.EXPECT().GetProductByID(gomock.Eq(ctx), gomock.Eq(gormDB), gomock.Eq(productID)).Return(persistence.Product{
					productDetail.ProductID,
					productDetail.Name,
					productDetail.Image,
					productDetail.Stock,
					productDetail.Price,
				}, nil)
				mockReservedMoneyRepository.EXPECT().GetReservedMoney(gomock.Eq(ctx), gomock.Eq(gormDB)).Return(map[int64]int64{
					1:    0,
					5:    0,
					10:   0,
					20:   0,
					50:   0,
					100:  0,
					500:  0,
					1000: 0,
				}, nil)

				_, err := testTarget.CheckoutProduct(ctx, productID, total)
				Expect(err).ToNot(BeNil())
				Expect(err.Error()).To(Equal("Sorry, don't have enough change."))
			})
		})
		When("Update Reserved Money fail", func() {
			It("should return error", func() {
				mockProductRepository.EXPECT().GetProductByID(gomock.Eq(ctx), gomock.Eq(gormDB), gomock.Eq(productID)).Return(persistence.Product{
					productDetail.ProductID,
					productDetail.Name,
					productDetail.Image,
					productDetail.Stock,
					productDetail.Price,
				}, nil)
				mockReservedMoneyRepository.EXPECT().GetReservedMoney(gomock.Eq(ctx), gomock.Eq(gormDB)).Return(map[int64]int64{}, nil)
				mockReservedMoneyRepository.EXPECT().UpdateReservedMoney(gomock.Eq(ctx), gomock.Any(), gomock.Any()).Return(errors.New("update err"))
				_, err := testTarget.CheckoutProduct(ctx, productID, total)

				Expect(err).To(HaveOccurred())

			})
		})
		When("Deduct Product Stock fail", func() {
			It("should return error", func() {
				mockProductRepository.EXPECT().GetProductByID(gomock.Eq(ctx), gomock.Eq(gormDB), gomock.Eq(productID)).Return(persistence.Product{
					productDetail.ProductID,
					productDetail.Name,
					productDetail.Image,
					productDetail.Stock,
					productDetail.Price,
				}, nil)
				mockReservedMoneyRepository.EXPECT().GetReservedMoney(gomock.Eq(ctx), gomock.Eq(gormDB)).Return(map[int64]int64{}, nil)
				mockReservedMoneyRepository.EXPECT().UpdateReservedMoney(gomock.Eq(ctx), gomock.Any(), gomock.Any()).Return(nil)
				mockProductRepository.EXPECT().DeductStockByProductID(gomock.Eq(ctx), gomock.Any(), productID, productDetail.Stock-1).Return(errors.New("deduct err"))
				_, err := testTarget.CheckoutProduct(ctx, productID, total)

				Expect(err).To(HaveOccurred())

			})
		})

		When("Checkout success", func() {
			It("should return total change", func() {
				mockProductRepository.EXPECT().GetProductByID(gomock.Eq(ctx), gomock.Eq(gormDB), gomock.Eq(productID)).Return(persistence.Product{
					productDetail.ProductID,
					productDetail.Name,
					productDetail.Image,
					productDetail.Stock,
					productDetail.Price,
				}, nil)
				mockReservedMoneyRepository.EXPECT().GetReservedMoney(gomock.Eq(ctx), gomock.Eq(gormDB)).Return(map[int64]int64{}, nil)
				mockReservedMoneyRepository.EXPECT().UpdateReservedMoney(gomock.Eq(ctx), gomock.Any(), gomock.Any()).Return(nil)
				mockProductRepository.EXPECT().DeductStockByProductID(gomock.Eq(ctx), gomock.Any(), productID, productDetail.Stock-1).Return(nil)
				res, _ := testTarget.CheckoutProduct(ctx, productID, total)
				Expect(res.TotalChange).To(Equal(int64(1686)))
				Expect(res.Coins1).To(Equal(total.Coins1))
				Expect(res.Coins5).To(Equal(total.Coins5))
				Expect(res.Coins10).To(Equal(total.Coins10))
				Expect(res.Bank20).To(Equal(total.Bank20))
				Expect(res.Bank50).To(Equal(total.Bank50))
				Expect(res.Bank100).To(Equal(total.Bank100 - 1))
				Expect(res.Bank1000).To(Equal(total.Bank1000))
				Expect(res.Bank500).To(Equal(total.Bank500))
				Expect(res.Bank1000).To(Equal(total.Bank1000))
			})
		})
	})

})
