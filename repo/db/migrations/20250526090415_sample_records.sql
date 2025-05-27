-- migrate:up
INSERT    INTO records (
          patient_id,
          doctor_id,
          type,
          primary_diagnosis,
          secondary_diagnosis,
          created_at,
          expired_at
          )
VALUES    (
          1100000000,
          1200000001,
          'Bệnh án nội khoa',
          'M47.9',
          'T78.4',
          '2025-05-26 09:04:15',
          '2025-06-26 09:04:15'
          );

-- migrate:down
DELETE    FROM records
WHERE     patient_id = 110000000
AND       doctor_id = 1200000001;