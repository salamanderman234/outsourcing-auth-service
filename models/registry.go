package model

func GetAllModel() []any {
	return []any{
		&Partner{},
		&Admin{},
	}
}