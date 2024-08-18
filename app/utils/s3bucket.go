// utils/s3bucket.go

package utils

import (
	"golang-template-api-service/app/config"
	"golang-template-api-service/app/internal/dto/common"
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func UploadFileToS3(fileKey string, file *multipart.FileHeader, objectKey string) (*common.UploadResponse, error) {
	s3bucket := config.LoadViperConfig().AmazonStorage
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(s3bucket.Region),
		Credentials: credentials.NewStaticCredentials(s3bucket.AccessKey, s3bucket.SecretKey, ""),
	})

	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)

	// Open the uploaded file
	openedFile, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer openedFile.Close()

	var buf bytes.Buffer

	// Create a sizeWriter to capture bytes written and calculate the size
	sw := &sizeWriter{Writer: &buf}

	// Use TeeReader with the sizeWriter
	if _, err := io.Copy(sw, openedFile); err != nil {
		return nil, err
	}

	// Validate file size
	if sw.size > int64(s3bucket.FileSizeLimit) {
		return nil, fmt.Errorf("File size exceeds the maximum allowed size")
	}
	// Validate file extension
	fileExtension := filepath.Ext(file.Filename)
	validExtension := false
	for _, ext := range s3bucket.AllowedExtensions {
		if ext == fileExtension {
			validExtension = true
			break
		}
	}
	if !validExtension {
		return nil, fmt.Errorf("Invalid file extension")
	}

	fullObjectKey := fmt.Sprintf("%s%s%s", objectKey, fileKey, fileExtension)

	// Create a buffer to read the file content
	_, err = io.Copy(&buf, openedFile)
	if err != nil {
		return nil, err
	}
	bufLen := int64(buf.Len())

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(s3bucket.BucketName),
		Key:           aws.String(fullObjectKey),
		Body:          bytes.NewReader(buf.Bytes()), // Use bytes.NewReader for seeking
		ContentLength: aws.Int64(bufLen),
	})

	if err != nil {
		return nil, err
	}

	objectURL := fmt.Sprintf("%s", fullObjectKey)
	//objectURL := fmt.Sprintf("%s%s", s3bucket.BucketUrl, objectKey)

	// Construct the custom response
	response := &common.UploadResponse{
		Message:   "File uploaded successfully to S3",
		ObjectURL: objectURL,
	}

	return response, nil
}

func UploadByteoS3(fileKey string, fileData []byte, objectKey string) (*common.UploadResponse, error) {
	s3bucket := config.LoadViperConfig().AmazonStorage
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(s3bucket.Region),
		Credentials: credentials.NewStaticCredentials(s3bucket.AccessKey, s3bucket.SecretKey, ""),
	})

	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)

	fullObjectKey := fmt.Sprintf("%s%s%s", objectKey, fileKey, ".pdf")

	// Create a buffer to read the file content
	buf := bytes.NewReader(fileData)
	bufLen := int64(len(fileData))

	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(s3bucket.BucketName),
		Key:           aws.String(fullObjectKey),
		Body:          buf,
		ContentLength: aws.Int64(bufLen),
	})

	if err != nil {
		return nil, err
	}

	objectURL := fmt.Sprintf("%s", fullObjectKey)

	// Construct the custom response
	response := &common.UploadResponse{
		Message:   "File uploaded successfully to S3",
		ObjectURL: objectURL,
	}

	return response, nil
}

func GeneratePresignedURL(fileName string) (*common.UploadResponse, error) {
	s3bucket := config.LoadViperConfig().AmazonStorage

	if fileName == "" {
		response := &common.UploadResponse{
			ObjectURL: "",
		}

		return response, nil
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(s3bucket.Region),
		Credentials: credentials.NewStaticCredentials(s3bucket.AccessKey, s3bucket.SecretKey, ""),
	})

	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)

	// Generate a pre-signed URL for the object key
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s3bucket.BucketName),
		Key:    aws.String(fileName),
	})
	objectURL, err := req.Presign(time.Hour * 4) // Pre-sign URL for 4 hours
	if err != nil {
		return nil, err
	}
	response := &common.UploadResponse{
		Message:   "File recieved successfully from S3",
		ObjectURL: objectURL,
	}

	return response, nil
}

// Create a custom writer to capture the bytes written and calculate the size
type sizeWriter struct {
	io.Writer
	size int64
}

func (sw *sizeWriter) Write(p []byte) (n int, err error) {
	n, err = sw.Writer.Write(p)
	sw.size += int64(n)
	return n, err
}

func DuplicateFileOnS3(fileName string, objectKey string, newFileName string) (*common.UploadResponse, error) {
	s3bucket := config.LoadViperConfig().AmazonStorage

	if fileName == "" {
		response := &common.UploadResponse{
			ObjectURL: "",
		}
		return response, nil
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(s3bucket.Region),
		Credentials: credentials.NewStaticCredentials(s3bucket.AccessKey, s3bucket.SecretKey, ""),
	})
	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)

	// Copy the object with the same key but a new file name

	fileExtension := filepath.Ext(fileName)
	fullObjectKey := fmt.Sprintf("%s%s%s", objectKey, newFileName, fileExtension)

	_, err = svc.CopyObject(&s3.CopyObjectInput{
		Bucket:     aws.String(s3bucket.BucketName),
		CopySource: aws.String(fmt.Sprintf("%s/%s", s3bucket.BucketName, fileName)),
		Key:        aws.String(fullObjectKey),
	})

	if err != nil {
		return nil, err
	}

	// Generate a pre-signed URL for the new object key
	objectURL := fmt.Sprintf("%s", fullObjectKey)

	response := &common.UploadResponse{
		Message:   "File duplicated successfully on S3",
		ObjectURL: objectURL,
	}

	return response, nil
}
