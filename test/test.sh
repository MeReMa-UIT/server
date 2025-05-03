curl -X POST -H "Content-Type: application/json" -d '{"citizen_id":"123456123456", "password":"haha"}'  http://localhost:8080/api/v1/accounts/login

curl -X POST -H "Content-Type: application/json" -d '{"citizen_id":"123456123456", "password":"test", "phone":"0123456789", "email":"23520199@gm.uit.edu.vn", "role":"patient"}'  http://localhost:8080/api/v1/accounts/register

curl -X POST -H "Content-Type: application/json" -d '{"citizen_id":"123456123456", "email":"23520199@gm.uit.edu.vn"}'  http://localhost:8080/api/v1/accounts/recovery

curl -X POST -H "Content-Type: application/json" -d '{"citizen_id":"123456123456", "otp":""}'  http://localhost:8080/api/v1/accounts/recovery/verify

curl -X PUT -H "Content-Type: application/json" -d '{"citizen_id":"123456123456", "new_password":"haha"}'  http://localhost:8080/api/v1/accounts/recovery/reset