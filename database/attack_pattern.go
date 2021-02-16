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

// GetAttackPatterns returns all attackPatterns.
// start, end int, order, sort string
func (d *TenDatabase) GetAttackPatterns(paging *model.Paging) []*model.AttackPattern {
	attackPatterns := []*model.AttackPattern{}
	cursor, err := d.DB.Collection("mitre_attack").
		Find(context.Background(), bson.M{"type": "attack-pattern"},
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
		attackPattern := &model.AttackPattern{}
		if err := cursor.Decode(attackPattern); err != nil {
			return nil
		}
		attackPatterns = append(attackPatterns, attackPattern)
	}

	return attackPatterns
}

// CreateAttackPattern creates a attackPattern.
func (d *TenDatabase) CreateAttackPattern(attackPattern *model.AttackPattern) *model.AttackPattern {
	_, result := d.DB.Collection("attackPatterns").
		InsertOne(context.Background(), attackPattern)
	if result != nil {
		return attackPattern
	}
	return nil
}

// GetAttackPatternByName returns the attackPattern by the given name or nil.
func (d *TenDatabase) GetAttackPatternByName(name string) *model.AttackPattern {
	var attackPattern *model.AttackPattern
	err := d.DB.Collection("attackPatterns").
		FindOne(context.Background(), bson.D{{Key: "name", Value: name}}).
		Decode(&attackPattern)
	if err != nil {
		return nil
	}
	return attackPattern
}

// GetAttackPatternByStixID returns the user by the given name or nil.
func (d *TenDatabase) GetAttackPatternByStixID(id string) *model.AttackPattern {
	var attackPattern *model.AttackPattern
	err := d.DB.Collection("attackPatterns").
		FindOne(context.Background(), bson.D{{Key: "id", Value: id}}).
		Decode(&attackPattern)
	if err != nil {
		return nil
	}
	return attackPattern
}

// GetAttackPatternByIDs returns the attackPattern by the given id or nil.
func (d *TenDatabase) GetAttackPatternByIDs(ids []primitive.ObjectID) []*model.AttackPattern {
	var attackPatterns []*model.AttackPattern
	cursor, err := d.DB.Collection("attackPatterns").
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
		attackPattern := &model.AttackPattern{}
		if err := cursor.Decode(attackPattern); err != nil {
			return nil
		}
		attackPatterns = append(attackPatterns, attackPattern)
	}

	return attackPatterns
}

// CountAttackPattern returns the attackPattern count
func (d *TenDatabase) CountAttackPattern() string {
	total, err := d.DB.Collection("attackPatterns").CountDocuments(context.Background(), bson.D{{}}, &options.CountOptions{})
	if err != nil {
		return "0"
	}
	return strconv.Itoa(int(total))
}

// DeleteAttackPatternByID deletes a attackPattern by its id.
func (d *TenDatabase) DeleteAttackPatternByID(id primitive.ObjectID) error {
	if d.CountPost(bson.D{{Key: "attackPatternId", Value: id}}) == "0" {
		_, err := d.DB.Collection("attackPatterns").DeleteOne(context.Background(), bson.D{{Key: "_id", Value: id}})
		return err
	}
	return errors.New("the current attackPattern has posts published")
}

// GetAttackPatternByID get a attackPattern by its id.
func (d *TenDatabase) GetAttackPatternByID(id primitive.ObjectID) *model.AttackPattern {
	var attackPattern *model.AttackPattern
	err := d.DB.Collection("attackPatterns").
		FindOne(context.Background(), bson.D{{Key: "_id", Value: id}}).
		Decode(&attackPattern)
	if err != nil {
		return nil
	}
	return attackPattern
}

// UpdateAttackPattern updates a attackPattern.
func (d *TenDatabase) UpdateAttackPattern(attackPattern *model.AttackPattern) *model.AttackPattern {
	result := d.DB.Collection("attackPatterns").
		FindOneAndReplace(context.Background(),
			bson.D{{Key: "_id", Value: attackPattern.ID}},
			attackPattern,
			&options.FindOneAndReplaceOptions{},
		)
	if result != nil {
		return attackPattern
	}
	return nil
}
