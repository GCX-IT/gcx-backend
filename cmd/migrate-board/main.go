package main

import (
	"log"
	"time"

	"gcx-cms/internal/cms/models"
	"gcx-cms/internal/shared/config"
)

func main() {
	// Initialize database
	config.InitDB()

	// Check if board members already exist
	var count int64
	config.DB.Model(&models.BoardMember{}).Count(&count)
	
	if count > 0 {
		log.Printf("Board members already exist in database (%d members). Skipping migration.", count)
		return
	}

	// Create board members from existing data
	boardMembers := []models.BoardMember{
		{
			Name:        "Mr. Kofi S. Yamoah",
			Position:    "Board Chairman",
			Image:       "/Board of directors/Mr. Kofi S. Yamoah 'Board Chairman'.png",
			Description: `Mr. Yamoah's experience from working at the Ghana Stock Exchange, Central Securities Depository Ltd, Absa Bank Ghana Ltd, Venture Capital Trust Fund and as a Member of the Board of Trustees of the Accra Institute of Technology are key knowledge and skills he brings along and are grounded on years of actual practice from both executive and non-executive roles.`,
			LinkedIn:    "#",
			Facebook:    "#",
			Instagram:   "#",
			OrderIndex:  1,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			Name:        "Mr. Robert Dowouna Owoo",
			Position:    "Acting Chief Executive Officer",
			Image:       "/Board of directors/Mr. Robert Dowouna Owoo 'Acting Chief Executive officer'.jpeg",
			Description: `Mr. Owoo is a Lawyer by profession and also holds a BSc in Mathematics from the University of Cape Coast, Ghana, and an MSc in Finance and Management from Exeter University, UK. He started his professional career with the Securities and Exchange Commission, Ghana.`,
			LinkedIn:    "#",
			Facebook:    "#",
			Instagram:   "#",
			OrderIndex:  2,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	// Insert board members
	for _, member := range boardMembers {
		if err := config.DB.Create(&member).Error; err != nil {
			log.Printf("Failed to create board member %s: %v", member.Name, err)
		} else {
			log.Printf("✅ Created board member: %s", member.Name)
		}
	}

	log.Println("✅ Board member migration completed!")
}
