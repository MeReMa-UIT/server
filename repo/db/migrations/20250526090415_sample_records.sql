-- migrate:up
INSERT    INTO record_types (type_id, type_name, template_path, schema_path)
VALUES    (
          '01/BV1',
          'BỆNH ÁN NỘI KHOA',
          './templates/01_BV1/01_BV1.template.json',
          './templates/01_BV1/01_BV1.schema.json'
          );

INSERT    INTO records (patient_id, doctor_id, type_id, primary_diagnosis, secondary_diagnosis, created_at, expired_at)
VALUES    (1100000000, 1200000001, '01/BV1', 'M47.9', 'T78.4', '2025-05-26 09:04:15', '2025-06-26 09:04:15');

INSERT    INTO records (patient_id, doctor_id, type_id, primary_diagnosis, secondary_diagnosis, created_at, expired_at)
VALUES    (1100000001, 1200000001, '01/BV1', 'M47.9', 'T78.4', '2025-05-26 09:04:15', '2025-06-26 09:04:15');

-- migrate:down
DELETE    FROM records
WHERE     doctor_id = 1200000001;

DELETE    FROM record_types
WHERE     type_id = '01/BV1';