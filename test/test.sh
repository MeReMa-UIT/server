# Auth and Registration stuff
curl -X POST -H "Content-Type: application/json" -d '{"id":"", "password":""}'  http://localhost:8080/api/v1/accounts/login 
curl -X POST -H "Content-Type: application/json" -d '{"citizen_id":"", "email":""}'  http://localhost:8080/api/v1/accounts/recovery
curl -X POST -H "Content-Type: application/json" -d '{"citizen_id":"", "otp":""}'  http://localhost:8080/api/v1/accounts/recovery/verify
curl -X PUT -H "Content-Type: application/json" -H "Authorization: Bearer " -d '{"new_password":""}'  http://localhost:8080/api/v1/accounts/recovery/reset
curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer " -d '{"citizen_id":""}'  http://localhost:8080/api/v1/accounts/register
curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer " -d '{"citizen_id":"", "phone":"", "email":"", "role":""}'  http://localhost:8080/api/v1/accounts/register/create
curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer " -d '{"acc_id":"", "full_name":"", "date_of_birth":"", "gender":"", "department":""}'  http://localhost:8080/api/v1/accounts/register/staffs

# Account stuff
curl -H "Authorization: Bearer " http://localhost:8080/api/v1/accounts
curl -H "Authorization: Bearer " http://localhost:8080/api/v1/accounts/profile

# Patients stuff
curl -H "Authorization: Bearer " http://localhost:8080/api/v1/patients

# Staffs stuff
curl -H "Authorization: Bearer " http://localhost:8080/api/v1/staffs

# Schedules stuff
curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer " -d '{"examination_date":"", "type":""}' http://localhost:8080/api/v1/schedules/book
curl -H "Authorization: Bearer " http://localhost:8080/api/v1/schedules?type[]=1&type[]=2&status[]=1&status[]=2&status[]=3
curl -X PUT -H "Content-Type: application/json" -H "Authorization: Bearer " -d '{"schedule_id":, "new_status":, "reception_time":""}' http://localhost:8080/api/v1/schedules/update-status
