package middlewares

import (
	"net/http"
)

func WithAuth(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// client := &http.Client{}
		// authURL := config.Endpoints.Auth
		// req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth/valid", authURL), nil)
		// req.Header.Set("Authorization", r.Header.Get("Authorization"))
		// if err != nil {
		// 	http.Error(w, "Error creating request", http.StatusInternalServerError)
		// 	return
		// }
		// res, err := client.Do(req)
		// if err != nil {
		//
		// 	http.Error(w, "Error sending request", http.StatusInternalServerError)
		// 	return
		// }
		// defer res.Body.Close()
		//
		// if res.StatusCode == http.StatusOK {
		// 	next.ServeHTTP(w, r)
		// } else {
		// 	http.Error(w, "Unauthorized", http.StatusUnauthorized)
		// }
		//
		next.ServeHTTP(w, r)
	})
}
