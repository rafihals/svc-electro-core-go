package entity

type BoilerplateKey struct {
	ForeignID uint64 `json:"-"`
}

type Boilerplate struct {
	BoilerplateKey
	CategoryID int `json:"category_id"`
	Name       string `json:"name"`
}
type ScholarshipEvaluation struct {
	BoilerplateKey
	Periode         string `json:"periode"`
	ScholarshipType string `json:"jenis_beasiswa"`
	Semester        string `json:"semester"`
	Kategori        string `json:"kategori"`
	NilaiMinimal    string `json:"nilai_minimal"`
}

type User struct {
	ID_user string `json:"ID_user"`
	User       string `json:"User"`
}