package data

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/relay"
)

var postType *graphql.Object
var queryType *graphql.Object
var Schema graphql.Schema

func init() {
	postType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("Post", nil),
			"text": &graphql.Field{
				Type: graphql.String,
			},
		},
	})

 	queryType = graphql.NewObject(graphql.ObjectConfig{
 		Name: "Query",
   	Fields: graphql.Fields{
    	"latestPost": &graphql.Field{
       Type: postType,
       Resolve: func(p graphql.ResolveParams) (interface{}, error) {
         return GetLatestPost(), nil
       },
    	},
   	},
  })


	Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
    Query: queryType,
  })
}

