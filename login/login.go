package Login

func Login(username string, password string) (string, *User) {
	if username == "" || password == "" {

		return "", nil
	}
	for _, user := range UserStorage {
		if user.Username == username && user.Password == password {

			return "ok", &user
		}
	}

	return "", nil
}
