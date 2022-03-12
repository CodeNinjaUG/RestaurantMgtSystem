package routes


import (
	"github.com/gin-gonic/gin"
	controllers "restaurant-management-system/controllers"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/invoices",controllers.GetInvoices())
	incomingRoutes.GET("/invoices/:id",controllers.GetInvoice())
	incomingRoutes.POST("/invoice/post",controllers.StoreInvoice())
	incomingRoutes.DELETE("/delete/invoice/:id",controllers.DeleteInvoice())
	incomingRoutes.PATCH("/update/invoice/:id",controllers.UpdateInvoice())
}
