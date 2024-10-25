package bytea

type ByteaSampleEntity struct {
	ID   string
	Data []byte
}

func (ByteaSampleEntity) TableName() string {
	return "bytea_samples"
}
