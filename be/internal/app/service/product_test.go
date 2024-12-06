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

var _ = Describe("ProductService", func() {

	var (
		ctx                   context.Context
		gormDB                *gorm.DB
		mockProductRepository *repository_mock.MockProductRepository

		testTarget service.ProductService
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
		testTarget = service.NewProductService(gormDB, mockProductRepository)
	})

	Describe("GetProductByID", func() {
		When("Product does not exist", func() {
			It("should return error", func() {
				productID := "1234"
				mockProductRepository.EXPECT().GetProductByID(gomock.Eq(ctx), gomock.Eq(gormDB), gomock.Eq(productID)).Return(persistence.Product{}, errors.New("product done not exit"))
				_, err := testTarget.GetProductByID(ctx, productID)
				Expect(err).To(HaveOccurred())
			})
		})
		When("Product does is exist", func() {
			It("should return product detail", func() {
				productID := "49030f6a-14a7-48db-9e00-39efd0f5e4f7"
				mockProductRepository.EXPECT().GetProductByID(
					gomock.Eq(ctx),
					gomock.Eq(gormDB),
					gomock.Eq(productID),
				).Return(persistence.Product{
					ProductID: uuid.MustParse(productID),
					Name:      "Product test ja",
					Image:     "http://image.png",
					Stock:     10,
					Price:     100,
				}, nil)

				res, _ := testTarget.GetProductByID(ctx, productID)

				Expect(res.ProductId).To(Equal(productID))
				Expect(res.Name).To(Equal("Product test ja"))
				Expect(res.ImageUrl).To(Equal("http://image.png"))
				Expect(res.Stock).To(Equal(int64(10)))
				Expect(res.Price).To(Equal(int64(100)))
			})
		})
	})
	Describe("GetProductList", func() {
		When("Product list does not exist", func() {
			It("should return error", func() {
				mockProductRepository.EXPECT().GetList(
					gomock.Eq(ctx),
					gomock.Eq(gormDB),
				).Return([]persistence.Product{}, errors.New("product list done not exist"))
				_, err := testTarget.GetProductList(ctx)
				Expect(err).To(HaveOccurred())
			})
		})
		When("Product list does is exist", func() {
			It("should return product list", func() {
				productID := "49030f6a-14a7-48db-9e00-39efd0f5e4f7"
				productID2 := "f95b325d-f0ed-4b2c-8467-78fb4dc0030c"
				mockProductRepository.EXPECT().GetList(
					gomock.Eq(ctx),
					gomock.Eq(gormDB),
				).Return([]persistence.Product{
					{
						ProductID: uuid.MustParse(productID),
						Name:      "Product test ja",
						Image:     "http://image.png",
						Stock:     10,
						Price:     100},
					{
						ProductID: uuid.MustParse(productID2),
						Name:      "Product test ja2",
						Image:     "http://image2.png",
						Stock:     0,
						Price:     1200,
					},
				}, nil)
				res, _ := testTarget.GetProductList(ctx)

				Expect(res).To(HaveLen(2))
				Expect(res[0].ProductId).To(Equal(productID))
				Expect(res[0].Name).To(Equal("Product test ja"))
				Expect(res[0].ImageUrl).To(Equal("http://image.png"))
				Expect(res[0].Stock).To(Equal(int64(10)))
				Expect(res[0].Price).To(Equal(int64(100)))
				Expect(res[1].ProductId).To(Equal(productID2))
				Expect(res[1].Name).To(Equal("Product test ja2"))
				Expect(res[1].ImageUrl).To(Equal("http://image2.png"))
				Expect(res[1].Stock).To(Equal(int64(0)))
				Expect(res[1].Price).To(Equal(int64(1200)))
			})
		})
	})
	Describe("CreateProduct", func() {
		productID := uuid.New()
		req := request_model.CreateProduct{
			Name:     "create product ja",
			ImageUrl: "",
			Stock:    0,
			Price:    0,
		}
		When("Creat product fail", func() {
			It("should return error", func() {
				mockProductRepository.EXPECT().CreateProduct(
					gomock.Eq(ctx),
					gomock.Eq(gormDB),
					request_model.CreateProduct{req.Name, req.ImageUrl, req.Stock, req.Price},
				).Return(persistence.Product{}, errors.New("create product fail"))

				_, err := testTarget.CreateProduct(ctx, req)
				Expect(err).To(HaveOccurred())
			})
		})
		When("Creat product success", func() {
			It("should return product detail", func() {

				mockProductRepository.EXPECT().CreateProduct(
					gomock.Eq(ctx),
					gomock.Eq(gormDB),
					request_model.CreateProduct{req.Name, req.ImageUrl, req.Stock, req.Price},
				).Return(persistence.Product{
					ProductID: uuid.MustParse(productID.String()),
					Name:      "create product ja",
					Image:     "",
					Stock:     0,
					Price:     0,
				}, nil)

				res, _ := testTarget.CreateProduct(ctx, req)
				Expect(res.ProductId).To(Equal(productID.String()))
				Expect(res.Name).To(Equal(req.Name))
				Expect(res.ImageUrl).To(Equal(req.ImageUrl))
				Expect(res.Stock).To(Equal(req.Stock))
				Expect(res.Price).To(Equal(req.Price))
			})
		})
	})
	Describe("Update Product Detail", func() {
		productID := "49030f6a-14a7-48db-9e00-39efd0f5e4f7"
		req := request_model.UpdateProductByID{
			ProductId: productID,
			Name:      "update product ja",
			ImageUrl:  "http://image.png",
			Stock:     0,
			Price:     0,
		}
		When("Update product fail", func() {
			It("should return error", func() {
				mockProductRepository.EXPECT().UpdateProductByID(
					gomock.Eq(ctx),
					gomock.Eq(gormDB),
					gomock.Eq(productID),
					request_model.UpdateProductByID{req.ProductId, req.Name, req.ImageUrl, req.Stock, req.Price},
				).Return(persistence.Product{}, errors.New("update product fail"))

				_, err := testTarget.UpdateProductByID(ctx, productID, req)
				Expect(err).To(HaveOccurred())
			})
		})
		When("Update product success", func() {
			It("should return product detail", func() {

				mockProductRepository.EXPECT().UpdateProductByID(
					gomock.Eq(ctx),
					gomock.Eq(gormDB),
					gomock.Eq(productID),
					request_model.UpdateProductByID{req.ProductId, req.Name, req.ImageUrl, req.Stock, req.Price},
				).Return(persistence.Product{
					ProductID: uuid.MustParse(productID),
					Name:      req.Name,
					Image:     req.ImageUrl,
					Stock:     req.Stock,
					Price:     req.Price,
				}, nil)

				res, _ := testTarget.UpdateProductByID(ctx, productID, req)
				Expect(res.ProductId).To(Equal(req.ProductId))
				Expect(res.Name).To(Equal(req.Name))
				Expect(res.ImageUrl).To(Equal(req.ImageUrl))
				Expect(res.Stock).To(Equal(req.Stock))
				Expect(res.Price).To(Equal(req.Price))

			})
		})
	})
	Describe("Delete Product", func() {
		productID := "49030f6a-14a7-48db-9e00-39efd0f5e4f7"
		When("Delete product fail", func() {
			It("should return error", func() {
				mockProductRepository.EXPECT().DeleteProductByID(
					gomock.Eq(ctx),
					gomock.Eq(gormDB),
					gomock.Eq(productID),
				).Return(errors.New("delete product fail"))

				err := testTarget.DeleteProductByID(ctx, productID)
				Expect(err).To(HaveOccurred())
			})
		})
		When("Delete product success", func() {
			It("should return error is nil", func() {

				mockProductRepository.EXPECT().DeleteProductByID(
					gomock.Eq(ctx),
					gomock.Eq(gormDB),
					gomock.Eq(productID),
				).Return(nil)

				err := testTarget.DeleteProductByID(ctx, productID)
				Expect(err).To(BeNil())

			})
		})
	})

})
