# Partners Data Setup

This document explains how to set up the partners data in your database.

## Prerequisites

1. Make sure your Go backend is running and the database is connected
2. Ensure the partners table has been created (it should be auto-migrated when you start the backend)

## Steps to Insert Partners Data

### Option 1: Using MySQL Command Line

```bash
# Connect to your MySQL database
mysql -u your_username -p your_database_name

# Run the SQL file
source partners_data.sql;
```

### Option 2: Using MySQL Workbench or phpMyAdmin

1. Open your MySQL management tool
2. Select your database
3. Open the `partners_data.sql` file
4. Execute the SQL commands

### Option 3: Using Go Backend (if you have a database admin endpoint)

You can also create a simple admin endpoint to run the SQL, but the above methods are recommended.

## What the SQL Does

The `partners_data.sql` file inserts all the existing partners from your website into the database, including:

- **Financial Institutions** (10 partners): Ghana EXIM Bank, Standard Chartered, Fidelity Bank, etc.
- **Donor Agencies** (9 partners): USAID, AGRA, IFC, GIZ, etc.
- **Government Agencies** (5 partners): Youth in Agriculture, GIRSAL, GIPC, etc.
- **NGOs** (9 partners): Ghana Made, PEF, Aflasafe, etc.
- **Private Agencies** (5 partners): DMT Collateral, Wienco, TATF, etc.
- **Tenders & Procurement** (2 partners): Ghana Procurement Authority, PPRA

Each partner includes:
- Name and description
- Category (financial, donor, government, ngo, private, tender)
- Logo path (using existing images from your public folder)
- Website, email, phone, address
- Status (all set to 'active')

## After Running the SQL

1. Restart your Go backend to ensure it picks up the new data
2. The frontend will automatically fetch and display the partners from the API
3. You can now manage partners through the CMS at `/cms/partners`

## Verification

To verify the data was inserted correctly:

1. Check the CMS partners page: `http://localhost:5173/cms/partners`
2. Check the public partnerships page: `http://localhost:5173/about/partnerships`
3. Check the landing page partners section

All partners should now be dynamic and manageable through the CMS!
