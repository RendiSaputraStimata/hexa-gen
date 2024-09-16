package out

type {{.Domain}}Response struct {
	//Define properties in here
}


type {{.Domain}}CreatedResponse = Response[{{.Domain}}Response]
type {{.Domain}}StandardResponse = Response[{{.Domain}}Response]
type {{.Domain}}ListResponse = Response[[]{{.Domain}}Response]
