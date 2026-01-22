package main

import (
	"log"

	"gcx-cms/internal/cms/models"
	"gcx-cms/internal/shared/config"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Initialize database
	config.InitDB()

	log.Println("🔄 Running team members migration...")

	// Create the team_members table
	if err := config.DB.AutoMigrate(&models.TeamMember{}); err != nil {
		log.Fatal("Failed to migrate team_members table:", err)
	}

	log.Println("✅ Team members table created successfully")

	// Check if we already have data
	var count int64
	config.DB.Model(&models.TeamMember{}).Count(&count)

	if count == 0 {
		log.Println("📝 Inserting sample team members data...")

		// Insert sample data
		sampleMembers := []models.TeamMember{
			// Board Members
			{
				Name:        "Mr. Kofi S. Yamoah",
				Title:       "Board Chairman",
				Description: "Experienced leader providing strategic direction and governance oversight.",
				Image:       "/Board of directors/Mr. Kofi S. Yamoah 'Board Chairman'.png",
				Type:        "board",
				OrderIndex:  1,
			},
			{
				Name:        "Mr. Kwame Daaku",
				Title:       "Non-Executive Director",
				Description: "Independent director bringing valuable industry expertise and governance experience.",
				Image:       "/Board of directors/Mr. Kwame Daaku 'Non-excutive director'.jpg",
				Type:        "board",
				OrderIndex:  2,
			},
			{
				Name:        "Mr. Stephen Antwi-Asimeng",
				Title:       "Non-Executive Director",
				Description: "Independent director with extensive experience in financial markets and corporate governance.",
				Image:       "/Board of directors/Mr. Stephen Antwi-Asimeng ' non-excutive director'.png",
				Type:        "board",
				OrderIndex:  3,
			},
			{
				Name:        "Mrs. Wendy Malm",
				Title:       "Board Secretary",
				Description: "Experienced professional ensuring effective board governance and compliance.",
				Image:       "/Board of directors/Mrs. Wendy Malm Board Secretary.png",
				Type:        "board",
				OrderIndex:  4,
			},
			// Executive Heads
			{
				Name:        "Mr. Robert Dowuona Owoo",
				Title:       "Acting Chief Executive Officer",
				Description: "Experienced leader with extensive background in commodity trading and financial markets.",
				Image:       "/Mr. Robert Dowuona Owoo.jpeg",
				Type:        "executive",
				OrderIndex:  1,
			},
			{
				Name:        "Ms. Ophelia Martekuo Atoklo",
				Title:       "Acting Deputy Chief Executive Officer",
				Description: "Strategic leader with deep expertise in operational excellence and stakeholder management.",
				Image:       "/Ms. Ophelia Martekuo Atoklo 'Deputy Chief Executive Officer'.jpg",
				Type:        "executive",
				OrderIndex:  2,
			},
			// Functional Heads
			{
				Name:        "Mr. Opoku Debrah",
				Title:       "Internal Auditor",
				Description: "Ensuring compliance and risk management across all organizational processes.",
				Image:       "/Functional Heads/Mr. Opoku Debrah (Internal Auditor).jpg",
				Type:        "functional",
				OrderIndex:  1,
			},
			{
				Name:        "Mr. Richard Ankrah",
				Title:       "Special Project",
				Description: "Leading strategic initiatives and special projects to drive organizational growth.",
				Image:       "/Functional Heads/Mr. Richard Ankrah ( Special Project).jpg",
				Type:        "functional",
				OrderIndex:  2,
			},
			{
				Name:        "Mr. Vitus Ninfaakang",
				Title:       "Risk, Membership and Partnership",
				Description: "Managing risk assessment and building strategic partnerships with key stakeholders.",
				Image:       "/Functional Heads/Mr. Vitus Ninfaakang (Risk, Membership and Partnership).jpg",
				Type:        "functional",
				OrderIndex:  3,
			},
			{
				Name:        "Mrs. Jemimah Naa Adjeley Oppong-Gyamfi",
				Title:       "Corporate Services",
				Description: "Overseeing corporate governance and administrative excellence.",
				Image:       "/Functional Heads/Mrs. Jemimah Naa Adjeley Oppong-Gyamfi ( Corporate Services).jpg",
				Type:        "functional",
				OrderIndex:  4,
			},
			{
				Name:        "Mrs. Wendy Malm",
				Title:       "Operations",
				Description: "Ensuring smooth operational processes and service delivery excellence.",
				Image:       "/Functional Heads/Mrs. Wendy Malm (Operations).jpg",
				Type:        "functional",
				OrderIndex:  5,
			},
			{
				Name:        "Mr. Godfred Kofi Nyamekye",
				Title:       "Value Chain and Product Development",
				Description: "Driving innovation in product development and value chain optimization.",
				Image:       "/Functional Heads/Mr. Godfred Kofi Nyamekye (Value Chain and Product Development).jpg",
				Type:        "functional",
				OrderIndex:  6,
			},
			{
				Name:        "Mr. Gabriel Aryeetey",
				Title:       "Warehouse & Quality",
				Description: "Maintaining quality standards and efficient warehouse operations.",
				Image:       "/Functional Heads/Mr. Gabriel Aryeetey (Warehouse & Quality).jpg",
				Type:        "functional",
				OrderIndex:  7,
			},
			{
				Name:        "Dr. Harold Okai-Tettey",
				Title:       "Information Technology and Information System",
				Description: "Leading digital transformation and technology infrastructure development.",
				Image:       "/Functional Heads/Dr. Harold Okai-Tettey ( Information Technology and Information System).jpg",
				Type:        "functional",
				OrderIndex:  8,
			},
			{
				Name:        "Mr. Albert Nii Ayi Tagoe",
				Title:       "Finance and Investments",
				Description: "Managing financial strategy and investment opportunities for sustainable growth.",
				Image:       "/Functional Heads/Mr. Albert Nii Ayi Tagoe (Finance and Investments).jpg",
				Type:        "functional",
				OrderIndex:  9,
			},
		}

		// Insert sample data
		for _, member := range sampleMembers {
			if err := config.DB.Create(&member).Error; err != nil {
				log.Printf("Failed to insert member %s: %v", member.Name, err)
			} else {
				log.Printf("✅ Inserted: %s", member.Name)
			}
		}

		log.Println("✅ Sample team members data inserted successfully")
	} else {
		log.Printf("ℹ️  Team members table already has %d records, skipping sample data insertion", count)
	}

	log.Println("🎉 Team members migration completed successfully!")
}
