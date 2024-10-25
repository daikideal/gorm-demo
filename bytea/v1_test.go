package bytea

import (
	"testing"

	"github.com/google/uuid"
)

func TestInsertWithV1(t *testing.T) {
	type args struct {
		entity ByteaSampleEntity
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
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
			name: "`data`にnilがセットされている場合、Insertできる",
			args: args{
				entity: ByteaSampleEntity{
					ID:   uuid.NewString(),
					Data: nil,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InsertWithV1(tt.args.entity); (err != nil) != tt.wantErr {
				t.Errorf("Insert() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
