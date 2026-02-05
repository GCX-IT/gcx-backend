package services

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

type S3Service struct {
	client     *s3.Client
	bucketName string
	region     string
}

func NewS3Service() (*S3Service, error) {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	// Get AWS credentials from environment
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := os.Getenv("AWS_REGION")
	bucketName := os.Getenv("AWS_BUCKET_NAME")

	log.Printf("AWS Configuration:")
	log.Printf("  Access Key: %s", maskString(accessKey))
	log.Printf("  Secret Key: %s", maskString(secretKey))
	log.Printf("  Region: %s", region)
	log.Printf("  Bucket: %s", bucketName)

	if accessKey == "" || secretKey == "" {
		return nil, fmt.Errorf("AWS credentials not found in environment variables")
	}

	if region == "" {
		region = "eu-north-1" // Default to your bucket region
		log.Printf("Using default region: %s", region)
	}

	if bucketName == "" {
		bucketName = "gcxwebsite" // Default to your bucket name
		log.Printf("Using default bucket: %s", bucketName)
	}

	// Create AWS config with static credentials
	cfg, err := config.LoadDefaultConfig(context.Background(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			accessKey,
			secretKey,
			"", // token
		)),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %v", err)
	}

	// Create S3 client
	s3Client := s3.NewFromConfig(cfg)

	// Test S3 connection by checking if bucket exists
	_, err = s3Client.HeadBucket(context.Background(), &s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		log.Printf("Warning: Could not access S3 bucket '%s': %v", bucketName, err)
		// Don't fail here, just log the warning
	} else {
		log.Printf("Successfully connected to S3 bucket: %s", bucketName)
	}

	return &S3Service{
		client:     s3Client,
		bucketName: bucketName,
		region:     region,
	}, nil
}

// UploadFile uploads a file to S3 and returns the URL
func (s *S3Service) UploadFile(file multipart.File, header *multipart.FileHeader, folder string) (string, error) {
	// Generate unique filename
	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%s/%d_%s%s", folder, time.Now().Unix(), strings.ReplaceAll(header.Filename, " ", "_"), ext)

	log.Printf("Uploading file to S3:")
	log.Printf("  Original filename: %s", header.Filename)
	log.Printf("  S3 key: %s", filename)
	log.Printf("  Folder: %s", folder)
	log.Printf("  Size: %d bytes", header.Size)
	log.Printf("  Content-Type: %s", header.Header.Get("Content-Type"))

	// Read file content
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	log.Printf("  Read %d bytes from file", len(fileContent))

	// Upload to S3
	putObjectInput := &s3.PutObjectInput{
		Bucket:      aws.String(s.bucketName),
		Key:         aws.String(filename),
		Body:        bytes.NewReader(fileContent),
		ContentType: aws.String(header.Header.Get("Content-Type")),
		// ACL removed - bucket has ACLs disabled
	}

	log.Printf("  Uploading to bucket: %s", s.bucketName)
	log.Printf("  Uploading with key: %s", filename)

	_, err = s.client.PutObject(context.Background(), putObjectInput)
	if err != nil {
		log.Printf("  Upload failed: %v", err)
		return "", fmt.Errorf("failed to upload file to S3: %v", err)
	}

	log.Printf("  Upload successful!")

	// Return public URL
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.bucketName, s.region, filename)
	log.Printf("  Generated URL: %s", url)
	return url, nil
}

// UploadFileFromPath uploads a file from local path to S3
func (s *S3Service) UploadFileFromPath(localPath, s3Key string) (string, error) {
	file, err := os.Open(localPath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	// Get file info (not used but kept for potential future use)
	_, err = file.Stat()
	if err != nil {
		return "", fmt.Errorf("failed to get file info: %v", err)
	}

	// Upload to S3
	_, err = s.client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket:      aws.String(s.bucketName),
		Key:         aws.String(s3Key),
		Body:        file,
		ContentType: aws.String(getContentType(localPath)),
		// ACL removed - bucket has ACLs disabled
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %v", err)
	}

	// Return public URL
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.bucketName, s.region, s3Key)
	return url, nil
}

// GetFile retrieves a file from S3 and returns its content
func (s *S3Service) GetFile(s3Key string) ([]byte, string, error) {
	// Get object from S3
	result, err := s.client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(s3Key),
	})
	if err != nil {
		return nil, "", fmt.Errorf("failed to get file from S3: %v", err)
	}
	defer result.Body.Close()

	// Read file content
	fileContent, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, "", fmt.Errorf("failed to read file content: %v", err)
	}

	// Get content type from result
	contentType := ""
	if result.ContentType != nil {
		contentType = *result.ContentType
	}

	log.Printf("Successfully retrieved file from S3: %s (Content-Type: %s, Size: %d bytes)", s3Key, contentType, len(fileContent))
	return fileContent, contentType, nil
}

// DeleteFile deletes a file from S3
func (s *S3Service) DeleteFile(s3Key string) error {
	_, err := s.client.DeleteObject(context.Background(), &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucketName),
		Key:    aws.String(s3Key),
	})
	if err != nil {
		return fmt.Errorf("failed to delete file from S3: %v", err)
	}
	return nil
}

// ListFiles lists files in a specific folder
func (s *S3Service) ListFiles(prefix string) ([]string, error) {
	result, err := s.client.ListObjectsV2(context.Background(), &s3.ListObjectsV2Input{
		Bucket: aws.String(s.bucketName),
		Prefix: aws.String(prefix),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list files from S3: %v", err)
	}

	var files []string
	for _, obj := range result.Contents {
		files = append(files, *obj.Key)
	}
	return files, nil
}

// GetFileURL returns the public URL for a file
func (s *S3Service) GetFileURL(s3Key string) string {
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.bucketName, s.region, s3Key)
}

// GetPresignedURL generates a presigned URL for a file with configurable expiration
func (s *S3Service) GetPresignedURL(s3Key string, expirationMinutes int) (string, error) {
	if expirationMinutes <= 0 {
		expirationMinutes = 60 // Default to 1 hour
	}

	// Create a presigner from the client
	presigner := s3.NewPresignClient(s.client)

	// Generate presigned GET URL
	presignedRequest, err := presigner.PresignGetObject(
		context.Background(),
		&s3.GetObjectInput{
			Bucket: aws.String(s.bucketName),
			Key:    aws.String(s3Key),
		},
		s3.WithPresignExpires(time.Duration(expirationMinutes)*time.Minute),
	)

	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %v", err)
	}

	log.Printf("Generated presigned URL for %s (expires in %d minutes)", s3Key, expirationMinutes)
	return presignedRequest.URL, nil
}

// Helper function to get content type based on file extension
func getContentType(filename string) string {
	ext := strings.ToLower(filepath.Ext(filename))
	switch ext {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	case ".gif":
		return "image/gif"
	case ".pdf":
		return "application/pdf"
	case ".doc":
		return "application/msword"
	case ".docx":
		return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	case ".mp4":
		return "video/mp4"
	case ".webm":
		return "video/webm"
	default:
		return "application/octet-stream"
	}
}

// maskString masks sensitive strings for logging
func maskString(s string) string {
	if len(s) <= 4 {
		return "****"
	}
	return s[:4] + "****"
}
