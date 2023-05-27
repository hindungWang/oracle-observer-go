package main

import (
	"compress/gzip"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
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

var DB *gorm.DB

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags) // set flags
	initDB(os.Getenv("DATABASE_URL"))
	lambda.Start(handleLambdaEvent)
}

func handleLambdaEvent(event events.S3Event) error {
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
	case "entropy":
		handler(fileBytes, &gen.EntropyReportV1{})
	case "radio_reward_share":
		handler(fileBytes, &gen.RadioRewardShare{})
	case "iot_beacon_ingest_report":
		handler(fileBytes, &gen.LoraBeaconIngestReportV1{})
	case "iot_witness_ingest_report":
		handler(fileBytes, &gen.LoraWitnessIngestReportV1{})
	case "packetreport":
		handler(fileBytes, &gen.PacketRouterPacketReportV1{})
	case "iot_reward_share":
		handler(fileBytes, &gen.IotRewardShare{})
	//case "iot_poc":
	//	handler(fileBytes, &gen.IotPoc{})
	//case "iot_invalid_beacon":
	//	handler(fileBytes, &gen.IotInvalidBeacon{})
	//case "iot_invalid_witness":
	//	handler(fileBytes, &gen.IotInvalidWitness{})
	//case "gateway_reward_share":
	//	handler(fileBytes, &gen.GatewayRewardShare{})
	case "reward_manifest":
		handler(fileBytes, &gen.RewardManifest{})
	case "heartbeat_report":
		handler(fileBytes, &gen.CellHeartbeatIngestReportV1{})
	case "speedtest_report":
		handler(fileBytes, &gen.SpeedtestIngestReportV1{})
	case "data_transfer_session_ingest_report":
		handler(fileBytes, &gen.DataTransferSessionIngestReportV1{})
	case "validated_heartbeat":
		handler(fileBytes, &gen.Heartbeat{})
	case "speedtest_avg":
		handler(fileBytes, &gen.SpeedtestAvg{})
	case "mobile_reward_share":
		handler(fileBytes, &gen.MobileRewardShare{})
	default:
		log.Println("no match handler")
	}
	return nil
}

func handler(body []byte, row proto.Message) {
	// Obtain a message type
	for offset := 0; offset < len(body); {
		// the first 4 bytes tell you how long the message is
		messageLength := int32(body[offset])<<24 | int32(body[offset+1])<<16 | int32(body[offset+2])<<8 | int32(body[offset+3])
		bufferMessage := body[offset+4 : offset+int(messageLength)+4]
		if err := proto.Unmarshal(bufferMessage, row); err != nil {
			log.Fatal(err)
		}
		re := DB.Create(row)
		if re.Error != nil {
			log.Println(re.Error.Error())
		}

		offset += int(messageLength) + 4
	}
}

func initDB(url string) {

	var err error

	DB, err = gorm.Open(gpostgres.Open(url), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	if !DB.Migrator().HasTable(&gen.EntropyReportV1{}) {
		if err := DB.Migrator().CreateTable(&gen.EntropyReportV1{}); err != nil {
			log.Fatal(err)
		}
	}

	if !DB.Migrator().HasTable(&gen.RadioRewardShare{}) {
		if err := DB.Migrator().CreateTable(&gen.RadioRewardShare{}); err != nil {
			log.Fatal(err)
		}
	}

	if !DB.Migrator().HasTable(&gen.LoraBeaconIngestReportV1{}) {
		if err := DB.Migrator().CreateTable(&gen.LoraBeaconIngestReportV1{}); err != nil {
			log.Fatal(err)
		}
	}

	if !DB.Migrator().HasTable(&gen.LoraWitnessIngestReportV1{}) {
		if err := DB.Migrator().CreateTable(&gen.LoraWitnessIngestReportV1{}); err != nil {
			log.Fatal(err)
		}
	}

	if !DB.Migrator().HasTable(&gen.PacketRouterPacketReportV1{}) {
		if err := DB.Migrator().CreateTable(&gen.PacketRouterPacketReportV1{}); err != nil {
			log.Fatal(err)
		}
	}

	if !DB.Migrator().HasTable(&gen.IotRewardShare{}) {
		if err := DB.Migrator().CreateTable(&gen.IotRewardShare{}); err != nil {
			log.Fatal(err)
		}
	}

	if !DB.Migrator().HasTable(&gen.RewardManifest{}) {
		if err := DB.Migrator().CreateTable(&gen.RewardManifest{}); err != nil {
			log.Fatal(err)
		}
	}

	if !DB.Migrator().HasTable(&gen.CellHeartbeatIngestReportV1{}) {
		if err := DB.Migrator().CreateTable(&gen.CellHeartbeatIngestReportV1{}); err != nil {
			log.Fatal(err)
		}
	}

	if !DB.Migrator().HasTable(&gen.SpeedtestIngestReportV1{}) {
		if err := DB.Migrator().CreateTable(&gen.SpeedtestIngestReportV1{}); err != nil {
			log.Fatal(err)
		}
	}

	if !DB.Migrator().HasTable(&gen.DataTransferSessionIngestReportV1{}) {
		if err := DB.Migrator().CreateTable(&gen.DataTransferSessionIngestReportV1{}); err != nil {
			log.Fatal(err)
		}
	}

	if !DB.Migrator().HasTable(&gen.Heartbeat{}) {
		if err := DB.Migrator().CreateTable(&gen.Heartbeat{}); err != nil {
			log.Fatal(err)
		}
	}

	if !DB.Migrator().HasTable(&gen.SpeedtestAvg{}) {
		if err := DB.Migrator().CreateTable(&gen.SpeedtestAvg{}); err != nil {
			log.Fatal(err)
		}
	}
	if !DB.Migrator().HasTable(&gen.MobileRewardShare{}) {
		if err := DB.Migrator().CreateTable(&gen.MobileRewardShare{}); err != nil {
			log.Fatal(err)
		}
	}
}
