-- migrate:up
CREATE    TABLE "accounts" (
          "acc_id" bigserial PRIMARY KEY NOT NULL,
          "citizen_id" CHAR(12) UNIQUE NOT NULL,
          "password_hash" CHAR(60) NOT NULL,
          "phone" CHAR(10) UNIQUE NOT NULL,
          "email" text UNIQUE NOT NULL,
          "role" VARCHAR(12) NOT NULL,
          "created_at" timestamptz NOT NULL DEFAULT (now ())
          );

CREATE    TABLE "staffs" (
          "staff_id" bigserial PRIMARY KEY NOT NULL,
          "acc_id" BIGINT UNIQUE NOT NULL,
          "department" text NOT NULL,
          "full_name" text NOT NULL,
          "date_of_birth" DATE NOT NULL,
          "gender" VARCHAR(3) NOT NULL
          );

CREATE    TABLE "patients" (
          "patient_id" bigserial PRIMARY KEY NOT NULL,
          "acc_id" BIGINT NOT NULL,
          "full_name" text NOT NULL,
          "date_of_birth" DATE NOT NULL,
          "gender" VARCHAR(3) NOT NULL,
          "ethnicity" text NOT NULL,
          "nationality" text NOT NULL,
          "address" text NOT NULL,
          "health_insurance_expired_date" DATE,
          "health_insurance_number" CHAR(15) UNIQUE,
          "emergency_contact_info" text NOT NULL
          );

CREATE    TABLE "records" (
          "record_id" bigserial PRIMARY KEY NOT NULL,
          "patient_id" BIGINT NOT NULL,
          "doctor_id" BIGINT NOT NULL,
          "type_id" CHAR(6) NOT NULL,
          "primary_diagnosis" VARCHAR(10),
          "secondary_diagnosis" VARCHAR(10),
          "created_at" timestamptz NOT NULL DEFAULT (now ()),
          "expired_at" timestamptz NOT NULL,
          "record_detail" jsonb
          );

CREATE    TABLE "record_types" (
          "type_id" CHAR(6) PRIMARY KEY NOT NULL,
          "type_name" text NOT NULL,
          "description" text,
          "template_path" text NOT NULL,
          "schema_path" text NOT NULL
          );

CREATE    TABLE "diagnoses" ("icd_code" VARCHAR(10) PRIMARY KEY NOT NULL, "name" text NOT NULL, "description" text);

CREATE    TABLE "record_attachments" (
          "attachment_id" bigserial PRIMARY KEY NOT NULL,
          "record_id" BIGINT NOT NULL,
          "type" text NOT NULL,
          "file_path" text NOT NULL,
          "uploaded_at" timestamptz NOT NULL DEFAULT (now ())
          );

CREATE    TABLE "prescriptions" (
          "prescription_id" bigserial PRIMARY KEY NOT NULL,
          "record_id" BIGINT UNIQUE NOT NULL,
          "is_insurance_covered" BOOLEAN NOT NULL,
          "prescription_note" text,
          "created_at" timestamptz NOT NULL DEFAULT (now ()),
          "received_at" timestamptz
          );

CREATE    TABLE "prescription_details" (
          "detail_id" bigserial PRIMARY KEY NOT NULL,
          "prescription_id" BIGINT NOT NULL,
          "med_id" BIGINT NOT NULL,
          "morning_dosage" DECIMAL(10, 2),
          "afternoon_dosage" DECIMAL(10, 2),
          "evening_dosage" DECIMAL(10, 2),
          "duration_days" INT NOT NULL,
          "total_dosage" DECIMAL(10, 2) NOT NULL,
          "dosage_unit" VARCHAR(20) NOT NULL,
          "instructions" text NOT NULL
          );

CREATE    TABLE "medications" (
          "med_id" bigserial PRIMARY KEY NOT NULL,
          "name" text NOT NULL,
          "generic_name" text,
          "med_type" VARCHAR(50) NOT NULL,
          "strength" VARCHAR(50),
          "route_of_administration" text NOT NULL,
          "manufacturer" text NOT NULL
          );

CREATE    TABLE "messages" (
          "from_acc_id" BIGINT NOT NULL,
          "to_acc_id" BIGINT NOT NULL,
          "sent_at" timestamptz NOT NULL DEFAULT (now ()),
          "content" text NOT NULL,
          PRIMARY KEY ("from_acc_id", "to_acc_id", "sent_at")
          );

CREATE    TABLE "schedules" (
          "schedule_id" bigserial PRIMARY KEY NOT NULL,
          "acc_id" BIGINT NOT NULL,
          "examination_date" DATE NOT NULL,
          "queue_number" INT NOT NULL,
          "type" INT NOT NULL,
          "expected_reception_time" timetz NOT NULL,
          "actual_reception_time" timetz,
          "status" INT NOT NULL DEFAULT 1
          );

CREATE INDEX ON "patients" ("acc_id");

CREATE INDEX ON "records" ("patient_id");

CREATE INDEX ON "records" ("doctor_id");

CREATE INDEX ON "records" ("type_id");

CREATE INDEX ON "records" ("primary_diagnosis");

CREATE INDEX ON "records" ("secondary_diagnosis");

CREATE INDEX ON "record_attachments" ("record_id");

CREATE INDEX ON "prescription_details" ("prescription_id");

CREATE INDEX ON "prescription_details" ("med_id");

CREATE INDEX ON "schedules" ("acc_id");

CREATE INDEX ON "schedules" ("type");

ALTER     TABLE "staffs" ADD FOREIGN KEY ("acc_id") REFERENCES "accounts" ("acc_id");

ALTER     TABLE "patients" ADD FOREIGN KEY ("acc_id") REFERENCES "accounts" ("acc_id");

ALTER     TABLE "records" ADD FOREIGN KEY ("patient_id") REFERENCES "patients" ("patient_id");

ALTER     TABLE "records" ADD FOREIGN KEY ("doctor_id") REFERENCES "staffs" ("staff_id");

ALTER     TABLE "records" ADD FOREIGN KEY ("type_id") REFERENCES "record_types" ("type_id");

ALTER     TABLE "records" ADD FOREIGN KEY ("primary_diagnosis") REFERENCES "diagnoses" ("icd_code");

ALTER     TABLE "records" ADD FOREIGN KEY ("secondary_diagnosis") REFERENCES "diagnoses" ("icd_code");

ALTER     TABLE "record_attachments" ADD FOREIGN KEY ("record_id") REFERENCES "records" ("record_id");

ALTER     TABLE "prescriptions" ADD FOREIGN KEY ("record_id") REFERENCES "records" ("record_id");

ALTER     TABLE "prescription_details" ADD FOREIGN KEY ("prescription_id") REFERENCES "prescriptions" ("prescription_id");

ALTER     TABLE "prescription_details" ADD FOREIGN KEY ("med_id") REFERENCES "medications" ("med_id");

ALTER     TABLE "messages" ADD FOREIGN KEY ("from_acc_id") REFERENCES "accounts" ("acc_id");

ALTER     TABLE "messages" ADD FOREIGN KEY ("to_acc_id") REFERENCES "accounts" ("acc_id");

ALTER     TABLE "schedules" ADD FOREIGN KEY ("acc_id") REFERENCES "accounts" ("acc_id");

-- migrate:down
DROP      TABLE if EXISTS "schedules";

DROP      TABLE if EXISTS "prescription_details";

DROP      TABLE if EXISTS "medications";

DROP      TABLE if EXISTS "prescriptions";

DROP      TABLE if EXISTS "messages";

DROP      TABLE if EXISTS "record_attachments";

DROP      TABLE if EXISTS "records";

DROP      TABLE if EXISTS "diagnoses";

DROP      TABLE if EXISTS "patients";

DROP      TABLE if EXISTS "staffs";

DROP      TABLE if EXISTS "accounts";