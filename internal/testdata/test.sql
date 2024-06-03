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

INSERT INTO users (full_name, hashed_password, email, phone_number, role) -- id 3
VALUES ('Nguyễn Ánh Vy', '123', 'anhvy@gmail.com', '0834042822', 'dentist');
INSERT INTO dentist_detail (dentist_id, sex, date_of_birth, specialty_id)
VALUES (3, 'female', '1978-03-09 00:00:00+07', 4);

INSERT INTO users (full_name, hashed_password, email, phone_number, role) -- id 4
VALUES ('Lại Ngọc Khánh Thư', '123', 'thulnk@gmail.com', '0806465865', 'dentist');
INSERT INTO dentist_detail (dentist_id, sex, date_of_birth, specialty_id)
VALUES (4, 'female', '2000-12-30 00:00:00+07', 6);

-- Insert examination schedules
INSERT INTO schedules (type, start_time, end_time, dentist_id, room_id)
VALUES ('examination', '2024-06-03 07:00:00+07', '2021-06-03 08:00:00+07', 1, 1),
       ('examination', '2024-06-03 08:00:00+07', '2021-06-03 09:00:00+07', 2, 2),
       ('examination', '2024-06-03 09:00:00+07', '2021-06-03 10:00:00+07', 3, 3),
       ('examination', '2024-06-03 10:00:00+07', '2021-06-03 11:00:00+07', 4, 4);

INSERT INTO examination_schedule_details (schedule_id, service_category_id)
VALUES (1, 1),
       (2, 2),
       (3, 3),
       (4, 4);
