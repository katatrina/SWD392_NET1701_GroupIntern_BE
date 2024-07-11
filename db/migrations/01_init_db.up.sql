CREATE TABLE "users"
(
    "id"              bigserial PRIMARY KEY,
    "full_name"       text        NOT NULL,
    "hashed_password" text        NOT NULL,
    "email"           text UNIQUE NOT NULL,
    "phone_number"    text UNIQUE NOT NULL,
    "date_of_birth"   DATE        NOT NULL,
    "gender"          text        NOT NULL,
    "role"            text        NOT NULL,
    "deleted_at"      timestamptz,
    "created_at"      timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "dentist_detail"
(
    "dentist_id"   bigint PRIMARY KEY,
    "specialty_id" bigint NOT NULL
);

CREATE TABLE "specialties"
(
    "id"         bigserial PRIMARY KEY,
    "name"       text        NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "rooms"
(
    "id"         bigserial PRIMARY KEY,
    "name"       text        NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "service_categories"
(
    "id"          bigserial PRIMARY KEY,
    "name"        text        NOT NULL,
    "icon_url"    text        NOT NULL,
    "banner_url"  text        NOT NULL,
    "description" text        NOT NULL,
    "slug"        text        NOT NULL,
    "created_at"  timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "services"
(
    "id"                bigserial PRIMARY KEY,
    "name"              text        NOT NULL,
    "category_id"       bigint      NOT NULL,
    "unit"              text        NOT NULL,
    "cost"              bigint      NOT NULL,
    "warranty_duration" text        NOT NULL,
    "created_at"        timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "schedules"
(
    "id"              bigserial PRIMARY KEY,
    "type"            text        NOT NULL,
    "start_time"      timestamptz NOT NULL,
    "end_time"        timestamptz NOT NULL,
    "dentist_id"      bigint      NOT NULL,
    "room_id"         bigint      NOT NULL,
    "slots_remaining" bigint      NOT NULL,
    "created_at"      timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "appointments"
(
    "id"          bigserial PRIMARY KEY,
    "booking_id"  bigint      NOT NULL,
    "schedule_id" bigint      NOT NULL,
    "patient_id"  bigint      NOT NULL,
    "status"      text        NOT NULL DEFAULT 'Đang chờ',
    "created_at"  timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "examination_appointment_detail"
(
    "appointment_id"      bigint PRIMARY KEY,
    "service_category_id" bigint,
    "created_at"          timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "treatment_appointment_detail"
(
    "appointment_id"   bigint PRIMARY KEY,
    "service_id"       bigint      NOT NULL,
    "service_quantity" bigint      NOT NULL,
    "created_at"       timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "bookings"
(
    "id"               bigserial PRIMARY KEY,
    "patient_id"       bigint      NOT NULL,
    "type"             text        NOT NULL,
    "payment_status"   text        NOT NULL DEFAULT 'Chưa thanh toán',
    "payment_id"       bigint,
    "total_cost"       bigint      NOT NULL DEFAULT 0,
    "appointment_date" DATE        NOT NULL,
    "status"           text        NOT NULL DEFAULT 'Đang chờ',
    "created_at"       timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "payments"
(
    "id"         bigserial PRIMARY KEY,
    "name"       text        NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "dentist_detail"
    ADD FOREIGN KEY ("specialty_id") REFERENCES "specialties" ("id");

ALTER TABLE "dentist_detail"
    ADD FOREIGN KEY ("dentist_id") REFERENCES "users" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "services"
    ADD FOREIGN KEY ("category_id") REFERENCES "service_categories" ("id") ON DELETE NO ACTION ON UPDATE CASCADE;

ALTER TABLE "schedules"
    ADD FOREIGN KEY ("dentist_id") REFERENCES "users" ("id");

ALTER TABLE "schedules"
    ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id") ON DELETE NO ACTION ON UPDATE CASCADE;

ALTER TABLE "appointments"
    ADD FOREIGN KEY ("booking_id") REFERENCES "bookings" ("id");

ALTER TABLE "appointments"
    ADD FOREIGN KEY ("schedule_id") REFERENCES "schedules" ("id");

ALTER TABLE "appointments"
    ADD FOREIGN KEY ("patient_id") REFERENCES "users" ("id");

ALTER TABLE "examination_appointment_detail"
    ADD FOREIGN KEY ("service_category_id") REFERENCES "service_categories" ("id");

ALTER TABLE "examination_appointment_detail"
    ADD FOREIGN KEY ("appointment_id") REFERENCES "appointments" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "treatment_appointment_detail"
    ADD FOREIGN KEY ("service_id") REFERENCES "services" ("id");

ALTER TABLE "treatment_appointment_detail"
    ADD FOREIGN KEY ("appointment_id") REFERENCES "appointments" ("id") ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "bookings"
    ADD FOREIGN KEY ("patient_id") REFERENCES "users" ("id");

ALTER TABLE "bookings"
    ADD FOREIGN KEY ("payment_id") REFERENCES "payments" ("id");
