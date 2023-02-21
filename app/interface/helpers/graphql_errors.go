package helpers

import (
	"context"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func GraphQLErrorResponse(ctx context.Context, err error, code int) error {
	graphql.AddError(ctx, &gqlerror.Error{
		Path:    graphql.GetPath(ctx),
		Message: err.Error(),
		Extensions: map[string]interface{}{
			"code": strconv.Itoa(code),
		},
	})

	return gqlerror.Errorf(err.Error())
}
