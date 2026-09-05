package bridges

import (
	"fmt"
	"io/fs"
	"net/http/httptest"
)

func Example() {
	err := fmt.Errorf("wrapped: %w", &fs.PathError{Op: "open", Path: "book.md", Err: fs.ErrNotExist})
	fmt.Println(pathFromError(err))
	fmt.Println(pathFromError(nil))
	response := httptest.NewRecorder()
	handler.ServeHTTP(response, httptest.NewRequest("GET", "/", nil))
	fmt.Println(response.Code)
	// Output:
	// book.md true
	//  false
	// 200
}
