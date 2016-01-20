// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudfront_test

import (
	"bytes"
	"fmt"
	"time"

	"bluet-deps/aws-sdk-go/aws"
	"bluet-deps/aws-sdk-go/aws/session"
	"bluet-deps/aws-sdk-go/service/cloudfront"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCloudFront_CreateCloudFrontOriginAccessIdentity() {
	svc := cloudfront.New(session.New())

	params := &cloudfront.CreateCloudFrontOriginAccessIdentityInput{
		CloudFrontOriginAccessIdentityConfig: &cloudfront.OriginAccessIdentityConfig{ // Required
			CallerReference: aws.String("string"), // Required
			Comment:         aws.String("string"), // Required
		},
	}
	resp, err := svc.CreateCloudFrontOriginAccessIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFront_CreateDistribution() {
	svc := cloudfront.New(session.New())

	params := &cloudfront.CreateDistributionInput{
		DistributionConfig: &cloudfront.DistributionConfig{ // Required
			CallerReference: aws.String("string"), // Required
			Comment:         aws.String("string"), // Required
			DefaultCacheBehavior: &cloudfront.DefaultCacheBehavior{ // Required
				ForwardedValues: &cloudfront.ForwardedValues{ // Required
					Cookies: &cloudfront.CookiePreference{ // Required
						Forward: aws.String("ItemSelection"), // Required
						WhitelistedNames: &cloudfront.CookieNames{
							Quantity: aws.Int64(1), // Required
							Items: []*string{
								aws.String("string"), // Required
								// More values...
							},
						},
					},
					QueryString: aws.Bool(true), // Required
					Headers: &cloudfront.Headers{
						Quantity: aws.Int64(1), // Required
						Items: []*string{
							aws.String("string"), // Required
							// More values...
						},
					},
				},
				MinTTL:         aws.Int64(1),         // Required
				TargetOriginId: aws.String("string"), // Required
				TrustedSigners: &cloudfront.TrustedSigners{ // Required
					Enabled:  aws.Bool(true), // Required
					Quantity: aws.Int64(1),   // Required
					Items: []*string{
						aws.String("string"), // Required
						// More values...
					},
				},
				ViewerProtocolPolicy: aws.String("ViewerProtocolPolicy"), // Required
				AllowedMethods: &cloudfront.AllowedMethods{
					Items: []*string{ // Required
						aws.String("Method"), // Required
						// More values...
					},
					Quantity: aws.Int64(1), // Required
					CachedMethods: &cloudfront.CachedMethods{
						Items: []*string{ // Required
							aws.String("Method"), // Required
							// More values...
						},
						Quantity: aws.Int64(1), // Required
					},
				},
				Compress:        aws.Bool(true),
				DefaultTTL:      aws.Int64(1),
				MaxTTL:          aws.Int64(1),
				SmoothStreaming: aws.Bool(true),
			},
			Enabled: aws.Bool(true), // Required
			Origins: &cloudfront.Origins{ // Required
				Quantity: aws.Int64(1), // Required
				Items: []*cloudfront.Origin{
					{ // Required
						DomainName: aws.String("string"), // Required
						Id:         aws.String("string"), // Required
						CustomHeaders: &cloudfront.CustomHeaders{
							Quantity: aws.Int64(1), // Required
							Items: []*cloudfront.OriginCustomHeader{
								{ // Required
									HeaderName:  aws.String("string"), // Required
									HeaderValue: aws.String("string"), // Required
								},
								// More values...
							},
						},
						CustomOriginConfig: &cloudfront.CustomOriginConfig{
							HTTPPort:             aws.Int64(1),                       // Required
							HTTPSPort:            aws.Int64(1),                       // Required
							OriginProtocolPolicy: aws.String("OriginProtocolPolicy"), // Required
							OriginSslProtocols: &cloudfront.OriginSslProtocols{
								Items: []*string{ // Required
									aws.String("SslProtocol"), // Required
									// More values...
								},
								Quantity: aws.Int64(1), // Required
							},
						},
						OriginPath: aws.String("string"),
						S3OriginConfig: &cloudfront.S3OriginConfig{
							OriginAccessIdentity: aws.String("string"), // Required
						},
					},
					// More values...
				},
			},
			Aliases: &cloudfront.Aliases{
				Quantity: aws.Int64(1), // Required
				Items: []*string{
					aws.String("string"), // Required
					// More values...
				},
			},
			CacheBehaviors: &cloudfront.CacheBehaviors{
				Quantity: aws.Int64(1), // Required
				Items: []*cloudfront.CacheBehavior{
					{ // Required
						ForwardedValues: &cloudfront.ForwardedValues{ // Required
							Cookies: &cloudfront.CookiePreference{ // Required
								Forward: aws.String("ItemSelection"), // Required
								WhitelistedNames: &cloudfront.CookieNames{
									Quantity: aws.Int64(1), // Required
									Items: []*string{
										aws.String("string"), // Required
										// More values...
									},
								},
							},
							QueryString: aws.Bool(true), // Required
							Headers: &cloudfront.Headers{
								Quantity: aws.Int64(1), // Required
								Items: []*string{
									aws.String("string"), // Required
									// More values...
								},
							},
						},
						MinTTL:         aws.Int64(1),         // Required
						PathPattern:    aws.String("string"), // Required
						TargetOriginId: aws.String("string"), // Required
						TrustedSigners: &cloudfront.TrustedSigners{ // Required
							Enabled:  aws.Bool(true), // Required
							Quantity: aws.Int64(1),   // Required
							Items: []*string{
								aws.String("string"), // Required
								// More values...
							},
						},
						ViewerProtocolPolicy: aws.String("ViewerProtocolPolicy"), // Required
						AllowedMethods: &cloudfront.AllowedMethods{
							Items: []*string{ // Required
								aws.String("Method"), // Required
								// More values...
							},
							Quantity: aws.Int64(1), // Required
							CachedMethods: &cloudfront.CachedMethods{
								Items: []*string{ // Required
									aws.String("Method"), // Required
									// More values...
								},
								Quantity: aws.Int64(1), // Required
							},
						},
						Compress:        aws.Bool(true),
						DefaultTTL:      aws.Int64(1),
						MaxTTL:          aws.Int64(1),
						SmoothStreaming: aws.Bool(true),
					},
					// More values...
				},
			},
			CustomErrorResponses: &cloudfront.CustomErrorResponses{
				Quantity: aws.Int64(1), // Required
				Items: []*cloudfront.CustomErrorResponse{
					{ // Required
						ErrorCode:          aws.Int64(1), // Required
						ErrorCachingMinTTL: aws.Int64(1),
						ResponseCode:       aws.String("string"),
						ResponsePagePath:   aws.String("string"),
					},
					// More values...
				},
			},
			DefaultRootObject: aws.String("string"),
			Logging: &cloudfront.LoggingConfig{
				Bucket:         aws.String("string"), // Required
				Enabled:        aws.Bool(true),       // Required
				IncludeCookies: aws.Bool(true),       // Required
				Prefix:         aws.String("string"), // Required
			},
			PriceClass: aws.String("PriceClass"),
			Restrictions: &cloudfront.Restrictions{
				GeoRestriction: &cloudfront.GeoRestriction{ // Required
					Quantity:        aws.Int64(1),                     // Required
					RestrictionType: aws.String("GeoRestrictionType"), // Required
					Items: []*string{
						aws.String("string"), // Required
						// More values...
					},
				},
			},
			ViewerCertificate: &cloudfront.ViewerCertificate{
				Certificate:                  aws.String("string"),
				CertificateSource:            aws.String("CertificateSource"),
				CloudFrontDefaultCertificate: aws.Bool(true),
				IAMCertificateId:             aws.String("string"),
				MinimumProtocolVersion:       aws.String("MinimumProtocolVersion"),
				SSLSupportMethod:             aws.String("SSLSupportMethod"),
			},
			WebACLId: aws.String("string"),
		},
	}
	resp, err := svc.CreateDistribution(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFront_CreateInvalidation() {
	svc := cloudfront.New(session.New())

	params := &cloudfront.CreateInvalidationInput{
		DistributionId: aws.String("string"), // Required
		InvalidationBatch: &cloudfront.InvalidationBatch{ // Required
			CallerReference: aws.String("string"), // Required
			Paths: &cloudfront.Paths{ // Required
				Quantity: aws.Int64(1), // Required
				Items: []*string{
					aws.String("string"), // Required
					// More values...
				},
			},
		},
	}
	resp, err := svc.CreateInvalidation(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFront_CreateStreamingDistribution() {
	svc := cloudfront.New(session.New())

	params := &cloudfront.CreateStreamingDistributionInput{
		StreamingDistributionConfig: &cloudfront.StreamingDistributionConfig{ // Required
			CallerReference: aws.String("string"), // Required
			Comment:         aws.String("string"), // Required
			Enabled:         aws.Bool(true),       // Required
			S3Origin: &cloudfront.S3Origin{ // Required
				DomainName:           aws.String("string"), // Required
				OriginAccessIdentity: aws.String("string"), // Required
			},
			TrustedSigners: &cloudfront.TrustedSigners{ // Required
				Enabled:  aws.Bool(true), // Required
				Quantity: aws.Int64(1),   // Required
				Items: []*string{
					aws.String("string"), // Required
					// More values...
				},
			},
			Aliases: &cloudfront.Aliases{
				Quantity: aws.Int64(1), // Required
				Items: []*string{
					aws.String("string"), // Required
					// More values...
				},
			},
			Logging: &cloudfront.StreamingLoggingConfig{
				Bucket:  aws.String("string"), // Required
				Enabled: aws.Bool(true),       // Required
				Prefix:  aws.String("string"), // Required
			},
			PriceClass: aws.String("PriceClass"),
		},
	}
	resp, err := svc.CreateStreamingDistribution(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFront_DeleteCloudFrontOriginAccessIdentity() {
	svc := cloudfront.New(session.New())

	params := &cloudfront.DeleteCloudFrontOriginAccessIdentityInput{
		Id:      aws.String("string"), // Required
		IfMatch: aws.String("string"),
	}
	resp, err := svc.DeleteCloudFrontOriginAccessIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFront_DeleteDistribution() {
	svc := cloudfront.New(session.New())

	params := &cloudfront.DeleteDistributionInput{
		Id:      aws.String("string"), // Required
		IfMatch: aws.String("string"),
	}
	resp, err := svc.DeleteDistribution(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFront_DeleteStreamingDistribution() {
	svc := cloudfront.New(session.New())

	params := &cloudfront.DeleteStreamingDistributionInput{
		Id:      aws.String("string"), // Required
		IfMatch: aws.String("string"),
	}
	resp, err := svc.DeleteStreamingDistribution(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFront_GetCloudFrontOriginAccessIdentity() {
	svc := cloudfront.New(session.New())

	params := &cloudfront.GetCloudFrontOriginAccessIdentityInput{
		Id: aws.String("string"), // Required
	}
	resp, err := svc.GetCloudFrontOriginAccessIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFront_GetCloudFrontOriginAccessIdentityConfig() {
	svc := cloudfront.New(session.New())

	params := &cloudfront.GetCloudFrontOriginAccessIdentityConfigInput{
		Id: aws.String("string"), // Required
	}
	resp, err := svc.GetCloudFrontOriginAccessIdentityConfig(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFront_GetDistribution() {
	svc := cloudfront.New(session.New())

	params := &cloudfront.GetDistributionInput{
		Id: aws.String("string"), // Required
	}
	resp, err := svc.GetDistribution(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFront_GetDistributionConfig() {
	svc := cloudfront.New(session.New())

	params := &cloudfront.GetDistributionConfigInput{
		Id: aws.String("string"), // Required
	}
	resp, err := svc.GetDistributionConfig(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFront_GetInvalidation() {
	svc := cloudfront.New(session.New())

	params := &cloudfront.GetInvalidationInput{
		DistributionId: aws.String("string"), // Required
		Id:             aws.String("string"), // Required
	}
	resp, err := svc.GetInvalidation(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFront_GetStreamingDistribution() {
	svc := cloudfront.New(session.New())

	params := &cloudfront.GetStreamingDistributionInput{
		Id: aws.String("string"), // Required
	}
	resp, err := svc.GetStreamingDistribution(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFront_GetStreamingDistributionConfig() {
	svc := cloudfront.New(session.New())

	params := &cloudfront.GetStreamingDistributionConfigInput{
		Id: aws.String("string"), // Required
	}
	resp, err := svc.GetStreamingDistributionConfig(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFront_ListCloudFrontOriginAccessIdentities() {
	svc := cloudfront.New(session.New())

	params := &cloudfront.ListCloudFrontOriginAccessIdentitiesInput{
		Marker:   aws.String("string"),
		MaxItems: aws.Int64(1),
	}
	resp, err := svc.ListCloudFrontOriginAccessIdentities(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFront_ListDistributions() {
	svc := cloudfront.New(session.New())

	params := &cloudfront.ListDistributionsInput{
		Marker:   aws.String("string"),
		MaxItems: aws.Int64(1),
	}
	resp, err := svc.ListDistributions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFront_ListDistributionsByWebACLId() {
	svc := cloudfront.New(session.New())

	params := &cloudfront.ListDistributionsByWebACLIdInput{
		WebACLId: aws.String("string"), // Required
		Marker:   aws.String("string"),
		MaxItems: aws.Int64(1),
	}
	resp, err := svc.ListDistributionsByWebACLId(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFront_ListInvalidations() {
	svc := cloudfront.New(session.New())

	params := &cloudfront.ListInvalidationsInput{
		DistributionId: aws.String("string"), // Required
		Marker:         aws.String("string"),
		MaxItems:       aws.Int64(1),
	}
	resp, err := svc.ListInvalidations(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFront_ListStreamingDistributions() {
	svc := cloudfront.New(session.New())

	params := &cloudfront.ListStreamingDistributionsInput{
		Marker:   aws.String("string"),
		MaxItems: aws.Int64(1),
	}
	resp, err := svc.ListStreamingDistributions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFront_UpdateCloudFrontOriginAccessIdentity() {
	svc := cloudfront.New(session.New())

	params := &cloudfront.UpdateCloudFrontOriginAccessIdentityInput{
		CloudFrontOriginAccessIdentityConfig: &cloudfront.OriginAccessIdentityConfig{ // Required
			CallerReference: aws.String("string"), // Required
			Comment:         aws.String("string"), // Required
		},
		Id:      aws.String("string"), // Required
		IfMatch: aws.String("string"),
	}
	resp, err := svc.UpdateCloudFrontOriginAccessIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFront_UpdateDistribution() {
	svc := cloudfront.New(session.New())

	params := &cloudfront.UpdateDistributionInput{
		DistributionConfig: &cloudfront.DistributionConfig{ // Required
			CallerReference: aws.String("string"), // Required
			Comment:         aws.String("string"), // Required
			DefaultCacheBehavior: &cloudfront.DefaultCacheBehavior{ // Required
				ForwardedValues: &cloudfront.ForwardedValues{ // Required
					Cookies: &cloudfront.CookiePreference{ // Required
						Forward: aws.String("ItemSelection"), // Required
						WhitelistedNames: &cloudfront.CookieNames{
							Quantity: aws.Int64(1), // Required
							Items: []*string{
								aws.String("string"), // Required
								// More values...
							},
						},
					},
					QueryString: aws.Bool(true), // Required
					Headers: &cloudfront.Headers{
						Quantity: aws.Int64(1), // Required
						Items: []*string{
							aws.String("string"), // Required
							// More values...
						},
					},
				},
				MinTTL:         aws.Int64(1),         // Required
				TargetOriginId: aws.String("string"), // Required
				TrustedSigners: &cloudfront.TrustedSigners{ // Required
					Enabled:  aws.Bool(true), // Required
					Quantity: aws.Int64(1),   // Required
					Items: []*string{
						aws.String("string"), // Required
						// More values...
					},
				},
				ViewerProtocolPolicy: aws.String("ViewerProtocolPolicy"), // Required
				AllowedMethods: &cloudfront.AllowedMethods{
					Items: []*string{ // Required
						aws.String("Method"), // Required
						// More values...
					},
					Quantity: aws.Int64(1), // Required
					CachedMethods: &cloudfront.CachedMethods{
						Items: []*string{ // Required
							aws.String("Method"), // Required
							// More values...
						},
						Quantity: aws.Int64(1), // Required
					},
				},
				Compress:        aws.Bool(true),
				DefaultTTL:      aws.Int64(1),
				MaxTTL:          aws.Int64(1),
				SmoothStreaming: aws.Bool(true),
			},
			Enabled: aws.Bool(true), // Required
			Origins: &cloudfront.Origins{ // Required
				Quantity: aws.Int64(1), // Required
				Items: []*cloudfront.Origin{
					{ // Required
						DomainName: aws.String("string"), // Required
						Id:         aws.String("string"), // Required
						CustomHeaders: &cloudfront.CustomHeaders{
							Quantity: aws.Int64(1), // Required
							Items: []*cloudfront.OriginCustomHeader{
								{ // Required
									HeaderName:  aws.String("string"), // Required
									HeaderValue: aws.String("string"), // Required
								},
								// More values...
							},
						},
						CustomOriginConfig: &cloudfront.CustomOriginConfig{
							HTTPPort:             aws.Int64(1),                       // Required
							HTTPSPort:            aws.Int64(1),                       // Required
							OriginProtocolPolicy: aws.String("OriginProtocolPolicy"), // Required
							OriginSslProtocols: &cloudfront.OriginSslProtocols{
								Items: []*string{ // Required
									aws.String("SslProtocol"), // Required
									// More values...
								},
								Quantity: aws.Int64(1), // Required
							},
						},
						OriginPath: aws.String("string"),
						S3OriginConfig: &cloudfront.S3OriginConfig{
							OriginAccessIdentity: aws.String("string"), // Required
						},
					},
					// More values...
				},
			},
			Aliases: &cloudfront.Aliases{
				Quantity: aws.Int64(1), // Required
				Items: []*string{
					aws.String("string"), // Required
					// More values...
				},
			},
			CacheBehaviors: &cloudfront.CacheBehaviors{
				Quantity: aws.Int64(1), // Required
				Items: []*cloudfront.CacheBehavior{
					{ // Required
						ForwardedValues: &cloudfront.ForwardedValues{ // Required
							Cookies: &cloudfront.CookiePreference{ // Required
								Forward: aws.String("ItemSelection"), // Required
								WhitelistedNames: &cloudfront.CookieNames{
									Quantity: aws.Int64(1), // Required
									Items: []*string{
										aws.String("string"), // Required
										// More values...
									},
								},
							},
							QueryString: aws.Bool(true), // Required
							Headers: &cloudfront.Headers{
								Quantity: aws.Int64(1), // Required
								Items: []*string{
									aws.String("string"), // Required
									// More values...
								},
							},
						},
						MinTTL:         aws.Int64(1),         // Required
						PathPattern:    aws.String("string"), // Required
						TargetOriginId: aws.String("string"), // Required
						TrustedSigners: &cloudfront.TrustedSigners{ // Required
							Enabled:  aws.Bool(true), // Required
							Quantity: aws.Int64(1),   // Required
							Items: []*string{
								aws.String("string"), // Required
								// More values...
							},
						},
						ViewerProtocolPolicy: aws.String("ViewerProtocolPolicy"), // Required
						AllowedMethods: &cloudfront.AllowedMethods{
							Items: []*string{ // Required
								aws.String("Method"), // Required
								// More values...
							},
							Quantity: aws.Int64(1), // Required
							CachedMethods: &cloudfront.CachedMethods{
								Items: []*string{ // Required
									aws.String("Method"), // Required
									// More values...
								},
								Quantity: aws.Int64(1), // Required
							},
						},
						Compress:        aws.Bool(true),
						DefaultTTL:      aws.Int64(1),
						MaxTTL:          aws.Int64(1),
						SmoothStreaming: aws.Bool(true),
					},
					// More values...
				},
			},
			CustomErrorResponses: &cloudfront.CustomErrorResponses{
				Quantity: aws.Int64(1), // Required
				Items: []*cloudfront.CustomErrorResponse{
					{ // Required
						ErrorCode:          aws.Int64(1), // Required
						ErrorCachingMinTTL: aws.Int64(1),
						ResponseCode:       aws.String("string"),
						ResponsePagePath:   aws.String("string"),
					},
					// More values...
				},
			},
			DefaultRootObject: aws.String("string"),
			Logging: &cloudfront.LoggingConfig{
				Bucket:         aws.String("string"), // Required
				Enabled:        aws.Bool(true),       // Required
				IncludeCookies: aws.Bool(true),       // Required
				Prefix:         aws.String("string"), // Required
			},
			PriceClass: aws.String("PriceClass"),
			Restrictions: &cloudfront.Restrictions{
				GeoRestriction: &cloudfront.GeoRestriction{ // Required
					Quantity:        aws.Int64(1),                     // Required
					RestrictionType: aws.String("GeoRestrictionType"), // Required
					Items: []*string{
						aws.String("string"), // Required
						// More values...
					},
				},
			},
			ViewerCertificate: &cloudfront.ViewerCertificate{
				Certificate:                  aws.String("string"),
				CertificateSource:            aws.String("CertificateSource"),
				CloudFrontDefaultCertificate: aws.Bool(true),
				IAMCertificateId:             aws.String("string"),
				MinimumProtocolVersion:       aws.String("MinimumProtocolVersion"),
				SSLSupportMethod:             aws.String("SSLSupportMethod"),
			},
			WebACLId: aws.String("string"),
		},
		Id:      aws.String("string"), // Required
		IfMatch: aws.String("string"),
	}
	resp, err := svc.UpdateDistribution(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleCloudFront_UpdateStreamingDistribution() {
	svc := cloudfront.New(session.New())

	params := &cloudfront.UpdateStreamingDistributionInput{
		Id: aws.String("string"), // Required
		StreamingDistributionConfig: &cloudfront.StreamingDistributionConfig{ // Required
			CallerReference: aws.String("string"), // Required
			Comment:         aws.String("string"), // Required
			Enabled:         aws.Bool(true),       // Required
			S3Origin: &cloudfront.S3Origin{ // Required
				DomainName:           aws.String("string"), // Required
				OriginAccessIdentity: aws.String("string"), // Required
			},
			TrustedSigners: &cloudfront.TrustedSigners{ // Required
				Enabled:  aws.Bool(true), // Required
				Quantity: aws.Int64(1),   // Required
				Items: []*string{
					aws.String("string"), // Required
					// More values...
				},
			},
			Aliases: &cloudfront.Aliases{
				Quantity: aws.Int64(1), // Required
				Items: []*string{
					aws.String("string"), // Required
					// More values...
				},
			},
			Logging: &cloudfront.StreamingLoggingConfig{
				Bucket:  aws.String("string"), // Required
				Enabled: aws.Bool(true),       // Required
				Prefix:  aws.String("string"), // Required
			},
			PriceClass: aws.String("PriceClass"),
		},
		IfMatch: aws.String("string"),
	}
	resp, err := svc.UpdateStreamingDistribution(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
