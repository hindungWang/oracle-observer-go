package main

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	gpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestHandleLambdaEvent(t *testing.T) {
	type args struct {
		event events.S3Event
	}
	var e = events.S3Event{}
	if err := json.Unmarshal([]byte(TestCase), &e); err != nil {
		t.Fatal(err)
	}

	DB, _ = gorm.Open(gpostgres.Open(os.Getenv("DATABASE_URL")), &gorm.Config{})

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "0", args: args{event: e}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := handleLambdaEvent(tt.args.event); (err != nil) != tt.wantErr {
				t.Errorf("HandleLambdaEvent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

var TestCase = `{
  "Records": [
    {
      "awsRegion": "us-east-2",
      "eventName": "ObjectCreated:Copy",
      "eventSource": "aws:s3",
      "eventTime": "2023-05-18T16:04:02.210Z",
      "eventVersion": "2.1",
      "requestParameters": {
        "sourceIPAddress": "120.230.139.90"
      },
      "responseElements": {
        "x-amz-id-2": "xxx",
        "x-amz-request-id": "xxx"
      },
      "s3": {
        "bucket": {
          "arn": "xxx",
          "name": "foundation-poc-data-requester-pays-backup",
          "ownerIdentity": {
            "principalId": "xxx"
          }
        },
        "configurationId": "xxx",
        "object": {
          "eTag": "xxx",
          "key": "foundation-mobile-verified/radio_reward_share.1684373599999.gz",
          "sequencer": "xxx",
          "size": 358061
        },
        "s3SchemaVersion": "1.0"
      },
      "userIdentity": {
        "principalId": "AWS:xxx"
      }
    }
  ]
}`

func Test_initDB(t *testing.T) {
	initDB(os.Getenv("DATABASE_URL"))
}
