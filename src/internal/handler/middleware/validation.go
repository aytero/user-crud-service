package middleware

import "golang.org/x/crypto/bcrypt"

//func HashPassword(password string) (string, error) {
//    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
//    return string(bytes), err
//}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

/*
func parseDate(t *testing.T, dateString string) time.Time {
    rfc3339MilliLayout := "2006-01-02T15:04:05.999Z07:00" // layout defined with Go reference time
    parsedDate, err := time.Parse(rfc3339MilliLayout, dateString)

    require.NoError(t, err)
    return parsedDate
}
*/
