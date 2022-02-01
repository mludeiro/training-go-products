package database

var instance Database

func InitInstance() {
	instance.InitializeMySql().Migrate().CreateTestData()
}

func GetInstance() *Database {
	if instance.gormDB == nil {
		InitInstance()
	}
	return &instance
}
