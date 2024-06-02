-- Insert specialties
INSERT INTO specialties (name)
VALUES ('Trợ thủ nha khoa'),
       ('Nha chu'),
       ('Nội nha'),
       ('Phục hình răng'),
       ('Nhổ răng - Tiểu phẩu'),
       ('Điều dưỡng nha khoa'),
       ('Nắn chỉnh răng');

-- Insert rooms
INSERT INTO rooms (name)
VALUES ('01'),
       ('02'),
       ('03'),
       ('04'),
       ('05'),
       ('06'),
       ('07');

-- Insert service categories
INSERT INTO service_categories (name, image_url, slug, price)
VALUES ('Bọc răng sứ', '', 'boc-rang-su', 500000),
       ('Cấy ghép Implant', '', 'cay-ghep-implant', 600000),
       ('Niềng răng thẩm mỹ', '', 'nieng-rang-tham-my', 700000),
       ('Tẩy trắng răng', '', 'tay-trang-rang', 800000),
       ('Nhổ răng khôn', '', 'nho-rang-khon', 900000),
       ('Bệnh lý nha chu', '', 'benh-ly-nha-chu', 1000000),
       ('Điều trị tủy răng', '', 'dieu-tri-tuy-rang', 1100000);

-- Insert payments
INSERT INTO payments (name)
VALUES ('Tiền mặt');

-- Insert dentists
INSERT INTO users (full_name, hashed_password, email, phone_number, role) -- id 1
VALUES ('Nguyễn Anh Dũng', '123', 'dungan@gmail.com', '0347836802', 'dentist');
INSERT INTO dentist_detail (dentist_id, sex, date_of_birth, specialty_id)
VALUES (1, 'male', '1992-11-22 00:00:00+07', 1);

INSERT INTO users (full_name, hashed_password, email, phone_number, role) -- id 2
VALUES ('Trần Văn Lâm', '123', 'lamtv@gmail.com', '0333686702', 'dentist');
INSERT INTO dentist_detail (dentist_id, sex, date_of_birth, specialty_id)
VALUES (2, 'male', '1992-11-22 00:00:00+07', 2);

INSERT INTO users (full_name, hashed_password, email, phone_number, role) -- id 2
VALUES ('Nguyễn Ánh Vy', '123', 'anhvy@gmail.com', '0834042822', 'dentist');
INSERT INTO dentist_detail (dentist_id, sex, date_of_birth, specialty_id)
VALUES (3, 'female', '1978-03-09 00:00:00+07', 4);

INSERT INTO users (full_name, hashed_password, email, phone_number, role) -- id 2
VALUES ('Lại Ngọc Khánh Thư', '123', 'thulnk@gmail.com', '0806465865', 'dentist');
INSERT INTO dentist_detail (dentist_id, sex, date_of_birth, specialty_id)
VALUES (4, 'female', '2000-12-30 00:00:00+07', 6);

-- Make examination schedules
INSERT INTO examination_schedules (booking_id, dentist_id, start_time, end_time, customer_id, service_category_id,
                                   room_id, slot)
VALUES (NULL, 1, '2024-6-3 07:00:00+07', '2024-6-3 08:00:00+07', NULL, 1, 1, 1),
       (NULL, 2, '2024-6-3 08:00:00+07', '2024-6-3 09:00:00+07', NULL, 2, 2, 2),
       (NULL, 3, '2024-6-3 09:00:00+07', '2024-6-3 10:00:00+07', NULL, 3, 3, 3),
       (NULL, 4, '2024-6-3 10:00:00+07', '2024-6-3 11:00:00+07', NULL, 4, 4, 4),
       (NULL, 1, '2024-6-5 07:00:00+07', '2024-6-5 08:00:00+07', NULL, 1, 1, 1),
       (NULL, 2, '2024-6-5 08:00:00+07', '2024-6-5 09:00:00+07', NULL, 2, 2, 2),
       (NULL, 3, '2024-6-5 09:00:00+07', '2024-6-5 10:00:00+07', NULL, 3, 3, 3),
       (NULL, 4, '2024-6-5 10:00:00+07', '2024-6-5 11:00:00+07', NULL, 4, 4, 4),
       (NULL, 1, '2024-6-7 07:00:00+07', '2024-6-7 08:00:00+07', NULL, 1, 1, 1),
       (NULL, 2, '2024-6-7 08:00:00+07', '2024-6-7 09:00:00+07', NULL, 2, 2, 2),
       (NULL, 3, '2024-6-7 09:00:00+07', '2024-6-7 10:00:00+07', NULL, 3, 3, 3),
       (NULL, 4, '2024-6-7 10:00:00+07', '2024-6-7 11:00:00+07', NULL, 4, 4, 4),
       (NULL, 1, '2024-6-4 13:00:00+07', '2024-6-9 14:00:00+07', NULL, 1, 1, 1),
       (NULL, 2, '2024-6-4 14:00:00+07', '2024-6-9 15:00:00+07', NULL, 2, 2, 2),
       (NULL, 3, '2024-6-4 15:00:00+07', '2024-6-9 16:00:00+07', NULL, 3, 3, 3),
       (NULL, 4, '2024-6-4 16:00:00+07', '2024-6-9 17:00:00+07', NULL, 4, 4, 4),
       (NULL, 1, '2024-6-6 13:00:00+07', '2024-6-9 14:00:00+07', NULL, 1, 1, 1),
       (NULL, 2, '2024-6-6 14:00:00+07', '2024-6-9 15:00:00+07', NULL, 2, 2, 2),
       (NULL, 3, '2024-6-6 15:00:00+07', '2024-6-9 16:00:00+07', NULL, 3, 3, 3),
       (NULL, 4, '2024-6-6 16:00:00+07', '2024-6-9 17:00:00+07', NULL, 4, 4, 4);

-- Insert customers
INSERT INTO users (full_name, hashed_password, email, phone_number, role) -- id 5
VALUES ('Nguyễn Thị Hương', '123', '123', '03478362323', 'customer');