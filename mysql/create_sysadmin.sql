-- -- 创建学院表（如果尚未存在）
-- CREATE TABLE IF NOT EXISTS cadm_departments (
--   id INT AUTO_INCREMENT PRIMARY KEY,
--   name VARCHAR(100) NOT NULL UNIQUE COMMENT '学院名称',
--   description TEXT COMMENT '学院描述'
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='学院表';

-- -- 创建用户表（如果尚未存在）
-- CREATE TABLE IF NOT EXISTS cadm_users (
--   user_id varchar(50) NOT NULL COMMENT '用户ID',
--   password varchar(255) NOT NULL COMMENT '密码(加密存储)',
--   name varchar(50) NOT NULL COMMENT '姓名',
--   role varchar(50) NOT NULL DEFAULT 'cadre' COMMENT '用户角色',
--   department_id INT DEFAULT NULL COMMENT '所属学院ID',
--   PRIMARY KEY (user_id),
--   FOREIGN KEY (department_id) REFERENCES cadm_departments(id)
--     ON DELETE SET NULL ON UPDATE CASCADE
-- ) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 插入系统管理员用户（密码为 bcrypt 加密的 "admin123"）
INSERT INTO cadm_users (user_id, password, name, role, department_id) VALUES 
('school_admin', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iKTVEFDa', '系统管理员', 'school_admin', NULL)
ON DUPLICATE KEY UPDATE 
  password = VALUES(password),
  name = VALUES(name),
  role = VALUES(role),
  department_id = VALUES(department_id);

-- 插入示例院系
-- INSERT INTO cadm_departments (name, description) VALUES 
-- ('计算机科学与技术学院', '计算机科学与技术相关专业'),
-- ('机械工程学院', '机械工程相关专业'),
-- ('电气工程学院', '电气工程相关专业'),
-- ('经济管理学院', '经济管理相关专业'),
-- ('外国语学院', '外语相关专业')
-- ON DUPLICATE KEY UPDATE 
--   description = VALUES(description);

-- -- 显示创建结果
-- SELECT '系统管理员创建完成' AS message;
-- SELECT user_id, name, role FROM cadm_users WHERE user_id = 'sysadmin';
-- SELECT id, name, description FROM cadm_departments;
