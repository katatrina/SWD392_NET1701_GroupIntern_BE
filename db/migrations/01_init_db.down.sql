-- Drop dependent tables first
DROP TABLE IF EXISTS "dentist_categories";
DROP TABLE IF EXISTS "room_categories";
DROP TABLE IF EXISTS "treatment_schedules";
DROP TABLE IF EXISTS "examination_schedules";
DROP TABLE IF EXISTS "bookings";
DROP TABLE IF EXISTS "services";

-- Drop remaining tables
DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "dentist_detail";
DROP TABLE IF EXISTS "rooms";
DROP TABLE IF EXISTS "service_categories";
DROP TABLE IF EXISTS "payments";
