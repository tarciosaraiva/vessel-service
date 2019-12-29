package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/micro/go-micro"
	pb "github.com/tarciosaraiva/vessel-service/proto/vessel"
)

func loadVessels() ([]*pb.Vessel, error) {
	var vessels []*pb.Vessel
	data, err := ioutil.ReadFile("data/vessels.json")

	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &vessels)
	return vessels, err
}

func main() {
	vessels, _ := loadVessels()
	for _, v := range vessels {
		fmt.Println(v)
	}
	repo := &VesselRepository{vessels}

	srv := micro.NewService(
		micro.Name("vessel.service"),
	)

	srv.Init()

	// Register our implementation with
	pb.RegisterVesselServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
