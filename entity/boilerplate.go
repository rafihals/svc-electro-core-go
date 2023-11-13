package entity

type BoilerplateKey struct {
	ForeignID uint64 `json:"-"`
}

type Boilerplate struct {
	BoilerplateKey
	CategoryID string `json:"category_id"`
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
