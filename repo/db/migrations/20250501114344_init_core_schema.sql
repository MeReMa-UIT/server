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
          "acc_id" BIGINT NOT NULL,
          "full_name" text NOT NULL,
          "date_of_birth" DATE NOT NULL,
          "gender" VARCHAR(3) NOT NULL,
          "department" text NOT NULL
          );

CREATE    TABLE "patients" (
          "patient_id" bigserial PRIMARY KEY NOT NULL,
          "acc_id" BIGINT NOT NULL,
          "full_name" VARCHAR(50) NOT NULL,
          "date_of_birth" DATE NOT NULL,
          "gender" VARCHAR(3) NOT NULL,
          "ethnicity" VARCHAR(15) NOT NULL,
          "nationality" VARCHAR(30),
          "address" text NOT NULL,
          "health_insurance_expired_date" DATE,
          "health_insurance_number" CHAR(15),
          "emergency_contact_info" VARCHAR NOT NULL
          );

CREATE    TABLE "records" (
          "record_id" bigserial PRIMARY KEY NOT NULL,
          "patient_id" BIGINT NOT NULL,
          "doctor_id" BIGINT NOT NULL,
          "type" text NOT NULL,
          "main_diagnosis" text,
          "secondary_diagnosis" text,
          "record_detail_path" text NOT NULL,
          "discharged_at" timestamptz,
          "created_at" timestamptz NOT NULL DEFAULT (now ()),
          "expired_at" timestamptz NOT NULL
          );

CREATE    TABLE "prescriptions" (
          "prescription_id" bigserial PRIMARY KEY NOT NULL,
          "record_id" BIGINT NOT NULL,
          "is_insurance_covered" BOOLEAN NOT NULL,
          "prescription_note" text,
          "created_at" timestamptz NOT NULL DEFAULT (now ()),
          "received_at" timestamptz
          );

CREATE    TABLE "prescription_details" (
          "detail_id" bigserial PRIMARY KEY NOT NULL,
          "prescription_id" BIGINT NOT NULL,
          "medication_name" text NOT NULL,
          "dosage" INT NOT NULL,
          "dosage_unit" VARCHAR(20) NOT NULL,
          "duration_days" INT NOT NULL,
          "quantity" INT NOT NULL,
          "frequency" text NOT NULL,
          "instructions" text NOT NULL
          );

CREATE    TABLE "messages" (
          "from_acc_id" BIGINT NOT NULL,
          "to_acc_id" BIGINT NOT NULL,
          "sent_at" timestamptz NOT NULL,
          "content" text NOT NULL,
          PRIMARY KEY ("from_acc_id", "to_acc_id", "sent_at")
          );

CREATE    TABLE "schedules" (
          "schedule_id" bigserial PRIMARY KEY NOT NULL,
          "patient_id" BIGINT NOT NULL,
          "queue_number" INT NOT NULL,
          "examination_date" DATE NOT NULL,
          "expected_examination_time" timetz NOT NULL,
          "status" VARCHAR(30) NOT NULL
          );

ALTER     TABLE "staffs" ADD FOREIGN KEY ("acc_id") REFERENCES "accounts" ("acc_id");

ALTER     TABLE "patients" ADD FOREIGN KEY ("acc_id") REFERENCES "accounts" ("acc_id");

ALTER     TABLE "records" ADD FOREIGN KEY ("patient_id") REFERENCES "patients" ("patient_id");

ALTER     TABLE "records" ADD FOREIGN KEY ("doctor_id") REFERENCES "staffs" ("staff_id");

ALTER     TABLE "prescriptions" ADD FOREIGN KEY ("record_id") REFERENCES "records" ("record_id");

ALTER     TABLE "prescription_details" ADD FOREIGN KEY ("prescription_id") REFERENCES "prescriptions" ("prescription_id");

ALTER     TABLE "messages" ADD FOREIGN KEY ("from_acc_id") REFERENCES "accounts" ("acc_id");

ALTER     TABLE "messages" ADD FOREIGN KEY ("to_acc_id") REFERENCES "accounts" ("acc_id");

ALTER     TABLE "schedules" ADD FOREIGN KEY ("patient_id") REFERENCES "patients" ("patient_id");

-- migrate:down
DROP      TABLE if EXISTS "schedules";

DROP      TABLE if EXISTS "prescription_details";

DROP      TABLE if EXISTS "prescriptions";

DROP      TABLE if EXISTS "messages";

DROP      TABLE if EXISTS "records";

DROP      TABLE if EXISTS "patients";

DROP      TABLE if EXISTS "staffs";

DROP      TABLE if EXISTS "accounts";