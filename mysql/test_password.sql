-- 测试密码加密
-- 这个文件用于验证密码加密是否正确


-- 插入测试用户
-- 密码: admin123 (使用Go的bcrypt.GenerateFromPassword生成的哈希)
INSERT INTO cadm_users (user_id, password, name) VALUES 
('testuser', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iKTVEFDa', '测试用户');

-- 插入测试用户角色
INSERT INTO cadm_user_roles (user_id, role) VALUES 
('testuser', 'cadre')
ON DUPLICATE KEY UPDATE role = 'cadre';

-- 显示测试用户
SELECT user_id, name FROM cadm_users WHERE user_id = 'testuser';
SELECT user_id, role FROM cadm_user_roles WHERE user_id = 'testuser';

-- 测试登录信息
-- 用户名: testuser
-- 密码: admin123
-- 角色: cadre 