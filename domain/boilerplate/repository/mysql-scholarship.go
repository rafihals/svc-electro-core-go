package repository

import (
	"svc-boilerplate-golang/utils/database"
	"svc-boilerplate-golang/valueobject"
)

func (db *mysqlBoilerplateRepository) GetAllEvaluation(param map[string]interface{}) (response []valueobject.ScholarshipEvaluation, err error) {
	var result valueobject.ScholarshipEvaluation

	builder := database.New(MYSQL, MYSQL_SCHOLARSHIP_EVALUATION, SELECT)

	builder.OnSelect = database.OnSelect{
		Column: []string{
			"uuid",
			"id",
			"periode",
			"jenis_beasiswa",
			"semester",
			"kategori",
			"nilai_minimal",
		},
		Where: param,
	}

	err = builder.QueryBuilder()

	if err != nil {
		return
	}

	query, err := db.sqlDB.Query(builder.Result.Query, builder.Result.Value...)

	if err != nil {
		return
	}

	defer query.Close()

	for query.Next() {
		err = query.Scan(
			&result.UUID,
			&result.ID,
			&result.Periode,
			&result.ScholarshipType,
			&result.Semester,
			&result.Kategori,
			&result.NilaiMinimal,
		)

		if err != nil {
			return
		}

		response = append(response, result)
	}

	return
}
