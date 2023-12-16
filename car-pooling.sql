CREATE database my_db;

USE my_db;
CREATE TABLE Persons (ID varchar (5) NOT NULL PRIMARY KEY, FirstName VARCHAR(30), LastName VARCHAR(30), Age INT); 

INSERT INTO Persons (ID, FirstName, LastName, Age) VALUES ("0001", "Rob", "Pike", 66);
CREATE USER 'car'@'localhost' IDENTIFIED BY 'password';
GRANT ALL ON *.* TO 'car'@'localhost';
-- users table
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    mobile_number VARCHAR(15) NOT NULL,
    email_address VARCHAR(255) NOT NULL,
    driver_license VARCHAR(20),
    car_plate VARCHAR(15),
    is_car_owner BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- trips table
CREATE database carpooling;

use carpool;
CREATE TABLE trips (
    id INT AUTO_INCREMENT PRIMARY KEY,
    car_owner_id INT NOT NULL,
    pickup_location VARCHAR(255) NOT NULL,
    alt_pickup_location VARCHAR(255),
    start_time DATETIME NOT NULL,
    destination VARCHAR(255) NOT NULL,
    max_passengers INT NOT NULL,
    seats_available INT NOT NULL,
    FOREIGN KEY (car_owner_id) REFERENCES users(id)
);

