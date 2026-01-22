-- SQL Schema to insert existing partners data
-- Run this after the partners table has been created

-- Financial Institutions
INSERT INTO partners (name, description, category, logo, website, email, phone, address, status, created_at, updated_at) VALUES
('Ghana Export-Import Bank', 'Ghana''s premier export-import bank providing financial services to support international trade and economic development.', 'financial', '/Partners/ghana-exim-bank.jpg', 'https://www.ghanaeximbank.com', 'info@ghanaeximbank.com', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('Ghana Export Promotion Authority', 'Promoting Ghana''s exports and facilitating trade relationships with international partners.', 'financial', '/Partners/ghana-export-promotion-authority-gepa.png', 'https://www.gepa.gov.gh', 'info@gepa.gov.gh', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('Ghana Standards Authority', 'Ensuring quality standards and certification for Ghana''s products and services.', 'financial', '/Partners/ghana-standard-authority-gsa.png', 'https://www.gsa.gov.gh', 'info@gsa.gov.gh', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('Standard Chartered', 'International banking services supporting Ghana''s economic growth and development.', 'financial', '/Partners/standard-chartered.jpg', 'https://www.sc.com/gh', 'ghana@sc.com', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('Fidelity Bank', 'Leading Ghanaian bank providing comprehensive financial services.', 'financial', '/Partners/fidelity-bank.png', 'https://www.fidelitybank.com.gh', 'info@fidelitybank.com.gh', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('Ecobank', 'Pan-African bank supporting economic development across the continent.', 'financial', '/Partners/ecobank.png', 'https://www.ecobank.com/gh', 'ghana@ecobank.com', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('Ghana Grains Council', 'Promoting the development of Ghana''s grain industry and market access.', 'financial', '/Partners/ghana-grains-council-ggc.png', 'https://www.ghanagrainscouncil.org', 'info@ghanagrainscouncil.org', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('IPMC', 'Information technology and business solutions provider.', 'financial', '/Partners/ipmc.jpg', 'https://www.ipmc.com.gh', 'info@ipmc.com.gh', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('Africa Cashew Alliance', 'Promoting sustainable cashew production and trade across Africa.', 'financial', '/Partners/africa-cashew-alliance.png', 'https://www.africacashewalliance.org', 'info@africacashewalliance.org', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('CIAG', 'Supporting agricultural development and food security initiatives.', 'financial', '/Partners/6-ciag.jpg', 'https://www.ciag.org.gh', 'info@ciag.org.gh', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW());

-- Donor Agencies
INSERT INTO partners (name, description, category, logo, website, email, phone, address, status, created_at, updated_at) VALUES
('USAID', 'United States Agency for International Development supporting development programs in Ghana.', 'donor', '/Donors/usaid.png', 'https://www.usaid.gov/ghana', 'ghana@usaid.gov', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('AGRA', 'Alliance for a Green Revolution in Africa promoting sustainable agricultural development.', 'donor', '/Donors/agra.png', 'https://www.agra.org', 'info@agra.org', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('IFC', 'International Finance Corporation supporting private sector development in Ghana.', 'donor', '/Donors/ifc.png', 'https://www.ifc.org/ghana', 'ghana@ifc.org', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('GIZ', 'German development agency supporting sustainable development initiatives.', 'donor', '/Donors/giz-logo.gif', 'https://www.giz.de/en/worldwide/ghana.html', 'info@giz.de', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('SNV', 'Netherlands Development Organisation supporting inclusive development.', 'donor', '/Donors/snv.png', 'https://www.snv.org/country/ghana', 'ghana@snv.org', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('UNDP', 'United Nations Development Programme supporting sustainable development goals.', 'donor', '/Donors/undp_100.png', 'https://www.gh.undp.org', 'ghana@undp.org', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('UNIDO', 'United Nations Industrial Development Organization promoting industrial development.', 'donor', '/Donors/unido.png', 'https://www.unido.org/ghana', 'ghana@unido.org', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('IFAD', 'International Fund for Agricultural Development supporting rural development.', 'donor', '/Donors/ifad-a-edit.png', 'https://www.ifad.org/en/ghana', 'ghana@ifad.org', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('World Food Programme', 'United Nations agency fighting hunger and promoting food security.', 'donor', '/Donors/world-food-programme-wfp.jpg', 'https://www.wfp.org/countries/ghana', 'ghana@wfp.org', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW());

-- Government Agencies
INSERT INTO partners (name, description, category, logo, website, email, phone, address, status, created_at, updated_at) VALUES
('Youth in Agriculture Programme', 'Government initiative supporting youth participation in agricultural development.', 'government', '/government/youth-in-agriculture-programme-yiap.png', 'https://www.moagd.gov.gh', 'yiap@moagd.gov.gh', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('Ghana Incentive-Based Risk-Sharing System for Agricultural Lending', 'Government program facilitating agricultural financing and risk management.', 'government', '/government/ghana-incentive-based-risk-sharing-system-for-agricultural-lending-girsal.png', 'https://www.girsal.com', 'info@girsal.com', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('Ghana Investment Promotion Centre', 'Government agency promoting investment opportunities in Ghana.', 'government', '/government/ghana-investment-promotion-centregipc.png', 'https://www.gipc.gov.gh', 'info@gipc.gov.gh', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('Ministry of Finance and Economic Planning', 'Government ministry overseeing economic policy and financial planning.', 'government', '/government/ministry-of-finance-and-economic-planning-100.png', 'https://www.mofep.gov.gh', 'info@mofep.gov.gh', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('Ministry of Trade and Industry', 'Government ministry promoting trade and industrial development.', 'government', '/government/ministry-of-trade-and-industry-moti-100.png', 'https://www.moti.gov.gh', 'info@moti.gov.gh', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW());

-- NGOs
INSERT INTO partners (name, description, category, logo, website, email, phone, address, status, created_at, updated_at) VALUES
('Ghana Made', 'Promoting locally made products and supporting Ghanaian manufacturers.', 'ngo', '/NGO/ghana-made.png', 'https://www.ghanamade.org', 'info@ghanamade.org', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('PEF', 'Private Enterprise Federation supporting private sector development.', 'ngo', '/NGO/pef.png', 'https://www.pef.org.gh', 'info@pef.org.gh', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('Aflasafe', 'Promoting aflatoxin control and food safety in agricultural products.', 'ngo', '/NGO/aflasafe2.png', 'https://www.aflasafe.com', 'info@aflasafe.com', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('Socodevi', 'Supporting cooperative development and rural entrepreneurship.', 'ngo', '/NGO/socodevi.jpeg', 'https://www.socodevi.org', 'ghana@socodevi.org', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('Ghana National Chamber of Commerce and Industry', 'Representing the interests of Ghana''s business community.', 'ngo', '/NGO/ghana-national-chamber-of-commerce-and-industry.jpg', 'https://www.ghanachamber.org', 'info@ghanachamber.org', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('Ghana Rice Inter-Professional Body', 'Promoting rice production and processing in Ghana.', 'ngo', '/NGO/ghana-rice-inter-professional-body-grib.png', 'https://www.grib.org.gh', 'info@grib.org.gh', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('Peasant Farmers Association of Ghana', 'Representing smallholder farmers and promoting their interests.', 'ngo', '/NGO/peasant-farmers-association-of-ghana-pfag.jpg', 'https://www.pfag.org.gh', 'info@pfag.org.gh', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('The John A. Kufuor Foundation', 'Promoting good governance and development initiatives.', 'ngo', '/NGO/the-john-a-kufuor-foundation.jpg', 'https://www.kufuorfoundation.org', 'info@kufuorfoundation.org', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('IITA', 'International Institute of Tropical Agriculture supporting agricultural research.', 'ngo', '/NGO/iita.png', 'https://www.iita.org', 'ghana@iita.org', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW());

-- Private Agencies
INSERT INTO partners (name, description, category, logo, website, email, phone, address, status, created_at, updated_at) VALUES
('DMT Collateral', 'Providing collateral management services for agricultural commodities.', 'private', '/Private/dmt-collateral.png', 'https://www.dmtcollateral.com', 'info@dmtcollateral.com', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('Intervalle Geneve', 'International trading company supporting agricultural commodity trade.', 'private', '/Private/intervalle-geneve.png', 'https://www.intervalle-geneve.com', 'info@intervalle-geneve.com', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('Wienco', 'Agricultural input supplier and commodity trading company.', 'private', '/Private/wienco.jpg', 'https://www.wienco.com', 'info@wienco.com', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('TATF', 'Technology and agricultural training foundation supporting farmers.', 'private', '/Private/tatf.png', 'https://www.tatf.org.gh', 'info@tatf.org.gh', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('DESS Inc', 'Development and environmental services supporting sustainable agriculture.', 'private', '/Private/dess-inc.png', 'https://www.dessinc.com', 'info@dessinc.com', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW());

-- Tenders & Procurement (using some existing partners as examples)
INSERT INTO partners (name, description, category, logo, website, email, phone, address, status, created_at, updated_at) VALUES
('Ghana Procurement Authority', 'Government agency managing public procurement processes.', 'tender', '/government/ghana-investment-promotion-centregipc.png', 'https://www.ppa.gov.gh', 'info@ppa.gov.gh', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW()),

('Public Procurement Regulatory Authority', 'Regulating and monitoring public procurement in Ghana.', 'tender', '/government/ministry-of-finance-and-economic-planning-100.png', 'https://www.ppra.gov.gh', 'info@ppra.gov.gh', '+233 302 500 000', 'Accra, Ghana', 'active', NOW(), NOW());
