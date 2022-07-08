package main

import (
	"fmt"
	"server/database"
	"server/database/repositories"
	"server/utils"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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

func newProject() *Project {
	return &Project{
		DTO: database.DTO{
			CollectionName: "Project",
		},
		ProjectName:        utils.RandomString(utils.RandomNumber(20, 50)),
		ProjectDescription: utils.RandomString(utils.RandomNumber(100, 1000)),
		ProjectStatus:      utils.RandomNumber(0, 10)%2 == 0,
		ProjectTime:        utils.RandomNumber(0, 25602),
		Features:           getProjectFeatures(utils.RandomNumber(50, 250)),
	}
}

func newProjectFeature() *ProjectFeature {
	cl := utils.RandomNumber(0, 10)
	caps := make([]string, cl)
	for i := 0; i < cl; i++ {
		caps[i] = utils.RandomString(utils.RandomNumber(20, 50))
	}
	return &ProjectFeature{
		FeatureName:         utils.RandomString(utils.RandomNumber(10, 20)),
		FeatureCapabilities: caps,
	}
}

func getProjectFeatures(qt int) []*ProjectFeature {
	b := make([]*ProjectFeature, qt)
	for i := 0; i < qt; i++ {
		b[i] = newProjectFeature()
	}
	return b
}

func newProjEmp(empID primitive.ObjectID, projID primitive.ObjectID) *Proj_Emp {
	return &Proj_Emp{
		DTO: database.DTO{
			CollectionName: "Proj_Emp",
		},
		EmpID:           empID,
		ProjID:          projID,
		TaskDescription: utils.RandomString(utils.RandomNumber(1000, 5000)),
	}
}

func newEmploye() *Employe {
	return &Employe{
		DTO: database.DTO{
			CollectionName: "Employe",
		},
		EmployeName:    utils.RandomString(utils.RandomNumber(10, 20)),
		EmployeAddress: utils.RandomString(utils.RandomNumber(30, 50)),
	}
}

func randomObjects(qt int) {
	for i := 0; i < qt; i++ {
		empQt := utils.RandomNumber(2, 10)
		projQt := utils.RandomNumber(2, 10)

		empIDS := make([]*primitive.ObjectID, empQt)
		projIDS := make([]*primitive.ObjectID, projQt)

		for i := 0; i < empQt; i++ {
			emp := newEmploye()
			empID, err := emp.Save(emp)
			if err != nil {
				panic(err)
			}
			empIDS[i] = empID
		}

		for i := 0; i < projQt; i++ {
			proj := newProject()
			projID, err := proj.Save(proj)
			if err != nil {
				panic(err)
			}
			projIDS[i] = projID
		}

		for i := 0; i < projQt; i++ {
			for j := 0; j < empQt; j++ {
				pe := newProjEmp(*empIDS[j], *projIDS[i])
				_, err := pe.Save(pe)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

func randomLogs(l int) {
	database.Init()

	start := time.Now()
	fmt.Println("Allocating memory")
	buff := make([]interface{}, l)
	fmt.Println("Filling buffer memory")
	for i := 0; i < l; i++ {
		buff[i] = repositories.LogMessageDTO{
			Level:      utils.RandomString(utils.RandomNumber(6, 12)),
			Message:    utils.RandomString(utils.RandomNumber(10, 30)),
			Detail:     utils.RandomString(utils.RandomNumber(8, 16)),
			Location:   utils.RandomString(utils.RandomNumber(10, 30)),
			StackTrace: utils.RandomString(utils.RandomNumber(20, 60)),
			Time:       utils.RandomString(10),
		}
	}
	fmt.Println("Buffer Filled")
	database.Database.Collection("LogMessages").InsertMany(*database.ClientContext, buff)
	end := time.Now()
	dur := end.Sub(start)

	fmt.Println("Inserted in: " + strconv.FormatInt(dur.Milliseconds(), 10))
}
