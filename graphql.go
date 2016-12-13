package main

import (
	"fmt"
	"github.com/graphql-go/graphql"
)

var studentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Student",
		Fields: graphql.Fields{
			"authorized": &graphql.Field{
				Type: graphql.Boolean,
			},
			"coursework": &graphql.Field{
				Type: graphql.String,
			},
			"midtermsGrades": &graphql.Field{
				Type: graphql.String,
			},
			"absenceLevels": &graphql.Field{
				Type: graphql.String,
			},
			"examsSchedule": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var queryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"student": &graphql.Field{
				Type: studentType,
				Args: graphql.FieldConfigArgument{
					"username": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"password": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					username, isUsernameOK := p.Args["username"].(string)
					password, isPasswordOK := p.Args["password"].(string)

					if isUsernameOK && isPasswordOK {
						authorized := IsUserAuthorized(username, password)
						coursework, _ := GetUserCoursework(username, password)
						midtermsGrades, _ := GetUserMidterms(username, password)
						absenceLevels, _ := GetUserAbsenceReports(username, password)
						examsSchedule, _ := GetUserExams(username, password)

						return NewStudentAPI(authorized, coursework, midtermsGrades, absenceLevels, examsSchedule), nil
					}

					return nil, nil
				},
			},
		},
	})

var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: queryType,
	},
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("Wrong result, unexpected errors: %v", result.Errors)
	}

	return result
}