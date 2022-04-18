package waterlink

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lukasl-dev/waterlink/v2/internal/pkgerror"
	"github.com/lukasl-dev/waterlink/v2/track"
	"github.com/lukasl-dev/waterlink/v2/track/query"
	"io"
	"net/http"
	"net/url"
)

// Client is used to perform HTTP requests to the server node.
type Client struct {
	// url is the base url to perform requests to.
	url url.URL

	// creds
	creds Credentials
}

// NewClient creates a new client that performs rest actions at the given base
// url authorized by the given credentials. For instance, a base url can be:
// "http://localhost:2333".
func NewClient(baseURL string, creds Credentials) (*Client, error) {
	u, err := url.ParseRequestURI(baseURL)
	if err != nil {
		return nil, pkgerror.Wrap("client: invalid base url", err)
	}
	return &Client{url: *u, creds: creds}, nil
}

// request represents an dispatchable HTTP action.
type request struct {
	// method is the HTTP method to use.
	method string

	// path is the relative path to the endpoint.
	path string

	// params is a map of query parameters to send with.
	params map[string]string

	// body is the request body to send as JSON.
	body interface{}

	// result is the pointer to the result object. After a successful request,
	// the response body will be unmarshalled into this object.
	result interface{}

	// validate is the function to validate the response's status code.
	validate func(code int) error
}

// LoadTracks resolves the given query to a track.LoadResult and preloads the
// resulting tracks.
func (c Client) LoadTracks(q query.Query) (res *track.LoadResult, err error) {
	return res, c.dispatch("load tracks", request{
		method: http.MethodGet,
		path:   "/loadtracks",
		params: map[string]string{
			"identifier": string(q),
		},
		result:   &res,
		validate: validateStatus(http.StatusOK),
	})
}

// DecodeTrack decodes the given track ID into track.Info to gain more detailed
// information about a track.
func (c Client) DecodeTrack(trackID string) (res *track.Info, err error) {
	return res, c.dispatch("decode track", request{
		method: http.MethodGet,
		path:   "/decodetrack",
		params: map[string]string{
			"track": trackID,
		},
		result:   &res,
		validate: validateStatus(http.StatusOK),
	})
}

// DecodeTracks decodes the given track IDs into an array of track.Info to gain
// more detailed information about the given preloaded tracks.
func (c Client) DecodeTracks(trackIDs []string) (res []track.Track, err error) {
	return res, c.dispatch("decode tracks", request{
		method:   http.MethodPost,
		path:     "/decodetracks",
		body:     trackIDs,
		result:   &res,
		validate: validateStatus(http.StatusOK),
	})
}

// dispatch performs the request and wraps any errors into a package error.
func (c Client) dispatch(action string, req request) error {
	return pkgerror.Wrap(fmt.Sprintf("client: %s", action), c.doRequest(req))
}

// doRequest performs the request and returns the error.
func (c Client) doRequest(req request) error {
	r, err := c.createRequest(req)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer resp.Body.Close()

	if err = c.validate(req, resp); err != nil {
		return err
	}
	return c.unmarshalResponse(req, resp)
}

// createRequest creates a new HTTP request.
func (c Client) createRequest(req request) (*http.Request, error) {
	body, err := c.createBody(req)
	if err != nil {
		return nil, err
	}

	u := c.createURL(req)
	r, err := http.NewRequest(req.method, u.String(), body)
	if err != nil {
		return nil, err
	}
	r.Header = c.creds.header()
	r.Header.Set("Content-Type", "application/json")

	return r, nil
}

// createBody returns the request body to send to the server.
func (c Client) createBody(req request) (io.Reader, error) {
	if req.body != nil {
		b, err := json.Marshal(req.body)
		if err != nil {
			return nil, err
		}
		return bytes.NewReader(b), nil
	}
	return nil, nil
}

// createURL creates a URL from the request's path and query parameters.
func (c Client) createURL(req request) url.URL {
	u := c.url
	u.Path = req.path

	if req.params != nil {
		q := u.Query()
		for key, val := range req.params {
			q.Set(key, val)
		}
		u.RawQuery = q.Encode()
	}

	return u
}

// validate validates the response's status code against the request's
// validation function.
func (c Client) validate(req request, resp *http.Response) error {
	if req.validate == nil {
		return nil
	}
	return req.validate(resp.StatusCode)
}

// unmarshalResponse unmarshals the response body into the request's result
// object.
func (c Client) unmarshalResponse(req request, resp *http.Response) error {
	if resp.Body == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(req.result)
}

func validateStatus(allowed ...int) func(code int) error {
	set := make(map[int]bool)
	for _, code := range allowed {
		set[code] = true
	}

	return func(code int) error {
		if !set[code] {
			return fmt.Errorf("unexpected status code %d", code)
		}
		return nil
	}
}
