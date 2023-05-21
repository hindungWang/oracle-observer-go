package main

type RadioRewardShare struct {
	OwnerKey   []byte `protobuf:"bytes,1,opt,name=owner_key,json=ownerKey,proto3" json:"owner_key,omitempty"`
	HotspotKey []byte `protobuf:"bytes,2,opt,name=hotspot_key,json=hotspotKey,proto3" json:"hotspot_key,omitempty"`
	CbsdId     string `protobuf:"bytes,3,opt,name=cbsd_id,json=cbsdId,proto3" json:"cbsd_id,omitempty"`
	Amount     uint64 `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
	StartEpoch uint64 `protobuf:"varint,5,opt,name=start_epoch,json=startEpoch,proto3" json:"start_epoch,omitempty"`
	EndEpoch   uint64 `protobuf:"varint,6,opt,name=end_epoch,json=endEpoch,proto3" json:"end_epoch,omitempty"`
}

func (c *RadioRewardShare) Reset() {
}

func (c *RadioRewardShare) String() string {
	return ""
}

func (c *RadioRewardShare) ProtoMessage() {
}
