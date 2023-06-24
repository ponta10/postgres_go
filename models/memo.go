package models

import (
	"gorm.io/gorm"
)

type Memo struct {
	gorm.Model
	Content string
}

type MemoModel struct {
	DB *gorm.DB
}

func NewMemoModel(db *gorm.DB) *MemoModel {
	return &MemoModel{DB: db}
}

func (m *MemoModel) GetAll() ([]Memo, error) {
	var memos []Memo
	if err := m.DB.Find(&memos).Error; err != nil {
		return nil, err
	}
	return memos, nil
}

func (m *MemoModel) GetByID(id uint) (Memo, error) {
	var memo Memo
	if err := m.DB.Where("id = ?", id).First(&memo).Error; err != nil {
		return Memo{}, err
	}
	return memo, nil
}

func (m *MemoModel) Create(content string) (Memo, error) {
	memo := Memo{Content: content}
	if err := m.DB.Create(&memo).Error; err != nil {
		return Memo{}, err
	}
	return memo, nil
}

func (m *MemoModel) Update(id uint, content string) (Memo, error) {
	memo, err := m.GetByID(id)
	if err != nil {
		return Memo{}, err
	}
	if err := m.DB.Model(&memo).Update("content", content).Error; err != nil {
		return Memo{}, err
	}
	return memo, nil
}

func (m *MemoModel) Delete(id uint) error {
	memo, err := m.GetByID(id)
	if err != nil {
		return err
	}
	return m.DB.Delete(&memo).Error
}