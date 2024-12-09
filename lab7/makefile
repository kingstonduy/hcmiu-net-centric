create-user:
	curl -X POST http://localhost:8080/create-user -H "Content-Type: application/json" -d '{"first_name": "John", "last_name": "Doe", "username": "johndoe", "email": "johndoe@example.com", "phone_number": "1234567890", "date_of_birth": "1990-01-01T00:00:00Z", "address_country": "USA", "address_city": "New York", "address_street_name": "5th Ave", "address_street_address": "1234"}'

create-random:
	curl -X POST http://localhost:8080/create-user -H "Content-Type: application/json" -d '{}'
get-user:
	curl -X POST http://localhost:8080/get-user -H "Content-Type: application/json" -d '{"user_id": 1}'

update-user:
	curl -X POST http://localhost:8080/update-user -H "Content-Type: application/json" -d '{"user_id": 1, "first_name": "Jane", "last_name": "Doe", "username": "janedoe", "email": "janedoe@example.com", "phone_number": "9876543210", "date_of_birth": "1990-01-01T00:00:00Z", "address_country": "USA", "address_city": "New York", "address_street_name": "5th Ave", "address_street_address": "5678"}'

delete-user:
	curl -X POST http://localhost:8080/delete-user -H "Content-Type: application/json" -d '{"user_id": 1}'
	
get-user-with-name:
	curl "http://localhost:8080/get-users?first_name=Tequila&sort_by=last_name&sort_order=desc"
