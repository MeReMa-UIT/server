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

DROP      TABLE IF EXISTS test;

CREATE    TABLE test (test jsonb);

INSERT    INTO test (test)
VALUES    ('');

SELECT    *
FROM      queue_number;

SELECT    *
FROM      schedules;

SELECT    *
FROM      records;

SELECT    *
FROM      record_attachments;

SELECT    *
FROM      prescriptions;

SELECT    *
FROM      conversations;

SELECT    *
FROM      prescription_details;

SELECT    *
FROM      messages;

SELECT    *
FROM      medications;

SELECT    *
FROM      record_types;

DELETE    FROM prescription_details;

DELETE    FROM prescriptions;

DROP      DATABASE IF EXISTS merema;

DELETE    FROM record_attachments;