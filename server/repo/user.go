package repo

import (
	"fmt"

	"github.com/Tonmoy404/project/service"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
)

type UserRepo interface {
	service.UserRepo
}

type userRepo struct {
	svc       *dynamodb.DynamoDB
	tableName string
}

func NewUserRepo(svc *dynamodb.DynamoDB, tableName string) UserRepo {
	return &userRepo{
		svc:       svc,
		tableName: tableName,
	}
}

func (r *userRepo) Create(user *service.User) error {
	user.ID = uuid.New().String()

	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(user.ID),
			},
			"Username": {
				S: aws.String(user.Username),
			},
			"Email": {
				S: aws.String(user.Email),
			},
			"Password": {
				S: aws.String(user.Password),
			},
		},
	}
	_, err := r.svc.PutItem(input)
	if err != nil {
		return fmt.Errorf("cannot create item:  %v", err)
	}
	return nil
}

func (r *userRepo) Get(id string) (*service.User, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(id),
			},
		},
	}

	output, err := r.svc.GetItem(input)
	if err != nil {
		return nil, fmt.Errorf("cannot get item: %v", err)
	}
	if output.Item == nil {
		return nil, nil
	}

	pass := aws.StringValue(output.Item["Password"].S)
	fmt.Println("the password is->", pass)

	user := &service.User{
		ID:       aws.StringValue(output.Item["ID"].S),
		Username: aws.StringValue(output.Item["Username"].S),
		Email:    aws.StringValue(output.Item["Email"].S),
		Password: aws.StringValue(output.Item["Password"].S),
	}
	return user, nil
}

func (r *userRepo) Update(id string, user *service.User) error {
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(user.ID),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":e": {
				S: aws.String(user.Email),
			},
			":p": {
				S: aws.String(user.Password),
			},
		},
		UpdateExpression: aws.String("SET #e = :e, #p = :p"),
		ExpressionAttributeNames: map[string]*string{
			"#e": aws.String("Email"),
			"#p": aws.String("Password"),
		},
	}

	_, err := r.svc.UpdateItem(input)
	if err != nil {
		return fmt.Errorf("cannot update item: %v", err)
	}
	return nil
}

func (r *userRepo) Delete(id string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(r.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(id),
			},
		},
	}

	_, err := r.svc.DeleteItem(input)
	if err != nil {
		return fmt.Errorf("cannot delete item: %v", err)
	}
	return nil
}
