curl -X POST -H "Content-Type: application/json" -d '{"token":"", "citizen_id":"123456123456", "password":"test", "phone":"0123456789", "email":"23520199@gm.uit.edu.vn", "role":"patient"}'  http://localhost:8080/api/v1/accounts/register

curl -X POST -H "Content-Type: application/json" -d '{"id":"123456123456", "password":"test"}'  http://localhost:8080/api/v1/accounts/login

curl -X POST -H "Content-Type: application/json" -d '{"token":"", "citizen_id":"123123123123", "password":"test", "phone":"0111222333", "email":"hello@gm.uit.edu.vn", "role":"admin"}'  http://localhost:8080/api/v1/accounts/register

curl -X POST -H "Content-Type: application/json" -d '{"id":"123123123123", "password":"test"}'  http://localhost:8080/api/v1/accounts/login

curl -X POST -H "Content-Type: application/json" -d '{"citizen_id":"123456123456", "email":"23520199@gm.uit.edu.vn"}'  http://localhost:8080/api/v1/accounts/recovery

curl -X POST -H "Content-Type: application/json" -d '{"citizen_id":"123456123456", "otp":""}'  http://localhost:8080/api/v1/accounts/recovery/verify

curl -X PUT -H "Content-Type: application/json" -d '{"citizen_id":"123456123456", "new_password":"haha"}'  http://localhost:8080/api/v1/accounts/recovery/reset

curl -X POST -H "Content-Type: application/json" -d '{"id":"123456123456", "password":"haha"}'  http://localhost:8080/api/v1/accounts/login


curl -X POST \
  http://localhost:8080/api/v1/accounts/register/patients \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer " \
  -d '{
    "token": "",
    "citizen_id": "123456789",
    "password": "securePassword123",
    "phone": "1234567890",
    "email": "patient@example.com",
    "role": "patient",
    "full_name": "Nguyễn Văn A",
    "date_of_birth": "1990-01-15T00:00:00Z",
    "gender": "Nam",
    "ethnicity": "Kinh",
    "nationality": "Việt Nam",
    "address": "123 Đường Lê Lợi, Quận 1, TP.HCM",
    "health_insurance_expired_date": "2025-12-31T00:00:00Z",
    "health_insurance_number": "HI-987654321",
    "emergency_contact_info": "Nguyễn Thị B - 0987654321 - Mẹ"
  }'