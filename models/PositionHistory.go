package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Positionhistory struct {
	ID           int    `gorm:"primaryKey;autoIncrement" json:"id"`
	CadreID      string `gorm:"not null;column:user_id" json:"user_id"`                   // Associates with the cadre's basic info
	Name         string `gorm:"size:100;not null" json:"name"`                            // 姓名
	PhoneNumber  string `gorm:"size:20" json:"phone_number"`                              // 电话号码
	Email        string `gorm:"size:100" json:"email"`                                    // 电子邮件
	Department   string `gorm:"size:100;not null" json:"department"`                      // 院系
	Category     string `gorm:"size:50;not null" json:"category"`                         // 类别: 专职团干部/兼职团干部/教师/学生
	Office       string `gorm:"type:ENUM('校团委内设部门','学生会','研究生会');not null" json:"office"` // 任职部门
	AcademicYear string `gorm:"size:50;not null" json:"academic_year"`                    // 任职年度格式: "2023-2024第一学期"
	Positions    string `gorm:"size:200" json:"positions"`                                // 职位名称
	Year         uint   `gorm:"not null;column:applied_at_year;type:int unsigned" json:"applied_at_year"`
	Month        uint   `gorm:"column:applied_at_month;type:tinyint unsigned" json:"applied_at_month"`
	Day          uint   `gorm:"column:applied_at_day;type:tinyint unsigned" json:"applied_at_day"`
	IsAudited    int    `gorm:"default:0;column:is_audited" json:"is_audited"`
	DepartmentID int    `gorm:"type:int;not null;" json:"department_id"`
}

type Posexp struct {
	ID         int    `gorm:"primaryKey;autoIncrement" json:"id"`
	CadreID    string `gorm:"size:50;column:user_id" json:"user_id"` // ✅ 用空格分隔两个标签
	Posyear    string `gorm:"size:20" json:"year"`
	Department string `gorm:"size:100" json:"department"`
	Pos        string `gorm:"size:50" json:"position"`
	PosID      int    `gorm:"not null;column:pos_id" json:"pos_id"`
	IsAudited  bool   `gorm:"default:false;column:is_audited"`
}

func (Positionhistory) TableName() string {
	return "cadm_position_histories"
}

func (Posexp) TableName() string {
	return "cadm_Posexp"
}

func AddPositionhistory(data map[string]interface{}) (int, error) {
	// 从 data 中解析直接提供的字段
	cadreID, ok := data["user_id"].(string)
	if !ok {
		return -1, errors.New("invalid or missing cadre ID")
	}

	// 查询 cadre_info 表获取其他信息
	var cadre Cadre
	if err := db.Where("user_id = ?", cadreID).First(&cadre).Error; err != nil {
		return -1, fmt.Errorf("请先编辑基本信息或等管理员审核信息")
	}

	// 创建 Positionhistory 对象
	positionHistory := Positionhistory{
		CadreID:      cadreID,
		Name:         cadre.Name,
		PhoneNumber:  cadre.Phone,
		Email:        cadre.Email,
		Department:   data["department"].(string),
		Category:     data["category"].(string),
		Office:       data["office"].(string),
		AcademicYear: data["academic_year"].(string),
		Year:         data["applied_at_year"].(uint),
		Month:        data["applied_at_month"].(uint),
		Day:          data["applied_at_day"].(uint),
		DepartmentID: data["department_id"].(int),
	}

	// 检查是否存在相同 CadreID 和 AcademicYear 的记录
	var existing Positionhistory
	err := db.Where("user_id = ? AND academic_year = ?", cadreID, data["academic_year"].(string)).First(&existing).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return -1, fmt.Errorf("failed to check existing position history: %v", err)
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// 不存在旧记录，执行插入
		if err := db.Create(&positionHistory).Error; err != nil {
			return -1, fmt.Errorf("failed to create position history: %v", err)
		}
	} else {
		// 存在旧记录，更新
		if err := db.Save(&positionHistory).Error; err != nil {
			return -1, fmt.Errorf("failed to update position history: %v", err)
		}
	}

	return positionHistory.ID, nil
}

func Addyearpositon(data map[string]interface{}) error {
	posexp := Posexp{
		CadreID:    data["user_id"].(string),
		Posyear:    data["year"].(string),
		Department: data["department"].(string),
		Pos:        data["position"].(string),
		PosID:      data["posid"].(int),
	}

	// 3. 数据库操作
	if err := db.Create(&posexp).Error; err != nil {
		return fmt.Errorf("database create failed: %v", err)
	}

	return nil
}

func GetPositionHistoryModByID(id int) (*Positionhistory, error) {
	var positionHistoryMod Positionhistory
	if err := db.Where("id = ?", id).First(&positionHistoryMod).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &positionHistoryMod, nil
}

// GetPositionHistories gets a list of position histories based on paging and constraints
func GetPositionHistories(pageNum int, pageSize int, maps interface{}) ([]Positionhistory, error) {
	var (
		positionHistories []Positionhistory
		err               error
	)

	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(&positionHistories).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&positionHistories).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return positionHistories, nil
}

// GetPositionHistoryTotal counts the total number of position histories based on the constraint
func GetPositionHistoryModTotal(maps interface{}) (int64, error) {
	var count int64
	if err := db.Model(&Positionhistory{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// DeletePositionHistoryByID delete a single position history
func DeletePositionHistoryByID(id int) error {
	if err := db.Where("id = ?", id).Delete(Positionhistory{}).Error; err != nil {
		return err
	}

	return nil
}

func ExistPositionHistoryByID(id int) (bool, int, error) {
	var positionHistory Positionhistory
	err := db.Select("id").Where("id = ?  and is_audited = ?", id, 0).First(&positionHistory).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, -1, err
	}

	if positionHistory.ID > 0 {
		return true, -1, nil
	}

	return false, -1, nil
}

func ExistPosexpByID(id int) (bool, error) {
	var posexpMod Posexp
	err := db.Select("id").Where("id = ?", id).First(&posexpMod).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if posexpMod.ID > 0 {
		return true, nil
	}
	return false, nil
}

func GetPosexpByID(id int) (*Posexp, error) {
	var posexpMod Posexp
	err := db.Where("id = ?", id).First(&posexpMod).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &posexpMod, nil
}

func EditPositionHistoryMod(id int, data map[string]interface{}) error {
	var positionHistoryMod Positionhistory
	if err := db.Model(&positionHistoryMod).Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func ExistPoexpModByCadreID(cadreID string) (bool, error) {
	var poexpMod Posexp
	err := db.Select("id").Where("user_id = ?", cadreID).First(&poexpMod).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if poexpMod.ID > 0 {
		return true, nil
	}
	return false, nil
}

// GetPoexpModByCadreID 根据 CadreID 获取 PoexpMod 记录
func GetPoexpModByCadreID(cadreID string) ([]Posexp, error) {
	var poexpMods []Posexp
	err := db.Where("user_id = ?", cadreID).Find(&poexpMods).Error
	if err != nil {
		return nil, err
	}
	return poexpMods, nil
}

func Comfirmpoexp(cadreID string) error {
	var mod Posexp
	// 查询待审核的岗位经历修改记录
	result := db.Where("user_id = ? AND is_audited = 0", cadreID).First(&mod)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return fmt.Errorf("未找到待审核的岗位经历信息: %s", cadreID)
		}
		return fmt.Errorf("查询待审核岗位经历信息失败: %v", result.Error)
	}

	// 创建正式的岗位经历记录
	poexp := Posexp{
		CadreID:    mod.CadreID,
		Posyear:    mod.Posyear,
		Department: mod.Department,
		Pos:        mod.Pos,
	}
	if err := db.Create(&poexp).Error; err != nil {
		return err
	}

	// 更新修改记录的审核状态
	mod.IsAudited = true
	if err := db.Save(&mod).Error; err != nil {
		return fmt.Errorf("更新审核状态失败: %v", err)
	}

	return nil
}

func DeletePosexpByID(id int) error {
	if id <= 0 {
		return errors.New("无效的岗位经历记录 ID")
	}
	result := db.Where("id = ?", id).Delete(&Posexp{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("未找到匹配的岗位经历记录")
	}
	return nil
}

func GetPositionHistoryModsByUserID(userID string, pageNum int, pageSize int) ([]Positionhistory, error) {
	var positionHistoryMods []Positionhistory
	offset := (pageNum - 1) * pageSize
	err := db.Where("user_id = ?", userID).Offset(offset).Limit(pageSize).Find(&positionHistoryMods).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return positionHistoryMods, nil
}

func GetPosExpTotalByPosID(posID int) (int64, error) {
	var count int64
	err := db.Model(&Posexp{}).Where("pos_id = ?", posID).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetPosExpByPosID(posID int) ([]Posexp, error) {
	var posExps []Posexp
	err := db.Where("pos_id= ?", posID).Find(&posExps).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return posExps, nil
}

func ComfirmPositionhistory(id int) error {
	result := db.Model(&Positionhistory{}).
		Where("id = ? and is_audited = 0", id, 0).
		Updates(map[string]interface{}{
			"is_audited": 1, // Using column name from struct tag
		})

	if result.Error != nil {
		return result.Error
	}

	// Check if any record was actually updated
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
