package grafine

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

func ExecuteQuery(q string, s graphql.Schema) *graphql.Result {
	reesult := graphql.Do(graphql.Params{
		Schema: s,
		RequestString: q,
	})
	if len(reesult.Errors) > 0 {
		fmt.Printf("unexpected errors: %v", reesult.Errors)
	}
	return reesult
}