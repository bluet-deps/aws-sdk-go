// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package lambda_test

import (
	"bytes"
	"fmt"
	"time"

	"bluet-deps/aws-sdk-go/aws"
	"bluet-deps/aws-sdk-go/aws/session"
	"bluet-deps/aws-sdk-go/service/lambda"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleLambda_AddPermission() {
	svc := lambda.New(session.New())

	params := &lambda.AddPermissionInput{
		Action:        aws.String("Action"),       // Required
		FunctionName:  aws.String("FunctionName"), // Required
		Principal:     aws.String("Principal"),    // Required
		StatementId:   aws.String("StatementId"),  // Required
		Qualifier:     aws.String("Qualifier"),
		SourceAccount: aws.String("SourceOwner"),
		SourceArn:     aws.String("Arn"),
	}
	resp, err := svc.AddPermission(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_CreateAlias() {
	svc := lambda.New(session.New())

	params := &lambda.CreateAliasInput{
		FunctionName:    aws.String("FunctionName"), // Required
		FunctionVersion: aws.String("Version"),      // Required
		Name:            aws.String("Alias"),        // Required
		Description:     aws.String("Description"),
	}
	resp, err := svc.CreateAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_CreateEventSourceMapping() {
	svc := lambda.New(session.New())

	params := &lambda.CreateEventSourceMappingInput{
		EventSourceArn:   aws.String("Arn"),                 // Required
		FunctionName:     aws.String("FunctionName"),        // Required
		StartingPosition: aws.String("EventSourcePosition"), // Required
		BatchSize:        aws.Int64(1),
		Enabled:          aws.Bool(true),
	}
	resp, err := svc.CreateEventSourceMapping(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_CreateFunction() {
	svc := lambda.New(session.New())

	params := &lambda.CreateFunctionInput{
		Code: &lambda.FunctionCode{ // Required
			S3Bucket:        aws.String("S3Bucket"),
			S3Key:           aws.String("S3Key"),
			S3ObjectVersion: aws.String("S3ObjectVersion"),
			ZipFile:         []byte("PAYLOAD"),
		},
		FunctionName: aws.String("FunctionName"), // Required
		Handler:      aws.String("Handler"),      // Required
		Role:         aws.String("RoleArn"),      // Required
		Runtime:      aws.String("Runtime"),      // Required
		Description:  aws.String("Description"),
		MemorySize:   aws.Int64(1),
		Publish:      aws.Bool(true),
		Timeout:      aws.Int64(1),
	}
	resp, err := svc.CreateFunction(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_DeleteAlias() {
	svc := lambda.New(session.New())

	params := &lambda.DeleteAliasInput{
		FunctionName: aws.String("FunctionName"), // Required
		Name:         aws.String("Alias"),        // Required
	}
	resp, err := svc.DeleteAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_DeleteEventSourceMapping() {
	svc := lambda.New(session.New())

	params := &lambda.DeleteEventSourceMappingInput{
		UUID: aws.String("String"), // Required
	}
	resp, err := svc.DeleteEventSourceMapping(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_DeleteFunction() {
	svc := lambda.New(session.New())

	params := &lambda.DeleteFunctionInput{
		FunctionName: aws.String("FunctionName"), // Required
		Qualifier:    aws.String("Qualifier"),
	}
	resp, err := svc.DeleteFunction(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_GetAlias() {
	svc := lambda.New(session.New())

	params := &lambda.GetAliasInput{
		FunctionName: aws.String("FunctionName"), // Required
		Name:         aws.String("Alias"),        // Required
	}
	resp, err := svc.GetAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_GetEventSourceMapping() {
	svc := lambda.New(session.New())

	params := &lambda.GetEventSourceMappingInput{
		UUID: aws.String("String"), // Required
	}
	resp, err := svc.GetEventSourceMapping(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_GetFunction() {
	svc := lambda.New(session.New())

	params := &lambda.GetFunctionInput{
		FunctionName: aws.String("FunctionName"), // Required
		Qualifier:    aws.String("Qualifier"),
	}
	resp, err := svc.GetFunction(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_GetFunctionConfiguration() {
	svc := lambda.New(session.New())

	params := &lambda.GetFunctionConfigurationInput{
		FunctionName: aws.String("FunctionName"), // Required
		Qualifier:    aws.String("Qualifier"),
	}
	resp, err := svc.GetFunctionConfiguration(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_GetPolicy() {
	svc := lambda.New(session.New())

	params := &lambda.GetPolicyInput{
		FunctionName: aws.String("FunctionName"), // Required
		Qualifier:    aws.String("Qualifier"),
	}
	resp, err := svc.GetPolicy(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_Invoke() {
	svc := lambda.New(session.New())

	params := &lambda.InvokeInput{
		FunctionName:   aws.String("FunctionName"), // Required
		ClientContext:  aws.String("String"),
		InvocationType: aws.String("InvocationType"),
		LogType:        aws.String("LogType"),
		Payload:        []byte("PAYLOAD"),
		Qualifier:      aws.String("Qualifier"),
	}
	resp, err := svc.Invoke(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_InvokeAsync() {
	svc := lambda.New(session.New())

	params := &lambda.InvokeAsyncInput{
		FunctionName: aws.String("FunctionName"),         // Required
		InvokeArgs:   bytes.NewReader([]byte("PAYLOAD")), // Required
	}
	resp, err := svc.InvokeAsync(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_ListAliases() {
	svc := lambda.New(session.New())

	params := &lambda.ListAliasesInput{
		FunctionName:    aws.String("FunctionName"), // Required
		FunctionVersion: aws.String("Version"),
		Marker:          aws.String("String"),
		MaxItems:        aws.Int64(1),
	}
	resp, err := svc.ListAliases(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_ListEventSourceMappings() {
	svc := lambda.New(session.New())

	params := &lambda.ListEventSourceMappingsInput{
		EventSourceArn: aws.String("Arn"),
		FunctionName:   aws.String("FunctionName"),
		Marker:         aws.String("String"),
		MaxItems:       aws.Int64(1),
	}
	resp, err := svc.ListEventSourceMappings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_ListFunctions() {
	svc := lambda.New(session.New())

	params := &lambda.ListFunctionsInput{
		Marker:   aws.String("String"),
		MaxItems: aws.Int64(1),
	}
	resp, err := svc.ListFunctions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_ListVersionsByFunction() {
	svc := lambda.New(session.New())

	params := &lambda.ListVersionsByFunctionInput{
		FunctionName: aws.String("FunctionName"), // Required
		Marker:       aws.String("String"),
		MaxItems:     aws.Int64(1),
	}
	resp, err := svc.ListVersionsByFunction(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_PublishVersion() {
	svc := lambda.New(session.New())

	params := &lambda.PublishVersionInput{
		FunctionName: aws.String("FunctionName"), // Required
		CodeSha256:   aws.String("String"),
		Description:  aws.String("Description"),
	}
	resp, err := svc.PublishVersion(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_RemovePermission() {
	svc := lambda.New(session.New())

	params := &lambda.RemovePermissionInput{
		FunctionName: aws.String("FunctionName"), // Required
		StatementId:  aws.String("StatementId"),  // Required
		Qualifier:    aws.String("Qualifier"),
	}
	resp, err := svc.RemovePermission(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_UpdateAlias() {
	svc := lambda.New(session.New())

	params := &lambda.UpdateAliasInput{
		FunctionName:    aws.String("FunctionName"), // Required
		Name:            aws.String("Alias"),        // Required
		Description:     aws.String("Description"),
		FunctionVersion: aws.String("Version"),
	}
	resp, err := svc.UpdateAlias(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_UpdateEventSourceMapping() {
	svc := lambda.New(session.New())

	params := &lambda.UpdateEventSourceMappingInput{
		UUID:         aws.String("String"), // Required
		BatchSize:    aws.Int64(1),
		Enabled:      aws.Bool(true),
		FunctionName: aws.String("FunctionName"),
	}
	resp, err := svc.UpdateEventSourceMapping(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_UpdateFunctionCode() {
	svc := lambda.New(session.New())

	params := &lambda.UpdateFunctionCodeInput{
		FunctionName:    aws.String("FunctionName"), // Required
		Publish:         aws.Bool(true),
		S3Bucket:        aws.String("S3Bucket"),
		S3Key:           aws.String("S3Key"),
		S3ObjectVersion: aws.String("S3ObjectVersion"),
		ZipFile:         []byte("PAYLOAD"),
	}
	resp, err := svc.UpdateFunctionCode(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleLambda_UpdateFunctionConfiguration() {
	svc := lambda.New(session.New())

	params := &lambda.UpdateFunctionConfigurationInput{
		FunctionName: aws.String("FunctionName"), // Required
		Description:  aws.String("Description"),
		Handler:      aws.String("Handler"),
		MemorySize:   aws.Int64(1),
		Role:         aws.String("RoleArn"),
		Timeout:      aws.Int64(1),
	}
	resp, err := svc.UpdateFunctionConfiguration(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
