package user

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"reflect"
	"testing"
)

func TestNewUserDataBase(t *testing.T) {
	type args struct{
		db *gorm.DB
	}

	db,_ := mockDb()


	testArgs := args{
		db: db,
	}
	tests := []struct{
		name string
		args args
		want *UserDatabase
	}{
		{
			name : "NewUserDataBaseTest",
			args: testArgs,
			want: &UserDatabase{db},
		},
	}


	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := NewUserDataBase(tt.args.db);  !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserDataBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func mockDb() (*gorm.DB, sqlmock.Sqlmock) {
	var (
		db     *sql.DB
		gormDb *gorm.DB
		mock   sqlmock.Sqlmock
	)
	db, mock, _ = sqlmock.New()
	gormDb, _ = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	return gormDb, mock
}