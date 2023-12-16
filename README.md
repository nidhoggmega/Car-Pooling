Car-pooling Platform Microservices
Design Consideration

The car-pooling platform is designed as a microservices architecture with two main services: one for handling user-related functionalities and trip-related operations. Another is for managing database-related functions
User & Trip Microservice

This microservice handles user creation, updating, deletion, and retrieval. It supports both passengers and car owners, allowing users to switch between passenger and car owner profiles.
Another section of the microservice manages car-pooling trips, including trip creation, updates, and retrieval.
Trip Microservice

This microservice handles data for both users and trips. User data is stored in a MySQL database. Trip data is stored in the same MySQL database used by the microservice.
Architecture Diagram

Architecture Diagram
![Architecture Diagram](https://github.com/nidhoggmega/Car-Pooling/assets/92702429/2af750d5-99dd-4a5b-8445-dcf7d1f36ae2)


Instructions for Setting Up and Running Microservices

    Clone the repository:

    git clone https://github.com/nidhoggmega/Car-Pooling/
    cd repository

    Set up the MySQL database:

    Create a MySQL database named "carpool".
    Update the MySQL connection details in database.go with your username and password.

3.Run the microservices in terminal:
go run main.go
This will start the microservices on localhost:8080.

4.Use the console application to interact with the microservices:

Refer to the available API endpoints in main.go.
Use tools like curl or Postman to make requests to the endpoints.

5.Clean up:

Stop the microservices when done.

API Endpoints
User Microservice

Create User: POST /user
Example:

curl -X POST -H "Content-Type: application/json" -d '{"first_name": "John", "last_name": "Doe", "mobile_number": "1234567890", "email_address": "john.doe@example.com"}' http://localhost:8080/user

Update User: PUT /user
Example:

curl -X PUT -H "Content-Type: application/json" -d '{"id": 1, "first_name": "UpdatedJohn", "last_name": "Doe", "mobile_number": "1234567890", "email_address": "john.doe@example.com"}' http://localhost:8080/user

Delete User: DELETE /user
Example:

curl -X DELETE -H "Content-Type: application/json" -d '{"id": 1}' http://localhost:8080/user

Get User: GET /user
Example:

curl http://localhost:8080/user

Trip Microservice

Create Trip: POST /trip
Example:

curl -X POST -H "Content-Type: application/json" -d '{"car_owner_id": 1, "pickup_loc
