-- Set timezone
SET GLOBAL time_zone = '+8:00';
-- Create Database
CREATE DATABASE IF NOT EXISTS operations_ecosystem;
USE operations_ecosystem;

-- Drop tables of old database if needed
-- DROP TABLE IF EXISTS `broadcast_recepients`;
-- DROP TABLE IF EXISTS `broadcast`;
-- DROP TABLE IF EXISTS `user`;
-- DROP TABLE IF EXISTS `client`;
-- DROP TABLE IF EXISTS `schedule`;
-- DROP TABLE IF EXISTS `schedule_detail`;
-- DROP TABLE IF EXISTS `aifs_client_schedule`;
-- DROP TABLE IF EXISTS `availability`;
-- DROP TABLE IF EXISTS `default_rostering`;

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
	is_part_timer BOOLEAN DEFAULT false NOT NULL,
	tele_chat_id VARCHAR(250) NOT NULL
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
    content VARCHAR(1000) NOT NULL,
    creation_date DATETIME NOT NULL,
    deadline DATETIME NOT NULL,
    creator INT NOT NULL,
	urgency ENUM('Low', 'Medium', 'High') NOT NULL,

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
    aifs_id INT NOT NULL,
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
	confirmation BOOLEAN,
	attended BOOLEAN DEFAULT false NOT NULL,
	attendance_time DATETIME, 
	is_assigned BOOLEAN NOT NULL,
    rejected BOOLEAN NOT NULL, 

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
    patrol_order INT NOT NULL, 
    
    FOREIGN KEY (schedule)
        REFERENCES `schedule` (schedule_id)
        ON UPDATE RESTRICT ON DELETE CASCADE,
	FOREIGN KEY (related_client) 
		REFERENCES `user` (user_id)
        ON UPDATE RESTRICT ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS `default_rostering` (
    default_rostering_id INT PRIMARY KEY AUTO_INCREMENT,
    day_of_week INT NOT NULL,
    aifs1_schedule INT NOT NULL,
    aifs2_schedule INT NOT NULL,
    aifs3_schedule INT NOT NULL,
 
    FOREIGN KEY (aifs1_schedule)
        REFERENCES `schedule` (schedule_id)
        ON UPDATE RESTRICT ON DELETE CASCADE,
	FOREIGN KEY (aifs2_schedule)
        REFERENCES `schedule` (schedule_id)
        ON UPDATE RESTRICT ON DELETE CASCADE,
	FOREIGN KEY (aifs3_schedule)
        REFERENCES `schedule` (schedule_id)
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
	next_sunday JSON,
    
	FOREIGN KEY (guard) 
		REFERENCES `user` (user_id)
        ON UPDATE RESTRICT ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS `incident_report_content` (
    report_content_id INT PRIMARY KEY AUTO_INCREMENT,
    last_modified_date DATETIME NOT NULL, 
    last_modifed_user INT NOT NULL,
    address VARCHAR(500) NOT NULL,
    incident_time DATETIME NOT NULL, 
    title VARCHAR(500) NOT NULL,
    is_police_notified BOOLEAN NOT NULL DEFAULT FALSE,
    description VARCHAR(1000) NOT NULL,
    has_action_taken BOOLEAN NOT NULL DEFAULT FALSE,
    action_taken VARCHAR(1000),
    has_injury BOOLEAN NOT NULL DEFAULT FALSE,
    injury_description VARCHAR(1000),
    has_stolen_item BOOLEAN NOT NULL DEFAULT FALSE,
    stolen_item_description VARCHAR(1000),
    report_image VARCHAR(1000)
);


CREATE TABLE IF NOT EXISTS `incident_report` (
    report_id INT PRIMARY KEY AUTO_INCREMENT,
    report_type  ENUM('Fire Alarm','Intruder', 'Others') NOT NULL,
    original_content INT NOT NULL,
	modified_content INT NOT NULL,
    is_approved BOOLEAN NOT NULL DEFAULT FALSE,
    signature INT,
    approval_date DATETIME, 
    aifs_id INT NOT NULL DEFAULT -1, 

	FOREIGN KEY (original_content) 
		REFERENCES `incident_report_content` (report_content_id)
        ON UPDATE RESTRICT ON DELETE CASCADE,
    FOREIGN KEY (modified_content) 
		REFERENCES `incident_report_content` (report_content_id)
        ON UPDATE RESTRICT,
    FOREIGN KEY (signature) 
		REFERENCES `user` (user_id)
        ON UPDATE RESTRICT ON DELETE CASCADE
);


CREATE TABLE IF NOT EXISTS `camera_iot` (
    camera_iot_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(1000) NOT NULL,
    camera_url VARCHAR(500) NOT NULL,
    gate_id VARCHAR(500) NOT NULL,
    gate_access_key VARCHAR(100) NOT NULL,
    fire_id VARCHAR(500) NOT NULL,
    fire_access_key	VARCHAR(100) NOT NULL,
    cpu_id VARCHAR(500) NOT NULL,
    cpu_access_key VARCHAR(100) NOT NULL
);