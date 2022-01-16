package user

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository struct {
	db     *gorm.DB
	logger zap.SugaredLogger
}

func (repo *UserRepository) GetById(id string) (*User, error) {
	return nil, nil
}

func (repo *UserRepository) List(search string, limit int, offset int) ([]User, *int, error) {
	return nil, nil, nil
}

func (repo *UserRepository) Create(user *User) error {
	return nil
}

func (repo *UserRepository) Update(user *User) error {
	return nil
}

func (repo *UserRepository) Delete(id string) error {
	return nil
}

func (repo *UserRepository) Migrate() error {
	repo.db.AutoMigrate(&User{})
	return nil
}

func registerHooks(
	lifecycle fx.Lifecycle,
	repo *UserRepository,
	logger *zap.SugaredLogger,
) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				logger.Info("UserRepository Initialized")
				logger.Info("Migrating AppRepository")
				return repo.Migrate()
			},
			OnStop: func(ctx context.Context) error {
				return nil
			},
		},
	)
}

func NewUserRepository(db *gorm.DB, logger *zap.SugaredLogger) UserRepository {
	return UserRepository{db, *logger}
}

var Module = fx.Options(
	fx.Invoke(NewUserRepository),
	fx.Invoke(registerHooks),
)
