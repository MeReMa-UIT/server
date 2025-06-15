-- migrate:up
INSERT    INTO record_types (type_id, type_name, template_path, schema_path)
VALUES    (
          '01/BV1',
          'BỆNH ÁN NỘI KHOA',
          './templates/01_BV1/01_BV1.template.json',
          './templates/01_BV1/01_BV1.schema.json'
          );

INSERT    INTO records (
          patient_id,
          doctor_id,
          type_id,
          primary_diagnosis,
          secondary_diagnosis,
          created_at,
          expired_at,
          record_detail
          )
VALUES    (
          1100000000,
          1200000001,
          '01/BV1',
          'M47.9',
          'T78.4',
          '2025-05-26 09:04:15',
          '2025-06-26 09:04:15',
          '
  
{
  "THÔNG TIN CHUNG": {
    "Thông tin bệnh án": {
      "Loại bệnh án": "BỆNH ÁN NỘI KHOA",
      "MS": "01/BV1",
      "Sở Y tế": "",
      "Bệnh viện": "",
      "Khoa": "",
      "Giường": "",
      "Mã YT": "",
      "Số lưu trữ": ""
    },
    "Hành chính": {
      "Họ và tên": "",
      "Sinh ngày": "",
      "Giới": "",
      "Nghề nghiệp": "",
      "Dân tộc": "",
      "Ngoại kiều": "",
      "Địa chỉ": {
        "Số nhà": "",
        "Thôn, phố": "",
        "Xã, phường": "",
        "Huyện (Q, Tx)": ""
      },
      "Nơi làm việc": "",
      "Đối tượng": "",
      "BHYT": {
        "Số thẻ": "",
        "Giá trị đến ngày": ""
      },
      "Liên hệ": {
        "Họ và tên": "",
        "Địa chỉ": "",
        "Số điện thoại": ""
      }
    },
    "Quản lý người bệnh": {
      "Vào viện": "",
      "Trực tiếp vào": "",
      "Nơi giới thiệu": "",
      "Vào viện do bệnh này lần thứ": 0,
      "Vào khoa": {
        "Khoa": "",
        "Thời gian": ""
      },
      "Chuyển khoa": [
        {
          "Khoa": "",
          "Thời gian": ""
        },
        {
          "Khoa": "",
          "Thời gian": ""
        },
        {
          "Khoa": "",
          "Thời gian": ""
        }
      ],
      "Chuyển viện": {
        "Loại": "",
        "Chuyển đến": ""
      },
      "Ra viện": {
        "Thời gian": "",
        "Lý do": ""
      },
      "Tổng số ngày điều trị": 0
    },
    "Chẩn đoán": {
      "Nơi chuyển đến": "",
      "KKB, Cấp cứu": "",
      "Khi vào khoa điều trị": "",
      "Thủ thuật": false,
      "Phẫu thuật": false,
      "Ra viện": {
        "Bệnh chính": "",
        "Bệnh kèm theo": ""
      },
      "Tai biến": false,
      "Biến chứng": false
    },
    "Tình trạng ra viện": {
      "Kết quả điều trị": "",
      "Giải phẫu bệnh": "",
      "Tình hình tử vong": {
        "Thời gian": "",
        "Nguyên nhân": "",
        "Trong 24 giờ vào viện": false
      },
      "Nguyên nhân chính tử vong": "",
      "Khám nghiệm tử thi": false,
      "Chẩn đoán giải phẫu tử thi": ""
    }
  },
  "BỆNH ÁN": {
    "Lý do vào viện": {
      "Lý do": "",
      "Vào ngày thứ": ""
    },
    "Hỏi bệnh": {
      "Quá trình bệnh lý": "",
      "Tiền sử bệnh": {
        "Bản thân": "",
        "Đặc điểm liên quan bệnh": {
          "Dị ứng": "",
          "Ma túy": "",
          "Rượu bia": "",
          "Thuốc lá": "",
          "Thuốc lào": "",
          "Khác": ""
        },
        "Gia đình": ""
      }
    },
    "Khám bệnh": {
      "Toàn thân": {
        "Nội dung": "",
        "Mạch": 0,
        "Nhiệt độ": -0.1,
        "Huyết áp": {
          "Tâm thu": 0,
          "Tâm trương": 0
        },
        "Nhịp thở": 0,
        "Cân nặng": -0.1
      },
      "Các cơ quan": {
        "Tuần hoàn": "",
        "Hô hấp": "",
        "Tiêu hóa": "",
        "Thận - Tiết niệu - Sinh dục": "",
        "Thần kinh": "",
        "Cơ - Xương - Khớp": "",
        "Tai - Mũi - Họng": "",
        "Răng - Hàm - Mặt": "",
        "Mắt": "",
        "Nội tiết, dinh dưỡng và các bệnh lý khác": ""
      },
      "Các xét nghiệm cận lâm sàng cần làm": "",
      "Tóm tắt bệnh án": ""
    },
    "Chẩn đoán khi vào khoa điều trị": {
      "Bệnh chính": "",
      "Bệnh kèm theo": "",
      "Phân biệt": ""
    },
    "Tiên lượng": "",
    "Hướng điều trị": ""
  },
  "TỔNG KẾT BỆNH ÁN": {
    "Quá trình bệnh lý và diễn biến lâm sàng": "",
    "Tóm tắt kết quả xét nghiệm cận lâm sàng có giá trị chẩn đoán": "",
    "Phương pháp điều trị": "",
    "Tình trạng người bệnh ra viện": "",
    "Hướng điều trị và các chế độ tiếp theo": "",
    "Hồ sơ, phim, ảnh": {
      "X - quang": 0,
      "CT Scanner": 0,
      "Siêu âm": 0,
      "Xét nghiệm": 0,
      "Khác": 0
    }
  }
}
'
          ),
          (
          1100000001,
          1200000001,
          '01/BV1',
          'M47.9',
          'T78.4',
          '2025-05-26 09:04:15',
          '2025-06-26 09:04:15',
          '
          
{
  "THÔNG TIN CHUNG": {
    "Thông tin bệnh án": {
      "Loại bệnh án": "BỆNH ÁN NỘI KHOA",
      "MS": "01/BV1",
      "Sở Y tế": "",
      "Bệnh viện": "",
      "Khoa": "",
      "Giường": "",
      "Mã YT": "",
      "Số lưu trữ": ""
    },
    "Hành chính": {
      "Họ và tên": "",
      "Sinh ngày": "",
      "Giới": "",
      "Nghề nghiệp": "",
      "Dân tộc": "",
      "Ngoại kiều": "",
      "Địa chỉ": {
        "Số nhà": "",
        "Thôn, phố": "",
        "Xã, phường": "",
        "Huyện (Q, Tx)": ""
      },
      "Nơi làm việc": "",
      "Đối tượng": "",
      "BHYT": {
        "Số thẻ": "",
        "Giá trị đến ngày": ""
      },
      "Liên hệ": {
        "Họ và tên": "",
        "Địa chỉ": "",
        "Số điện thoại": ""
      }
    },
    "Quản lý người bệnh": {
      "Vào viện": "",
      "Trực tiếp vào": "",
      "Nơi giới thiệu": "",
      "Vào viện do bệnh này lần thứ": 0,
      "Vào khoa": {
        "Khoa": "",
        "Thời gian": ""
      },
      "Chuyển khoa": [
        {
          "Khoa": "",
          "Thời gian": ""
        },
        {
          "Khoa": "",
          "Thời gian": ""
        },
        {
          "Khoa": "",
          "Thời gian": ""
        }
      ],
      "Chuyển viện": {
        "Loại": "",
        "Chuyển đến": ""
      },
      "Ra viện": {
        "Thời gian": "",
        "Lý do": ""
      },
      "Tổng số ngày điều trị": 0
    },
    "Chẩn đoán": {
      "Nơi chuyển đến": "",
      "KKB, Cấp cứu": "",
      "Khi vào khoa điều trị": "",
      "Thủ thuật": false,
      "Phẫu thuật": false,
      "Ra viện": {
        "Bệnh chính": "",
        "Bệnh kèm theo": ""
      },
      "Tai biến": false,
      "Biến chứng": false
    },
    "Tình trạng ra viện": {
      "Kết quả điều trị": "",
      "Giải phẫu bệnh": "",
      "Tình hình tử vong": {
        "Thời gian": "",
        "Nguyên nhân": "",
        "Trong 24 giờ vào viện": false
      },
      "Nguyên nhân chính tử vong": "",
      "Khám nghiệm tử thi": false,
      "Chẩn đoán giải phẫu tử thi": ""
    }
  },
  "BỆNH ÁN": {
    "Lý do vào viện": {
      "Lý do": "",
      "Vào ngày thứ": ""
    },
    "Hỏi bệnh": {
      "Quá trình bệnh lý": "",
      "Tiền sử bệnh": {
        "Bản thân": "",
        "Đặc điểm liên quan bệnh": {
          "Dị ứng": "",
          "Ma túy": "",
          "Rượu bia": "",
          "Thuốc lá": "",
          "Thuốc lào": "",
          "Khác": ""
        },
        "Gia đình": ""
      }
    },
    "Khám bệnh": {
      "Toàn thân": {
        "Nội dung": "",
        "Mạch": 0,
        "Nhiệt độ": -0.1,
        "Huyết áp": {
          "Tâm thu": 0,
          "Tâm trương": 0
        },
        "Nhịp thở": 0,
        "Cân nặng": -0.1
      },
      "Các cơ quan": {
        "Tuần hoàn": "",
        "Hô hấp": "",
        "Tiêu hóa": "",
        "Thận - Tiết niệu - Sinh dục": "",
        "Thần kinh": "",
        "Cơ - Xương - Khớp": "",
        "Tai - Mũi - Họng": "",
        "Răng - Hàm - Mặt": "",
        "Mắt": "",
        "Nội tiết, dinh dưỡng và các bệnh lý khác": ""
      },
      "Các xét nghiệm cận lâm sàng cần làm": "",
      "Tóm tắt bệnh án": ""
    },
    "Chẩn đoán khi vào khoa điều trị": {
      "Bệnh chính": "",
      "Bệnh kèm theo": "",
      "Phân biệt": ""
    },
    "Tiên lượng": "",
    "Hướng điều trị": ""
  },
  "TỔNG KẾT BỆNH ÁN": {
    "Quá trình bệnh lý và diễn biến lâm sàng": "",
    "Tóm tắt kết quả xét nghiệm cận lâm sàng có giá trị chẩn đoán": "",
    "Phương pháp điều trị": "",
    "Tình trạng người bệnh ra viện": "",
    "Hướng điều trị và các chế độ tiếp theo": "",
    "Hồ sơ, phim, ảnh": {
      "X - quang": 0,
      "CT Scanner": 0,
      "Siêu âm": 0,
      "Xét nghiệm": 0,
      "Khác": 0
    }
  }
}
          '
          );

INSERT    INTO conversations (acc_id_1, acc_id_2)
VALUES    (
          (
          SELECT    acc_id
          FROM      accounts
          WHERE     citizen_id = '000000001113'
          ),
          (
          SELECT    acc_id
          FROM      accounts
          WHERE     citizen_id = '000000001112'
          )
          );

-- migrate:down
DELETE    FROM records
WHERE     doctor_id = 1200000001;

DELETE    FROM record_types
WHERE     type_id = '01/BV1';