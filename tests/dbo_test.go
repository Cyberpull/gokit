package tests

import (
	"os"
	"testing"

	"github.com/Cyberpull/gokit/dbo"
	"github.com/Cyberpull/gokit/tests/models"

	"github.com/joho/godotenv"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DBOTestSuite struct {
	suite.Suite

	ins dbo.Instance
}

func (x *DBOTestSuite) SetupSuite() {
	var err error

	godotenv.Load()

	x.ins, err = dbo.Connect(&dbo.Options{
		Driver:   "mysql",
		Host:     "localhost",
		Port:     "3306",
		DBName:   os.Getenv("DB_DATABASE"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Config: &gorm.Config{
			CreateBatchSize: 3000,
			Logger:          logger.Default.LogMode(logger.Silent),
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		},
	})

	require.NoError(x.T(), err)

	x.ins.AddMigrations(
		&models.Actor{},
		&models.Person{},
		&models.Car{},
		&models.Movie{},
	)

	require.NoError(x.T(), x.ins.Migrate())

	x.ins.AddSeeders(func(db *gorm.DB) (err error) {
		return db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoNothing: true,
		}).Create([]*models.Actor{
			{ID: 1, Name: "Christian Ezeani", Age: 25},
			{ID: 2, Name: "John Doe", Age: 15},
		}).Error
	})

	x.ins.AddSeeders(func(db *gorm.DB) (err error) {
		return db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoNothing: true,
		}).Create(&models.Person{ID: 1, Name: "Christian Ezeani"}).Error
	})

	x.ins.AddSeeders(func(db *gorm.DB) (err error) {
		return db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoNothing: true,
		}).Create([]*models.Car{
			{ID: 1, Brand: "Toyota", Color: "Red", OwnerID: 1},
			{ID: 2, Brand: "Honda", Color: "Black", OwnerID: 1},
		}).Error
	})

	x.ins.AddSeeders(func(db *gorm.DB) (err error) {
		return db.Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoNothing: true,
		}).Create([]*models.Movie{
			{ID: 1, Name: "Jumong", IsSeries: true, OwnerID: 1},
			{ID: 2, Name: "King The Land", IsSeries: true, OwnerID: 1},
			{ID: 3, Name: "Avengers", OwnerID: 1},
		}).Error
	})

	require.NoError(x.T(), x.ins.Seed())
}

func (x *DBOTestSuite) TearDownSuite() {
	// x.db
}

func (x *DBOTestSuite) TestPluginPreload() {
	db := x.ins.New()

	var entry models.Person

	result := db.First(&entry)
	require.NoError(x.T(), result.Error)
	require.EqualValues(x.T(), int64(1), result.RowsAffected)
	require.EqualValues(x.T(), uint64(1), entry.ID)

	// Cars
	require.Len(x.T(), entry.Cars, 2)

	// Movies
	require.Len(x.T(), entry.Movies, 1)
}

func (x *DBOTestSuite) TestPluginScope() {
	db := x.ins.New()

	entries := []models.Actor{}

	result := db.Find(&entries)
	require.NoError(x.T(), result.Error)
	require.EqualValues(x.T(), int64(1), result.RowsAffected)
	require.EqualValues(x.T(), uint64(1), entries[0].ID)
}

func (x *DBOTestSuite) TestSet() {
	var data dbo.Set[string]

	err := data.Scan("a,b,c,d,e")
	require.NoError(x.T(), err)
	require.EqualValues(x.T(), []string{"a", "b", "c", "d", "e"}, data.Data)

	value, err := data.Value()
	require.NoError(x.T(), err)
	require.EqualValues(x.T(), "a,b,c,d,e", value)
}

// ===============================

func TestDBO(t *testing.T) {
	suite.Run(t, new(DBOTestSuite))
}
