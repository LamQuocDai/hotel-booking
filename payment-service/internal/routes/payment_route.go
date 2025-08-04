package routes

import (
	"payment-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupPaymentRoute(r *gin.Engine, h *handlers.PaymentHandler) {
	paymentGroup := r.Group("/payments")
	{
		paymentGroup.GET("", h.GetAllPayments)
		paymentGroup.GET("/:id", h.GetPaymentByID)
		paymentGroup.POST("", h.CreatePayment)
		paymentGroup.PUT("/:id", h.UpdatedPayment)
		paymentGroup.DELETE("/:id", h.DeletePayment)
	}
}
