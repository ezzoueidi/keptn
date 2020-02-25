// Code generated by go-swagger; DO NOT EDIT.

package project_resource

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// GetProjectProjectNameResourceURL generates an URL for the get project project name resource operation
type GetProjectProjectNameResourceURL struct {
	ProjectName string

	DisableUpstreamSync *bool
	NextPageKey         *string
	PageSize            *int64

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetProjectProjectNameResourceURL) WithBasePath(bp string) *GetProjectProjectNameResourceURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *GetProjectProjectNameResourceURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *GetProjectProjectNameResourceURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/project/{projectName}/resource"

	projectName := o.ProjectName
	if projectName != "" {
		_path = strings.Replace(_path, "{projectName}", projectName, -1)
	} else {
		return nil, errors.New("projectName is required on GetProjectProjectNameResourceURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var disableUpstreamSyncQ string
	if o.DisableUpstreamSync != nil {
		disableUpstreamSyncQ = swag.FormatBool(*o.DisableUpstreamSync)
	}
	if disableUpstreamSyncQ != "" {
		qs.Set("disableUpstreamSync", disableUpstreamSyncQ)
	}

	var nextPageKeyQ string
	if o.NextPageKey != nil {
		nextPageKeyQ = *o.NextPageKey
	}
	if nextPageKeyQ != "" {
		qs.Set("nextPageKey", nextPageKeyQ)
	}

	var pageSizeQ string
	if o.PageSize != nil {
		pageSizeQ = swag.FormatInt64(*o.PageSize)
	}
	if pageSizeQ != "" {
		qs.Set("pageSize", pageSizeQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *GetProjectProjectNameResourceURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *GetProjectProjectNameResourceURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *GetProjectProjectNameResourceURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on GetProjectProjectNameResourceURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on GetProjectProjectNameResourceURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *GetProjectProjectNameResourceURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
