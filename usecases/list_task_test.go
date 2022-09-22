package usecases

import (
	"context"
	"reflect"
	"testing"
	"undefeated-davout/echo-api-sample/entities"
	"undefeated-davout/echo-api-sample/interface_adapters/gateways/repositories"
)

func TestListTaskUsecase_ListTasks(t *testing.T) {
	wantUID := entities.UserID(10)
	wantTasks := []entities.Task{{
		UserID: entities.UserID(10),
		Title:  "1",
		Status: entities.TaskStatusTodo,
	}, {
		UserID: entities.UserID(10),
		Title:  "2",
		Status: entities.TaskStatusTodo,
	}}
	ctx := context.Background()
	moqDB := &DBerMock{}
	moqRepo := &TaskListerMock{}
	moqRepo.ListTasksFunc = func(ctx context.Context, db repositories.DBer, id entities.UserID) ([]entities.Task, error) {
		if id == wantUID {
			return wantTasks, nil
		}
		return []entities.Task{}, nil
	}

	type fields struct {
		DB   repositories.DBer
		Repo TaskLister
	}
	type args struct {
		ctx    context.Context
		userID entities.UserID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []entities.Task
		wantErr bool
	}{
		{
			name:    "正常取得",
			fields:  fields{DB: moqDB, Repo: moqRepo},
			args:    args{ctx: ctx, userID: wantUID},
			want:    wantTasks,
			wantErr: false,
		},
		{
			name:    "正常取得：該当なし",
			fields:  fields{DB: moqDB, Repo: moqRepo},
			args:    args{ctx: ctx, userID: entities.UserID(99)},
			want:    []entities.Task{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &ListTaskUsecase{
				DB:   tt.fields.DB,
				Repo: tt.fields.Repo,
			}
			got, err := l.ListTasks(tt.args.ctx, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListTaskUsecase.ListTasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListTaskUsecase.ListTasks() = %v, want %v", got, tt.want)
			}
		})
	}
}
