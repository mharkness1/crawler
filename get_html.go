package main

func getHTML(rawURL string) (string, error) {
	return "", nil
}

/*
Use http.Get to fetch the webpage of the rawURL
Return an error if the HTTP status code is an error-level code (400+)
Return an error if the response content-type header is not text/html
Return any other possible errors
Return the webpage's HTML if successful
*/
