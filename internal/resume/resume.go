package resume

import "fmt"

type Resume struct {
	Id         uint
	DocumentId uint
}

func New() *Resume {
	return &Resume{
		Id:         0,
		DocumentId: 0,
	}
}

func MustNew() Resume {
	return Resume{
		Id:         0,
		DocumentId: 0,
	}
}

func (r *Resume) Init(id uint, documentId uint) {
	r.Id = id
	r.DocumentId = documentId
}

func (r *Resume) Close() {
	fmt.Printf("Achievement '%q' closed\n", r)
}

func (r *Resume) String() string {
	return fmt.Sprintf("Achievement (id=%d, DocumentId=%d", r.Id, r.DocumentId)
}
