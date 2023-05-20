package user

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"rest/conn"
	"rest/models"
	"time"
)

const UserCollection string = "user"

var (
	errNotExist        = errors.New("users are not exist")
	errInvalid         = errors.New("invalid ID")
	errInvalidBody     = errors.New("invalid request body")
	errInsertionFailed = errors.New("error in the user insertion")
	errUpdateFailed    = errors.New("error in the user update")
	errDeletionField   = errors.New("error in the user deletion")
)

func GetAllUser(ctx *gin.Context) {
	db := conn.GetMongoDb()
	users := models.Users{}
	err := db.C(UserCollection).Find(bson.M{}).All(&users)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errNotExist.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "users": &users})
}

func GetUser(ctx *gin.Context) {
	var id bson.ObjectId = bson.ObjectIdHex(ctx.Param("id"))
	user, err := models.UserInfo(id, UserCollection)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "msg": errInvalid.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "user": &user})
}

func CreateUser(ctx *gin.Context) {
	db := conn.GetMongoDb()
	user := models.User{}

	err := ctx.Bind(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "msg": errInsertionFailed.Error()})
		return
	}
	user.ID = bson.NewObjectId()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	err = db.C(UserCollection).Insert(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "msg": errInsertionFailed.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "user": &user})
}

func UpdateUser(ctx *gin.Context) {
	db := conn.GetMongoDb()

	var id bson.ObjectId = bson.ObjectIdHex(ctx.Param("id"))

	existingUser, err := models.UserInfo(id, UserCollection)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "msg": errInvalid.Error()})
		return
	}

	err = ctx.Bind(&existingUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "msg": errInvalidBody.Error()})
		return
	}

	existingUser.ID = id
	existingUser.UpdatedAt = time.Now()
	err = db.C(UserCollection).Update(bson.M{"_id": &id}, existingUser)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "msg": errUpdateFailed.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "user": &existingUser})
}

func DeleteUser(ctx *gin.Context) {
	db := conn.GetMongoDb()
	var id bson.ObjectId = bson.ObjectIdHex(ctx.Param("id"))
	err := db.C(UserCollection).Remove(bson.M{"_id": id})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "failed", "msg": errDeletionField.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "msg": "User deleted successfully"})
}
