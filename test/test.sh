curl -X POST -H "Content-Type: application/json" -d '{"citizen_id":"123456123457", "password":"test"}'  http://localhost:8080/api/v1/accounts/login

curl -X POST -H "Content-Type: application/json" -d '{"citizen_id":"123456123456", "password":"test", "phone":"0123456789", "email":"test@gmail.com", "role":"patient"}'  http://localhost:8080/api/v1/accounts/register