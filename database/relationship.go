package database

import (
	"context"
	"errors"
	"strconv"

	"github.com/lotteryjs/ten-minutes-app/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetRelationships returns all relationships.
// start, end int, order, sort string
func (d *TenDatabase) GetRelationships(paging *model.Paging) []*model.Relationship {
	relationships := []*model.Relationship{}
	cursor, err := d.DB.Collection("mitre_attack").
		Find(context.Background(), bson.M{"type": "relationship"},
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
		relationship := &model.Relationship{}
		if err := cursor.Decode(relationship); err != nil {
			return nil
		}
		relationships = append(relationships, relationship)
	}

	return relationships
}

// CreateRelationship creates a relationship.
func (d *TenDatabase) CreateRelationship(relationship *model.Relationship) *model.Relationship {
	_, result := d.DB.Collection("relationships").
		InsertOne(context.Background(), relationship)
	if result != nil {
		return relationship
	}
	return nil
}

// GetRelationshipByName returns the relationship by the given name or nil.
func (d *TenDatabase) GetRelationshipByName(name string) *model.Relationship {
	var relationship *model.Relationship
	err := d.DB.Collection("relationships").
		FindOne(context.Background(), bson.D{{Key: "name", Value: name}}).
		Decode(&relationship)
	if err != nil {
		return nil
	}
	return relationship
}

// GetRelationshipByStixID returns the user by the given name or nil.
func (d *TenDatabase) GetRelationshipByStixID(id string) *model.Relationship {
	var relationship *model.Relationship
	err := d.DB.Collection("relationships").
		FindOne(context.Background(), bson.D{{Key: "id", Value: id}}).
		Decode(&relationship)
	if err != nil {
		return nil
	}
	return relationship
}

// GetRelationshipByIDs returns the relationship by the given id or nil.
func (d *TenDatabase) GetRelationshipByIDs(ids []primitive.ObjectID) []*model.Relationship {
	var relationships []*model.Relationship
	cursor, err := d.DB.Collection("relationships").
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
		relationship := &model.Relationship{}
		if err := cursor.Decode(relationship); err != nil {
			return nil
		}
		relationships = append(relationships, relationship)
	}

	return relationships
}

// CountRelationship returns the relationship count
func (d *TenDatabase) CountRelationship() string {
	total, err := d.DB.Collection("relationships").CountDocuments(context.Background(), bson.D{{}}, &options.CountOptions{})
	if err != nil {
		return "0"
	}
	return strconv.Itoa(int(total))
}

// DeleteRelationshipByID deletes a relationship by its id.
func (d *TenDatabase) DeleteRelationshipByID(id primitive.ObjectID) error {
	if d.CountPost(bson.D{{Key: "relationshipId", Value: id}}) == "0" {
		_, err := d.DB.Collection("relationships").DeleteOne(context.Background(), bson.D{{Key: "_id", Value: id}})
		return err
	}
	return errors.New("the current relationship has posts published")
}

// GetRelationshipByID get a relationship by its id.
func (d *TenDatabase) GetRelationshipByID(id primitive.ObjectID) *model.Relationship {
	var relationship *model.Relationship
	err := d.DB.Collection("relationships").
		FindOne(context.Background(), bson.D{{Key: "_id", Value: id}}).
		Decode(&relationship)
	if err != nil {
		return nil
	}
	return relationship
}

// UpdateRelationship updates a relationship.
func (d *TenDatabase) UpdateRelationship(relationship *model.Relationship) *model.Relationship {
	result := d.DB.Collection("relationships").
		FindOneAndReplace(context.Background(),
			bson.D{{Key: "_id", Value: relationship.ID}},
			relationship,
			&options.FindOneAndReplaceOptions{},
		)
	if result != nil {
		return relationship
	}
	return nil
}
