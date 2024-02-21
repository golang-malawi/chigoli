package problemdetail

var ProblemDetailRootSchema = "schema.example.com/http/types/problemdetail/"

type ProblemDetail struct {
	Type    string `json:"type"`
	Title   string `json:"title"`
	Detail  string `json:"detail"`
	Context any    `json:"context"`
}
