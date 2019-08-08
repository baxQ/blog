package blog_test

import (
	"log"
	"os"

	"baxQ/blog"
)

func Example() {
	blog.Info("hello, world")
}

func ExampleNewLogger() {
	w := os.Stdout
	flag := log.Llongfile

	l := blog.NewWriterLogger(w, flag, 3)
	l.Info("hello, world")
}
