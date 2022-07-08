package grid

import (
	"server/database"
	"server/server/router/routes"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProjectFeature struct {
	FeatureName         string   `bson:"featureName" json:"featureName"`
	FeatureCapabilities []string `bson:"featureCapabilities" json:"featureCapabilities"`
}

type Project struct {
	database.DTO
	ProjectName        string            `bson:"projectName" json:"projectName"`
	ProjectDescription string            `bson:"projectDescription" json:"projectDescription"`
	ProjectStatus      bool              `bson:"projectStatus" json:"projectStatus"`
	ProjectTime        int               `bson:"projectTime" json:"projectTime"`
	Features           []*ProjectFeature `bson:"features" json:"features"`
}

type Proj_Emp struct {
	database.DTO
	EmpID           primitive.ObjectID `bson:"empID" json:"empID"`
	ProjID          primitive.ObjectID `bson:"projID" json:"projID"`
	TaskDescription string             `bson:"taskDescription" json:"taskDescription"`
}

type Employe struct {
	database.DTO
	EmployeName    string `bson:"employeName" json:"employeName"`
	EmployeAddress string `bson:"employeAddress" json:"employeAddress"`
}

func Test(route *routes.Route, filter bson.M, order bson.M) (*GridResponse, error) {
	var orderOpts *options.FindOptions
	results := make([]*Project, 0)

	if order != nil {
		orderOpts = options.Find().SetSort(order)
	}
	opts := options.Find()
	opts.SetLimit(int64(15000))

	err := database.GetManyDTOs(&results, "Project", filter, opts, orderOpts)
	if err != nil {
		return nil, err
	}

	result := GridResponse{
		Rows:     results,
		RowCount: len(results),
	}
	return &result, nil
}
