package assistservice

import (
	"context"
	"time"

	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client         *mongo.Client
	PatientProfile *mongo.Collection
}

func New() (*Repository, error) {
	c, err := extractConfigFromEnv()
	if err != nil {
		return nil, err
	}

	client, err := mongo.NewClient(options.Client().ApplyURI(c.Uri))
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	PatientProfile := client.Database("covid19").Collection("Profile")

	return &Repository{
		client:         client,
		PatientProfile: PatientProfile,
	}, nil

}

func (a *API) NewPatientProfile(dni, phone, email string, record *PatientRecord) (*PatientProfile, error) {
	id := xid.New().String()

	profile := new(PatientProfile)
	profile.DNI = dni
	profile.Email = email
	profile.ID = id
	profile.Records = append(profile.Records, *record)
	profile.Phone = phone

	_, err := a.repo.PatientProfile.InsertOne(context.Background(), profile)

	if err != nil {
		return nil, err
	}
	return profile, nil

}

func (a *API) AddNewRecord(id string, record *PatientRecord) (*PatientProfile, error) {
	profilePatient := new(PatientProfile)

	if err := a.repo.PatientProfile.FindOne(context.Background(), bson.M{"id": id}).Decode(profilePatient); err != nil {
		return nil, err
	}

	profilePatient.Records = append(profilePatient.Records, *record)

	_, err := a.repo.PatientProfile.UpdateOne(context.Background(), bson.M{"id": id}, profilePatient)
	if err != nil {
		return nil, err
	}

	return profilePatient, nil
}

func (a *API) NewPatientRecord(id string, payload DiseasesPayload) *PatientRecord {
	newId := xid.New().String()
	created := time.Now()
	c := NewDiseasesWeight()
	result := GetStatus(payload, *c)

	return &PatientRecord{
		ID:               newId,
		CreatedAt:        created,
		PatientID:        id,
		Input:            payload,
		EvaluatedWeight:  *c,
		EvaluationResult: result,
	}

}
