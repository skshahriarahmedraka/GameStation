package controller

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func TestDatabaseCollections_SveltekitLogin(t *testing.T) {
	type fields struct {
		Mongo *mongo.Database
	}
	tests := []struct {
		name   string
		fields fields
		want   gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			H := &DatabaseCollections{
				Mongo: tt.fields.Mongo,
			}
			if got := H.SveltekitLogin(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SveltekitLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}
