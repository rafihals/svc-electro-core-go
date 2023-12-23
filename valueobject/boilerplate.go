package valueobject

import "svc-boilerplate-golang/entity"

type Boilerplate struct {
	entity.Boilerplate
	entity.StandardKey
	entity.Time
}

type User struct {
	entity.User
	entity.StandardKey
	entity.Time
}

type BoilerplatePayloadInsert struct {
	Data []Boilerplate `json:"data" binding:"required"`
	User string
}

type BoilerplatePayloadUpdate struct {
	Data []BoilerplateDataUpdate `json:"data" binding:"required"`
	User string
}

type BoilerplateDataUpdate struct {
	Param Boilerplate `json:"param" binding:"required"`
	Body  Boilerplate `json:"body" binding:"required"`
}

type BoilerplatePayloadDelete struct {
	Param []Boilerplate `json:"param" binding:"required"`
}

type ScholarshipEvaluation struct {
	entity.ScholarshipEvaluation
	entity.StandardKey
	entity.Time
}

type ScholarshipPayloadInsert struct {
	Data []ScholarshipEvaluation `json:"data" binding:"required"`
	User string
}

type ScholarshipPayloadUpdate struct {
	Data []ScholarshipDataUpdate `json:"data" binding:"required"`
	User string
}

type ScholarshipDataUpdate struct {
	Param ScholarshipEvaluation `json:"param" binding:"required"`
	Body  ScholarshipEvaluation `json:"body" binding:"required"`
}

type ScholarshipPayloadDelete struct {
	Param []ScholarshipEvaluation `json:"param" binding:"required"`
}

