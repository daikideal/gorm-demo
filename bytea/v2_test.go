package bytea

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func TestInsertWithV2(t *testing.T) {
	type args struct {
		entity ByteaSampleEntity
	}

	tests := []struct {
		name     string
		args     args
		wantErr  bool
		checkErr func(t *testing.T, err error)
	}{
		{
			name: "`data`に値がセットされている場合、Insertできる",
			args: args{
				entity: ByteaSampleEntity{
					ID:   uuid.NewString(),
					Data: []byte("random_bytes"),
				},
			},
			wantErr: false,
		},
		{
			name: "`data`空のバイト列がセットされている場合、Insertできる",
			args: args{
				entity: ByteaSampleEntity{
					ID:   uuid.NewString(),
					Data: []byte{},
				},
			},
			wantErr: false,
		},
		{
			name: "`data`にnilがセットされている場合、NOT NULL制約エラーが発生する",
			args: args{
				entity: ByteaSampleEntity{
					ID:   uuid.NewString(),
					Data: nil,
				},
			},
			wantErr: true,
			checkErr: func(t *testing.T, err error) {
				t.Helper()

				if !isNotNullViolation(t, err) {
					t.Error("unexpected error:", err)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := InsertWithV2(tt.args.entity)
			if (err != nil) != tt.wantErr {
				t.Errorf("InsertWithV2() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				tt.checkErr(t, err)
			}
		})
	}
}

// =========================
// MARK: private functions
// =========================

// 参考: https://github.com/go-gorm/gorm/issues/4037#issuecomment-1376821364
func isNotNullViolation(t *testing.T, err error) bool {
	t.Helper()

	var pgErr *pgconn.PgError

	return errors.As(err, &pgErr) && pgErr.Code == pgerrcode.NotNullViolation
}
