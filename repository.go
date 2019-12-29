package main

import (
	"errors"

	pb "github.com/tarciosaraiva/vessel-service/proto/vessel"
)

type repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

// VesselRepository struct representing a response from the repo
type VesselRepository struct {
	vessels []*pb.Vessel
}

// FindAvailable ensures that we can find a vessel given a spec
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}

	return nil, errors.New("No vessel found for the spec")
}
