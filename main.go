package main

import (
	"compress/gzip"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang/protobuf/proto"
	_ "github.com/lib/pq"
	gpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"io"
	"log"
	"os"
	"strings"

	"github.com/hindungWang/oracle-observer-go/proto/gen"
)

func HandleLambdaEvent(event events.S3Event) error {
	if len(event.Records) == 0 {
		log.Println("nil event records")
		return nil
	}

	region := event.Records[0].AWSRegion
	bucket := event.Records[0].S3.Bucket.Name
	key := event.Records[0].S3.Object.Key

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		log.Fatal(err)
	}

	svc := s3.New(sess)
	input := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	result, err := svc.GetObject(input)
	if err != nil {
		log.Fatal(err)
	}

	gz, err := gzip.NewReader(result.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer gz.Close()

	fileBytes, err := io.ReadAll(gz)
	if err != nil {
		log.Fatal(err)
	}

	i := strings.Index(key, "/") + 1
	prefix := strings.Split(key[i:], ".")[0]

	switch prefix {
	case "radio_reward_share":
		// Obtain a message type
		for offset := 0; offset < len(fileBytes); {
			// the first 4 bytes tell you how long the message is
			messageLength := int32(fileBytes[offset])<<24 | int32(fileBytes[offset+1])<<16 | int32(fileBytes[offset+2])<<8 | int32(fileBytes[offset+3])
			bufferMessage := fileBytes[offset+4 : offset+int(messageLength)+4]

			row := &gen.RadioRewardShare{}
			if err = proto.Unmarshal(bufferMessage, row); err != nil {
				log.Fatal(err)
			}
			re := DB.Create(row)
			if re.Error != nil {
				log.Println(re.Error.Error())
			}

			offset += int(messageLength) + 4
		}
	case "gateway_reward_share":

	}
	return nil
}

var DB *gorm.DB

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags) // set flags
	db := os.Getenv("DATABASE_URL")
	initDB(db)
	lambda.Start(HandleLambdaEvent)
}

func initDB(url string) {
	m, err := migrate.New("file://migrations", url)
	if err != nil {
		log.Fatal(err)
	}
	if err = m.Up(); err != nil {
		if !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal(err)
		}
	}

	DB, err = gorm.Open(gpostgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}
