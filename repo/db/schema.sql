SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: accounts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.accounts (
    acc_id bigint NOT NULL,
    citizen_id character(12) NOT NULL,
    password_hash character(60) NOT NULL,
    phone character(10) NOT NULL,
    email text NOT NULL,
    role character varying(12) NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL
);


--
-- Name: accounts_acc_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.accounts_acc_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: accounts_acc_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.accounts_acc_id_seq OWNED BY public.accounts.acc_id;


--
-- Name: diagnoses; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.diagnoses (
    icd_code character varying(10) NOT NULL,
    name text NOT NULL,
    description text
);


--
-- Name: medications; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.medications (
    med_id bigint NOT NULL,
    name text NOT NULL,
    generic_name text,
    med_type character varying(50) NOT NULL,
    strength character varying(50),
    route_of_administration text NOT NULL,
    manufacturer text NOT NULL
);


--
-- Name: medications_med_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.medications_med_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: medications_med_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.medications_med_id_seq OWNED BY public.medications.med_id;


--
-- Name: messages; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.messages (
    from_acc_id bigint NOT NULL,
    to_acc_id bigint NOT NULL,
    sent_at timestamp with time zone DEFAULT now() NOT NULL,
    content text NOT NULL
);


--
-- Name: patients; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.patients (
    patient_id bigint NOT NULL,
    acc_id bigint NOT NULL,
    full_name text NOT NULL,
    date_of_birth date NOT NULL,
    gender character varying(3) NOT NULL,
    ethnicity text NOT NULL,
    nationality text NOT NULL,
    address text NOT NULL,
    health_insurance_expired_date date,
    health_insurance_number character(15),
    emergency_contact_info text NOT NULL
);


--
-- Name: patients_patient_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.patients_patient_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: patients_patient_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.patients_patient_id_seq OWNED BY public.patients.patient_id;


--
-- Name: prescription_details; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.prescription_details (
    detail_id bigint NOT NULL,
    prescription_id bigint NOT NULL,
    med_id bigint NOT NULL,
    morning_dosage numeric(10,2),
    afternoon_dosage numeric(10,2),
    evening_dosage numeric(10,2),
    duration_days integer NOT NULL,
    total_dosage numeric(10,2) NOT NULL,
    dosage_unit character varying(20) NOT NULL,
    instructions text NOT NULL
);


--
-- Name: prescription_details_detail_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.prescription_details_detail_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: prescription_details_detail_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.prescription_details_detail_id_seq OWNED BY public.prescription_details.detail_id;


--
-- Name: prescriptions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.prescriptions (
    prescription_id bigint NOT NULL,
    record_id bigint NOT NULL,
    is_insurance_covered boolean NOT NULL,
    prescription_note text,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    received_at timestamp with time zone
);


--
-- Name: prescriptions_prescription_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.prescriptions_prescription_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: prescriptions_prescription_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.prescriptions_prescription_id_seq OWNED BY public.prescriptions.prescription_id;


--
-- Name: queue_number; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.queue_number (
    date date NOT NULL,
    number integer NOT NULL
);


--
-- Name: record_attachments; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.record_attachments (
    attachment_id bigint NOT NULL,
    record_id bigint NOT NULL,
    type text NOT NULL,
    file_path text NOT NULL,
    uploaded_at timestamp with time zone DEFAULT now() NOT NULL
);


--
-- Name: record_attachments_attachment_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.record_attachments_attachment_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: record_attachments_attachment_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.record_attachments_attachment_id_seq OWNED BY public.record_attachments.attachment_id;


--
-- Name: record_types; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.record_types (
    type_id character(6) NOT NULL,
    type_name text NOT NULL,
    description text,
    template_path text NOT NULL,
    schema_path text NOT NULL
);


--
-- Name: records; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.records (
    record_id bigint NOT NULL,
    patient_id bigint NOT NULL,
    doctor_id bigint NOT NULL,
    type_id character(6) NOT NULL,
    primary_diagnosis character varying(10) NOT NULL,
    secondary_diagnosis character varying(10),
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    expired_at timestamp with time zone DEFAULT (now() + '10 years'::interval) NOT NULL,
    record_detail jsonb
);


--
-- Name: records_record_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.records_record_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: records_record_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.records_record_id_seq OWNED BY public.records.record_id;


--
-- Name: schedules; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schedules (
    schedule_id bigint NOT NULL,
    acc_id bigint NOT NULL,
    examination_date date NOT NULL,
    queue_number integer NOT NULL,
    type integer NOT NULL,
    expected_reception_time time with time zone NOT NULL,
    actual_reception_time time with time zone,
    status integer DEFAULT 1 NOT NULL
);


--
-- Name: schedules_schedule_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.schedules_schedule_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: schedules_schedule_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.schedules_schedule_id_seq OWNED BY public.schedules.schedule_id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying(128) NOT NULL
);


--
-- Name: staffs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.staffs (
    staff_id bigint NOT NULL,
    acc_id bigint NOT NULL,
    department text NOT NULL,
    full_name text NOT NULL,
    date_of_birth date NOT NULL,
    gender character varying(3) NOT NULL
);


--
-- Name: staffs_staff_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.staffs_staff_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: staffs_staff_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.staffs_staff_id_seq OWNED BY public.staffs.staff_id;


--
-- Name: accounts acc_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.accounts ALTER COLUMN acc_id SET DEFAULT nextval('public.accounts_acc_id_seq'::regclass);


--
-- Name: medications med_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.medications ALTER COLUMN med_id SET DEFAULT nextval('public.medications_med_id_seq'::regclass);


--
-- Name: patients patient_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patients ALTER COLUMN patient_id SET DEFAULT nextval('public.patients_patient_id_seq'::regclass);


--
-- Name: prescription_details detail_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.prescription_details ALTER COLUMN detail_id SET DEFAULT nextval('public.prescription_details_detail_id_seq'::regclass);


--
-- Name: prescriptions prescription_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.prescriptions ALTER COLUMN prescription_id SET DEFAULT nextval('public.prescriptions_prescription_id_seq'::regclass);


--
-- Name: record_attachments attachment_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.record_attachments ALTER COLUMN attachment_id SET DEFAULT nextval('public.record_attachments_attachment_id_seq'::regclass);


--
-- Name: records record_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.records ALTER COLUMN record_id SET DEFAULT nextval('public.records_record_id_seq'::regclass);


--
-- Name: schedules schedule_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schedules ALTER COLUMN schedule_id SET DEFAULT nextval('public.schedules_schedule_id_seq'::regclass);


--
-- Name: staffs staff_id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.staffs ALTER COLUMN staff_id SET DEFAULT nextval('public.staffs_staff_id_seq'::regclass);


--
-- Name: accounts accounts_citizen_id_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_citizen_id_key UNIQUE (citizen_id);


--
-- Name: accounts accounts_email_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_email_key UNIQUE (email);


--
-- Name: accounts accounts_phone_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_phone_key UNIQUE (phone);


--
-- Name: accounts accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (acc_id);


--
-- Name: diagnoses diagnoses_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.diagnoses
    ADD CONSTRAINT diagnoses_pkey PRIMARY KEY (icd_code);


--
-- Name: medications medications_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.medications
    ADD CONSTRAINT medications_pkey PRIMARY KEY (med_id);


--
-- Name: messages messages_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT messages_pkey PRIMARY KEY (from_acc_id, to_acc_id, sent_at);


--
-- Name: patients patients_health_insurance_number_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patients
    ADD CONSTRAINT patients_health_insurance_number_key UNIQUE (health_insurance_number);


--
-- Name: patients patients_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patients
    ADD CONSTRAINT patients_pkey PRIMARY KEY (patient_id);


--
-- Name: prescription_details prescription_details_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.prescription_details
    ADD CONSTRAINT prescription_details_pkey PRIMARY KEY (detail_id);


--
-- Name: prescriptions prescriptions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.prescriptions
    ADD CONSTRAINT prescriptions_pkey PRIMARY KEY (prescription_id);


--
-- Name: prescriptions prescriptions_record_id_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.prescriptions
    ADD CONSTRAINT prescriptions_record_id_key UNIQUE (record_id);


--
-- Name: queue_number queue_number_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.queue_number
    ADD CONSTRAINT queue_number_pkey PRIMARY KEY (date);


--
-- Name: record_attachments record_attachments_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.record_attachments
    ADD CONSTRAINT record_attachments_pkey PRIMARY KEY (attachment_id);


--
-- Name: record_types record_types_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.record_types
    ADD CONSTRAINT record_types_pkey PRIMARY KEY (type_id);


--
-- Name: records records_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.records
    ADD CONSTRAINT records_pkey PRIMARY KEY (record_id);


--
-- Name: schedules schedules_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schedules
    ADD CONSTRAINT schedules_pkey PRIMARY KEY (schedule_id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: staffs staffs_acc_id_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.staffs
    ADD CONSTRAINT staffs_acc_id_key UNIQUE (acc_id);


--
-- Name: staffs staffs_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.staffs
    ADD CONSTRAINT staffs_pkey PRIMARY KEY (staff_id);


--
-- Name: patients_acc_id_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX patients_acc_id_idx ON public.patients USING btree (acc_id);


--
-- Name: prescription_details_med_id_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX prescription_details_med_id_idx ON public.prescription_details USING btree (med_id);


--
-- Name: prescription_details_prescription_id_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX prescription_details_prescription_id_idx ON public.prescription_details USING btree (prescription_id);


--
-- Name: record_attachments_record_id_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX record_attachments_record_id_idx ON public.record_attachments USING btree (record_id);


--
-- Name: records_doctor_id_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX records_doctor_id_idx ON public.records USING btree (doctor_id);


--
-- Name: records_patient_id_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX records_patient_id_idx ON public.records USING btree (patient_id);


--
-- Name: records_primary_diagnosis_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX records_primary_diagnosis_idx ON public.records USING btree (primary_diagnosis);


--
-- Name: records_secondary_diagnosis_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX records_secondary_diagnosis_idx ON public.records USING btree (secondary_diagnosis);


--
-- Name: records_type_id_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX records_type_id_idx ON public.records USING btree (type_id);


--
-- Name: schedules_acc_id_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX schedules_acc_id_idx ON public.schedules USING btree (acc_id);


--
-- Name: schedules_type_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX schedules_type_idx ON public.schedules USING btree (type);


--
-- Name: messages messages_from_acc_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT messages_from_acc_id_fkey FOREIGN KEY (from_acc_id) REFERENCES public.accounts(acc_id);


--
-- Name: messages messages_to_acc_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.messages
    ADD CONSTRAINT messages_to_acc_id_fkey FOREIGN KEY (to_acc_id) REFERENCES public.accounts(acc_id);


--
-- Name: patients patients_acc_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.patients
    ADD CONSTRAINT patients_acc_id_fkey FOREIGN KEY (acc_id) REFERENCES public.accounts(acc_id);


--
-- Name: prescription_details prescription_details_med_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.prescription_details
    ADD CONSTRAINT prescription_details_med_id_fkey FOREIGN KEY (med_id) REFERENCES public.medications(med_id);


--
-- Name: prescription_details prescription_details_prescription_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.prescription_details
    ADD CONSTRAINT prescription_details_prescription_id_fkey FOREIGN KEY (prescription_id) REFERENCES public.prescriptions(prescription_id);


--
-- Name: prescriptions prescriptions_record_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.prescriptions
    ADD CONSTRAINT prescriptions_record_id_fkey FOREIGN KEY (record_id) REFERENCES public.records(record_id);


--
-- Name: record_attachments record_attachments_record_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.record_attachments
    ADD CONSTRAINT record_attachments_record_id_fkey FOREIGN KEY (record_id) REFERENCES public.records(record_id);


--
-- Name: records records_doctor_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.records
    ADD CONSTRAINT records_doctor_id_fkey FOREIGN KEY (doctor_id) REFERENCES public.staffs(staff_id);


--
-- Name: records records_patient_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.records
    ADD CONSTRAINT records_patient_id_fkey FOREIGN KEY (patient_id) REFERENCES public.patients(patient_id);


--
-- Name: records records_primary_diagnosis_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.records
    ADD CONSTRAINT records_primary_diagnosis_fkey FOREIGN KEY (primary_diagnosis) REFERENCES public.diagnoses(icd_code);


--
-- Name: records records_secondary_diagnosis_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.records
    ADD CONSTRAINT records_secondary_diagnosis_fkey FOREIGN KEY (secondary_diagnosis) REFERENCES public.diagnoses(icd_code);


--
-- Name: records records_type_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.records
    ADD CONSTRAINT records_type_id_fkey FOREIGN KEY (type_id) REFERENCES public.record_types(type_id);


--
-- Name: schedules schedules_acc_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schedules
    ADD CONSTRAINT schedules_acc_id_fkey FOREIGN KEY (acc_id) REFERENCES public.accounts(acc_id);


--
-- Name: staffs staffs_acc_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.staffs
    ADD CONSTRAINT staffs_acc_id_fkey FOREIGN KEY (acc_id) REFERENCES public.accounts(acc_id);


--
-- PostgreSQL database dump complete
--


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20250501114344'),
    ('20250502100638'),
    ('20250506063225'),
    ('20250514151657'),
    ('20250516100948'),
    ('20250526090415'),
    ('20250605162654');
