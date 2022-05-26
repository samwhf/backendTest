package tests

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/samwhf/backendTest/config"
	"github.com/samwhf/backendTest/database/postgres"
	"github.com/samwhf/backendTest/objects"
	us "github.com/samwhf/backendTest/services/user"
)

var (
	envFile  = "../.env.test"
	ctx      = context.Background()
	database *postgres.Client
)

func init() {
	// 读取测试环境配置文件
	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal("error: ", err)
	}
	// configuration
	configuration, err := config.LoadEnvFile()
	if err != nil {
		log.Println("error: ", err)
	}
	log.Println("Configuration Variables:", configuration)
	// db
	log.Println("Connecting to postgres")
	database = postgres.New(*configuration)

	err = us.CreateTable(context.Background())
	if err != nil {
		log.Fatal("error: ", err)
	}
}

var tests = []struct {
	name    string
	input   *objects.User
	wantErr bool
}{
	{
		name: "success case",
		input: &objects.User{
			Name:        "sanzhang",
			DateOfBirth: time.Now(),
			Address:     "xxx",
			Description: "test",
		},
		wantErr: false,
	},
	{
		name: "wrong case",
		input: &objects.User{
			Name:        "sanzhang",
			DateOfBirth: time.Now(),
			Address:     "xxx",
			Description: "test",
		},
		wantErr: true,
	},
}

func TestDBCreate(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := us.Create(ctx, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
			} else {
				tt.input.ID = id
			}
		})
	}
}

func TestDBGet(t *testing.T) {
	successId := ""
	for _, tt := range tests {
		if tt.input.ID != "" {
			successId = tt.input.ID
		}
	}
	fmt.Println(successId)
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "success case",
			input:   successId,
			wantErr: false,
		},
		{
			name:    "success case",
			input:   "nothing",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := us.Get(ctx, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBUpdate(t *testing.T) {
	successId := ""
	for _, tt := range tests {
		if tt.input.ID != "" {
			successId = tt.input.ID
		}
	}
	fmt.Println(successId)
	user, err := us.Get(ctx, successId)
	if err != nil {
		t.Errorf("Get() error = %v", err)
	}
	user.Description = "updated"
	tests := []struct {
		name    string
		input   *objects.User
		wantErr bool
	}{
		{
			name:    "success case",
			input:   user,
			wantErr: false,
		},
		{
			name:    "wrong case",
			input:   &objects.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := us.Update(ctx, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDBDelete(t *testing.T) {
	successId := ""
	for _, tt := range tests {
		if tt.input.ID != "" {
			successId = tt.input.ID
		}
	}
	fmt.Println(successId)

	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "success case",
			input:   successId,
			wantErr: false,
		},
		{
			name:    "success case",
			input:   "nothing",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := us.Delete(ctx, tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func cleanuptest() {
	defer database.Close()
}

func TestClean(t *testing.T) {
	t.Cleanup(cleanuptest)
	t.Log("testing done.")
}
