-- Drop tables with foreign key references first
DROP TABLE IF EXISTS "treatment_schedules" CASCADE;
DROP TABLE IF EXISTS "examination_schedules" CASCADE;
DROP TABLE IF EXISTS "bookings" CASCADE;
DROP TABLE IF EXISTS "services" CASCADE;
DROP TABLE IF EXISTS "dentist_detail" CASCADE;

-- Drop remaining tables
DROP TABLE IF EXISTS "payments" CASCADE;
DROP TABLE IF EXISTS "users" CASCADE;
DROP TABLE IF EXISTS "specialties" CASCADE;
DROP TABLE IF EXISTS "rooms" CASCADE;
DROP TABLE IF EXISTS "service_categories" CASCADE;
