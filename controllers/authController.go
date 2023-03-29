package controllers

import (
	"context"
	"ecommerce/database"
	"ecommerce/helpers"
	"ecommerce/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserCollection *mongo.Collection = database.UserData(database.Client, "Users")
var Validate = validator.New()

func Signin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := UserCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(*foundUser.Password))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		token, refreshToken, _ = helpers.UpdateAllTokens(*foundUser.Email, foundUser.First_Name, *foundUser.Last_Name, foundUser.User_ID)

		c.JSON(http.StatusFound, foundUser)

	}
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationError := Validate.Struct(user)

		if validationError != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationError})
			return
		}

		count, err := UserCollection.CountDocuments(ctx, bson.M{"$or": []bson.M{bson.M{"email": user.Email}, bson.M{"phone": user.Phone}}})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return

		}
		if count > 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
			return
		}

		password, err := bcrypt.GenerateFromPassword([]byte(*user.Password), 14)

		if err != nil {
			log.Panic(err)
		}

		hashedPassword := string(password)
		user.Password = &hashedPassword

		user.Created_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_At, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		user.ID = primitive.NewObjectID()
		user.User_ID = user.ID.Hex()

		token, refreshToken, _ = helpers.GenerateToken(*user.Email, user.First_Name, *user.Last_Name, user.User_ID)
		user.Token = &token
		user.Refresh_Token = &refreshToken

		user.User_Cart = make([]models.Product, 0)
		user.User_Cart = make([]models.Address, 0)

		user.Order_Status = make([]models.Order, 0)

		_, err = UserCollection.InsertOne(ctx, user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		}

		c.JSON(http.StatusCreated, "Successfully Signed Up!!")
	}
}
