package interfaces

import (
	"context"

	"github.com/msegeya56/eCommerce-smart_gads/pkg/domain"
	"github.com/msegeya56/eCommerce-smart_gads/pkg/utils/response"
)

type PaymentService interface {
	GetAllPaymentOptions(ctx context.Context) (PaymentOptions []response.PaymentOptionResp, err error)

	SavePaymentDetails(ctx context.Context, paymentData domain.PaymentDetails) error
	UpdatePaymentStatus(ctx context.Context, statusId, orderId uint) error
	GetPaymentDataByOrderId(ctx context.Context, orderId uint) (paymentData domain.PaymentDetails, err error)
}
