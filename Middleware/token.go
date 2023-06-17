package Middleware

var ValidTokens = make(map[string]uint)

// Fungsi untuk menghapus token dari daftar token yang valid
func DeleteToken(token string) {
	delete(ValidTokens, token)
}

// Fungsi untuk mendapatkan token sebelumnya berdasarkan ID pengguna
func GetToken(userID uint) string {
	for token, id := range ValidTokens {
		if id == userID {
			return token
		}
	}
	return ""
}
