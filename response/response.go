package response

import (
	"encoding/json"
	"log"
)

//ResponseBody is a simple struct suitable for
//use as a response to a HTTP request.
type Body struct {
	Message []string
	Errors  uint
	Failed  bool
	Data    interface{}
}

//NewResponseBody creates a new ResponseBody
func New() Body {
	return Body{}
}

//Append pushes a message onto a []string of response
//messages.
func (r *Body) Append(message string) {
	r.Message = append(r.Message, message)
}

//Error increments ResponseBody.Errors to indicate
//something went wrong. Whether the error was fatal
//is up to the client implementation.
func (r *Body) Error(message string) {
	r.Errors++
	log.Printf(message)
	r.Append(message)
}

//Fatal sets ResponseBody.Failed to true and
//calls ResponseBody.Error() to set errored state.
//Fatal should always indicate an unrecoverable
//failure.
func (r *Body) Fatal(message string) {
	r.Failed = true
	r.Error(message)
}

//Send marshalls the ResponseBody into JSON and
//sends the response back to the client.
//It takes a http.ResponseWriter and a HTTP status code as arguments.
func (r *Body) Json() (string, error) {
	json, err := json.Marshal(r)
	return string(json), err
}
