package util
import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("iknow")        //jwt密钥

/*在jwt中添加的自定义用户信息，有用户的ID和手机号,也可以加一些其它信息*/
type Claims struct {
	ID uint `json:"id"`

	jwt.StandardClaims
}

/*生成Token*/
func GeterateToken(id uint) (string,error) {
	claims := Claims{
		ID: id,

		StandardClaims:jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 60*3,   //过期时间 3min
			Issuer : "iknow",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret) //该方法内部生成签名字符串，再用于获取完整、已签名的token
	if err!=nil{fmt.Printf("%s",err)}
	fmt.Println(token)

	return token, err
}

/*校验和解析token*/
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err!=nil{fmt.Printf("%s",err)}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	fmt.Println(tokenClaims.Claims.(*Claims).ID)

	return nil, err
}