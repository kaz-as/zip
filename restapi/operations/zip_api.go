// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewZipAPI creates a new Zip instance
func NewZipAPI(spec *loads.Document) *ZipAPI {
	return &ZipAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		BinConsumer:  runtime.ByteStreamConsumer(),
		JSONConsumer: runtime.JSONConsumer(),

		BinProducer:  runtime.ByteStreamProducer(),
		JSONProducer: runtime.JSONProducer(),

		CheckChunksHandler: CheckChunksHandlerFunc(func(params CheckChunksParams) middleware.Responder {
			return middleware.NotImplemented("operation CheckChunks has not yet been implemented")
		}),
		CreateArchiveHandler: CreateArchiveHandlerFunc(func(params CreateArchiveParams) middleware.Responder {
			return middleware.NotImplemented("operation CreateArchive has not yet been implemented")
		}),
		GetFilesHandler: GetFilesHandlerFunc(func(params GetFilesParams) middleware.Responder {
			return middleware.NotImplemented("operation GetFiles has not yet been implemented")
		}),
		InitUploadArchiveHandler: InitUploadArchiveHandlerFunc(func(params InitUploadArchiveParams) middleware.Responder {
			return middleware.NotImplemented("operation InitUploadArchive has not yet been implemented")
		}),
		IsCompletedHandler: IsCompletedHandlerFunc(func(params IsCompletedParams) middleware.Responder {
			return middleware.NotImplemented("operation IsCompleted has not yet been implemented")
		}),
		UploadChunkHandler: UploadChunkHandlerFunc(func(params UploadChunkParams) middleware.Responder {
			return middleware.NotImplemented("operation UploadChunk has not yet been implemented")
		}),
	}
}

/*ZipAPI API for zip processing */
type ZipAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// BinConsumer registers a consumer for the following mime types:
	//   - application/octet-stream
	BinConsumer runtime.Consumer
	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// BinProducer registers a producer for the following mime types:
	//   - application/octet-stream
	BinProducer runtime.Producer
	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// CheckChunksHandler sets the operation handler for the check chunks operation
	CheckChunksHandler CheckChunksHandler
	// CreateArchiveHandler sets the operation handler for the create archive operation
	CreateArchiveHandler CreateArchiveHandler
	// GetFilesHandler sets the operation handler for the get files operation
	GetFilesHandler GetFilesHandler
	// InitUploadArchiveHandler sets the operation handler for the init upload archive operation
	InitUploadArchiveHandler InitUploadArchiveHandler
	// IsCompletedHandler sets the operation handler for the is completed operation
	IsCompletedHandler IsCompletedHandler
	// UploadChunkHandler sets the operation handler for the upload chunk operation
	UploadChunkHandler UploadChunkHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *ZipAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *ZipAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *ZipAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ZipAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *ZipAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *ZipAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ZipAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ZipAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ZipAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ZipAPI
func (o *ZipAPI) Validate() error {
	var unregistered []string

	if o.BinConsumer == nil {
		unregistered = append(unregistered, "BinConsumer")
	}
	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.BinProducer == nil {
		unregistered = append(unregistered, "BinProducer")
	}
	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.CheckChunksHandler == nil {
		unregistered = append(unregistered, "CheckChunksHandler")
	}
	if o.CreateArchiveHandler == nil {
		unregistered = append(unregistered, "CreateArchiveHandler")
	}
	if o.GetFilesHandler == nil {
		unregistered = append(unregistered, "GetFilesHandler")
	}
	if o.InitUploadArchiveHandler == nil {
		unregistered = append(unregistered, "InitUploadArchiveHandler")
	}
	if o.IsCompletedHandler == nil {
		unregistered = append(unregistered, "IsCompletedHandler")
	}
	if o.UploadChunkHandler == nil {
		unregistered = append(unregistered, "UploadChunkHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ZipAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ZipAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	return nil
}

// Authorizer returns the registered authorizer
func (o *ZipAPI) Authorizer() runtime.Authorizer {
	return nil
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *ZipAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/octet-stream":
			result["application/octet-stream"] = o.BinConsumer
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *ZipAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/octet-stream":
			result["application/octet-stream"] = o.BinProducer
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *ZipAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the zip API
func (o *ZipAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *ZipAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/files/upload"] = NewCheckChunks(o.context, o.CheckChunksHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/files/zip"] = NewCreateArchive(o.context, o.CreateArchiveHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/files"] = NewGetFiles(o.context, o.GetFilesHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/files"] = NewInitUploadArchive(o.context, o.InitUploadArchiveHandler)
	if o.handlers["HEAD"] == nil {
		o.handlers["HEAD"] = make(map[string]http.Handler)
	}
	o.handlers["HEAD"]["/files/upload"] = NewIsCompleted(o.context, o.IsCompletedHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/files/upload"] = NewUploadChunk(o.context, o.UploadChunkHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ZipAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *ZipAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *ZipAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *ZipAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *ZipAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
