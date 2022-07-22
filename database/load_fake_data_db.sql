-- Set timezone
SET GLOBAL time_zone = '+8:00';
-- Create Database
USE operations_ecosystem;
-- Load Fake Data if needed
-- To do this, one must put the files into the secure file priv folder
-- to find this folder run:
-- SHOW VARIABLES LIKE "secure_file_priv";
SHOW VARIABLES LIKE "secure_file_priv";

LOAD DATA INFILE 'P:/ProgramData/MySQL/MySQL Server 8.0/Uploads/operations_ecosys_fake_data/Fake Data - user.csv' 
INTO TABLE user 
FIELDS TERMINATED BY ','
LINES TERMINATED BY '\n'
IGNORE 1 ROWS;

LOAD DATA INFILE 'P:/ProgramData/MySQL/MySQL Server 8.0/Uploads/operations_ecosys_fake_data/Fake Data - client.csv' 
INTO TABLE client 
FIELDS TERMINATED BY ','
LINES TERMINATED BY '\n'
IGNORE 1 ROWS;

LOAD DATA INFILE 'P:/ProgramData/MySQL/MySQL Server 8.0/Uploads/operations_ecosys_fake_data/Fake Data - broadcast.csv' 
INTO TABLE broadcast 
FIELDS OPTIONALLY ENCLOSED BY '"' TERMINATED BY ','
LINES TERMINATED BY '\n'
IGNORE 1 ROWS;

LOAD DATA INFILE 'P:/ProgramData/MySQL/MySQL Server 8.0/Uploads/operations_ecosys_fake_data/Fake Data - broadcast_recepients.csv' 
INTO TABLE broadcast_recepients
FIELDS TERMINATED BY ','
LINES TERMINATED BY '\n'
IGNORE 1 ROWS;
