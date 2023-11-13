package usecase

import (
	"svc-boilerplate-golang/valueobject"
)

func (boilerplate boilerplateUsecase) GetAllEvaluation(param map[string]interface{}) (response []valueobject.ScholarshipEvaluation, err error) {
	response, err = boilerplate.mysqlRepository.GetAllEvaluation(param)
	return
}
