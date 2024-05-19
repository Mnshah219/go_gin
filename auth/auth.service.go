package auth

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mnshah219/go_gin/auth/dto"
	"github.com/mnshah219/go_gin/auth/utils"
)

func signup(ctx *gin.Context) {
	var user dto.SignupRequestDto
	if err := ctx.ShouldBindJSON(&user); err != nil {
		// https://blog.logrocket.com/gin-binding-in-go-a-tutorial-with-examples/
		// https://www.convictional.com/blog/gin-validation
		// https://pkg.go.dev/github.com/go-playground/validator/v10#section-readme
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	existingUser := findOneUser(map[string]any{"email": user.Email})
	if existingUser.ID != "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Err: User with given email exists!"})
		return
	}
	// hash passwd
	// for more std implementation https://snyk.io/blog/secure-password-hashing-in-go/
	salt := os.Getenv("SALT")
	user.Password = utils.GenerateHash([]byte(user.Password), []byte(salt))
	_, err := createUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

func login(ctx *gin.Context) {
	var cred dto.LoginRequestDto
	if err := ctx.ShouldBindJSON(&cred); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := findOneUser(map[string]any{"email": cred.Email})
	if user.ID == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Err: User with given email does not exists!"})
		return
	}
	salt := os.Getenv("SALT")
	err := utils.Compare(user.Password, []byte(salt), []byte(cred.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Err: Incorrect credentials!"})
		return
	}
	jwt := utils.IssueJWT(user.ID)
	ctx.JSON(http.StatusOK, dto.LoginResponseDto{Token: jwt})
}
