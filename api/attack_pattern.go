package api

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lotteryjs/ten-minutes-app/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// The AttackPatternDatabase interface for encapsulating database access.
type AttackPatternDatabase interface {
	GetAttackPatternByIDs(ids []primitive.ObjectID) []*model.AttackPattern
	GetAttackPatternByID(id primitive.ObjectID) *model.AttackPattern
	DeleteAttackPatternByID(id primitive.ObjectID) error
	CreateAttackPattern(attackPattern *model.AttackPattern) *model.AttackPattern
	GetAttackPatterns(paging *model.Paging) []*model.AttackPattern
	UpdateAttackPattern(attackPattern *model.AttackPattern) *model.AttackPattern
	CountAttackPattern() string
}

// The AttackPatternAPI provides handlers for managing attackPatterns.
type AttackPatternAPI struct {
	DB AttackPatternDatabase
}

// GetAttackPatternByIDs returns the attackPattern by id
func (a *AttackPatternAPI) GetAttackPatternByIDs(ctx *gin.Context) {
	withIDs(ctx, "id", func(ids []primitive.ObjectID) {
		ctx.JSON(200, a.DB.GetAttackPatternByIDs(ids))
	})
}

// DeleteAttackPatternByID deletes the attackPattern by id
func (a *AttackPatternAPI) DeleteAttackPatternByID(ctx *gin.Context) {
	withID(ctx, "id", func(id primitive.ObjectID) {
		if err := a.DB.DeleteAttackPatternByID(id); err == nil {
			ctx.JSON(200, http.StatusOK)
		} else {
			if err != nil {
				ctx.AbortWithError(500, err)
			} else {
				ctx.AbortWithError(404, errors.New("attackPattern does not exist"))
			}
		}
	})
}

// GetAttackPatterns returns all the attackPatterns
// _end=5&_order=DESC&_sort=id&_start=0 adapt react-admin
func (a *AttackPatternAPI) GetAttackPatterns(ctx *gin.Context) {
	var (
		start int64
		end   int64
		sort  string
		order int
	)
	id := ctx.DefaultQuery("id", "")
	if id != "" {
		a.GetAttackPatternByIDs(ctx)
		return
	}
	start, _ = strconv.ParseInt(ctx.DefaultQuery("_start", "0"), 10, 64)
	end, _ = strconv.ParseInt(ctx.DefaultQuery("_end", "10"), 10, 64)
	sort = ctx.DefaultQuery("_sort", "_id")
	order = 1

	if sort == "id" {
		sort = "_id"
	}

	if ctx.DefaultQuery("_order", "DESC") == "DESC" {
		order = -1
	}

	limit := end - start
	attackPatterns := a.DB.GetAttackPatterns(
		&model.Paging{
			Skip:      &start,
			Limit:     &limit,
			SortKey:   sort,
			SortVal:   order,
			Condition: nil,
		})

	ctx.Header("X-Total-Count", a.DB.CountAttackPattern())
	ctx.JSON(200, attackPatterns)
}

// CreateAttackPattern creates a attackPattern.
func (a *AttackPatternAPI) CreateAttackPattern(ctx *gin.Context) {
	var attackPattern = model.AttackPattern{}
	if err := ctx.ShouldBind(&attackPattern); err == nil {
		if result := a.DB.CreateAttackPattern(attackPattern.New()); result != nil {
			ctx.JSON(201, result)
		} else {
			ctx.AbortWithError(500, errors.New("CreateAttackPattern error"))
		}
	} else {
		ctx.AbortWithError(500, errors.New("ShouldBind error"))
	}
}

// GetAttackPatternByID returns the attackPattern by id
func (a *AttackPatternAPI) GetAttackPatternByID(ctx *gin.Context) {
	withID(ctx, "id", func(id primitive.ObjectID) {
		if attackPattern := a.DB.GetAttackPatternByID(id); attackPattern != nil {
			ctx.JSON(200, attackPattern)
		} else {
			ctx.AbortWithError(404, errors.New("attackPattern does not exist"))
		}
	})
}

// UpdateAttackPatternByID returns the attackPattern by id
func (a *AttackPatternAPI) UpdateAttackPatternByID(ctx *gin.Context) {
	withID(ctx, "id", func(id primitive.ObjectID) {
		var attackPattern = model.AttackPattern{}
		abort := errors.New("attackPattern does not exist")
		if err := ctx.ShouldBind(&attackPattern); err == nil {
			if result := a.DB.UpdateAttackPattern(&attackPattern); result != nil {
				ctx.JSON(200, result)
			} else {
				ctx.AbortWithError(404, abort)
			}
		} else {
			ctx.AbortWithError(404, abort)
		}
	})
}
