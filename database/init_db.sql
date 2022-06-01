-- Set timezone
SET GLOBAL time_zone = '+8:00';
-- Create Database
CREATE DATABASE IF NOT EXISTS operations_ecosystem;
USE operations_ecosystem;

-- Drop tables of old database if needed
-- DROP TABLE IF EXISTS `broadcast_recepients`;
-- DROP TABLE IF EXISTS `broadcast`;
-- DROP TABLE IF EXISTS `user`;
DROP TABLE IF EXISTS `client`;
DROP TABLE IF EXISTS `schedule`;
DROP TABLE IF EXISTS `schedule_detail`;
DROP TABLE IF EXISTS `aifs_client_schedule`;
DROP TABLE IF EXISTS `availability`;

-- Create Tables
-- Admin Tables 
CREATE TABLE IF NOT EXISTS `user` (
	user_id INT PRIMARY KEY AUTO_INCREMENT,
	user_type ENUM('I-Specialist','Security Guard', 'Controller', 'Manager') NOT NULL,
	name VARCHAR(200) NOT NULL,
	email VARCHAR(500) NOT NULL,
	phone_number VARCHAR(500) NOT NULL,
	telegram_handle VARCHAR(500) NOT NULL,
	user_security_img VARCHAR(1000) NOT NULL,
	is_part_timer BOOLEAN DEFAULT false NOT NULL
);

CREATE TABLE IF NOT EXISTS `client` (
	client_id INT PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(200) NOT NULL,
	abbreviation VARCHAR(200) NOT NULL,
	email VARCHAR(500) NOT NULL,
	address VARCHAR(500) NOT NULL,
    postal_code INT NOT NULL, 
	phone_number VARCHAR(500) NOT NULL
);

-- Broadcasting
CREATE TABLE IF NOT EXISTS `broadcast` (
	broadcast_id INT PRIMARY KEY AUTO_INCREMENT,
    type ENUM('announcement','assignment') NOT NULL,
    title VARCHAR(500) NOT NULL, 
    content VARCHAR(1000) NOT NULL,
    creation_date DATETIME NOT NULL,
    deadline DATETIME NOT NULL,
    creator INT NOT NULL,
	
	FOREIGN KEY (creator) 
		REFERENCES `user` (user_id)
        ON UPDATE RESTRICT ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `broadcast_recepients` (
    broadcast_recipients_id INT PRIMARY KEY AUTO_INCREMENT,
    related_broadcast INT NOT NULL,
    recipient INT NOT NULL, 
    acknowledged BOOLEAN DEFAULT false NOT NULL,
    rejected BOOLEAN DEFAULT false NOT NULL,
	last_replied DATETIME, 
    aids_id INT NOT NULL,
    FOREIGN KEY (related_broadcast)
        REFERENCES `broadcast` (broadcast_id)
        ON UPDATE RESTRICT ON DELETE CASCADE,
	FOREIGN KEY (recipient) 
		REFERENCES `user` (user_id)
        ON UPDATE RESTRICT ON DELETE CASCADE
);


-- Rostering
CREATE TABLE IF NOT EXISTS `schedule` (
    schedule_id INT PRIMARY KEY AUTO_INCREMENT,
    aifs_id INT NOT NULL,
	start_time DATETIME NOT NULL, 
	end_time DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS `schedule_detail` (
    schedule_detail_id INT PRIMARY KEY AUTO_INCREMENT,
    schedule INT NOT NULL,
    guard_assigned INT NOT NULL,
	custom_start_time DATETIME NOT NULL, 
	custom_end_time DATETIME NOT NULL, 
	confirmation BOOLEAN DEFAULT false NOT NULL,
	attended BOOLEAN DEFAULT false NOT NULL,
	attendance_time DATETIME, 

    FOREIGN KEY (schedule)
        REFERENCES `schedule` (schedule_id)
        ON UPDATE RESTRICT ON DELETE CASCADE,
	FOREIGN KEY (guard_assigned) 
		REFERENCES `user` (user_id)
        ON UPDATE RESTRICT ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `aifs_client_schedule` (
    aifs_client_schedule_id INT PRIMARY KEY AUTO_INCREMENT,
    schedule INT NOT NULL,
    related_client INT NOT NULL,

    FOREIGN KEY (schedule)
        REFERENCES `schedule` (schedule_id)
        ON UPDATE RESTRICT ON DELETE CASCADE,
	FOREIGN KEY (related_client) 
		REFERENCES `user` (user_id)
        ON UPDATE RESTRICT ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `availability` (
    availability_id INT PRIMARY KEY AUTO_INCREMENT,
    week INT NOT NULL,
    year INT NOT NULL,
	guard INT NOT NULL,
	sunday JSON,
	monday JSON,
	tuesday JSON,
	wednesday JSON,
	thursday JSON,
	friday JSON,
	saturday JSON,
    
	FOREIGN KEY (guard) 
		REFERENCES `user` (user_id)
        ON UPDATE RESTRICT ON DELETE CASCADE
);