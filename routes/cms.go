package routes

import (
	"gcx-cms/internal/cms/handlers"
	"gcx-cms/internal/shared/middleware"

	"github.com/gin-gonic/gin"
)

// SetupCMSRoutes configures all CMS related routes
func SetupCMSRoutes(r *gin.Engine) {
	// CMS API group
	cms := r.Group("/api")

	// Add database middleware to all CMS routes
	cms.Use(middleware.DatabaseMiddleware())

	// Public CMS routes (no authentication required)
	{
		// Public blog posts (for website)
		cms.GET("/posts", handlers.GetPublicPosts)
		cms.GET("/posts/:slug", handlers.GetPublicPost)

		// Public pages (for website)
		cms.GET("/pages/:slug", handlers.GetPageBySlug)

		// Public settings (for website)
		cms.GET("/settings/public", handlers.GetPublicSettings)

		// Public settings by group (for website)
		cms.GET("/settings/hero", handlers.GetSettingsByGroupPublic)
		cms.GET("/settings/services", handlers.GetSettingsByGroupPublic)
		cms.GET("/settings/why_join", handlers.GetSettingsByGroupPublic)
		cms.GET("/settings/cta", handlers.GetSettingsByGroupPublic)
		cms.GET("/settings/market_data", handlers.GetSettingsByGroupPublic)

		// Public menus (for website)
		cms.GET("/menus/location/:location", handlers.GetMenuByLocation)

		// Public team members (for website)
		cms.GET("/team-members", handlers.GetTeamMembers)    // GET /api/team-members (list all team members)
		cms.GET("/team-members/:id", handlers.GetTeamMember) // GET /api/team-members/{id}

		// Public traders and brokers (for website)
		cms.GET("/traders", handlers.GetTraders)    // GET /api/traders (list all traders)
		cms.GET("/traders/:id", handlers.GetTrader) // GET /api/traders/{id}
		cms.GET("/brokers", handlers.GetBrokers)    // GET /api/brokers (list all brokers)
		cms.GET("/brokers/:id", handlers.GetBroker) // GET /api/brokers/{id}

		// Public partners (for website)
		cms.GET("/partners", handlers.GetActivePartners)                        // GET /api/partners (list all active partners)
		cms.GET("/partners/:id", handlers.GetPartner)                           // GET /api/partners/{id}
		cms.GET("/partners/category/:category", handlers.GetPartnersByCategory) // GET /api/partners/category/{category}

		// Public publications (for website)
		cms.GET("/publications", handlers.GetPublications)    // GET /api/publications (list all publications)
		cms.GET("/publications/:id", handlers.GetPublication) // GET /api/publications/{id}

		// Public careers (for website)
		cms.GET("/careers", handlers.GetCareers)    // GET /api/careers (list all careers)
		cms.GET("/careers/:id", handlers.GetCareer) // GET /api/careers/{id}

		// Public commodities (for website)
		cms.GET("/commodities", handlers.GetCommodities)   // GET /api/commodities (list all commodities)
		cms.GET("/commodities/:id", handlers.GetCommodity) // GET /api/commodities/{id}

		// Contract file presigned URL (for website)
		cms.GET("/commodities/:commodityId/contract-url", handlers.GetContractFilePresignedURL) // GET /api/commodities/{commodityId}/contract-url

		// Public commodities with contract types (for website)
		cms.GET("/commodities-with-contract-types", handlers.GetCommoditiesWithContractTypes) // GET /api/commodities-with-contract-types

		// Public contract types (for website)
		cms.GET("/contract-types/commodity/:commodityId", handlers.GetAllCommodityContractTypes) // GET /api/contract-types/commodity/{commodityId}
		cms.GET("/contract-types/:id", handlers.GetCommodityContractType)                        // GET /api/contract-types/{id}

		// Public events (for website)
		cms.GET("/events", handlers.GetEvents)                      // GET /api/events (list all events with filters)
		cms.GET("/events/upcoming", handlers.GetUpcomingEvents)     // GET /api/events/upcoming
		cms.GET("/events/past", handlers.GetPastEvents)             // GET /api/events/past
		cms.GET("/events/:id", handlers.GetEvent)                   // GET /api/events/{id}
		cms.GET("/events/slug/:slug", handlers.GetEventBySlug)      // GET /api/events/slug/{slug}
		cms.POST("/events/:id/register", handlers.RegisterForEvent) // POST /api/events/{id}/register

		// Public RTI requests (for website)
		cms.POST("/rti/request", handlers.CreateRTIRequest) // POST /api/rti/request (submit RTI request)

		// Public RTI documents (for website)
		cms.GET("/rti/documents", handlers.GetRTIDocuments)                   // GET /api/rti/documents (list all RTI documents)
		cms.GET("/rti/documents/:id", handlers.GetRTIDocument)                // GET /api/rti/documents/{id}
		cms.POST("/rti/documents/:id/download", handlers.DownloadRTIDocument) // POST /api/rti/documents/{id}/download (track downloads)

		// Public photo galleries (for website)
		cms.GET("/galleries", handlers.GetGalleries)                // GET /api/galleries (list all galleries)
		cms.GET("/galleries/:id", handlers.GetGallery)              // GET /api/galleries/{id} (get gallery with photos)
		cms.GET("/galleries/slug/:slug", handlers.GetGalleryBySlug) // GET /api/galleries/slug/{slug}
		cms.GET("/galleries/:id/photos", handlers.GetGalleryPhotos) // GET /api/galleries/{id}/photos

		// Public video libraries (for website)
		cms.GET("/video-libraries", handlers.GetVideoLibraries)                // GET /api/video-libraries (list all libraries)
		cms.GET("/video-libraries/:id", handlers.GetVideoLibrary)              // GET /api/video-libraries/{id} (get library with videos)
		cms.GET("/video-libraries/slug/:slug", handlers.GetVideoLibraryBySlug) // GET /api/video-libraries/slug/{slug}
		cms.GET("/video-libraries/:id/videos", handlers.GetLibraryVideos)      // GET /api/video-libraries/{id}/videos
		cms.POST("/videos/:id/view", handlers.TrackVideoView)                  // POST /api/videos/{id}/view (track video views)

		// Public news ticker (for website)
		cms.GET("/news", handlers.GetNewsItems)                 // GET /api/news (list active news items)
		cms.GET("/news/breaking", handlers.GetBreakingNews)     // GET /api/news/breaking (list breaking news only)
		cms.GET("/news/:id", handlers.GetNewsItem)              // GET /api/news/{id} (get single news item)
		cms.GET("/news-categories", handlers.GetNewsCategories) // GET /api/news-categories (list news categories)
	}

	// Protected CMS routes (authentication required)
	protected := cms.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		// Dashboard routes
		protected.GET("/cms/dashboard/stats", handlers.GetDashboardStats)
		protected.GET("/cms/dashboard/activity", handlers.GetDashboardActivity)

		// User profile management
		protected.GET("/user/profile", handlers.GetProfile)
		protected.PUT("/user/profile", handlers.UpdateProfile)
		protected.POST("/user/change-password", handlers.ChangePassword)

		// CMS Blog management (requires blogger or admin role)
		cmsProtected := protected.Group("/cms")
		cmsProtected.Use(middleware.BloggerMiddleware())
		{
			cmsProtected.GET("/posts", handlers.GetAllPosts)       // Get all posts for CMS
			cmsProtected.POST("/posts", handlers.CreatePost)       // Create new post
			cmsProtected.GET("/posts/:id", handlers.GetPost)       // Get single post for editing
			cmsProtected.PUT("/posts/:id", handlers.UpdatePost)    // Update post
			cmsProtected.DELETE("/posts/:id", handlers.DeletePost) // Delete post
		}

		// Media management routes (requires authentication)
		protected.GET("/media", handlers.GetMedia)           // GET /api/media (list all media)
		protected.POST("/media", handlers.UploadFile)        // POST /api/media (upload file)
		protected.POST("/upload", handlers.UploadFile)       // POST /api/upload (alternative upload endpoint)
		protected.GET("/media/:id", handlers.GetMediaFile)   // GET /api/media/{id}
		protected.DELETE("/media/:id", handlers.DeleteMedia) // DELETE /api/media/{id}

		// Document management routes (requires authentication)
		protected.GET("/documents", handlers.GetDocuments)    // GET /api/documents (list all documents)
		protected.POST("/documents", handlers.UploadDocument) // POST /api/documents (upload document)

		// Page management routes (requires authentication)
		protected.GET("/pages", handlers.GetPages)             // GET /api/pages (list all pages)
		protected.POST("/pages", handlers.CreatePage)          // POST /api/pages (create page)
		protected.GET("/pages/id/:id", handlers.GetPage)       // GET /api/pages/id/{id} (get page by ID)
		protected.PUT("/pages/id/:id", handlers.UpdatePage)    // PUT /api/pages/id/{id} (update page)
		protected.DELETE("/pages/id/:id", handlers.DeletePage) // DELETE /api/pages/id/{id} (delete page)

		// Settings management routes (requires authentication)
		protected.GET("/settings", handlers.GetSettings)                     // GET /api/settings (list all settings)
		protected.GET("/settings/group/:group", handlers.GetSettingsByGroup) // GET /api/settings/group/{group}
		protected.GET("/settings/:key", handlers.GetSetting)                 // GET /api/settings/{key}
		protected.POST("/settings", handlers.CreateSetting)                  // POST /api/settings (create setting)
		protected.PUT("/settings/batch", handlers.UpdateSettingsBatch)       // PUT /api/settings/batch (update multiple)

		// Board Members management routes (requires authentication)
		protected.GET("/board-members", handlers.GetBoardMembers)             // GET /api/board-members (list all board members)
		protected.POST("/board-members", handlers.CreateBoardMember)          // POST /api/board-members (create board member)
		protected.GET("/board-members/:id", handlers.GetBoardMember)          // GET /api/board-members/{id}
		protected.PUT("/board-members/:id", handlers.UpdateBoardMember)       // PUT /api/board-members/{id}
		protected.DELETE("/board-members/:id", handlers.DeleteBoardMember)    // DELETE /api/board-members/{id}
		protected.PUT("/board-members/reorder", handlers.ReorderBoardMembers) // PUT /api/board-members/reorder

		// Team Members management routes (requires authentication) - CRUD operations only
		protected.POST("/team-members", handlers.CreateTeamMember)          // POST /api/team-members (create team member)
		protected.PUT("/team-members/:id", handlers.UpdateTeamMember)       // PUT /api/team-members/{id}
		protected.DELETE("/team-members/:id", handlers.DeleteTeamMember)    // DELETE /api/team-members/{id}
		protected.PUT("/team-members/reorder", handlers.ReorderTeamMembers) // PUT /api/team-members/reorder

		// Traders management routes (requires authentication) - CRUD operations
		protected.POST("/traders", handlers.CreateTrader)       // POST /api/traders (create trader)
		protected.PUT("/traders/:id", handlers.UpdateTrader)    // PUT /api/traders/{id}
		protected.DELETE("/traders/:id", handlers.DeleteTrader) // DELETE /api/traders/{id}

		// Brokers management routes (requires authentication) - CRUD operations
		protected.POST("/brokers", handlers.CreateBroker)       // POST /api/brokers (create broker)
		protected.PUT("/brokers/:id", handlers.UpdateBroker)    // PUT /api/brokers/{id}
		protected.DELETE("/brokers/:id", handlers.DeleteBroker) // DELETE /api/brokers/{id}

		// Partners management routes (requires authentication) - CRUD operations
		protected.GET("/cms/partners", handlers.GetAllPartners)       // GET /api/cms/partners (list all partners for CMS)
		protected.POST("/cms/partners", handlers.CreatePartner)       // POST /api/cms/partners (create partner)
		protected.GET("/cms/partners/:id", handlers.GetPartner)       // GET /api/cms/partners/{id}
		protected.PUT("/cms/partners/:id", handlers.UpdatePartner)    // PUT /api/cms/partners/{id}
		protected.DELETE("/cms/partners/:id", handlers.DeletePartner) // DELETE /api/cms/partners/{id}

		// Publications management routes (requires authentication) - CRUD operations
		protected.POST("/publications", handlers.CreatePublication)       // POST /api/publications (create publication)
		protected.PUT("/publications/:id", handlers.UpdatePublication)    // PUT /api/publications/{id}
		protected.DELETE("/publications/:id", handlers.DeletePublication) // DELETE /api/publications/{id}

		// Careers management routes (requires authentication) - CRUD operations
		protected.POST("/careers", handlers.CreateCareer)       // POST /api/careers (create career)
		protected.PUT("/careers/:id", handlers.UpdateCareer)    // PUT /api/careers/{id}
		protected.DELETE("/careers/:id", handlers.DeleteCareer) // DELETE /api/careers/{id}

		// Commodities management routes (requires authentication) - CRUD operations
		protected.POST("/commodities", handlers.CreateCommodity)       // POST /api/commodities (create commodity)
		protected.PUT("/commodities/:id", handlers.UpdateCommodity)    // PUT /api/commodities/{id}
		protected.DELETE("/commodities/:id", handlers.DeleteCommodity) // DELETE /api/commodities/{id}

		// Contract Types management routes (requires authentication) - CRUD operations
		protected.POST("/contract-types", handlers.CreateCommodityContractType)        // POST /api/contract-types (create contract type)
		protected.PUT("/contract-types/:id", handlers.UpdateCommodityContractType)     // PUT /api/contract-types/{id}
		protected.DELETE("/contract-types/:id", handlers.DeleteCommodityContractType)  // DELETE /api/contract-types/{id}
		protected.PUT("/contract-types/reorder", handlers.UpdateContractTypeSortOrder) // PUT /api/contract-types/reorder
		protected.PUT("/settings/:key", handlers.UpdateSetting)                        // PUT /api/settings/{key} (update setting)
		protected.DELETE("/settings/:key", handlers.DeleteSetting)                     // DELETE /api/settings/{key} (delete setting)

		// Events management routes (requires authentication) - CRUD operations
		protected.GET("/cms/events", handlers.GetAllEvents)                            // GET /api/cms/events (list all events for CMS)
		protected.POST("/cms/events", handlers.CreateEvent)                            // POST /api/cms/events (create event)
		protected.PUT("/cms/events/:id", handlers.UpdateEvent)                         // PUT /api/cms/events/{id}
		protected.DELETE("/cms/events/:id", handlers.DeleteEvent)                      // DELETE /api/cms/events/{id}
		protected.GET("/cms/events/stats", handlers.GetEventStats)                     // GET /api/cms/events/stats
		protected.GET("/cms/events/:id/registrations", handlers.GetEventRegistrations) // GET /api/cms/events/{id}/registrations

		// RTI management routes (requires authentication) - CRUD operations
		protected.GET("/cms/rti/requests", handlers.GetAllRTIRequests)                // GET /api/cms/rti/requests (list all RTI requests)
		protected.GET("/cms/rti/requests/:id", handlers.GetRTIRequest)                // GET /api/cms/rti/requests/{id}
		protected.PUT("/cms/rti/requests/:id", handlers.UpdateRTIRequest)             // PUT /api/cms/rti/requests/{id}
		protected.POST("/cms/rti/requests/:id/respond", handlers.RespondToRTIRequest) // POST /api/cms/rti/requests/{id}/respond
		protected.PUT("/cms/rti/requests/:id/status", handlers.UpdateRTIStatus)       // PUT /api/cms/rti/requests/{id}/status
		protected.DELETE("/cms/rti/requests/:id", handlers.DeleteRTIRequest)          // DELETE /api/cms/rti/requests/{id}
		protected.GET("/cms/rti/stats", handlers.GetRTIStats)                         // GET /api/cms/rti/stats

		// RTI documents management (requires authentication) - CRUD operations
		protected.GET("/cms/rti/documents", handlers.GetAllRTIDocuments)       // GET /api/cms/rti/documents (list all documents)
		protected.POST("/cms/rti/documents", handlers.CreateRTIDocument)       // POST /api/cms/rti/documents (create document)
		protected.PUT("/cms/rti/documents/:id", handlers.UpdateRTIDocument)    // PUT /api/cms/rti/documents/{id}
		protected.DELETE("/cms/rti/documents/:id", handlers.DeleteRTIDocument) // DELETE /api/cms/rti/documents/{id}

		// Photo gallery management (requires authentication) - CRUD operations
		protected.GET("/cms/galleries", handlers.GetAllGalleries)                  // GET /api/cms/galleries (list all galleries)
		protected.POST("/cms/galleries", handlers.CreateGallery)                   // POST /api/cms/galleries (create gallery)
		protected.PUT("/cms/galleries/:id", handlers.UpdateGallery)                // PUT /api/cms/galleries/{id}
		protected.DELETE("/cms/galleries/:id", handlers.DeleteGallery)             // DELETE /api/cms/galleries/{id}
		protected.POST("/cms/galleries/:id/photos", handlers.AddPhotoToGallery)    // POST /api/cms/galleries/{id}/photos (add photo)
		protected.PUT("/cms/galleries/photos/:id", handlers.UpdateGalleryPhoto)    // PUT /api/cms/galleries/photos/{id}
		protected.DELETE("/cms/galleries/photos/:id", handlers.DeleteGalleryPhoto) // DELETE /api/cms/galleries/photos/{id}

		// Video library management (requires authentication) - CRUD operations
		protected.GET("/cms/video-libraries", handlers.GetAllVideoLibraries)             // GET /api/cms/video-libraries (list all libraries)
		protected.POST("/cms/video-libraries", handlers.CreateVideoLibrary)              // POST /api/cms/video-libraries (create library)
		protected.PUT("/cms/video-libraries/:id", handlers.UpdateVideoLibrary)           // PUT /api/cms/video-libraries/{id}
		protected.DELETE("/cms/video-libraries/:id", handlers.DeleteVideoLibrary)        // DELETE /api/cms/video-libraries/{id}
		protected.POST("/cms/video-libraries/:id/videos", handlers.AddVideoToLibrary)    // POST /api/cms/video-libraries/{id}/videos (add video)
		protected.PUT("/cms/video-libraries/videos/:id", handlers.UpdateLibraryVideo)    // PUT /api/cms/video-libraries/videos/{id}
		protected.DELETE("/cms/video-libraries/videos/:id", handlers.DeleteLibraryVideo) // DELETE /api/cms/video-libraries/videos/{id}

		// Menu management routes (requires authentication)
		protected.GET("/menus", handlers.GetMenus)                       // GET /api/menus (list all menus)
		protected.POST("/menus", handlers.CreateMenu)                    // POST /api/menus (create menu)
		protected.GET("/menus/id/:id", handlers.GetMenu)                 // GET /api/menus/id/{id} (get menu by ID)
		protected.PUT("/menus/id/:id", handlers.UpdateMenu)              // PUT /api/menus/id/{id} (update menu)
		protected.DELETE("/menus/id/:id", handlers.DeleteMenu)           // DELETE /api/menus/id/{id} (delete menu)
		protected.GET("/menus/:menu_id/items", handlers.GetMenuItems)    // GET /api/menus/{menu_id}/items
		protected.POST("/menus/:menu_id/items", handlers.CreateMenuItem) // POST /api/menus/{menu_id}/items
		protected.PUT("/menu-items/:id", handlers.UpdateMenuItem)        // PUT /api/menu-items/{id}
		protected.DELETE("/menu-items/:id", handlers.DeleteMenuItem)     // DELETE /api/menu-items/{id}

		// News management routes (requires authentication) - CRUD operations
		protected.GET("/cms/news", handlers.GetAllNewsItems)              // GET /api/cms/news (list all news items for CMS)
		protected.POST("/cms/news", handlers.CreateNewsItem)              // POST /api/cms/news (create news item)
		protected.PUT("/cms/news/:id", handlers.UpdateNewsItem)           // PUT /api/cms/news/{id} (update news item)
		protected.DELETE("/cms/news/:id", handlers.DeleteNewsItem)        // DELETE /api/cms/news/{id} (delete news item)
		protected.POST("/cms/news/:id/publish", handlers.PublishNewsItem) // POST /api/cms/news/{id}/publish (publish news item)
		protected.POST("/cms/news/:id/archive", handlers.ArchiveNewsItem) // POST /api/cms/news/{id}/archive (archive news item)
		protected.PUT("/cms/news/:id/breaking", handlers.SetBreakingNews) // PUT /api/cms/news/{id}/breaking (set breaking news)

		// News categories management routes (requires authentication) - CRUD operations
		protected.POST("/cms/news-categories", handlers.CreateNewsCategory)       // POST /api/cms/news-categories (create news category)
		protected.PUT("/cms/news-categories/:id", handlers.UpdateNewsCategory)    // PUT /api/cms/news-categories/{id} (update news category)
		protected.DELETE("/cms/news-categories/:id", handlers.DeleteNewsCategory) // DELETE /api/cms/news-categories/{id} (delete news category)
	}
}
