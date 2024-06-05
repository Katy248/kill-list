package repository

import (
	"errors"
	"gorm.io/gorm"
	"kill-list/database"
	"strconv"
)

//type DbContext struct {
//	DSN                  string
//	PreferSimpleProtocol bool
//}

//func NewDbContext() *DbContext {
//	return &DbContext{
//		DSN: os.Getenv("DB_CONNECTION_STRING"),
//	}
//}

//func (ctx *DbContext) Open() *gorm.DB {
//	db, err := gorm.Open(postgres.New(postgres.Config{
//		DSN:                  ctx.DSN,
//		PreferSimpleProtocol: ctx.PreferSimpleProtocol,
//	}), &gorm.Config{})
//
//	if err != nil {
//		log.Fatalln(err)
//	}
//	return db
//}

// Pager - paginating object
type Pager struct {
	page  int
	count int
}

var (
	PagerCountMinValue = 1
	PagerCountMaxValue = 20
)

// NewPager creates new instance of type `Pager`
// count
// page
func NewPager(count int, page int) (*Pager, error) {

	if page < 0 {
		return nil, errors.New("page must be greater than zero")
	}

	if count < PagerCountMinValue || count > PagerCountMaxValue {
		return nil,
			errors.New(
				"count must be between " + strconv.Itoa(PagerCountMinValue) + " and " + strconv.Itoa(PagerCountMaxValue))
	}

	return &Pager{
		page:  page,
		count: count,
	}, nil
}

func (p *Pager) Offset() int {
	return p.page * p.count
}
func (p *Pager) Limit() int {
	return p.count
}

func GetLists(db *gorm.DB /*, userId uint*/, pager Pager) (lists *[]database.List) {
	db.Offset(pager.Offset()).Limit(pager.Limit()).Find(lists)
	return lists
}

func GetListItemsDataTypes(db *gorm.DB, listId uint, pager Pager) (types *[]database.ListItemDataType) {
	db.Where("list_id = ?", listId).Offset(pager.Offset()).Limit(pager.Limit()).Find(types)
	return types
}
