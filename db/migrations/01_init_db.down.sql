-- Drop tables with foreign key references first
DROP TABLE IF EXISTS "appointments";
DROP TABLE IF EXISTS "treatment_schedule_detail";
DROP TABLE IF EXISTS "examination_schedule_detail";
DROP TABLE IF EXISTS "schedules";
DROP TABLE IF EXISTS "services";
DROP TABLE IF EXISTS "bookings";
DROP TABLE IF EXISTS "dentist_detail";

-- Drop remaining tables
DROP TABLE IF EXISTS "service_categories";
DROP TABLE IF EXISTS "payments";
DROP TABLE IF EXISTS "rooms";
DROP TABLE IF EXISTS "specialties";
DROP TABLE IF EXISTS "users";
