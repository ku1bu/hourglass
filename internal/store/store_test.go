package store

import (
	"fmt"
	"reflect"
	"sync"
	"testing"

	"github.com/ku1bu/hourglass/internal/model"
)

func Test_memoryStore_GetGroup(t *testing.T) {
	type fields struct {
		cache sync.Map
	}
	type args struct {
		groupName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *model.Group
	}{
		{
			name: "test1",
			fields: fields{
				cache: sync.Map{},
			},
			args: args{
				groupName: "",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &memoryStore{
				cache: tt.fields.cache,
			}
			got := m.GetOrNewGroup(tt.args.groupName)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("memoryStore.GetGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLoadAndStore(t *testing.T) {
	cache := sync.Map{}

	serviceName := "aaa"

	g := &model.Group{}
	g2 := new(model.Group)

	cache.Store(serviceName, g)
	result, loaded := cache.LoadOrStore(serviceName, g2)

	fmt.Println(result == g2, loaded)
}
