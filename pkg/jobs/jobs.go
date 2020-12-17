package jobs

import (
	"torre-test-gql/graph/model"
	database "torre-test-gql/internal/db"
)

func CountPerStatus() ([]*model.JobStatus, error) {
	stmt, err := database.Db.Prepare(`
		SELECT
			count(id) count, status
		FROM
			job
		GROUP BY status
		;
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var query []*model.JobStatus
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var row model.JobStatus
		err := rows.Scan(&row.Value, &row.Label)
		if err != nil {
			return nil, err
		}
		query = append(query, &row)
	}

	return query, nil
}

func GetByAvgIncomeLastX(last int) ([]*model.JobsByAvgIncome, error) {
	stmt, err := database.Db.Prepare(`
		SELECT
			avg(max_amount) avg, locations
		FROM
			compensation
		INNER JOIN job j on compensation.id = j.compensation
		WHERE max_amount > 0
		and currency='USD$'
		and periodicity='yearly'
		GROUP BY locations
		ORDER BY avg desc LIMIT $1
		;
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var query []*model.JobsByAvgIncome
	rows, err := stmt.Query(last)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var row model.JobsByAvgIncome
		err := rows.Scan(&row.Value, &row.Label)
		if err != nil {
			return nil, err
		}
		query = append(query, &row)
	}

	return query, nil
}
