curl -X POST -H "Content-Type: application/json" -d '{"id":"", "password":""}'  http://localhost:8080/api/v1/accounts/login 

curl -X POST -H "Content-Type: application/json" -d '{"citizen_id":"", "email":""}'  http://localhost:8080/api/v1/accounts/recovery

curl -X POST -H "Content-Type: application/json" -d '{"citizen_id":"", "otp":""}'  http://localhost:8080/api/v1/accounts/recovery/verify

curl -X PUT -H "Content-Type: application/json" -H "Authorization: Bearer " -d '{"new_password":""}'  http://localhost:8080/api/v1/accounts/recovery/reset

curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer " -d '{"citizen_id":""}'  http://localhost:8080/api/v1/accounts/register

curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer " -d '{"phone":"", "email":"", "role":""}'  http://localhost:8080/api/v1/accounts/register/confirm

curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer " -d '{"full_name":"", "date_of_birth":"", "gender":"", "department":""}'  http://localhost:8080/api/v1/accounts/register/staffs