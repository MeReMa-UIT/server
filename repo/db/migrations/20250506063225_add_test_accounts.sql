-- migrate:up
INSERT    INTO accounts (citizen_id, password_hash, phone, email, role)
VALUES    (
          '123123123123',
          '$2a$10$bqcjrpmJ8qID2bkUscD15uFKzE0tiYFpbM3oex3GxCeqDy9IuhZ2K',
          '0111222333',
          '23520199@gm.uit.edu.vn',
          'admin'
          );

INSERT    INTO accounts (citizen_id, password_hash, phone, email, role)
VALUES    (
          '123412341234',
          '$2a$10$bqcjrpmJ8qID2bkUscD15uFKzE0tiYFpbM3oex3GxCeqDy9IuhZ2K',
          '0123456789',
          '23521734@gm.uit.edu.vn',
          'admin'
          );

INSERT    INTO accounts (citizen_id, password_hash, phone, email, role)
VALUES    (
          '000000001111',
          '$2a$10$bqcjrpmJ8qID2bkUscD15uFKzE0tiYFpbM3oex3GxCeqDy9IuhZ2K',
          '0987654321',
          'recep@merema.com',
          'receptionist'
          );

INSERT    INTO accounts (citizen_id, password_hash, phone, email, role)
VALUES    (
          '000000001112',
          '$2a$10$bqcjrpmJ8qID2bkUscD15uFKzE0tiYFpbM3oex3GxCeqDy9IuhZ2K',
          '0987654322',
          'doctor@merema.com',
          'doctor'
          );

INSERT    INTO accounts (citizen_id, password_hash, phone, email, role)
VALUES    (
          '000000001113',
          '$2a$10$bqcjrpmJ8qID2bkUscD15uFKzE0tiYFpbM3oex3GxCeqDy9IuhZ2K',
          '0987654323',
          'patient@merema.com',
          'patient'
          );

INSERT    INTO staffs (acc_id, full_name, date_of_birth, gender, department)
VALUES    (
          (
          SELECT    acc_id
          FROM      accounts
          WHERE     citizen_id = '000000001111'
          ),
          'Nguyễn Thị Hoa',
          '1995-04-06',
          'Nữ',
          'Phòng Hành chính'
          );

INSERT    INTO staffs (acc_id, full_name, date_of_birth, gender, department)
VALUES    (
          (
          SELECT    acc_id
          FROM      accounts
          WHERE     citizen_id = '000000001112'
          ),
          'Nguyễn Văn Tài',
          '1997-11-21',
          'Nam',
          'Khoa Tai Mũi Họng'
          );

INSERT    INTO patients (
          acc_id,
          full_name,
          date_of_birth,
          gender,
          ethnicity,
          nationality,
          address,
          health_insurance_expired_date,
          health_insurance_number,
          emergency_contact_info
          )
VALUES    (
          (
          SELECT    acc_id
          FROM      accounts
          WHERE     citizen_id = '000000001113'
          ),
          'Nguyễn Văn A',
          '1993-05-01',
          'Nam',
          'Kinh',
          'Việt Nam',
          '123 Đường ABC, Phường 1, Quận 1, TP.HCM',
          '2025-01-01',
          '123456789012',
          'Nguyễn Thị B, 0987654321'
          );

-- migrate:down
DELETE    FROM accounts
WHERE     citizen_id = '123123123123';

DELETE    FROM accounts
WHERE     citizen_id = '123412341234';

DELETE    FROM staffs
WHERE     acc_id = (
          SELECT    acc_id
          FROM      accounts
          WHERE     citizen_id = '000000001111'
          );

DELETE    FROM accounts
WHERE     citizen_id = '000000001111';

DELETE    FROM staffs
WHERE     acc_id = (
          SELECT    acc_id
          FROM      accounts
          WHERE     citizen_id = '000000001112'
          );

DELETE    FROM accounts
WHERE     citizen_id = '000000001112';

DELETE    FROM patients
WHERE     acc_id = (
          SELECT    acc_id
          FROM      accounts
          WHERE     citizen_id = '000000001113'
          );

DELETE    FROM accounts
WHERE     citizen_id = '000000001113';