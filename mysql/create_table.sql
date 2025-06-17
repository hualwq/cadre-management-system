-- -- DROP TABLE IF EXISTS `cadm_cadreinfo`;
-- CREATE TABLE `cadm_cadreinfo` (
--   `user_id` varchar(50) NOT NULL COMMENT '干部ID',
--   `name` varchar(50) NOT NULL COMMENT '姓名',
--   `gender` ENUM('男','女') NOT NULL COMMENT '性别',
--   `birth_date` varchar(50) NOT NULL COMMENT '出生日期',
--   `age` tinyint unsigned DEFAULT NULL COMMENT '年龄',
--   `ethnic_group` varchar(20) NOT NULL COMMENT '民族',
--   `native_place` varchar(100) NOT NULL COMMENT '籍贯',
--   `birth_place` varchar(100) DEFAULT NULL COMMENT '出生地',
--   `political_status` ENUM('中共党员','中共预备党员','共青团员') DEFAULT NULL COMMENT '政治面貌',
--   `work_start_date` varchar(50) NOT NULL COMMENT '参加工作时间',
--   `health_status` varchar(20) DEFAULT NULL COMMENT '健康状况',
--   `professional_title` varchar(100) DEFAULT NULL COMMENT '专业技术职务',
--   `specialty` varchar(200) DEFAULT NULL COMMENT '特长',
--   `phone` varchar(20) NOT NULL COMMENT '联系电话',
--   `current_position` varchar(200) NOT NULL COMMENT '现任职务',
--   `resume` text DEFAULT NULL COMMENT '简历',
--   `awards_and_punishments` text DEFAULT NULL COMMENT '奖惩情况',
--   `annual_assessment` text DEFAULT NULL COMMENT '年度考核情况',
--   `email` varchar(50) DEFAULT NULL COMMENT '电子邮箱',
--   `filled_by` varchar(50) DEFAULT NULL COMMENT '填表人',
--   `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
--   `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
--   `full_time_education_degree` varchar(50) DEFAULT NULL COMMENT '全日制教育学历',
--   `full_time_education_school` varchar(200) DEFAULT NULL COMMENT '全日制教育毕业院校',
--   `on_the_job_education_degree` varchar(50) DEFAULT NULL COMMENT '在职教育学历',
--   `on_the_job_education_school` varchar(200) DEFAULT NULL COMMENT '在职教育毕业院校',
--   `reporting_unit` varchar(200) DEFAULT NULL COMMENT '呈报单位',
--   `approval_authority` text DEFAULT NULL COMMENT '审批机关',
--   `administrative_appointment` text DEFAULT NULL COMMENT '行政职务任命',
--   INDEX `idx_name` (`name`),
--   INDEX `idx_phone` (`phone`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='干部信息表';

-- -- DROP TABLE IF EXISTS `cadm_cadreinfo_mod`;
-- CREATE TABLE `cadm_cadreinfo_mod` (
--   `user_id` varchar(50) NOT NULL COMMENT '干部ID',
--   `name` varchar(50) NOT NULL COMMENT '姓名',
--   `gender` ENUM('男','女') NOT NULL COMMENT '性别',
--   `birth_date` varchar(50) NOT NULL COMMENT '出生日期',
--   `age` tinyint unsigned DEFAULT NULL COMMENT '年龄',
--   `ethnic_group` varchar(20) NOT NULL COMMENT '民族',
--   `native_place` varchar(100) NOT NULL COMMENT '籍贯',
--   `birth_place` varchar(100) DEFAULT NULL COMMENT '出生地',
--   `political_status` ENUM('中共党员','中共预备党员','共青团员') DEFAULT NULL COMMENT '政治面貌',
--   `work_start_date` varchar(50) NOT NULL COMMENT '参加工作时间',
--   `health_status` varchar(20) DEFAULT NULL COMMENT '健康状况',
--   `professional_title` varchar(100) DEFAULT NULL COMMENT '专业技术职务',
--   `specialty` varchar(200) DEFAULT NULL COMMENT '特长',
--   `phone` varchar(20) NOT NULL COMMENT '联系电话',
--   `current_position` varchar(200) NOT NULL COMMENT '现任职务',
--   `resume` text DEFAULT NULL COMMENT '简历',
--   `awards_and_punishments` text DEFAULT NULL COMMENT '奖惩情况',
--   `annual_assessment` text DEFAULT NULL COMMENT '年度考核情况',
--   `email` varchar(50) DEFAULT NULL COMMENT '电子邮箱',
--   `filled_by` varchar(50) DEFAULT NULL COMMENT '填表人',
--   `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
--   `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
--   `full_time_education_degree` varchar(50) DEFAULT NULL COMMENT '全日制教育学历',
--   `full_time_education_school` varchar(200) DEFAULT NULL COMMENT '全日制教育毕业院校',
--   `on_the_job_education_degree` varchar(50) DEFAULT NULL COMMENT '在职教育学历',
--   `on_the_job_education_school` varchar(200) DEFAULT NULL COMMENT '在职教育毕业院校',
--   `reporting_unit` varchar(200) DEFAULT NULL COMMENT '呈报单位',
--   `approval_authority` text DEFAULT NULL COMMENT '审批机关',
--   `administrative_appointment` text DEFAULT NULL COMMENT '行政职务任命',
--   INDEX `idx_name` (`name`),
--   INDEX `idx_phone` (`phone`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='干部信息表';

-- -- DROP TABLE IF EXISTS `cadm_assessments`;
-- CREATE TABLE `cadm_assessments` (
--   `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
--   `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
--   `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
--   `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
--   `name` varchar(50) NOT NULL COMMENT '姓名',
--   `user_id` varchar(20) DEFAULT NULL COMMENT '学(工)号',
--   `phone` varchar(20) DEFAULT NULL COMMENT '手机号码',
--   `email` varchar(100) DEFAULT NULL COMMENT '电子邮箱',
--   `department` varchar(100) NOT NULL COMMENT '院系',
--   `category` varchar(20) NOT NULL COMMENT '类别(专职团干部/兼职团干部/教师/学生)',
--   `assess_dept` varchar(100) NOT NULL COMMENT '考核部门(多选)',
--   `year` int NOT NULL COMMENT '考核年度',
--   `work_summary` text NOT NULL COMMENT '工作说明(1000字以内)',
--   `grade` varchar(10) NOT NULL COMMENT '考核等级(优秀/合格/不合格)',
--   PRIMARY KEY (`id`),
--   INDEX `idx_user_id` (`user_id`),
--   INDEX `idx_name` (`name`),
--   INDEX `idx_department` (`department`),
--   INDEX `idx_year` (`year`),
--   INDEX `idx_category` (`category`),
--   INDEX `idx_deleted_at` (`deleted_at`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='考核表';

-- -- DROP TABLE IF EXISTS `cadm_assessments_mod`;
-- CREATE TABLE `cadm_assessments_mod` (
--   `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
--   `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
--   `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
--   `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
--   `name` varchar(50) NOT NULL COMMENT '姓名',
--   `user_id` varchar(20) DEFAULT NULL COMMENT '学(工)号',
--   `phone` varchar(20) DEFAULT NULL COMMENT '手机号码',
--   `email` varchar(100) DEFAULT NULL COMMENT '电子邮箱',
--   `department` varchar(100) NOT NULL COMMENT '院系',
--   `category` varchar(20) NOT NULL COMMENT '类别(专职团干部/兼职团干部/教师/学生)',
--   `assess_dept` varchar(100) NOT NULL COMMENT '考核部门(多选)',
--   `year` int NOT NULL COMMENT '考核年度',
--   `work_summary` text NOT NULL COMMENT '工作说明(1000字以内)',
--   `grade` varchar(10) NOT NULL COMMENT '考核等级(优秀/合格/不合格)',
--   PRIMARY KEY (`id`),
--   INDEX `idx_user_id` (`user_id`),
--   INDEX `idx_name` (`name`),
--   INDEX `idx_department` (`department`),
--   INDEX `idx_year` (`year`),
--   INDEX `idx_category` (`category`),
--   INDEX `idx_deleted_at` (`deleted_at`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='考核表';

-- -- DROP TABLE IF EXISTS `cadm_users`;
-- CREATE TABLE `cadm_users` (
--   `user_id` varchar(50) NOT NULL COMMENT '用户ID',
--   `password` varchar(255) NOT NULL COMMENT '密码(加密存储)',
--   `name` varchar(50) NOT NULL COMMENT '姓名',
--   `role` ENUM('admin','cadre','sysadmin') NOT NULL DEFAULT 'cadre' COMMENT '角色',
--   PRIMARY KEY (`user_id`),
--   INDEX `idx_name` (`name`),
--   INDEX `idx_role` (`role`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- -- DROP TABLE IF EXISTS `cadm_roles`;
-- CREATE TABLE `cadm_roles` (
--   `name` varchar(255) NOT NULL COMMENT '角色名称',
--   `description` text DEFAULT NULL COMMENT '角色描述',
--   PRIMARY KEY (`name`),
--   UNIQUE KEY `uk_name` (`name`)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='角色表';

-- CREATE TABLE `cadm_resume_entries` (
--   `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
--   `user_id` varchar(255) NOT NULL COMMENT '关联用户ID',
--   `start_date` varchar(50) NOT NULL COMMENT '开始日期(格式:2007.09)',
--   `end_date` varchar(50) DEFAULT NULL COMMENT '结束日期(格式:2011.07或NULL表示至今)',
--   `organization` varchar(255) NOT NULL COMMENT '工作单位/学校',
--   `department` varchar(100) DEFAULT NULL COMMENT '学院/部门',
--   `position` varchar(100) DEFAULT NULL COMMENT '职务/身份',
  
--   PRIMARY KEY (`id`),
--   INDEX `idx_user_id` (`user_id`),
--   INDEX `idx_organization` (`organization`),
--   INDEX `idx_department` (`department`),
  
--   CONSTRAINT `fk_resume_entry_user` 
--     FOREIGN KEY (`user_id`) 
--     REFERENCES `cadm_cadreinfo` (`user_id`) 
--     ON DELETE CASCADE 
--     ON UPDATE CASCADE
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='简历经历表';

-- CREATE TABLE `cadm_resume_entries_mod` (
--   `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
--   `user_id` varchar(255) NOT NULL COMMENT '关联用户ID',
--   `start_date` varchar(50) NOT NULL COMMENT '开始日期(格式:2007.09)',
--   `end_date` varchar(50) DEFAULT NULL COMMENT '结束日期(格式:2011.07或NULL表示至今)',
--   `organization` varchar(255) NOT NULL COMMENT '工作单位/学校',
--   `department` varchar(100) DEFAULT NULL COMMENT '学院/部门',
--   `position` varchar(100) DEFAULT NULL COMMENT '职务/身份',

--   PRIMARY KEY (`id`),
--   INDEX `idx_user_id` (`user_id`),
--   INDEX `idx_organization` (`organization`),
--   INDEX `idx_department` (`department`),
  
--   CONSTRAINT `fk_resume_entry_user_mod` 
--     FOREIGN KEY (`user_id`) 
--     REFERENCES `cadm_cadreinfo_mod` (`user_id`) 
--     ON DELETE CASCADE 
--     ON UPDATE CASCADE
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='简历经历表';


-- CREATE TABLE cadm_position_histories (
--     id INT AUTO_INCREMENT PRIMARY KEY,
--     user_id varchar(50) NOT NULL,
--     name VARCHAR(100) NOT NULL,
--     phone_number VARCHAR(20),
--     email VARCHAR(100),
--     department VARCHAR(100) NOT NULL,
--     category VARCHAR(50) NOT NULL,
--     office VARCHAR(100) NOT NULL,
--     academic_year VARCHAR(50) NOT NULL,
--     positions VARCHAR(200),
    
--     -- 定义外键约束
--     FOREIGN KEY (user_id) REFERENCES cadm_cadreinfo(user_id),
    
--     -- 为常用查询字段添加索引
--     INDEX idx_user_id (user_id),
--     INDEX idx_department (department),
--     INDEX idx_category (category),
--     INDEX idx_office (office)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- CREATE TABLE cadm_position_histories_mod (
--     id INT AUTO_INCREMENT PRIMARY KEY,
--     user_id varchar(50) NOT NULL,
--     name VARCHAR(100) NOT NULL,
--     phone_number VARCHAR(20),
--     email VARCHAR(100),
--     department VARCHAR(100) NOT NULL,
--     category VARCHAR(50) NOT NULL,
--     office VARCHAR(100) NOT NULL,
--     academic_year VARCHAR(50) NOT NULL,
--     positions VARCHAR(200),
    
--     -- 定义外键约束
--     FOREIGN KEY (user_id) REFERENCES cadm_cadreinfo(user_id),
    
--     -- 为常用查询字段添加索引
--     INDEX idx_user_id (user_id)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- CREATE TABLE cadm_family_members (
--     id INT PRIMARY KEY AUTO_INCREMENT,
--     user_id varchar(50) NOT NULL,
--     relation VARCHAR(20) NOT NULL,
--     name VARCHAR(50) NOT NULL,
--     birth_date VARCHAR(50),
--     political_status VARCHAR(50),
--     work_unit VARCHAR(200),
    
--     -- 定义外键约束
--     FOREIGN KEY (user_id) REFERENCES cadm_cadreinfo(user_id),
    
--     -- 为常用查询字段添加索引
--     INDEX idx_user_id (user_id)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- CREATE TABLE cadm_family_members_mod (
--     id INT PRIMARY KEY AUTO_INCREMENT,
--     user_id varchar(50) NOT NULL,
--     relation VARCHAR(20) NOT NULL,
--     name VARCHAR(50) NOT NULL,
--     birth_date VARCHAR(50),
--     political_status VARCHAR(50),
--     work_unit VARCHAR(200),
    
--     -- 定义外键约束
--     FOREIGN KEY (user_id) REFERENCES cadm_cadreinfo_mod(user_id),
    
--     -- 为常用查询字段添加索引
--     INDEX idx_user_id (user_id)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- CREATE TABLE `cadm_posexp_mod` (
--   `id` int unsigned NOT NULL AUTO_INCREMENT,
--   `user_id` varchar(50) NOT NULL,
--   `posyear` varchar(20) NOT NULL,
--   `department` varchar(100) NOT NULL,
--   `pos` varchar(50) NOT NULL,
--   PRIMARY KEY (`id`),
--   FOREIGN KEY (user_id) REFERENCES cadm_cadreinfo_mod(user_id)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

-- CREATE TABLE `cadm_posexp` (
--   `id` int unsigned NOT NULL AUTO_INCREMENT,
--   `user_id` varchar(50) NOT NULL,
--   `posyear` varchar(20) NOT NULL,
--   `department` varchar(100) NOT NULL,
--   `pos` varchar(50) NOT NULL,
--   PRIMARY KEY (`id`),
--   FOREIGN KEY (user_id) REFERENCES cadm_cadreinfo(user_id)
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci

-- ALTER TABLE cadm_assessments_mod 
-- ADD COLUMN is_audited BOOLEAN DEFAULT FALSE;
-- ALTER TABLE cadm_cadreinfo_mod
-- ADD COLUMN is_audited BOOLEAN DEFAULT FALSE;
-- ALTER TABLE cadm_family_members_mod
-- ADD COLUMN is_audited BOOLEAN DEFAULT FALSE;
-- ALTER TABLE cadm_position_histories_mod
-- ADD COLUMN is_audited BOOLEAN DEFAULT FALSE;
-- ALTER TABLE cadm_resume_entries_mod
-- ADD COLUMN is_audited BOOLEAN DEFAULT FALSE;
-- ALTER TABLE cadm_posexp_mod
-- ADD COLUMN is_audited BOOLEAN DEFAULT FALSE;

-- ALTER TABLE cadm_position_histories_mod
-- ADD COLUMN applied_at_year INT UNSIGNED;

-- ALTER TABLE cadm_position_histories
-- ADD COLUMN applied_at_year INT UNSIGNED;

-- ALTER TABLE cadm_position_histories_mod
-- ADD COLUMN applied_at_month INT UNSIGNED;

-- ALTER TABLE cadm_position_histories
-- ADD COLUMN applied_at_month INT UNSIGNED;

-- ALTER TABLE cadm_position_histories_mod
-- ADD COLUMN applied_at_day INT UNSIGNED;

-- ALTER TABLE cadm_position_histories
-- ADD COLUMN applied_at_day INT UNSIGNED;

ALTER TABLE cadm_posexp_mod
ADD COLUMN pos_id INT,
ADD CONSTRAINT fk_cadm_posexp_mod_pos_id
FOREIGN KEY (pos_id) 
REFERENCES cadm_position_histories_mod(id);

-- 修改 cadm_posexp 表
ALTER TABLE cadm_posexp
ADD COLUMN pos_id INT,
ADD CONSTRAINT fk_cadm_posexp_pos_id
FOREIGN KEY (pos_id) 
REFERENCES cadm_position_histories(id);  

-- ALTER TABLE cadm_cadreinfo_mod ADD COLUMN photourl VARCHAR(100);  
-- ALTER TABLE cadm_cadreinfo ADD COLUMN photourl VARCHAR(100);     