package repository

import (
	"database/sql"
	"fmt"
	"github.com/juvoinc/exposure-service/config"
	"github.com/juvoinc/exposure-service/model"
	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	Database *sql.DB
	Metrics  []*model.Metric
}

var Instance *Repository

const (
	DRIVER_NAME        = "sqlite3"
	SQLITE_FILE_PATH   = "./metrics.db"
	CREATE_TABLE_QUERY = `CREATE TABLE IF NOT EXISTS metrics (id INTEGER PRIMARY KEY, domain TEXT, carrier TEXT, 
							selector TEXT, match TEXT, count INTEGER, amount NUMERIC,query TEXT, predicate TEXT);`
	SELECT_ALL_QUERY        = `SELECT domain FROM metrics`
	SELECT_AMOUNT_QUERY     = `SELECT amount FROM metrics WHERE domain=$1 AND carrier=$2 AND selector=$3 AND match=$4;`
	SELECT_COUNT_QUERY      = `SELECT count FROM metrics WHERE domain=$1 AND carrier=$2 AND selector=$3 AND match=$4;`
	DELETE_TABLE_QUERY      = `DELETE FROM metrics;`
	INSERT_ONE_METRIC_QUERY = `INSERT INTO metrics (domain, carrier, selector, match, query, predicate, amount, count) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`
	UPDATE_AMOUNT_QUERY     = `UPDATE metrics SET amount=$1 WHERE domain=$2 AND carrier=$3 AND selector=$4 AND match=$5;`
	UPDATE_COUNT_QUERY      = `UPDATE metrics SET count=$1 WHERE domain=$2 AND carrier=$3 AND selector=$4 AND match=$5;`
)

func (repository *Repository) CreateTableUnlessExist() (err error) {
	err = repository.SilentExecute(CREATE_TABLE_QUERY)
	return
}

func (repository *Repository) DeleteTable() (err error) {
	err = repository.SilentExecute(DELETE_TABLE_QUERY)
	return
}

func (repository *Repository) SyncDataWithConfig(config *model.Config) (err error) {
	err = repository.CreateTableUnlessExist()
	if err != nil {
		return
	}
	err = repository.DeleteTable()
	if err != nil {
		return
	}
	repository.Metrics = ConvertConfig(config)
	err = repository.InsertMetrics(repository.Metrics)
	return
}

func (repository *Repository) InsertMetric(metric *model.Metric) (err error) {
	err = repository.SilentExecute(INSERT_ONE_METRIC_QUERY, metric.Domain, metric.Carrier, metric.Selector, metric.Match, metric.Query, metric.Predicate)
	return
}

func (repository *Repository) InsertMetrics(metrics []*model.Metric) (err error) {
	for _, metric := range metrics {
		err = repository.SilentExecute(INSERT_ONE_METRIC_QUERY, metric.Domain, metric.Carrier, metric.Selector, metric.Match, metric.Query, metric.Predicate, 9.0, 10111)
		if err != nil {
			return
		}
	}
	return
}

func (repository *Repository) SilentExecute(queryString string, args ...interface{}) (err error) {
	statement, err := repository.Database.Prepare(queryString)
	if err != nil {
		return
	}
	if len(args) > 0 {
		_, err = statement.Exec(args...)
	} else {
		_, err = statement.Exec()
	}
	return
}

func (repository *Repository) QueryAmount(request *model.MetricRequest) (amount float32, err error) {
	row := repository.Database.QueryRow(SELECT_AMOUNT_QUERY, request.Domain, request.Carrier, request.Selector, request.Match)
	err = row.Scan(&amount)
	return
}

func (repository *Repository) QueryCount(request *model.MetricRequest) (count int, err error) {
	row := repository.Database.QueryRow(SELECT_COUNT_QUERY, request.Domain, request.Carrier, request.Selector, request.Match)
	err = row.Scan(&count)
	return
}

func (repository *Repository) UpdateAmount(request *model.MetricRequest, amount float32) (err error) {
	err = repository.SilentExecute(UPDATE_AMOUNT_QUERY, amount, request.Domain, request.Carrier, request.Selector, request.Match)
	return
}

func (repository *Repository) UpdateCount(request *model.MetricRequest, count int) (err error) {
	err = repository.SilentExecute(UPDATE_COUNT_QUERY, count, request.Domain, request.Carrier, request.Selector, request.Match)
	return
}

func (repository *Repository) SelectAll() (err error) {
	rows, err := repository.Database.Query(SELECT_ALL_QUERY)
	if err != nil {
		return
	}
	var domain string
	for rows.Next() {
		err = rows.Scan(&domain)
		if err != nil {
			return err
		}
	}
	return
}

func (repository *Repository) InitDB() (err error) {
	repository.Database, err = sql.Open(DRIVER_NAME, SQLITE_FILE_PATH)
	return

	//rows, _ := database.Query(SELECT_ALL_QUERY)
	//var id int
	//var firstname string
	//var lastname string
	//for rows.Next() {
	//	rows.Scan(&id, &firstname, &lastname)
	//	fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	//}
}

func init() {
	Instance = &Repository{}
	err := Instance.InitDB()
	if err != nil {
		panic(err)
	}
	err = Instance.SyncDataWithConfig(config.Instance)
	if err != nil {
		panic(err)
	}
	amount, err := Instance.QueryAmount(&model.MetricRequest{Domain: "exposure", Carrier: "bsnl", Selector: "hour", Match: "1"})
	if err != nil {
		panic(err)
	}

	err = Instance.UpdateAmount(&model.MetricRequest{Domain: "exposure", Carrier: "bsnl", Selector: "hour", Match: "2"}, 1039.0)
	if err != nil {
		panic(err)
	}
	fmt.Println(amount)
}
