package main

import (
	"context"
	"github.com/gocraft/dbr/v2"
	"hal9000/pkg/client/database"
	"log"
	"time"
)

type App struct {
	AppId       string
	Name        string
	RepoId      string
	Description string
	Status      string
	Home        string
	Icon        string
	Screenshots string
	Maintainers string
	Sources     string
	Readme      string
	Owner       string
	ChartName   string
	CreateTime  time.Time
	StatusTime  time.Time
	UpdateTime  *time.Time
}

var AppColumns = database.GetColumnsFromStruct(&App{})


func main() {
	opt := database.NewDatabaseOptions()
	opt.Host = "192.168.234.137"
	opt.Port = "3306"
	opt.Database = "user"
	opt.Username = "root"
	opt.Password = "123456"
	var apps []*App
	database, err := database.NewDataBase(opt, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer database.Conn.Close()
	dbHandle := database.NewConn(context.Background())

	err = execTx(database)
	if err != nil {
		log.Println(err)
		return
	}

	//_, err = dbHandle.Update("app").Set("status", "deleted").Where(db.Eq("app_id", "2")).Exec()
	//if err != nil {
	//	log.Println(err)
	//	return
	//}

	query := dbHandle.Select(AppColumns...).From("app").Offset(0).Limit(10)

	_, err = query.Load(&apps)
	if err != nil {
		log.Println(err)
		return
	}

	count, err := query.Count()
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(count)
	for _, app := range apps {
		log.Println(app)
	}

}

func execTx(d *database.Database) error {
	dbHandle := d.NewConn(context.TODO())

	return dbHandle.ExecTransaction(func(tx *dbr.Tx) error {
		_, err := tx.Update("app").Set("status", "delete").Where("app_id = ?", "2").Exec()
		if err != nil {
			log.Println(err)
			return err
		}

		newApp := &App{
			AppId:       "3",
			Name:        "22",
			RepoId:      "22",
			Description: "22",
			Status:      "22",
			Home:        "22",
			Icon:        "22",
			Screenshots: "22",
			Maintainers: "22",
			Sources:     "22",
			Readme:      "22",
			Owner:       "22",
			ChartName:   "22",
			CreateTime:  time.Now(),
			StatusTime:  time.Now(),
			UpdateTime:  nil,
		}

		_, err = tx.InsertInto("app").Columns(database.GetColumnsFromStruct(newApp)...).Record(newApp).Exec()
		if err != nil {
			log.Println(err)
			return err
		}
		return nil
	})



	//tx, err := dbHandle.Session.Begin()
	//if err != nil {
	//	log.Println(err)
	//	return err
	//}
	//defer tx.RollbackUnlessCommitted()

	//_, err = tx.Update("app").Set("status", "deleted").Where("app_id = ?", "2").Exec()
	//if err != nil {
	//	log.Println(err)
	//	return err
	//}

	//newApp := &App{
	//	AppId:       "2",
	//	Name:        "22",
	//	RepoId:      "22",
	//	Description: "22",
	//	Status:      "22",
	//	Home:        "22",
	//	Icon:        "22",
	//	Screenshots: "22",
	//	Maintainers: "22",
	//	Sources:     "22",
	//	Readme:      "22",
	//	Owner:       "22",
	//	ChartName:   "22",
	//	CreateTime:  time.Now(),
	//	StatusTime:  time.Now(),
	//	UpdateTime:  nil,
	//}
	//
	//_, err = tx.InsertInto("app").Columns(db.GetColumnsFromStruct(newApp)...).Record(newApp).Exec()
	//if err != nil {
	//	log.Println(err)
	//	return err
	//}
	//tx.Commit()
	//return nil
}