package main

import (
	"fmt"
	"google.golang.org/protobuf/runtime/protoimpl"
	"os"
	"reflect"
)

func PatchTailwindConfig(pathToPatch string) error {
	patch, err := os.ReadFile("./classes.json")
	if err != nil {
		return err
	}

	err = os.WriteFile(pathToPatch, patch, 0644)
	if err != nil {
		return err
	}

	return nil
}

type GetOneDistribution_Response_Level_Gates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GateId string `protobuf:"bytes,1,opt,name=gate_id,json=gateId,proto3" json:"gate_id,omitempty" form:"gate_id,omitempty"`
	// Active status in current group set
	IsActive bool `protobuf:"varint,2,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty" form:"is_active,omitempty"`
	// Percent of distribution
	Percentage int32 `protobuf:"varint,3,opt,name=percentage,proto3" json:"percentage,omitempty" form:"percentage,omitempty"`
}

type GetListGate_Response_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of current level
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" form:"name,omitempty"`
	// Description of current level
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" form:"description,omitempty"`
	// Active status in current level
	IsActive bool                                       `protobuf:"varint,3,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty" form:"is_active,omitempty"`
	Gates    []*GetOneDistribution_Response_Level_Gates `protobuf:"bytes,4,rep,name=gates,proto3" json:"gates,omitempty" form:"gates,omitempty"`
}

func main() {
	service := &GetListGate_Response_Item{
		Name:     "asdasdasd",
		IsActive: true,
		Gates: []*GetOneDistribution_Response_Level_Gates{
			{
				GateId:     "123",
				IsActive:   true,
				Percentage: 0,
			},
		},
	}
	v := reflect.ValueOf(service)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		panic("Not a struct!")
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if t.Field(i).IsExported() {
			fmt.Println(t.Field(i).Name,
				t.Field(i).Tag.Get("json"),
				v.Field(i).Interface())

		}
	}
}
