-- ALTER TABLE cadm_cadreinfo MODIFY COLUMN birth_date VARCHAR(50);
-- ALTER TABLE cadm_cadreinfo MODIFY COLUMN work_start_date varchar(50);
-- ALTER TABLE cadm_cadreinfo_mod MODIFY COLUMN birth_date VARCHAR(50);
-- ALTER TABLE cadm_cadreinfo_mod MODIFY COLUMN work_start_date varchar(50);
-- ALTER TABLE cadm_resume_entries MODIFY COLUMN start_date VARCHAR(50);
-- ALTER TABLE cadm_resume_entries MODIFY COLUMN end_date varchar(50);
-- ALTER TABLE cadm_resume_entries_mod MODIFY COLUMN start_date VARCHAR(50);
-- ALTER TABLE cadm_resume_entries_mod MODIFY COLUMN end_date varchar(50);
-- ALTER TABLE cadm_resume_entries DROP FOREIGN KEY fk_resume_entry_user;
-- ALTER TABLE cadm_resume_entries_mod DROP FOREIGN KEY fk_resume_entry_user_mod;
-- ALTER TABLE cadm_resume_entries_mod MODIFY COLUMN user_id varchar(50);
-- ALTER TABLE cadm_resume_entries MODIFY COLUMN user_id varchar(50);

-- ALTER TABLE cadm_resume_entries
-- ADD CONSTRAINT fk_resume_entry_user
-- FOREIGN KEY (user_id)
-- REFERENCES cadm_cadreinfo(user_id)
-- ON DELETE CASCADE
-- ON UPDATE CASCADE;

-- ALTER TABLE cadm_resume_entries_mod
-- ADD CONSTRAINT fk_resume_entry_user_mod
-- FOREIGN KEY (user_id)
-- REFERENCES cadm_cadreinfo_mod(user_id)
-- ON DELETE CASCADE
-- ON UPDATE CASCADE;

-- ALTER TABLE cadm_resume_entries DROP COLUMN created_at;
-- ALTER TABLE cadm_resume_entries DROP COLUMN updated_at;
-- ALTER TABLE cadm_resume_entries DROP COLUMN deleted_at;
-- ALTER TABLE cadm_resume_entries DROP COLUMN entry_type;
-- ALTER TABLE cadm_resume_entries DROP COLUMN description;
-- ALTER TABLE cadm_resume_entries DROP COLUMN is_verified;
-- ALTER TABLE cadm_resume_entries DROP COLUMN verified_at;

-- ALTER TABLE cadm_resume_entries_mod DROP COLUMN created_at;
-- ALTER TABLE cadm_resume_entries_mod DROP COLUMN updated_at;
-- ALTER TABLE cadm_resume_entries_mod DROP COLUMN deleted_at;
-- ALTER TABLE cadm_resume_entries_mod DROP COLUMN entry_type;
-- ALTER TABLE cadm_resume_entries_mod DROP COLUMN description;
-- ALTER TABLE cadm_resume_entries_mod DROP COLUMN is_verified;
-- ALTER TABLE cadm_resume_entries_mod DROP COLUMN verified_at;

-- ALTER TABLE cadm_assessments CHANGE cadre_id user_id varchar(50);
-- ALTER TABLE cadm_assessments_mod CHANGE cadre_id user_id varchar(50);

-- ALTER TABLE cadm_position_histories CHANGE cadre_id user_id varchar(50);
-- ALTER TABLE cadm_position_histories_mod CHANGE cadre_id user_id varchar(50);

-- ALTER TABLE cadm_family_members CHANGE cadre_id user_id varchar(50);
-- ALTER TABLE cadm_family_members_mod CHANGE cadre_id user_id varchar(50);