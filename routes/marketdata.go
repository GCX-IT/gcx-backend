package routes

import (
	"gcx-cms/internal/marketdata/handlers"
	"gcx-cms/internal/shared/middleware"

	"github.com/gin-gonic/gin"
)

// SetupMarketDataRoutes configures all market data related routes
func SetupMarketDataRoutes(r *gin.Engine) {
	// Market Data API group
	marketData := r.Group("/api/marketdata")

	// Public routes (no authentication required)
	{
		// Get current market prices
		marketData.GET("/prices", handlers.GetCurrentPrices)

		// Get commodity prices
		marketData.GET("/prices/:commodity", handlers.GetCommodityPrices)

		// Get historical prices
		marketData.GET("/history", handlers.GetHistoricalPrices)

		// Get price summary
		marketData.GET("/summary/:commodity", handlers.GetPriceSummary)

		// Get subscription plans (public pricing)
		marketData.GET("/plans", handlers.GetSubscriptionPlans)
	}

	// Protected routes (authentication required)
	protected := marketData.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		// User subscription management
		protected.GET("/subscription", handlers.GetUserSubscription)
		protected.POST("/subscription", handlers.CreateSubscription)
		protected.PUT("/subscription/:id", handlers.UpdateSubscription)
		protected.DELETE("/subscription/:id", handlers.CancelSubscription)

		// Advanced market data (requires subscription)
		protected.GET("/analytics", handlers.GetMarketAnalytics)
		protected.GET("/alerts", handlers.GetPriceAlerts)
		protected.POST("/alerts", handlers.CreatePriceAlert)
		protected.PUT("/alerts/:id", handlers.UpdatePriceAlert)
		protected.DELETE("/alerts/:id", handlers.DeletePriceAlert)

		// Real-time data (requires premium subscription)
		protected.GET("/realtime", handlers.GetRealTimeData)
		protected.GET("/stream", handlers.GetDataStream)
	}
}

// SetupMarketDataAdminRoutes configures admin-only market data routes
func SetupMarketDataAdminRoutes(r *gin.Engine) {
	admin := r.Group("/api/admin/marketdata")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.AdminMiddleware())
	{
		// Admin can manage all market data
		admin.POST("/prices", handlers.AdminCreatePrice)
		admin.PUT("/prices/:id", handlers.AdminUpdatePrice)
		admin.DELETE("/prices/:id", handlers.AdminDeletePrice)

		// Admin can manage subscription plans
		admin.POST("/plans", handlers.AdminCreatePlan)
		admin.PUT("/plans/:id", handlers.AdminUpdatePlan)
		admin.DELETE("/plans/:id", handlers.AdminDeletePlan)

		// Admin can manage commodities
		admin.POST("/commodities", handlers.AdminCreateCommodity)
		admin.PUT("/commodities/:id", handlers.AdminUpdateCommodity)
		admin.DELETE("/commodities/:id", handlers.AdminDeleteCommodity)
	}
}
