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
INSERT INTO service_categories (name, image_url, slug, cost, short_description)
VALUES ('Bọc răng sứ', '', 'boc-rang-su', 500000,
        'Bọc răng sứ (phục hình cố định răng sứ) là sử dụng răng sứ được làm hoàn toàn từ sứ hoặc sứ kết hợp cùng kim loại để chụp lên phần răng khiếm khuyết hoặc hư tổn để tái tạo hình dáng, kích thước và màu sắc như răng thật.'),
       ('Cấy ghép Implant', '', 'cay-ghep-implant', 600000,
        'Cấy ghép Implant là phương pháp thay thế răng bị mất bằng cách cấy ghép vào xương hàm một cọc titan hoặc hợp kim titan.'),
       ('Niềng răng thẩm mỹ', '', 'nieng-rang-tham-my', 700000,
        'Niềng răng thẩm mỹ là phương pháp chỉnh hình răng mà không cần phải đeo nhiều phụ kiện ngoại vi.'),
       ('Tẩy trắng răng', '', 'tay-trang-rang', 800000,
        'Tẩy trắng răng là phương pháp giúp làm sáng răng mà không cần phải mài hoặc phục hình răng.'),
       ('Nhổ răng khôn', '', 'nho-rang-khon', 900000,
        'Nhổ răng khôn là phương pháp nhổ răng khôn bị hỏng hoặc gây đau nhức.'),
       ('Bệnh lý nha chu', '', 'benh-ly-nha-chu', 1000000,
        'Bệnh lý nha chu là phương pháp điều trị các bệnh lý nha chu.'),
       ('Điều trị tủy răng', '', 'dieu-tri-tuy-rang', 1100000,
        'Điều trị tủy răng là phương pháp điều trị các bệnh lý tủy răng.');

-- Insert payments.sql
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

INSERT INTO users (full_name, hashed_password, email, phone_number, role) -- id 5
VALUES ('Nguyễn Thị Lan Hương', '123', 'lanhuong.nguyen@gmail.com', '0987654321', 'dentist');
INSERT INTO dentist_detail (dentist_id, sex, date_of_birth, specialty_id)
VALUES (5, 'female', '1985-06-30 00:00:00+07', 7);

INSERT INTO users (full_name, hashed_password, email, phone_number, role) -- id 6
VALUES ('Trần Văn Minh', '123', 'minhtvan.tran@gmail.com', '0912345678', 'dentist');
INSERT INTO dentist_detail (dentist_id, sex, date_of_birth, specialty_id)
VALUES (5, 'male', '1985-06-30 00:00:00+07', 1);

-- Insert examination schedules
INSERT INTO schedules (type, start_time, end_time, dentist_id, room_id)
VALUES ('examination', '2024-06-03 07:00:00+07', '2021-06-03 08:00:00+07', 1, 1),
       ('examination', '2024-06-03 08:00:00+07', '2021-06-03 09:00:00+07', 2, 2),
       ('examination', '2024-06-03 09:00:00+07', '2021-06-03 10:00:00+07', 3, 3),
       ('examination', '2024-06-03 10:00:00+07', '2021-06-03 11:00:00+07', 4, 4);

INSERT INTO examination_schedule_detail (schedule_id, service_category_id)
VALUES (1, 1),
       (2, 2),
       (3, 3),
       (4, 4);

-- Insert services
INSERT INTO services (name, category_id, unit, cost, warranty_duration)
VALUES ('Bọc răng sứ cao cấp', 1, 'cái', 500000, '1 năm'),
       ('Bọc răng sứ titanium', 1, 'cái', 700000, '2 năm'),
       ('Bọc răng xứ composite', 1, 'cái', 300000, '6 tháng'),
       ('Cấy ghép Implant cao cấp', 2, 'cái', 600000, '1 năm'),
       ('Cấy ghép Implant titanium', 2, 'cái', 800000, '2 năm'),
       ('Cấy ghép Implant composite', 2, 'cái', 400000, '6 tháng'),
       ('Niềng răng thẩm mỹ cao cấp', 3, 'cái', 700000, '1 năm'),
       ('Niềng răng thẩm mỹ titanium', 3, 'cái', 900000, '2 năm'),
       ('Niềng răng thẩm mỹ composite', 3, 'cái', 500000, '6 tháng'),
       ('Tẩy trắng răng cao cấp', 4, 'lần', 800000, '1 năm'),
       ('Tẩy trắng răng titanium', 4, 'lần', 1000000, '2 năm'),
       ('Tẩy trắng răng composite', 4, 'lần', 600000, '6 tháng'),
       ('Nhổ răng khôn cao cấp', 5, 'cái', 900000, '1 năm'),
       ('Nhổ răng khôn titanium', 5, 'cái', 1100000, '2 năm'),
       ('Nhổ răng khôn composite', 5, 'cái', 700000, '6 tháng'),
       ('Bệnh lý nha chu cao cấp', 6, 'cái', 1000000, '1 năm'),
       ('Bệnh lý nha chu titanium', 6, 'cái', 1200000, '2 năm'),
       ('Bệnh lý nha chu composite', 6, 'cái', 800000, '6 tháng'),
       ('Điều trị tủy răng cao cấp', 7, 'cái', 1100000, '1 năm'),
       ('Điều trị tủy răng titanium', 7, 'cái', 1300000, '2 năm'),
       ('Điều trị tủy răng composite', 7, 'cái', 900000, '6 tháng');
