package db

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"
	
	"github.com/katatrina/SWD392/internal/util"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

var (
	testQueries *Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open("postgres", "postgresql://root:secret@localhost:5432/dental_clinic?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}
	
	testQueries = New(testDB)
	
	m.Run()
}

func TestInitDB(t *testing.T) {
	// Insert specialties
	specialties := []string{"Trợ thủ nha khoa", "Nha chu", "Nội nha", "Phục hình răng", "Nhổ răng - Tiểu phẩu", "Điều dưỡng nha khoa", "Nắn chỉnh răng"}
	for _, specialty := range specialties {
		_, err := testQueries.CreateSpecialty(context.Background(), specialty)
		require.NoError(t, err)
	}
	
	// Insert rooms
	rooms := []string{"01", "02", "03", "04", "05", "06", "07", "08", "09", "10"}
	for _, room := range rooms {
		_, err := testQueries.CreateRoom(context.Background(), room)
		require.NoError(t, err)
	}
	
	// Insert service categories
	serviceCategories := []map[string]interface{}{
		{
			"name":        "Bọc răng sứ",
			"icon_url":    "https://i.ibb.co/Xbqw0bc/icon-boc-rang-su-1.png",
			"banner_url":  "https://i.ibb.co/NFfcvS0/Banner-boc-rang-su.webp",
			"slug":        "boc-rang-su",
			"cost":        int64(100000),
			"description": "",
		},
		{
			"name":        "Cấy ghép Implant",
			"icon_url":    "https://i.ibb.co/5GqWy1d/trong-rang-implant.webp",
			"banner_url":  "https://i.ibb.co/3fyJKdL/Banner-Implant.jpg",
			"slug":        "cay-ghep-implant",
			"cost":        int64(100000),
			"description": "",
		},
		{
			"name":        "Niềng răng thẩm mỹ",
			"icon_url":    "https://i.ibb.co/qD6K82P/nieng-rang-tham-my.png",
			"banner_url":  "https://i.ibb.co/NW8dF2Y/Banner-nieng-rang-tham-my.jpg",
			"slug":        "nieng-rang-tham-my",
			"cost":        int64(100000),
			"description": "",
		},
		{
			"name":        "Tẩy trắng răng",
			"icon_url":    "https://i.ibb.co/b1wTr8L/icon-tay-trang-rang-1.png",
			"banner_url":  "https://i.ibb.co/xD2RBR3/Banner-Tay-Trang-Rang.webp",
			"slug":        "tay-trang-rang",
			"cost":        int64(100000),
			"description": "",
		},
		{
			"name":        "Điều trị tủy răng",
			"icon_url":    "https://i.ibb.co/pyBGXtp/dieu-tri-tuy.png",
			"banner_url":  "https://i.ibb.co/cQx67HW/Banner-Dieu-Tri-Tuy.jpg",
			"slug":        "dieu-tri-tuy-rang",
			"cost":        int64(100000),
			"description": "",
		},
		{
			"name":        "Nhổ răng khôn",
			"icon_url":    "https://i.ibb.co/R0Dy1Kg/icon-nho-rang-khon-1.png",
			"banner_url":  "https://i.ibb.co/7JVhBY0/Banner-Nho-Rang-Khon.webp",
			"slug":        "nho-rang-khon",
			"cost":        int64(100000),
			"description": "",
		},
	}
	for _, serviceCategory := range serviceCategories {
		arg := CreateServiceCategoryParams{
			Name:        serviceCategory["name"].(string),
			IconUrl:     serviceCategory["icon_url"].(string),
			BannerUrl:   serviceCategory["banner_url"].(string),
			Slug:        serviceCategory["slug"].(string),
			Cost:        serviceCategory["cost"].(int64),
			Description: serviceCategory["description"].(string),
		}
		
		_, err := testQueries.CreateServiceCategory(context.Background(), arg)
		require.NoError(t, err)
	}
	
	// Insert payments
	payments := []string{"Tiền mặt", "Banking", "Bitcoin"}
	for _, payment := range payments {
		_, err := testQueries.CreatePayment(context.Background(), payment)
		require.NoError(t, err)
	}
	
	// Insert dentists
	dentists := []map[string]interface{}{
		{
			"full_name":   "Nguyễn Anh Dũng",
			"email":       "dungan@gmail.com",
			"phone":       "0987654321",
			"dateOfBirth": time.Date(1990, 1, 1, 0, 0, 0, 0, time.Local),
			"sex":         "Nam",
		},
		{
			"full_name":   "Trần Văn Lâm",
			"email":       "lamtv@gmail.com",
			"phone":       "0987654322",
			"dateOfBirth": time.Date(1990, 1, 2, 0, 0, 0, 0, time.Local),
			"sex":         "Nam",
		},
		{
			"full_name":   "Nguyễn Thị Hương",
			"email":       "huongnt10@gmail.com",
			"phone":       "0987654323",
			"dateOfBirth": time.Time{},
			"sex":         "Nữ",
		},
		{
			"full_name":   "Lê Thị Hương",
			"email":       "huonglt20@gmail.com",
			"phone":       "0987654324",
			"dateOfBirth": time.Date(1990, 1, 4, 0, 0, 0, 0, time.Local),
			"sex":         "Nữ",
		},
		{
			"full_name":   "Nguyễn Văn Hùng",
			"email":       "hungnv@gmail.com",
			"phone":       "0987654325",
			"dateOfBirth": time.Date(1990, 1, 5, 0, 0, 0, 0, time.Local),
			"sex":         "Nam",
		},
		{
			"full_name":   "Phạm Thị Bích Ngọc",
			"email":       "ngocptb@gmail.com",
			"phone":       "0987654326",
			"dateOfBirth": time.Date(1990, 1, 6, 0, 0, 0, 0, time.Local),
			"sex":         "Nữ",
		},
		{
			"full_name":   "Lại Ngọc Khánh Thư",
			"email":       "thulnk@gmail.com",
			"phone":       "0987654327",
			"dateOfBirth": time.Date(1990, 1, 7, 0, 0, 0, 0, time.Local),
			"sex":         "Nữ",
		},
		{
			"full_name":   "Lê Hoàng Anh",
			"email":       "anhlh25@gmail.com",
			"phone":       "0987654328",
			"dateOfBirth": time.Date(1990, 1, 8, 0, 0, 0, 0, time.Local),
			"sex":         "Nam",
		},
	}
	for _, dentist := range dentists {
		hashedPassword, err := util.GenerateHashedPassword("123456")
		require.NoError(t, err)
		
		arg := CreateDentistParams{
			FullName:       dentist["full_name"].(string),
			HashedPassword: hashedPassword,
			Email:          dentist["email"].(string),
			PhoneNumber:    dentist["phone"].(string),
		}
		user, err := testQueries.CreateDentist(context.Background(), arg)
		require.NoError(t, err)
		
		argDetail := CreateDentistDetailParams{
			DentistID:   user.ID,
			DateOfBirth: dentist["dateOfBirth"].(time.Time),
			Sex:         dentist["sex"].(string),
			SpecialtyID: int64(util.RandomIndex(len(specialties))),
		}
		_, err = testQueries.CreateDentistDetail(context.Background(), argDetail)
		require.NoError(t, err)
	}
	
	// Insert examination schedules
	examinationSchedules := []map[string]interface{}{
		{
			"start_time":          time.Date(2024, 6, 18, 7, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":            time.Date(2024, 6, 18, 8, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id":          int64(1),
			"room_id":             int64(1),
			"service_category_id": int64(util.RandomIndex(len(serviceCategories))),
		},
		{
			"start_time":          time.Date(2024, 6, 18, 8, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":            time.Date(2024, 6, 18, 9, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id":          int64(2),
			"room_id":             int64(2),
			"service_category_id": int64(util.RandomIndex(len(serviceCategories))),
		},
		{
			"start_time":          time.Date(2024, 6, 18, 9, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":            time.Date(2024, 6, 18, 10, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id":          int64(3),
			"room_id":             int64(3),
			"service_category_id": int64(util.RandomIndex(len(serviceCategories))),
		},
		{
			"start_time":          time.Date(2024, 6, 18, 10, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":            time.Date(2024, 6, 18, 11, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id":          int64(4),
			"room_id":             int64(4),
			"service_category_id": int64(util.RandomIndex(len(serviceCategories))),
		},
		{
			"start_time":          time.Date(2024, 6, 18, 11, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":            time.Date(2024, 6, 18, 12, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id":          int64(5),
			"room_id":             int64(5),
			"service_category_id": int64(util.RandomIndex(len(serviceCategories))),
		},
		{
			"start_time":          time.Date(2024, 6, 19, 13, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":            time.Date(2024, 6, 19, 14, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id":          int64(6),
			"room_id":             int64(6),
			"service_category_id": int64(util.RandomIndex(len(serviceCategories))),
		},
		{
			"start_time":          time.Date(2024, 6, 19, 14, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":            time.Date(2024, 6, 19, 15, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id":          int64(7),
			"room_id":             int64(7),
			"service_category_id": int64(util.RandomIndex(len(serviceCategories))),
		},
		{
			"start_time":          time.Date(2024, 6, 19, 15, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":            time.Date(2024, 6, 19, 16, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id":          int64(8),
			"room_id":             int64(8),
			"service_category_id": int64(util.RandomIndex(len(serviceCategories))),
		},
		{
			"start_time":          time.Date(2024, 6, 19, 16, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":            time.Date(2024, 6, 19, 17, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id":          int64(4),
			"room_id":             int64(9),
			"service_category_id": int64(util.RandomIndex(len(serviceCategories))),
		},
		{
			"start_time":          time.Date(2024, 6, 20, 7, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":            time.Date(2024, 6, 20, 8, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id":          int64(3),
			"room_id":             int64(10),
			"service_category_id": int64(util.RandomIndex(len(serviceCategories))),
		},
		{
			"start_time":          time.Date(2024, 6, 20, 8, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":            time.Date(2024, 6, 20, 9, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id":          int64(5),
			"room_id":             int64(1),
			"service_category_id": int64(util.RandomIndex(len(serviceCategories))),
		},
		{
			"start_time":          time.Date(2024, 6, 20, 9, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":            time.Date(2024, 6, 20, 10, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id":          int64(7),
			"room_id":             int64(2),
			"service_category_id": int64(util.RandomIndex(len(serviceCategories))),
		},
	}
	for _, examinationSchedule := range examinationSchedules {
		arg := CreateScheduleParams{
			Type:      "Examination",
			StartTime: examinationSchedule["start_time"].(time.Time),
			EndTime:   examinationSchedule["end_time"].(time.Time),
			DentistID: examinationSchedule["dentist_id"].(int64),
			RoomID:    examinationSchedule["room_id"].(int64),
		}
		schedule, err := testQueries.CreateSchedule(context.Background(), arg)
		require.NoError(t, err)
		
		argDetail := CreateExaminationScheduleDetailParams{
			ScheduleID:        schedule.ID,
			ServiceCategoryID: examinationSchedule["service_category_id"].(int64),
		}
		_, err = testQueries.CreateExaminationScheduleDetail(context.Background(), argDetail)
		require.NoError(t, err)
	}
	
	// Create sample patient account
	hashedPassword, err := util.GenerateHashedPassword("123456")
	require.NoError(t, err)
	arg := CreateUserParams{
		FullName:       "nguyen thi anh thu",
		HashedPassword: hashedPassword,
		Email:          "thunt@gmail.com",
		PhoneNumber:    "0987654320",
		Role:           "Patient",
	}
	_, err = testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	
	// Create sample admin account
	hashedPassword, err = util.GenerateHashedPassword("123456")
	require.NoError(t, err)
	arg = CreateUserParams{
		FullName:       "Admin",
		HashedPassword: hashedPassword,
		Email:          "admin@gmail.com",
		PhoneNumber:    "0987692730",
		Role:           "Admin",
	}
	
	// Insert services
	services := []map[string]interface{}{
		// Service category: Bọc răng sứ
		{
			"name":                "Răng sứ kim loại Titan + Công nghệ SwiftPerfect",
			"service_category_id": int64(1),
			"unit":                "Răng",
			"cost":                int64(2_500_000),
			"warranty_duration":   "3 năm",
		},
		
		// Service category: Cấy ghép Implant
		{
			"name":                "Liệu trình Implant + Abutment Biotem + Máng định vị in 3D + Công nghệ Safest",
			"service_category_id": int64(2),
			"unit":                "Trụ",
			"cost":                int64(17_000_000),
			"warranty_duration":   "10 năm",
		},
		
		// Service category: Niềng răng thẩm mỹ
		{
			"name":                "Niềng răng mắc cài kim loại chuẩn – Đơn giản + Công nghệ Optimal Align",
			"service_category_id": int64(3),
			"unit":                "Liệu trình",
			"cost":                int64(35_000_000),
			"warranty_duration":   "",
		},
		
		// Service category: Tẩy trắng răng
		{
			"name":                "Tẩy trắng răng tại phòng khám",
			"service_category_id": int64(4),
			"unit":                "2 Hàm",
			"cost":                int64(3_000_000),
			"warranty_duration":   "",
		},
		
		// Service category: Điều trị tủy răng
		{
			"name":                "Điều trị tủy răng",
			"service_category_id": int64(5),
			"unit":                "Răng",
			"cost":                int64(1_500_000),
			"warranty_duration":   "",
		},
		
		// Service category: Nhổ răng khôn
		{
			"name":                "Nhổ răng khôn mọc thẳng",
			"service_category_id": int64(6),
			"unit":                "Răng",
			"cost":                int64(1_000_000),
			"warranty_duration":   "",
		},
	}
	for _, service := range services {
		arg := CreateServiceParams{
			Name:             service["name"].(string),
			CategoryID:       service["service_category_id"].(int64),
			Unit:             service["unit"].(string),
			Cost:             service["cost"].(int64),
			WarrantyDuration: service["warranty_duration"].(string),
		}
		_, err := testQueries.CreateService(context.Background(), arg)
		require.NoError(t, err)
	}
}
