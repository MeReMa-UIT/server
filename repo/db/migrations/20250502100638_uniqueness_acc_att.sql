-- migrate:up
ALTER     TABLE accounts ADD CONSTRAINT unique_citizen_id UNIQUE (citizen_id);

ALTER     TABLE accounts ADD CONSTRAINT unique_email UNIQUE (email);

ALTER     TABLE accounts ADD CONSTRAINT unique_phone UNIQUE (phone);

ALTER     SEQUENCE accounts_acc_id_seq
RESTART   WITH 1000000000;

ALTER     SEQUENCE patients_patient_id_seq
RESTART   WITH 1100000000;

ALTER     SEQUENCE staffs_staff_id_seq
RESTART   WITH 1200000000;

ALTER     SEQUENCE records_record_id_seq
RESTART   WITH 2000000000;

ALTER     SEQUENCE prescriptions_prescription_id_seq
RESTART   WITH 3000000000;

ALTER     SEQUENCE schedules_schedule_id_seq
RESTART   WITH 4000000000;

-- migrate:down
ALTER     TABLE accounts
DROP      CONSTRAINT IF EXISTS unique_citizen_id;

ALTER     TABLE accounts
DROP      CONSTRAINT IF EXISTS unique_email;

ALTER     TABLE accounts
DROP      CONSTRAINT IF EXISTS unique_phone;

ALTER     SEQUENCE accounts_acc_id_seq
RESTART   WITH 1;

ALTER     SEQUENCE patients_patient_id_seq
RESTART   WITH 1;

ALTER     SEQUENCE staffs_staff_id_seq
RESTART   WITH 1;

ALTER     SEQUENCE records_record_id_seq
RESTART   WITH 1;

ALTER     SEQUENCE prescriptions_prescription_id_seq
RESTART   WITH 1;

ALTER     SEQUENCE schedules_schedule_id_seq
RESTART   WITH 1;