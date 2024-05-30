CREATE TABLE "users"
(
    "id"              bigserial PRIMARY KEY,
    "full_name"       text        NOT NULL,
    "hashed_password" text        NOT NULL,
    "email"           text UNIQUE NOT NULL,
    "phone_number"    text UNIQUE NOT NULL,
    "role"            text        NOT NULL,
    "created_at"      timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "dentist_detail"
(
    "dentist_id"    bigint PRIMARY KEY,
    "date_of_birth" timestamptz NOT NULL,
    "sex"           text        NOT NULL,
    "specialty_id"  bigint      NOT NULL
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
    "id"         bigserial PRIMARY KEY,
    "name"       text        NOT NULL,
    "image_url"  text        NOT NULL,
    "slug"       text        NOT NULL,
    "price"      bigint      NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "services"
(
    "id"          bigserial PRIMARY KEY,
    "category_id" bigint      NOT NULL,
    "unit"        text        NOT NULL,
    "price"       bigint      NOT NULL,
    "warranty_duration" interval NOT NULL,
    "created_at"  timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "treatment_schedules"
(
    "id"               bigserial PRIMARY KEY,
    "booking_id"       bigint      NOT NULL,
    "start_time"       timestamptz NOT NULL,
    "end_time"         timestamptz NOT NULL,
    "customer_id"      bigint      NOT NULL,
    "dentist_id"       bigint      NOT NULL,
    "service_id"       bigint      NOT NULL,
    "service_quantity" bigint      NOT NULL,
    "room_id"          bigint      NOT NULL,
    "slot"             bigint      NOT NULL,
    "status"           text        NOT NULL DEFAULT 'Đang chờ',
    "created_at"       timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "examination_schedules"
(
    "id"                  bigserial PRIMARY KEY,
    "booking_id"          bigint,
    "start_time"          timestamptz NOT NULL,
    "end_time"            timestamptz NOT NULL,
    "customer_id"         bigint,
    "dentist_id"          bigint      NOT NULL,
    "service_category_id" bigint      NOT NULL,
    "room_id"             bigint      NOT NULL,
    "slot"                bigint      NOT NULL,
    "status"              text        NOT NULL DEFAULT '',
    "created_at"          timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "bookings"
(
    "id"              bigserial PRIMARY KEY,
    "type"            text        NOT NULL,
    "customer_id"     bigint      NOT NULL,
    "customer_reason" text        NOT NULL DEFAULT '',
    "payment_status"  text        NOT NULL DEFAULT 'Chưa thanh toán',
    "payment_id"      bigint      NOT NULL,
    "is_cancelled"    bool        NOT NULL DEFAULT false,
    "created_at"      timestamptz NOT NULL DEFAULT (now())
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
    ADD FOREIGN KEY ("category_id") REFERENCES "service_categories" ("id");

ALTER TABLE "treatment_schedules"
    ADD FOREIGN KEY ("booking_id") REFERENCES "bookings" ("id");

ALTER TABLE "treatment_schedules"
    ADD FOREIGN KEY ("customer_id") REFERENCES "users" ("id");

ALTER TABLE "treatment_schedules"
    ADD FOREIGN KEY ("dentist_id") REFERENCES "users" ("id");

ALTER TABLE "treatment_schedules"
    ADD FOREIGN KEY ("service_id") REFERENCES "services" ("id");

ALTER TABLE "treatment_schedules"
    ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");

ALTER TABLE "examination_schedules"
    ADD FOREIGN KEY ("booking_id") REFERENCES "bookings" ("id");

ALTER TABLE "examination_schedules"
    ADD FOREIGN KEY ("customer_id") REFERENCES "users" ("id");

ALTER TABLE "examination_schedules"
    ADD FOREIGN KEY ("dentist_id") REFERENCES "users" ("id");

ALTER TABLE "examination_schedules"
    ADD FOREIGN KEY ("service_category_id") REFERENCES "service_categories" ("id");

ALTER TABLE "examination_schedules"
    ADD FOREIGN KEY ("room_id") REFERENCES "rooms" ("id");

ALTER TABLE "bookings"
    ADD FOREIGN KEY ("customer_id") REFERENCES "users" ("id");

ALTER TABLE "bookings"
    ADD FOREIGN KEY ("payment_id") REFERENCES "payments" ("id");
