curl -X POST -H "Content-Type: application/json" -d '{"id":"", "password":""}'  http://localhost:8080/api/v1/accounts/login 

curl -X POST -H "Content-Type: application/json" -d '{"citizen_id":"", "email":""}'  http://localhost:8080/api/v1/accounts/recovery

curl -X POST -H "Content-Type: application/json" -d '{"citizen_id":"", "otp":""}'  http://localhost:8080/api/v1/accounts/recovery/verify

curl -X PUT -H "Content-Type: application/json" -H "Authorization: Bearer " -d '{"new_password":""}'  http://localhost:8080/api/v1/accounts/recovery/reset


curl -X POST \
  http://localhost:8080/api/v1/accounts/register/patients \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer " \
  -d '{
    "citizen_id": "123456789123",
    "password": "adu",
    "phone": "1234567890",
    "email": "23520199@gm.uit.edu.vn",
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