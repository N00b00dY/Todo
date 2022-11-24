package main

//Test not working because i dont know how to fake the db-service host

// func Test_handler(t *testing.T) {
// 	testApp := Config{}

// 	req, _ := http.NewRequest("GET", "/", nil)
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(testApp.Home)

// 	handler.ServeHTTP(rr, req)

// 	if rr.Code != http.StatusOK {
// 		t.Errorf("expected http.StatusAccepted but got %d", rr.Code)
// 	}
// }
