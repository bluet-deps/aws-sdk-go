// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudsearchdomain_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/bluet-deps/aws-sdk-go/aws"
	"github.com/bluet-deps/aws-sdk-go/aws/session"
	"github.com/bluet-deps/aws-sdk-go/service/cloudsearchdomain"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCloudSearchDomain_Search() {
	svc := cloudsearchdomain.New(session.New())

	params := &cloudsearchdomain.SearchInput{
		Query:        aws.String("Query"), // Required
		Cursor:       aws.String("Cursor"),
		Expr:         aws.String("Expr"),
		Facet:        aws.String("Facet"),
		FilterQuery:  aws.String("FilterQuery"),
		Highlight:    aws.String("Highlight"),
		Partial:      aws.Bool(true),
		QueryOptions: aws.String("QueryOptions"),
		QueryParser:  aws.String("QueryParser"),
		Return:       aws.String("Return"),
		Size:         aws.Int64(1),
		Sort:         aws.String("Sort"),
		Start:        aws.Int64(1),
	}
	resp, err := svc.Search(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearchDomain_Suggest() {
	svc := cloudsearchdomain.New(session.New())

	params := &cloudsearchdomain.SuggestInput{
		Query:     aws.String("Query"),     // Required
		Suggester: aws.String("Suggester"), // Required
		Size:      aws.Int64(1),
	}
	resp, err := svc.Suggest(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudSearchDomain_UploadDocuments() {
	svc := cloudsearchdomain.New(session.New())

	params := &cloudsearchdomain.UploadDocumentsInput{
		ContentType: aws.String("ContentType"),          // Required
		Documents:   bytes.NewReader([]byte("PAYLOAD")), // Required
	}
	resp, err := svc.UploadDocuments(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
