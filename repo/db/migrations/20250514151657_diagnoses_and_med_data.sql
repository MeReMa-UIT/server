-- migrate:up
INSERT    INTO diagnoses (icd_code, name)
VALUES    ('A09.9', 'Viêm dạ dày ruột nhiễm trùng'),
          ('A16.9', 'Lao phổi không xác định'),
          ('A41.9', 'Nhiễm trùng huyết không xác định'),
          ('B15.9', 'Viêm gan A'),
          ('B20', 'HIV/AIDS'),
          ('B34.9', 'Nhiễm virus không xác định'),
          ('D50.9', 'Thiếu máu do thiếu sắt'),
          ('D57.9', 'Bệnh hồng cầu hình liềm'),
          ('D61.9', 'Suy tủy xương'),
          ('D64.9', 'Thiếu máu không rõ nguyên nhân'),
          ('D69.9', 'Rối loạn đông máu'),
          ('D84.9', 'Suy giảm miễn dịch'),
          ('E03.9', 'Suy giáp'),
          ('E04.9', 'Bướu giáp không độc'),
          ('E11.9', 'Đái tháo đường type 2'),
          ('E78.5', 'Rối loạn lipid máu'),
          ('E88.9', 'Rối loạn chuyển hóa khác'),
          ('G20', 'Bệnh Parkinson'),
          ('G25.9', 'Rối loạn vận động'),
          ('G35', 'Đa xơ cứng'),
          ('G40.9', 'Động kinh không xác định'),
          ('G43.9', 'Đau nửa đầu (Migraine)'),
          ('I10', 'Tăng huyết áp nguyên phát'),
          ('I25.1', 'Bệnh mạch vành xơ vữa'),
          ('I48', 'Rung nhĩ'),
          ('I50.9', 'Suy tim không rõ nguyên nhân'),
          ('I73.9', 'Bệnh mạch máu ngoại biên'),
          ('J15.9', 'Viêm phổi do vi khuẩn'),
          ('J18.9', 'Viêm phổi không xác định'),
          ('J20.9', 'Viêm phế quản cấp'),
          ('J44.9', 'COPD (Bệnh phổi tắc nghẽn mạn tính)'),
          ('J45.9', 'Hen phế quản không dị ứng'),
          ('K21.9', 'Trào ngược dạ dày-thực quản (GERD)'),
          ('K29.7', 'Viêm dạ dày mạn'),
          ('K57.3', 'Viêm túi thừa đại tràng'),
          ('K58.9', 'Hội chứng ruột kích thích (IBS)'),
          ('K76.0', 'Gan nhiễm mỡ không do rượu'),
          ('L40.9', 'Vẩy nến'),
          ('M06.9', 'Viêm khớp dạng thấp'),
          ('M15.9', 'Thoái hóa đa khớp'),
          ('M17.9', 'Thoái hóa khớp gối'),
          ('M32.9', 'Lupus ban đỏ hệ thống'),
          ('M47.9', 'Thoái hóa cột sống'),
          ('M54.5', 'Đau thắt lưng'),
          ('M79.7', 'Đau cơ xơ hóa (Fibromyalgia)'),
          ('N18.9', 'Bệnh thận mạn'),
          ('N20.9', 'Sỏi thận'),
          ('N39.0', 'Nhiễm trùng tiểu'),
          ('N40', 'Phì đại tuyến tiền liệt'),
          ('T78.4', 'Dị ứng không xác định');

INSERT    INTO medications (
          name,
          generic_name,
          med_type,
          strength,
          manufacturer
          )
VALUES   
          -- KHÁNG SINH (15 loại)
          (
          'Amoxicillin',
          'Amoxicilin',
          'viên nang',
          '500mg',
          'VN Pharma'
          ),
          (
          'Augmentin',
          'Amoxicilin/Clavulanat',
          'viên nén',
          '625mg',
          'GlaxoSmithKline'
          ),
          (
          'Azithromycin',
          'Azithromycin',
          'viên nén',
          '500mg',
          'Pfizer'
          ),
          (
          'Cefixime',
          'Cefixim',
          'viên nén',
          '200mg',
          'Sandoz'
          ),
          (
          'Ciprofloxacin',
          'Ciprofloxacin',
          'viên nén',
          '500mg',
          'Bayer'
          ),
          (
          'Erythromycin',
          'Erythromycin',
          'viên nén',
          '250mg',
          'Abbott'
          ),
          (
          'Metronidazole',
          'Metronidazol',
          'viên nén',
          '500mg',
          'Sanofi'
          ),
          (
          'Doxycycline',
          'Doxycyclin',
          'viên nén',
          '100mg',
          'Pfizer'
          ),
          (
          'Clarithromycin',
          'Clarithromycin',
          'viên nén',
          '500mg',
          'Abbott'
          ),
          (
          'Ceftriaxone',
          'Ceftriaxon',
          'lọ tiêm',
          '1g',
          'Roche'
          ),
          (
          'Amikacin',
          'Amikacin',
          'lọ tiêm',
          '500mg',
          'Fresenius Kabi'
          ),
          (
          'Levofloxacin',
          'Levofloxacin',
          'viên nén',
          '500mg',
          'Sanofi'
          ),
          (
          'Vancomycin',
          'Vancomycin',
          'lọ tiêm',
          '500mg',
          'Pfizer'
          ),
          (
          'Fluconazole',
          'Fluconazol',
          'viên nén',
          '150mg',
          'Pfizer'
          ),
          (
          'Acyclovir',
          'Acyclovir',
          'viên nén',
          '400mg',
          'GSK'
          ),
          -- GIẢM ĐAU/HẠ SỐT (10 loại)
          (
          'Paracetamol',
          'Paracetamol',
          'viên nén',
          '500mg',
          'DHG Pharma'
          ),
          (
          'Panadol',
          'Paracetamol',
          'viên nén',
          '500mg',
          'GSK'
          ),
          (
          'Efferalgan',
          'Paracetamol',
          'viên sủi',
          '500mg',
          'Sanofi'
          ),
          (
          'Ibuprofen',
          'Ibuprofen',
          'viên nén',
          '400mg',
          'Bayer'
          ),
          (
          'Diclofenac',
          'Diclofenac',
          'viên nén',
          '50mg',
          'Novartis'
          ),
          (
          'Meloxicam',
          'Meloxicam',
          'viên nén',
          '7.5mg',
          'Boehringer'
          ),
          (
          'Celecoxib',
          'Celecoxib',
          'viên nang',
          '200mg',
          'Pfizer'
          ),
          (
          'Tramadol',
          'Tramadol',
          'viên nén',
          '50mg',
          'Grunenthal'
          ),
          ('Aspirin', 'Aspirin', 'viên nén', '81mg', 'Bayer'),
          (
          'Arcoxia',
          'Etoricoxib',
          'viên nén',
          '60mg',
          'MSD'
          ),
          -- TIM MẠCH (8 loại)
          ('Losartan', 'Losartan', 'viên nén', '50mg', 'MSD'),
          (
          'Amlodipine',
          'Amlodipin',
          'viên nén',
          '5mg',
          'Pfizer'
          ),
          (
          'Bisoprolol',
          'Bisoprolol',
          'viên nén',
          '5mg',
          'AstraZeneca'
          ),
          (
          'Atorvastatin',
          'Atorvastatin',
          'viên nén',
          '20mg',
          'Pfizer'
          ),
          (
          'Simvastatin',
          'Simvastatin',
          'viên nén',
          '20mg',
          'Merck'
          ),
          (
          'Furosemide',
          'Furosemid',
          'viên nén',
          '40mg',
          'Sanofi'
          ),
          (
          'Spironolactone',
          'Spironolacton',
          'viên nén',
          '25mg',
          'Pfizer'
          ),
          (
          'Warfarin',
          'Warfarin',
          'viên nén',
          '5mg',
          'Bristol-Myers'
          ),
          -- TIÊU HÓA (7 loại)
          (
          'Omeprazole',
          'Omeprazol',
          'viên nang',
          '20mg',
          'AstraZeneca'
          ),
          (
          'Lansoprazole',
          'Lansoprazol',
          'viên nang',
          '30mg',
          'Takeda'
          ),
          (
          'Domperidone',
          'Domperidon',
          'viên nén',
          '10mg',
          'Janssen'
          ),
          ('Smecta', 'Diosmectit', 'gói bột', '3g', 'Ipsen'),
          (
          'Buscopan',
          'Hyoscine',
          'viên nén',
          '10mg',
          'Boehringer'
          ),
          (
          'Gaviscon',
          'Alginat',
          'hỗn dịch',
          '500mg',
          'Reckitt'
          ),
          (
          'Loperamide',
          'Loperamid',
          'viên nang',
          '2mg',
          'Johnson & Johnson'
          ),
          -- TIỂU ĐƯỜNG (5 loại)
          (
          'Metformin',
          'Metformin',
          'viên nén',
          '500mg',
          'Merck'
          ),
          (
          'Gliclazide',
          'Gliclazid',
          'viên nén',
          '80mg',
          'Servier'
          ),
          (
          'Insulin Mixtard',
          'Insulin người',
          'ống tiêm',
          '100IU/mL',
          'Novo Nordisk'
          ),
          (
          'Glucophage',
          'Metformin',
          'viên nén',
          '850mg',
          'Merck'
          ),
          (
          'Januvia',
          'Sitagliptin',
          'viên nén',
          '100mg',
          'MSD'
          ),
          -- HÔ HẤP (5 loại)
          (
          'Salbutamol',
          'Salbutamol',
          'bình xịt',
          '100mcg/liều',
          'GSK'
          ),
          (
          'Seretide',
          'Salmeterol/Fluticason',
          'bình xịt',
          '25/125mcg',
          'GSK'
          ),
          (
          'Budesonide',
          'Budesonid',
          'dung dịch khí dung',
          '0.5mg/mL',
          'AstraZeneca'
          ),
          (
          'Theophylline',
          'Theophyllin',
          'viên nén',
          '100mg',
          'Sanofi'
          ),
          (
          'Mucosolvan',
          'Ambroxol',
          'siro',
          '30mg/5mL',
          'Boehringer'
          );

-- migrate:down
DELETE    FROM diagnoses;

DELETE    FROM medications;