package purchase

type listResponse struct {
	Count int           `json:"count"`
	Rows  []ResponseDTO `json:"rows"`
}

func MountListResponse(count int, rows []ResponseDTO) listResponse {
	return listResponse{
		Count: count,
		Rows:  rows,
	}
}
