-- migrate:up
INSERT    INTO prescriptions (record_id, is_insurance_covered, prescription_note)
VALUES    (2000000000, TRUE, 'Sample prescription note for record 1');

INSERT    INTO prescription_details (
          prescription_id,
          med_id,
          morning_dosage,
          afternoon_dosage,
          evening_dosage,
          duration_days,
          total_dosage,
          dosage_unit,
          instructions
          )
VALUES    (
          (
          SELECT    prescription_id
          FROM      prescriptions
          WHERE     record_id = 2000000000
          ),
          5000000000,
          1.0,
          1.0,
          1.0,
          3,
          9.0,
          'Viên',
          '1 viên một ngày nhá ku'
          ),
          (
          (
          SELECT    prescription_id
          FROM      prescriptions
          WHERE     record_id = 2000000000
          ),
          5000000003,
          2.0,
          1.0,
          2.0,
          2,
          10.0,
          'Viên',
          '2 viên sáng, tối; trưa 1 viên nhá cưng'
          );

-- migrate:down
DELETE    FROM prescription_details
WHERE     prescription_id = (
          SELECT    prescription_id
          FROM      prescriptions
          WHERE     record_id = 2000000000
          );

DELETE    FROM prescriptions
WHERE     prescription_id = (
          SELECT    prescription_id
          FROM      prescriptions
          WHERE     record_id = 2000000000
          );