package subject

import (
	"io"
	"net/http"
	"os"
	"sync"
)

var uploadM sync.Mutex

// Subject have profiling targert processes.
type Subject interface {
	// UploadAndZip
	UploadAndZip(w http.ResponseWriter, r *http.Request)
}

type subject struct {
}

// NewSubject generate subject instance.
func NewSubject() Subject {
	subject := subject{}
	return &subject
}

func (s *subject) UploadAndZip(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Open("./image.jpg")
	defer file.Close()

	writer, _ := os.Create("./out")
	defer writer.Close()
	buf3(file, writer)
}

func buf3(r io.Reader, w *os.File) {
	buf := make([]byte, 1000)
	buf2 := make([]byte, 100000000)
	for {
		n, _ := r.Read(buf)
		r.Read(buf2)
		// fmt.Println(string(buf[:n]), n)
		if n == 0 {
			break
		}
		w.Write(buf[:n])
	}

	io.Copy(w, r)
}
