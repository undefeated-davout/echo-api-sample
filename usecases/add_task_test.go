package usecases

import (
	"context"
	"reflect"
	"testing"
	"undefeated-davout/echo-api-sample/entities"
	"undefeated-davout/echo-api-sample/interface_adapters/gateways/repositories"
)

func TestAddTaskUsecase_AddTask(t *testing.T) {
	wantUID := entities.UserID(10)
	wantTitle := "test title"
	wantTask := &entities.Task{
		UserID: wantUID,
		Title:  wantTitle,
		Status: entities.TaskStatusTodo,
	}
	ctx := context.Background()
	moqDB := &DBerMock{}
	moqRepo := &TaskAdderMock{}
	moqRepo.AddTaskFunc = func(pctx context.Context, db repositories.DBer, task *entities.Task) error {
		return nil
	}

	type fields struct {
		DB   repositories.DBer
		Repo TaskAdder
	}
	type args struct {
		ctx    context.Context
		userID entities.UserID
		title  string
		status entities.TaskStatus
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entities.Task
		wantErr bool
	}{
		{
			name:    "正常登録",
			fields:  fields{DB: moqDB, Repo: moqRepo},
			args:    args{ctx: ctx, userID: wantUID, title: wantTitle, status: entities.TaskStatusTodo},
			want:    wantTask,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			a := &AddTaskUsecase{
				DB:   tt.fields.DB,
				Repo: tt.fields.Repo,
			}
			got, err := a.AddTask(tt.args.ctx, tt.args.userID, tt.args.title, tt.args.status)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTaskUsecase.AddTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddTaskUsecase.AddTask() = %v, want %v", got, tt.want)
			}
		})
	}
}
