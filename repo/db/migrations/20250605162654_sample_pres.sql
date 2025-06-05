-- migrate:up
INSERT    INTO prescriptions (record_id, is_insurance_covered, prescription_note)
VALUES    (2000000000, TRUE, 'Sample prescription note for record 1');

INSERT    INTO prescription_details (
          prescription_id,
          med_id,
          morning_dosage,
          afternoon_dosage,
          evening_dosage,
          total_dosage,
          duration_days,
          dosage_unit,
          instructions
          )
VALUES    (3000000000, 1, 1.0, 1.0, 1.0, 9.0, 3, 'Viên', '1 viên một ngày nhá ku'),
          (3000000000, 4, 2.0, 1.0, 2.0, 9.0, 3, 'Viên', '2 viên sáng, tối; trưa 1 viên nhá cưng');

-- migrate:down
DELETE    FROM prescription_details
WHERE     prescription_id = 3000000000;

DELETE    FROM prescription_details
WHERE     prescription_id = 3000000000;