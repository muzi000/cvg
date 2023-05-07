package security

import "golang.org/x/crypto/bcrypt"

// 加密密码
func HashPasswd(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return ""
	}
	return string(hash)
}

//比较密码
func ComparePasswd(plainPwd []byte, HashPwd string) bool {
	if r := bcrypt.CompareHashAndPassword([]byte(HashPwd), plainPwd); r != nil {
		return false
	}
	return true
}
