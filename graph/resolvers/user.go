package resolvers

import (
	"simple-graphql-go/graph/models"

	"github.com/graphql-go/graphql"
)

func GetUserByID(params graphql.ResolveParams) (interface{}, error) {
	id, ok := params.Args["id"].(int)
	if ok {
		var user models.User
		if err := DB.First(&user, id).Error; err != nil {
			return nil, err
		}
		return user, nil
	}
	return nil, nil
}
