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
		err := rows.Scan(&row.Count, &row.Status)
		if err != nil {
			return nil, err
		}
		query = append(query, &row)
	}

	return query, nil
}
