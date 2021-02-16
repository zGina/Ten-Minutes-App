package database

import (
	"context"
	"errors"
	"strconv"

	"github.com/lotteryjs/ten-minutes-app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetUsers returns all users.
// start, end int, order, sort string
func (d *TenDatabase) GetUsers(paging *model.Paging) []*model.User {
	users := []*model.User{}
	cursor, err := d.DB.Collection("users").
		Find(context.Background(), bson.D{},
			&options.FindOptions{
				Skip:  paging.Skip,
				Sort:  bson.D{bson.E{Key: paging.SortKey, Value: paging.SortVal}},
				Limit: paging.Limit,
			})
	if err != nil {
		return nil
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		user := &model.User{}
		if err := cursor.Decode(user); err != nil {
			return nil
		}
		users = append(users, user)
	}

	return users
}

// CreateUser creates a user.
func (d *TenDatabase) CreateUser(user *model.User) *model.User {
	_, result := d.DB.Collection("users").
		InsertOne(context.Background(), user)

	if result != nil {
		return user
	}

	return nil
}

// GetUserByName returns the user by the given name or nil.
func (d *TenDatabase) GetUserByName(name string) *model.User {
	var user *model.User
	err := d.DB.Collection("users").
		FindOne(context.Background(), bson.D{{Key: "name", Value: name}}).
		Decode(&user)
	if err != nil {
		return nil
	}
	return user
}

// GetUserByIDs returns the user by the given id or nil.
func (d *TenDatabase) GetUserByIDs(ids []string) []*model.User {
	var users []*model.User
	cursor, err := d.DB.Collection("users").
		Find(context.Background(), bson.D{{
			Key: "_id",
			Value: bson.D{{
				Key:   "$in",
				Value: ids,
			}},
		}})
	if err != nil {
		return nil
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		user := &model.User{}
		if err := cursor.Decode(user); err != nil {
			return nil
		}
		users = append(users, user)
	}

	return users
}

// CountUser returns the user count
func (d *TenDatabase) CountUser() string {
	total, err := d.DB.Collection("users").CountDocuments(context.Background(), bson.D{{}}, &options.CountOptions{})
	if err != nil {
		return "0"
	}
	return strconv.Itoa(int(total))
}

// DeleteUserByID deletes a user by its id.
func (d *TenDatabase) DeleteUserByID(id string) error {
	if d.CountPost(bson.D{{Key: "userId", Value: id}}) == "0" {
		_, err := d.DB.Collection("users").DeleteOne(context.Background(), bson.D{{Key: "_id", Value: id}})
		return err
	}
	return errors.New("the current user has posts published")
}

// GetUserByID get a user by its id.
func (d *TenDatabase) GetUserByID(id string) *model.User {
	var user *model.User
	err := d.DB.Collection("users").
		FindOne(context.Background(), bson.D{{Key: "_id", Value: id}}).
		Decode(&user)
	if err != nil {
		return nil
	}
	return user
}

// UpdateUser updates a user.
func (d *TenDatabase) UpdateUser(user *model.User) *model.User {
	result := d.DB.Collection("users").
		FindOneAndReplace(context.Background(),
			bson.D{{Key: "_id", Value: user.ID}},
			user,
			&options.FindOneAndReplaceOptions{},
		)
	if result != nil {
		return user
	}
	return nil
}
