package http

type Hanlder interface {
	ServeHTTP(w ResponseWriter, r *Request)
}

func ListenAndServe(address string, h Hanlder) error
