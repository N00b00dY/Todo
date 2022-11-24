package main

/*
func Test_handler(t *testing.T) {
	testApp := Config{}
	testCases := []struct {
		Name     string
		Methode  string
		Url      string
		postBody map[string]interface{}
		function func(w http.ResponseWriter, r *http.Request)
	}{
		{
			Name:    "Add Todo Post",
			Methode: "Post",
			Url:     "/addToDo",
			postBody: map[string]interface{}{
				"todo": "First Test Todo",
			},
			function: testApp.AddToDo,
		},
		{
			Name:    "Check Todo",
			Methode: "Post",
			Url:     "/checkToDo",
			postBody: map[string]interface{}{
				"ID": 0,
			},
			function: testApp.CheckToDo,
		},
		{
			Name:    "Delete Todo",
			Methode: "Post",
			Url:     "/delete-todo",
			postBody: map[string]interface{}{
				"ID": 0,
			},
			function: testApp.DeleteToDo,
		},
	}

	// loop through testcases
	for _, c := range testCases {
		t.Run(c.Name, func(t *testing.T) {
			body, _ := json.Marshal(c.postBody)
			req, _ := http.NewRequest(c.Methode, c.Url, bytes.NewReader(body))
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(c.function)

			handler.ServeHTTP(rr, req)

			if rr.Code != http.StatusOK {
				t.Errorf("expected http.StatusAccepted but got %d", rr.Code)
			}
		})
	}

}
*/
