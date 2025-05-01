-- migrate:up
create    table "accounts" (
          "acc_id" serial primary KEY not null,
          "citizen_id" char(12) not null,
          "username" varchar(20) not null,
          "password_hash" varchar(20) not null,
          "phone" char(10) not null,
          "email" text,
          "role" varchar(10) not null,
          "created_at" timestamptz not null default (now ())
          );

create    table "staffs" (
          "staff_id" serial primary KEY not null,
          "acc_id" int not null,
          "full_name" text not null,
          "date_of_birth" date not null,
          "gender" varchar(3) not null,
          "department" text not null
          );

create    table "patients" (
          "patient_id" serial primary KEY not null,
          "acc_id" int not null,
          "full_name" varchar(50) not null,
          "date_of_birth" date not null,
          "gender" varchar(3) not null,
          "ethnicity" varchar(15) not null,
          "nationality" varchar(30),
          "address" text not null,
          "health_insurance_expired_date" date,
          "health_insurance_number" char(15),
          "emergency_contact_info" varchar not null
          );

create    table "records" (
          "record_id" serial primary KEY not null,
          "patient_id" int not null,
          "doctor_id" int not null,
          "type" text not null,
          "main_diagnosis" text,
          "secondary_diagnosis" text,
          "record_detail_path" text not null,
          "discharged_at" timestamptz,
          "created_at" timestamptz not null default (now ()),
          "expired_at" timestamptz not null
          );

create    table "prescriptions" (
          "prescription_id" serial primary KEY not null,
          "record_id" int not null,
          "is_insurance_covered" boolean not null,
          "prescription_note" text,
          "created_at" timestamptz not null default (now ()),
          "received_at" timestamptz
          );

create    table "prescription_details" (
          "detail_id" serial primary KEY not null,
          "prescription_id" serial not null,
          "medication_name" text not null,
          "dosage" int not null,
          "dosage_unit" varchar(20) not null,
          "duration_days" int not null,
          "quantity" int not null,
          "frequency" text not null,
          "instructions" text not null
          );

create    table "messages" (
          "from_acc_id" int not null,
          "to_acc_id" int not null,
          "content" text not null,
          "sent_at" timestamptz not null
          );

create    table "schedules" (
          "schedule_id" serial primary KEY not null,
          "patient_id" int not null,
          "queue_number" int not null,
          "examination_date" date not null,
          "expected_examination_time" timetz not null,
          "status" varchar(30) not null
          );

alter     table "staffs" ADD foreign KEY ("acc_id") references "accounts" ("acc_id");

alter     table "patients" ADD foreign KEY ("acc_id") references "accounts" ("acc_id");

alter     table "records" ADD foreign KEY ("patient_id") references "patients" ("patient_id");

alter     table "records" ADD foreign KEY ("doctor_id") references "staffs" ("staff_id");

alter     table "prescriptions" ADD foreign KEY ("record_id") references "records" ("record_id");

alter     table "prescription_details" ADD foreign KEY ("prescription_id") references "prescriptions" ("prescription_id");

alter     table "messages" ADD foreign KEY ("from_acc_id") references "accounts" ("acc_id");

alter     table "messages" ADD foreign KEY ("to_acc_id") references "accounts" ("acc_id");

alter     table "schedules" ADD foreign KEY ("patient_id") references "patients" ("patient_id");

-- migrate:down
drop      table if exists "schedules";

drop      table if exists "prescription_details";

drop      table if exists "prescriptions";

drop      table if exists "messages";

drop      table if exists "records";

drop      table if exists "patients";

drop      table if exists "staffs";

drop      table if exists "accounts";
