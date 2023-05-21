package main

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

func HandleLambdaEvent(event events.S3Event) error {
	// 下载文件存入内存

	// unzip
	fileBytes, err := ioutil.ReadFile("./radio_reward_share.1684373599999")
	if err != nil {
		log.Fatal(err)
	}

	// Obtain a message type
	row := &RadioRewardShare{}
	for offset := 0; offset < len(fileBytes); {
		// the first 4 bytes tell you how long the message is
		messageLength := int32(fileBytes[offset])<<24 | int32(fileBytes[offset+1])<<16 | int32(fileBytes[offset+2])<<8 | int32(fileBytes[offset+3])
		bufferMessage := fileBytes[offset+4 : offset+int(messageLength)+4]

		if err = proto.Unmarshal(bufferMessage, row); err != nil {
			log.Fatal(err)
		}

		log.Println(row.CbsdId, row.EndEpoch, row.StartEpoch, row.Amount, row.HotspotKey)

		// 根据类型插入不同的数据库表

		offset += int(messageLength) + 4
	}
	return nil
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags) // set flags

	lambda.Start(HandleLambdaEvent)
}
