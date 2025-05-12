SELECT    *
FROM      accounts;

SELECT    *
FROM      patients;

SELECT    *
FROM      staffs;

DELETE    FROM accounts
WHERE     acc_id > 1;

DELETE    FROM patients
WHERE     patient_id >= 1;

UPDATE    accounts
SET       role = 'admin'
WHERE     acc_id >= 1;

CREATE    TABLE test (path text);