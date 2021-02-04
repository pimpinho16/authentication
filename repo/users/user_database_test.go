package users

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
/*
func TestUserDatabase_IsUser(t *testing.T) {
	type fields struct {
		user     string
		password string
	}

	db, mock := mockDb()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "users"."bank_usr_users" WHERE 
										username=$1 and password=$2`)).
		WithArgs("aandrade", "12345").
		WillReturnRows(sqlmock.NewRows([]string{"id", "username", "password"}).
			AddRow("1", "aandrade", "12345"))

	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Test Is User Success",
			fields:fields{user: "aandrade",password: "12345"},
			want: true,
		},
		{
			name: "Test Is User Failed",
			fields:fields{user: "jsoriano",password: "67890"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := UserDatabase{db}
			if got := svc.IsUser(tt.fields.user,tt.fields.password);  !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserDataBase() = %v, want %v", got, tt.want)
			}
		})
	}

}
*/

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