package user

type Repository interface {
	AutoMigrate()
}
