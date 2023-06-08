package api

import (
	"io"

	"github.com/yomorun/yomo/core"
)

type Source interface {
	// Close will close the connection to YoMo-Zipper.
	Close() error
	// Connect to YoMo-Zipper.
	Connect() error
	// Write the data to directed downstream.
	Write(tag uint32, data []byte) error
	// Broadcast broadcast the data to all downstream.
	Broadcast(tag uint32, data []byte) error
	// SetErrorHandler set the error handler function when server error occurs
	SetErrorHandler(fn func(err error))
	// [Experimental] SetReceiveHandler set the observe handler function
	SetReceiveHandler(fn func(tag uint32, data []byte))
	// WriteFrom writes data from reader to source.
	WriteFrom(io.Reader) error
}

func NewSource(name string, zipperAddr string) Source

// StreamFunction defines serverless streaming functions.
type StreamFunction interface {
	// SetObserveDataTags set the data tag list that will be observed
	// Deprecated: use yomo.WithObserveDataTags instead
	SetObserveDataTags(tag ...uint32)
	// SetHandler set the handler function, which accept the raw bytes data and return the tag & response
	SetHandler(fn core.AsyncHandler) error
	// SetErrorHandler set the error handler function when server error occurs
	SetErrorHandler(fn func(err error))
	// SetPipeHandler set the pipe handler function
	SetPipeHandler(fn core.PipeHandler) error
	// Connect create a connection to the zipper
	Connect() error
	// Close will close the connection
	Close() error
	// SetStreamHandler set a handler function for handling stream data.
	SetStreamHandler(func(io.Reader)) error
}

// NewStreamFunction returns a new sfn runtime,
// The sourceID is be used to identify the source of stream.
func NewStreamFunction(
	name string,
	sourceID string, // The sourceID should be an optional params.
) StreamFunction
