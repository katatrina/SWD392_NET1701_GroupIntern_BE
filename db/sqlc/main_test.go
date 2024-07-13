package db

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"
	
	"github.com/katatrina/SWD392_NET1701_GroupIntern_BE/internal/util"
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
			"description": "Bọc răng sứ (phục hình cố định răng sứ) là sử dụng răng sứ được làm hoàn toàn từ sứ hoặc sứ kết hợp cùng kim loại để chụp lên phần răng khiếm khuyết hoặc hư tổn để tái tạo hình dáng, kích thước và màu sắc như răng thật.",
		},
		{
			"name":        "Cấy ghép Implant",
			"icon_url":    "https://i.ibb.co/5GqWy1d/trong-rang-implant.webp",
			"banner_url":  "https://i.ibb.co/3fyJKdL/Banner-Implant.jpg",
			"slug":        "cay-ghep-implant",
			"description": "Cấy ghép Implant (hay cắm Implant) là phương pháp dùng một trụ chân răng nhân tạo bằng Titanium đặt vào trong xương hàm tại vị trí răng đã mất. Trụ chân răng này sẽ thay thế chân răng thật, sau đó dùng răng sứ gắn lên trụ răng Implant tạo thành răng hoàn chỉnh.",
		},
		{
			"name":        "Niềng răng thẩm mỹ",
			"icon_url":    "https://i.ibb.co/qD6K82P/nieng-rang-tham-my.png",
			"banner_url":  "https://i.ibb.co/NW8dF2Y/Banner-nieng-rang-tham-my.jpg",
			"slug":        "nieng-rang-tham-my",
			"description": "Niềng răng là phương pháp sử dụng khí cụ chuyên dụng được gắn cố định hoặc tháo lắp trên răng để giúp dịch chuyển và sắp xếp răng về đúng vị trí. Từ đó, mang lại cho khách hàng một hàm răng đều, đẹp, đảm bảo chức năng ăn nhai, khớp cắn đúng…",
		},
		{
			"name":        "Tẩy trắng răng",
			"icon_url":    "https://i.ibb.co/b1wTr8L/icon-tay-trang-rang-1.png",
			"banner_url":  "https://i.ibb.co/xD2RBR3/Banner-Tay-Trang-Rang.webp",
			"slug":        "tay-trang-rang",
			"description": "Tẩy trắng răng là phương pháp dùng các hợp chất kết hợp với năng lượng ánh sáng sẽ tạo ra phản ứng oxy hóa cắt đứt các chuỗi phân tử màu trong ngà răng. Từ đó, giúp răng trắng sáng hơn so với màu răng ban đầu mà không làm tổn hại bề mặt răng hay bất kỳ yếu tố nào trong răng.",
		},
		{
			"name":        "Điều trị tủy răng",
			"icon_url":    "https://i.ibb.co/pyBGXtp/dieu-tri-tuy.png",
			"banner_url":  "https://i.ibb.co/cQx67HW/Banner-Dieu-Tri-Tuy.jpg",
			"slug":        "dieu-tri-tuy-rang",
			"description": "Trong cấu trúc răng, tủy răng đóng vai trò rất quan trọng là cung cấp dinh dưỡng nuôi sống và giúp răng luôn vững chắc, và khi tủy răng bị viêm nếu không được điều trị kịp thời gây ra nhiều hậu quả nghiêm trọng đối với sức khỏe của bạn.",
		},
		{
			"name":        "Nhổ răng khôn",
			"icon_url":    "https://i.ibb.co/R0Dy1Kg/icon-nho-rang-khon-1.png",
			"banner_url":  "https://i.ibb.co/7JVhBY0/Banner-Nho-Rang-Khon.webp",
			"slug":        "nho-rang-khon",
			"description": "Răng khôn mọc ngầm là tình trạng rạng răng mọc sai và cần được loại bỏ sớm nhằm hạn chế ảnh hưởng đến các răng cạnh. Tuy nhiên, không phải lúc nào dấu hiệu răng khôn mọc ngầm cũng rõ ràng và dễ nhận biết. Có những trường hợp răng mọc không gây đau đơn nên người bệnh rất khó phát hiện. Tham khảo bài viết sau để nhận biết sớm hơn sự xuất hiện của răng khôn mọc ngầm và cách xử lý chúng.",
		},
	}
	for _, serviceCategory := range serviceCategories {
		arg := CreateServiceCategoryParams{
			Name:        serviceCategory["name"].(string),
			IconUrl:     serviceCategory["icon_url"].(string),
			BannerUrl:   serviceCategory["banner_url"].(string),
			Slug:        serviceCategory["slug"].(string),
			Description: serviceCategory["description"].(string),
		}
		
		_, err := testQueries.CreateServiceCategory(context.Background(), arg)
		require.NoError(t, err)
	}
	
	// Insert payments
	payments := []string{"Tiền mặt", "Banking"}
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
			"gender":      "Nam",
		},
		{
			"full_name":   "Trần Văn Lâm",
			"email":       "lamtv@gmail.com",
			"phone":       "0987654322",
			"dateOfBirth": time.Date(1990, 1, 2, 0, 0, 0, 0, time.Local),
			"gender":      "Nam",
		},
		{
			"full_name":   "Nguyễn Thị Hương",
			"email":       "huongnt10@gmail.com",
			"phone":       "0987654323",
			"dateOfBirth": time.Date(1990, 1, 10, 0, 0, 0, 0, time.Local),
			"gender":      "Nữ",
		},
		{
			"full_name":   "Lê Thị Hồng",
			"email":       "huonglt20@gmail.com",
			"phone":       "0987654324",
			"dateOfBirth": time.Date(1990, 1, 4, 0, 0, 0, 0, time.Local),
			"gender":      "Nữ",
		},
		{
			"full_name":   "Nguyễn Văn Hùng",
			"email":       "hungnv@gmail.com",
			"phone":       "0987654325",
			"dateOfBirth": time.Date(1990, 1, 5, 0, 0, 0, 0, time.Local),
			"gender":      "Nam",
		},
		{
			"full_name":   "Phạm Thị Bích Ngọc",
			"email":       "ngocptb@gmail.com",
			"phone":       "0987654326",
			"dateOfBirth": time.Date(1990, 1, 6, 0, 0, 0, 0, time.Local),
			"gender":      "Nữ",
		},
		{
			"full_name":   "Lại Ngọc Khánh Thư",
			"email":       "thulnk@gmail.com",
			"phone":       "0987654327",
			"dateOfBirth": time.Date(1990, 1, 7, 0, 0, 0, 0, time.Local),
			"gender":      "Nữ",
		},
		{
			"full_name":   "Lê Hoàng Anh",
			"email":       "anhlh25@gmail.com",
			"phone":       "0987654328",
			"dateOfBirth": time.Date(1990, 1, 8, 0, 0, 0, 0, time.Local),
			"gender":      "Nam",
		},
	}
	for _, dentist := range dentists {
		hashedPassword, err := util.GenerateHashedPassword("12345")
		require.NoError(t, err)
		
		arg := CreateUserParams{
			FullName:       dentist["full_name"].(string),
			HashedPassword: hashedPassword,
			Email:          dentist["email"].(string),
			PhoneNumber:    dentist["phone"].(string),
			Role:           "Dentist",
			DateOfBirth:    dentist["dateOfBirth"].(time.Time),
			Gender:         dentist["gender"].(string),
		}
		user, err := testQueries.CreateUser(context.Background(), arg)
		require.NoError(t, err)
		
		argDetail := CreateDentistDetailParams{
			DentistID: user.ID,
			
			SpecialtyID: int64(util.RandomIndex(len(specialties))),
		}
		_, err = testQueries.CreateDentistDetail(context.Background(), argDetail)
		require.NoError(t, err)
	}
	
	// Insert examination schedules
	examinationSchedules := []map[string]interface{}{
		// 15/07/2024
		{
			"start_time": time.Date(2024, 7, 15, 7, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 15, 8, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(1),
			"room_id":    int64(1),
		},
		{
			"start_time": time.Date(2024, 7, 15, 8, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 15, 9, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(2),
			"room_id":    int64(2),
		},
		{
			"start_time": time.Date(2024, 7, 15, 9, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 15, 10, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(3),
			"room_id":    int64(3),
		},
		{
			"start_time": time.Date(2024, 7, 15, 10, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 15, 11, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(4),
			"room_id":    int64(4),
		},
		{
			"start_time": time.Date(2024, 7, 15, 11, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 15, 12, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(5),
			"room_id":    int64(5),
		},
		
		// 16/07/2024
		{
			"start_time": time.Date(2024, 7, 16, 13, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 16, 14, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(6),
			"room_id":    int64(6),
		},
		{
			"start_time": time.Date(2024, 7, 16, 14, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 16, 15, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(7),
			"room_id":    int64(7),
		},
		{
			"start_time": time.Date(2024, 7, 16, 15, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 16, 16, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(8),
			"room_id":    int64(8),
		},
		{
			"start_time": time.Date(2024, 7, 16, 16, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 16, 17, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(4),
			"room_id":    int64(9),
		},
		{
			"start_time": time.Date(2024, 7, 16, 17, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 16, 18, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(2),
			"room_id":    int64(10),
		},
		
		// 17/07/2024
		{
			"start_time": time.Date(2024, 7, 17, 7, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 17, 8, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(1),
			"room_id":    int64(1),
		},
		{
			"start_time": time.Date(2024, 7, 17, 8, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 17, 9, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(2),
			"room_id":    int64(2),
		},
		{
			"start_time": time.Date(2024, 7, 17, 9, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 17, 10, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(3),
			"room_id":    int64(3),
		},
		{
			"start_time": time.Date(2024, 7, 17, 10, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 17, 11, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(4),
			"room_id":    int64(4),
		},
		{
			"start_time": time.Date(2024, 7, 17, 11, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 17, 12, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(5),
			"room_id":    int64(5),
		},
		
		// 18/07/2024
		{
			"start_time": time.Date(2024, 7, 18, 13, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 18, 14, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(6),
			"room_id":    int64(6),
		},
		{
			"start_time": time.Date(2024, 7, 18, 14, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 18, 15, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(7),
			"room_id":    int64(7),
		},
		{
			"start_time": time.Date(2024, 7, 18, 15, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 18, 16, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(8),
			"room_id":    int64(8),
		},
		{
			"start_time": time.Date(2024, 7, 18, 16, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 18, 17, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(4),
			"room_id":    int64(9),
		},
		{
			"start_time": time.Date(2024, 7, 18, 17, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 18, 18, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(2),
			"room_id":    int64(10),
		},
		
		// 19/07/2024
		{
			"start_time": time.Date(2024, 7, 19, 7, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 19, 8, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(1),
			"room_id":    int64(1),
		},
		{
			"start_time": time.Date(2024, 7, 19, 8, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 19, 9, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(2),
			"room_id":    int64(2),
		},
		{
			"start_time": time.Date(2024, 7, 19, 9, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 19, 10, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(3),
			"room_id":    int64(3),
		},
		{
			"start_time": time.Date(2024, 7, 19, 10, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 19, 11, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(4),
			"room_id":    int64(4),
		},
		{
			"start_time": time.Date(2024, 7, 19, 11, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 19, 12, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(5),
			"room_id":    int64(5),
		},
		
		// 20/07/2024
		{
			"start_time": time.Date(2024, 7, 20, 13, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 20, 14, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(6),
			"room_id":    int64(6),
		},
		{
			"start_time": time.Date(2024, 7, 20, 14, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 20, 15, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(7),
			"room_id":    int64(7),
		},
		{
			"start_time": time.Date(2024, 7, 20, 15, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 20, 16, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(8),
			"room_id":    int64(8),
		},
		{
			"start_time": time.Date(2024, 7, 20, 16, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 20, 17, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(4),
			"room_id":    int64(9),
		},
		{
			"start_time": time.Date(2024, 7, 20, 17, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 20, 18, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(2),
			"room_id":    int64(10),
		},
		
		// 22/07/2024
		{
			"start_time": time.Date(2024, 7, 22, 7, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 22, 8, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(1),
			"room_id":    int64(1),
		},
		{
			"start_time": time.Date(2024, 7, 22, 8, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 22, 9, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(2),
			"room_id":    int64(2),
		},
		{
			"start_time": time.Date(2024, 7, 22, 9, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 22, 10, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(3),
			"room_id":    int64(3),
		},
		{
			"start_time": time.Date(2024, 7, 22, 10, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 22, 11, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(4),
			"room_id":    int64(4),
		},
		{
			"start_time": time.Date(2024, 7, 22, 11, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 22, 12, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(5),
			"room_id":    int64(5),
		},
		
		// 23/07/2024
		{
			"start_time": time.Date(2024, 7, 23, 13, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 23, 14, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(6),
			"room_id":    int64(6),
		},
		{
			"start_time": time.Date(2024, 7, 23, 14, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 23, 15, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(7),
			"room_id":    int64(7),
		},
		{
			"start_time": time.Date(2024, 7, 23, 15, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 23, 16, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(8),
			"room_id":    int64(8),
		},
		{
			"start_time": time.Date(2024, 7, 23, 16, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 23, 17, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(4),
			"room_id":    int64(9),
		},
		{
			"start_time": time.Date(2024, 7, 23, 17, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 23, 18, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(2),
			"room_id":    int64(10),
		},
		
		// 24/07/2024
		{
			"start_time": time.Date(2024, 7, 24, 7, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 24, 8, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(1),
			"room_id":    int64(1),
		},
		{
			"start_time": time.Date(2024, 7, 24, 8, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 24, 9, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(2),
			"room_id":    int64(2),
		},
		{
			"start_time": time.Date(2024, 7, 24, 9, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 24, 10, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(3),
			"room_id":    int64(3),
		},
		{
			"start_time": time.Date(2024, 7, 24, 10, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 24, 11, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(4),
			"room_id":    int64(4),
		},
		{
			"start_time": time.Date(2024, 7, 24, 11, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 24, 12, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(5),
			"room_id":    int64(5),
		},
		
		// 25/07/2024
		{
			"start_time": time.Date(2024, 7, 25, 13, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 25, 14, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(6),
			"room_id":    int64(6),
		},
		{
			"start_time": time.Date(2024, 7, 25, 14, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 25, 15, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(7),
			"room_id":    int64(7),
		},
		{
			"start_time": time.Date(2024, 7, 25, 15, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 25, 16, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(8),
			"room_id":    int64(8),
		},
		{
			"start_time": time.Date(2024, 7, 25, 16, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 25, 17, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(4),
			"room_id":    int64(9),
		},
		{
			"start_time": time.Date(2024, 7, 25, 17, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 25, 18, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(2),
			"room_id":    int64(10),
		},
		
		// 26/07/2024
		{
			"start_time": time.Date(2024, 7, 26, 7, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 26, 8, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(1),
			"room_id":    int64(1),
		},
		{
			"start_time": time.Date(2024, 7, 26, 8, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 26, 9, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(2),
			"room_id":    int64(2),
		},
		{
			"start_time": time.Date(2024, 7, 26, 9, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 26, 10, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(3),
			"room_id":    int64(3),
		},
		{
			"start_time": time.Date(2024, 7, 26, 10, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 26, 11, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(4),
			"room_id":    int64(4),
		},
		{
			"start_time": time.Date(2024, 7, 26, 11, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 26, 12, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(5),
			"room_id":    int64(5),
		},
		
		// 27/07/2024
		{
			"start_time": time.Date(2024, 7, 27, 13, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 27, 14, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(6),
			"room_id":    int64(6),
		},
		{
			"start_time": time.Date(2024, 7, 27, 14, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 27, 15, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(7),
			"room_id":    int64(7),
		},
		{
			"start_time": time.Date(2024, 7, 27, 15, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 27, 16, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(8),
			"room_id":    int64(8),
		},
		{
			"start_time": time.Date(2024, 7, 27, 16, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 27, 17, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(4),
			"room_id":    int64(9),
		},
		{
			"start_time": time.Date(2024, 7, 27, 17, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"end_time":   time.Date(2024, 7, 27, 18, 0, 0, 0, time.FixedZone("UTC+7", 7*60*60)),
			"dentist_id": int64(2),
			"room_id":    int64(10),
		},
	}
	for _, examinationSchedule := range examinationSchedules {
		arg := CreateScheduleParams{
			Type:           "Examination",
			StartTime:      examinationSchedule["start_time"].(time.Time),
			EndTime:        examinationSchedule["end_time"].(time.Time),
			DentistID:      examinationSchedule["dentist_id"].(int64),
			RoomID:         examinationSchedule["room_id"].(int64),
			MaxPatients:    3,
			SlotsRemaining: 3,
		}
		_, err := testQueries.CreateSchedule(context.Background(), arg)
		require.NoError(t, err)
	}
	
	// Create sample patient accounts
	hashedPassword, err := util.GenerateHashedPassword("12345")
	require.NoError(t, err)
	arg := CreateUserParams{
		FullName:       "Nguyễn Thị Anh Thư",
		HashedPassword: hashedPassword,
		Email:          "thunt@gmail.com",
		PhoneNumber:    "0987654320",
		Role:           "Patient",
		DateOfBirth:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.Local),
		Gender:         "Nữ",
	}
	_, err = testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	
	hashedPassword, err = util.GenerateHashedPassword("12345")
	require.NoError(t, err)
	arg = CreateUserParams{
		FullName:       "Nguyễn Văn Sang",
		HashedPassword: hashedPassword,
		Email:          "sangnv@gmail.com",
		PhoneNumber:    "0989654321",
		Role:           "Patient",
		DateOfBirth:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.Local),
		Gender:         "Nam",
	}
	_, err = testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	
	hashedPassword, err = util.GenerateHashedPassword("12345")
	require.NoError(t, err)
	arg = CreateUserParams{
		FullName:       "Lương Văn Lâm",
		HashedPassword: hashedPassword,
		Email:          "lamlv@gmail.com",
		PhoneNumber:    "0987624320",
		Role:           "Patient",
		DateOfBirth:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.Local),
		Gender:         "Nam",
	}
	_, err = testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	
	hashedPassword, err = util.GenerateHashedPassword("12345")
	require.NoError(t, err)
	arg = CreateUserParams{
		FullName:       "Tố Hữu",
		HashedPassword: hashedPassword,
		Email:          "tohuu@gmail.com",
		PhoneNumber:    "0987657320",
		Role:           "Patient",
		DateOfBirth:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.Local),
		Gender:         "Nam",
	}
	_, err = testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	
	// Create sample admin account
	hashedPassword, err = util.GenerateHashedPassword("admin")
	require.NoError(t, err)
	arg = CreateUserParams{
		FullName:       "Admin",
		HashedPassword: hashedPassword,
		Email:          "admin@gmail.com",
		PhoneNumber:    "0987834320",
		Role:           "Admin",
		DateOfBirth:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.Local),
		Gender:         "Nam",
	}
	_, err = testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	
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
		{
			"name":                "Răng sứ kim loại Cr-co + Công nghệ SwiftPerfect",
			"service_category_id": int64(1),
			"unit":                "Răng",
			"cost":                int64(3_500_000),
			"warranty_duration":   "4 năm",
		},
		{
			"name":                "Răng sứ Sage + Công nghệ SwiftPerfect",
			"service_category_id": int64(1),
			"unit":                "Răng",
			"cost":                int64(5_500_000),
			"warranty_duration":   "5 năm",
		},
		{
			"name":                "Răng sứ Bio + Công nghệ SwiftPerfect",
			"service_category_id": int64(1),
			"unit":                "Răng",
			"cost":                int64(6_500_000),
			"warranty_duration":   "7 năm",
		},
		{
			"name":                "Dán sứ Viva Shine + Công nghệ SwiftPerfect",
			"service_category_id": int64(1),
			"unit":                "Răng",
			"cost":                int64(8_800_000),
			"warranty_duration":   "10 năm",
		},
		{
			"name":                "Dán sứ Viva Ultrathin + Công nghệ SwiftPerfect\t",
			"service_category_id": int64(1),
			"unit":                "Răng",
			"cost":                int64(12_800_000),
			"warranty_duration":   "15 năm",
		},
		
		// Service category: Cấy ghép Implant
		{
			"name":                "Liệu trình Implant + Abutment Biotem + Máng định vị in 3D + Công nghệ Safest",
			"service_category_id": int64(2),
			"unit":                "Trụ",
			"cost":                int64(17_000_000),
			"warranty_duration":   "10 năm",
		},
		{
			"name":                "Liệu trình Implant + Abutment Megagen Anyridge + Máng định vị in 3D + Công nghệ Safest",
			"service_category_id": int64(2),
			"unit":                "Trụ",
			"cost":                int64(22_000_000),
			"warranty_duration":   "10 năm",
		},
		{
			"name":                "Liệu trình Implant + Abutment Straumann SLA + Máng định vị in 3D + Công nghệ Safest",
			"service_category_id": int64(2),
			"unit":                "Trụ",
			"cost":                int64(30_000_000),
			"warranty_duration":   "Vĩnh viễn",
		},
		{
			"name":                "Liệu trình Implant + Abutment Nobel Active + Máng định vị in 3D + Công nghệ Safest",
			"service_category_id": int64(2),
			"unit":                "Trụ",
			"cost":                int64(34_000_000),
			"warranty_duration":   "Vĩnh viễn",
		},
		{
			"name":                "Răng sứ trên Implant Cera SuperBright + Công nghệ SwiftPerfect",
			"service_category_id": int64(2),
			"unit":                "Răng",
			"cost":                int64(8_400_000),
			"warranty_duration":   "7 năm",
		},
		
		// Service category: Niềng răng thẩm mỹ
		{
			"name":                "Niềng răng mắc cài kim loại chuẩn – Đơn giản + Công nghệ Optimal Align",
			"service_category_id": int64(3),
			"unit":                "Liệu trình",
			"cost":                int64(35_000_000),
			"warranty_duration":   "",
		},
		{
			"name":                "Kế hoạch mô phỏng di chuyển răng 3D (Clincheck)",
			"service_category_id": int64(3),
			"unit":                "Liệu trình",
			"cost":                int64(8_000_000),
			"warranty_duration":   "",
		},
		{
			"name":                "Niềng răng Invisalign – Phức tạp cấp II + Công nghệ Optimal Align",
			"service_category_id": int64(3),
			"unit":                "Liệu trình",
			"cost":                int64(90_000_000),
			"warranty_duration":   "",
		},
		{
			"name":                "Niềng răng Invisalign – Đơn giản + Công nghệ Optimal Align",
			"service_category_id": int64(3),
			"unit":                "Liệu trình",
			"cost":                int64(34_000_000),
			"warranty_duration":   "",
		},
		{
			"name":                "Niềng răng mắc cài kim loại có khóa – Đơn giản + Công nghệ Optimal Align",
			"service_category_id": int64(4),
			"unit":                "Liệu trình",
			"cost":                int64(40_000_000),
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
		{
			"name":                "Tẩy trắng răng tại nhà",
			"service_category_id": int64(4),
			"unit":                "2 Hàm",
			"cost":                int64(1_300_000),
			"warranty_duration":   "",
		},
		{
			"name":                "Tẩy trắng răng tại phòng khám và tại nhà",
			"service_category_id": int64(4),
			"unit":                "Liệu trình",
			"cost":                int64(3_500_000),
			"warranty_duration":   "",
		},
		{
			"name":                "Thuốc tẩy trắng tại nhà",
			"service_category_id": int64(4),
			"unit":                "Ống",
			"cost":                int64(500_000),
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
		{
			"name":                "Điều trị tủy lại",
			"service_category_id": int64(5),
			"unit":                "Răng",
			"cost":                int64(2_500_000),
			"warranty_duration":   "",
		},
		{
			"name":                "Chốt sợi không kim loại mức 1-2",
			"service_category_id": int64(5),
			"unit":                "Răng",
			"cost":                int64(800_000),
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
		{
			"name":                "Nhổ răng khôn mọc lệch mức 1",
			"service_category_id": int64(6),
			"unit":                "Răng",
			"cost":                int64(1_800_000),
			"warranty_duration":   "",
		},
		{
			"name":                "Nhổ răng khôn mọc lệch mức 2",
			"service_category_id": int64(6),
			"unit":                "Răng",
			"cost":                int64(2_500_000),
			"warranty_duration":   "",
		},
	}
	for _, service := range services {
		arg := CreateServiceParams{
			Name:             service["name"].(string),
			CategoryID:       service["service_category_id"].(int64),
			Unit:             service["unit"].(string),
			Cost:             service["cost"].(int64),
			Currency:         "VND",
			WarrantyDuration: service["warranty_duration"].(string),
		}
		_, err := testQueries.CreateService(context.Background(), arg)
		require.NoError(t, err)
	}
}
