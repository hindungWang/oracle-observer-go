package main

import (
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	gpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"testing"
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
			if err := HandleLambdaEvent(tt.args.event); (err != nil) != tt.wantErr {
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
        "x-amz-id-2": "c7RR5oGoSyvL0XdUuBsobnk9cnw2TDZxz/p+6lSq7S9A/iV9Wnv36krO/ekd8kb4X4ljGsHxRoOcGE/W8qc+z5O1HwN/Yz7tT6SwOgzRmGw=",
        "x-amz-request-id": "97BZYCMVAMHDJSNE"
      },
      "s3": {
        "bucket": {
          "arn": "arn:aws:s3:::foundation-poc-data-requester-pays-backup",
          "name": "foundation-poc-data-requester-pays-backup",
          "ownerIdentity": {
            "principalId": "A1KNZZZ4GMSHA2"
          }
        },
        "configurationId": "e31b764b-db6f-440a-9364-8b1b0ee4feb6",
        "object": {
          "eTag": "90f62ee1e882a49290f4fe23b793150c",
          "key": "foundation-mobile-verified/radio_reward_share.1684373599999.gz",
          "sequencer": "0064664C716A0563FA",
          "size": 358061
        },
        "s3SchemaVersion": "1.0"
      },
      "userIdentity": {
        "principalId": "AWS:AIDAXVNYZ5KHHRMGQUJIJ"
      }
    }
  ]
}`
