package crud

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/m00nreal/go-jwt-auth/api/models"
	"github.com/m00nreal/go-jwt-auth/api/utils/channels"
	"time"
)

type RepositoryUsersCrud struct {
	db *gorm.DB
}

func NewRepositoryUsersCrud(db *gorm.DB) *RepositoryUsersCrud{
	return &RepositoryUsersCrud{db}
}

func (r *RepositoryUsersCrud) Save(user models.User) (models.User, error)  {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}
	return models.User{}, err
}

func (r *RepositoryUsersCrud) FindAll() ([]models.User, error)  {
	var err error
	users := []models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Limit(100).Find(&users).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return users, nil
	}
	return nil, err
}

func (r *RepositoryUsersCrud) FindById(uid uint32) (models.User, error)  {
	var err error
	user := models.User{}
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Where("id = ?", uid).Find(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
	}(done)
	if channels.OK(done) {
		return user, nil
	}
	if gorm.IsRecordNotFoundError(err) {
		return models.User{}, errors.New("user not found")
	}
	return models.User{}, err
}

func (r *RepositoryUsersCrud) Update(uid uint32, user models.User) (int64, error)  {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		defer close(ch)
		rs = r.db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&models.User{}).UpdateColumns(
			map[string]interface{}{
				"nickname":   user.Nickname,
				"email":      user.Email,
				"updated_at": time.Now(),
			},
		)
		ch <- true
	}(done)

	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}

		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}

func (r *RepositoryUsersCrud) Delete(uid uint32) (int64, error)  {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&models.User{}).Delete(&models.User{})
		ch <- true
	}(done)

	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}

		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}