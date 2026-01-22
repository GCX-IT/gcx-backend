// Migration script to add existing board members to the database via API
const axios = require('axios');

const API_BASE = 'http://localhost:8080/api';

// Login credentials
const loginData = {
  email: 'admin@gcx.com',
  password: 'admin123'
};

// Board members data
const boardMembers = [
  {
    name: "Mr. Kofi S. Yamoah",
    position: "Board Chairman",
    image: "/Board of directors/Mr. Kofi S. Yamoah 'Board Chairman'.png",
    description: `Mr. Yamoah's experience from working at the Ghana Stock Exchange, Central Securities Depository Ltd, Absa Bank Ghana Ltd, Venture Capital Trust Fund and as a Member of the Board of Trustees of the Accra Institute of Technology are key knowledge and skills he brings along and are grounded on years of actual practice from both executive and non-executive roles. They come in especially handy in current times when attention to strategic direction in a post-Covid new normal, challenging and highly competitive environment is very much a necessity. Furthermore, his deep understanding of corporate governance, risk mitigation, regulatory compliance and sustainability issues will be highly beneficial to any entity he may serve whether in an executive role or otherwise, in both private and public sectors.`,
    linkedin: "#",
    facebook: "#",
    instagram: "#",
    order_index: 1
  },
  {
    name: "Mr. Stephen Antwi-Asimeng",
    position: "Non-Executive Director",
    image: "/Board of directors/Mr. Stephen Antwi-Asimeng ' non-excutive director'.png",
    description: `Mr. Stephen Antwi-Asimeng is an independent financial advisor providing banking, investment, corporate finance, and private equity advisory services to large corporate and SME firms. He has over twenty-eight (28) years' experience in the banking, financial services, and SME business community including over eighteen (18) years in senior executive management positions. Stephen is a valued member and chair of local and international boards of directors, offering strategic leadership and corporate governance support. He served on the board of Africa Venture Capital & Private Equity Association for six (6) years from 2004. Stephen has extensive domain experience in several sectors including banking & financial services, agribusiness, education, pharmaceutical and healthcare. He was selected by the Economic Commission of Africa (ECA) as a member of an ad hoc team of experts on private equity and venture capital as an asset class in SME finance in Africa in 2013 and 2015. Mr. Antwi-Asimeng holds a Master of Arts degree in Banking and Finance from University of Wales, United Kingdom and a Bachelor of Arts degree in Economics and Sociology from the University of Ghana, Legon. He has received executive education from several reputable institutions including INSEAD, Amsterdam Institute of Finance and Harvard Business School.`,
    linkedin: "#",
    facebook: "#",
    instagram: "#",
    order_index: 2
  },
  {
    name: "Mr. Kwame Daaku",
    position: "Non-Executive Director",
    image: "/Board of directors/Mr. Kwame Daaku 'Non-excutive director'.jpg",
    description: `Mr. Daaku is a Banker and a financial consultant with over 18 years of experience. He is a Director of Afintek Limited- a financial consulting and management firm. He is also the Executive Director and Cofounder of Nature's Treasure Foods – a food tech startup and an early investor and advisor to Intelocate in Toronto, Canada- a tech company in Toronto in the field of task management. He currently consults on Infrastructure and Project Finance, Derivatives and Mortgage Financing. He began his career with a boutique consulting firm in Accra, Ghana, working with SMEs on strategy and cost optimization. He then joined Ecobank Ghana Limited where he managed credit relationships within the supply chain of regional multinational customers of the bank. He later joined Bank of Nova Scotia in Toronto, Canada in 2006 as a Manager in Special Accounts Management where he helped manage the bank's global corporate credit debt recovery portfolio. He subsequently joined their Global Risk Market team as a Senior Manager managing the Specialized Lending, Asset Managers and Municipal Finance Credit Risk portfolio for the US Market. He then moved to Global Capital Markets Banking as Associate Director for Precious Metals, Base Metals and Foreign Exchange managing an over US$45Bn client relationship portfolio, which involved Asset Managers, Reinsurance, Banks, Sovereign Wealth Funds and Corporates. Kwame is a graduate of the University of Ghana and has an MBA from the Ivey School of Business at Western University in Canada and a Certificate in Emerging Technologies and Renewable Energy from Stanford University. Kwame is an avid volunteer and has raised funds for United Way in Toronto and help establish the United Way GenNEXT on Bay Street, Toronto. He was instrumental in setting up Golden Futures in Canada, an NGO that engaged tutoring young minds in townships in Cape Town together students from the University of Windsor, York University in Toronto as well as Wayne State University in Detroit. He is an avid cyclist and runner. Over the years, he has used his passion for outdoor activities to raise funds for cancer research through the Ride to Conquer Cancer, as well for Juvenile Diabetes Research, Multiple Sclerosis, Becel Ride for Heart for Heart and Stroke Research.`,
    linkedin: "#",
    facebook: "#",
    instagram: "#",
    order_index: 3
  },
  {
    name: "Mr. Robert Dowouna Owoo",
    position: "Acting Chief Executive Officer",
    image: "/Board of directors/Mr. Robert Dowouna Owoo 'Acting Chief Executive officer'.jpeg",
    description: `Mr. Owoo is a Lawyer by profession and also holds a BSc in Mathematics from the University of Cape Coast, Ghana, and an MSc in Finance and Management from Exeter University, UK. He started his professional career with the Securities and Exchange Commission, Ghana, where he worked in various capacities and has thirteen (13) years' experience in Securities (Capital Market) Regulations and Supervision. Mr. Owoo served as the project director for the Ghana Commodity Exchange project from 2013 and led the Ghana team to design the current Ghana Commodity Exchange (GCX) model for Ghana. One of the founding members of the GCX, he is currently the Acting Chief Executive Officer of the GCX.`,
    linkedin: "#",
    facebook: "#",
    instagram: "#",
    order_index: 4
  },
  {
    name: "Mrs. Wendy Malm",
    position: "Board Secretary",
    image: "/Board of directors/Mrs. Wendy Malm Board Secretary.png",
    description: `Wendy Malm is the head of Trading, Central depository and surveillance at the Ghana Commodity Exchange. In this role she has developed commodity contracts, executed several trading strategies and has continually supervised the processing and execution of all market transactions also ensured that all trades conform to the regulatory requirements of the GCX and the Securities and Exchange Commission. She holds a BA (Hons.) in Economics & Geography from KNUST, Kumasi, Ghana and an MBA in Banking and Finance from the Paris Graduate School of Management, France. She is also certified by the Ghana Stock Exchange as an Authorized Dealing Officer and is a member of the Ghana Stockbrokers Association and a student member of the Institute of Chartered accountants Ghana. Wendy is part of the team that developed the systems and processes towards the setup of the GCX. She represents GCX on the Capital Market Master Plan Group for the development of markets and products. She also serves on the Internal Audit, Health and Safety and Brokers technical Committees at GCX. Prior to joining GCX, she served in several capacities within the SIC group for thirteen years attaining a rank of head of trading. She has Twelve years' experience in Stock trading and operations, Eight (8) years' experience in treasury and fixed income trading and operations and a years' experience in Fund/Portfolio Management. At SIC Brokerage Limited, she also served as a member of the investment committee and shortly as a compliance officer.`,
    linkedin: "#",
    facebook: "#",
    instagram: "#",
    order_index: 5
  }
];

async function migrateBoard() {
  try {
    console.log('🚀 Starting board members migration...');
    
    // Step 1: Login to get authentication token
    console.log('🔐 Logging in...');
    const loginResponse = await axios.post(`${API_BASE}/auth/login`, loginData);
    const token = loginResponse.data.token;
    
    if (!token) {
      throw new Error('Failed to get authentication token');
    }
    
    console.log('✅ Login successful');
    
    // Step 2: Check if board members already exist
    console.log('🔍 Checking existing board members...');
    const headers = { 'Authorization': `Bearer ${token}` };
    
    try {
      const existingResponse = await axios.get(`${API_BASE}/board-members`, { headers });
      if (existingResponse.data.success && existingResponse.data.data.length > 0) {
        console.log(`⚠️  Found ${existingResponse.data.data.length} existing board members. Skipping migration.`);
        return;
      }
    } catch (error) {
      console.log('📝 No existing board members found, proceeding with migration...');
    }
    
    // Step 3: Create board members
    console.log('📥 Creating board members...');
    
    for (const member of boardMembers) {
      try {
        const response = await axios.post(`${API_BASE}/board-members`, member, { headers });
        if (response.data.success) {
          console.log(`✅ Created: ${member.name}`);
        } else {
          console.log(`❌ Failed to create: ${member.name} - ${response.data.error}`);
        }
      } catch (error) {
        console.log(`❌ Error creating ${member.name}:`, error.response?.data?.error || error.message);
      }
    }
    
    console.log('🎯 Board members migration completed!');
    console.log('🎉 You can now manage your board members through the CMS Board Manager!');
    
  } catch (error) {
    console.error('💥 Migration failed:', error.response?.data?.error || error.message);
  }
}

// Run the migration
migrateBoard();
