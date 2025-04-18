// Code generated by go-swagger; DO NOT EDIT.

package repository_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostRepositoriesByRepositoryNameInvalidateCacheReader is a Reader for the PostRepositoriesByRepositoryNameInvalidateCache structure.
type PostRepositoriesByRepositoryNameInvalidateCacheReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRepositoriesByRepositoryNameInvalidateCacheReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPostRepositoriesByRepositoryNameInvalidateCacheNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostRepositoriesByRepositoryNameInvalidateCacheBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostRepositoriesByRepositoryNameInvalidateCacheUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostRepositoriesByRepositoryNameInvalidateCacheForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostRepositoriesByRepositoryNameInvalidateCacheNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /v1/repositories/{repositoryName}/invalidate-cache] PostRepositoriesByRepositoryNameInvalidateCache", response, response.Code())
	}
}

// NewPostRepositoriesByRepositoryNameInvalidateCacheNoContent creates a PostRepositoriesByRepositoryNameInvalidateCacheNoContent with default headers values
func NewPostRepositoriesByRepositoryNameInvalidateCacheNoContent() *PostRepositoriesByRepositoryNameInvalidateCacheNoContent {
	return &PostRepositoriesByRepositoryNameInvalidateCacheNoContent{}
}

/*
PostRepositoriesByRepositoryNameInvalidateCacheNoContent describes a response with status code 204, with default header values.

Repository cache invalidated
*/
type PostRepositoriesByRepositoryNameInvalidateCacheNoContent struct {
}

// IsSuccess returns true when this post repositories by repository name invalidate cache no content response has a 2xx status code
func (o *PostRepositoriesByRepositoryNameInvalidateCacheNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post repositories by repository name invalidate cache no content response has a 3xx status code
func (o *PostRepositoriesByRepositoryNameInvalidateCacheNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories by repository name invalidate cache no content response has a 4xx status code
func (o *PostRepositoriesByRepositoryNameInvalidateCacheNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this post repositories by repository name invalidate cache no content response has a 5xx status code
func (o *PostRepositoriesByRepositoryNameInvalidateCacheNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories by repository name invalidate cache no content response a status code equal to that given
func (o *PostRepositoriesByRepositoryNameInvalidateCacheNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the post repositories by repository name invalidate cache no content response
func (o *PostRepositoriesByRepositoryNameInvalidateCacheNoContent) Code() int {
	return 204
}

func (o *PostRepositoriesByRepositoryNameInvalidateCacheNoContent) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/invalidate-cache][%d] postRepositoriesByRepositoryNameInvalidateCacheNoContent", 204)
}

func (o *PostRepositoriesByRepositoryNameInvalidateCacheNoContent) String() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/invalidate-cache][%d] postRepositoriesByRepositoryNameInvalidateCacheNoContent", 204)
}

func (o *PostRepositoriesByRepositoryNameInvalidateCacheNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesByRepositoryNameInvalidateCacheBadRequest creates a PostRepositoriesByRepositoryNameInvalidateCacheBadRequest with default headers values
func NewPostRepositoriesByRepositoryNameInvalidateCacheBadRequest() *PostRepositoriesByRepositoryNameInvalidateCacheBadRequest {
	return &PostRepositoriesByRepositoryNameInvalidateCacheBadRequest{}
}

/*
PostRepositoriesByRepositoryNameInvalidateCacheBadRequest describes a response with status code 400, with default header values.

Repository is not of proxy or group type
*/
type PostRepositoriesByRepositoryNameInvalidateCacheBadRequest struct {
}

// IsSuccess returns true when this post repositories by repository name invalidate cache bad request response has a 2xx status code
func (o *PostRepositoriesByRepositoryNameInvalidateCacheBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories by repository name invalidate cache bad request response has a 3xx status code
func (o *PostRepositoriesByRepositoryNameInvalidateCacheBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories by repository name invalidate cache bad request response has a 4xx status code
func (o *PostRepositoriesByRepositoryNameInvalidateCacheBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories by repository name invalidate cache bad request response has a 5xx status code
func (o *PostRepositoriesByRepositoryNameInvalidateCacheBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories by repository name invalidate cache bad request response a status code equal to that given
func (o *PostRepositoriesByRepositoryNameInvalidateCacheBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post repositories by repository name invalidate cache bad request response
func (o *PostRepositoriesByRepositoryNameInvalidateCacheBadRequest) Code() int {
	return 400
}

func (o *PostRepositoriesByRepositoryNameInvalidateCacheBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/invalidate-cache][%d] postRepositoriesByRepositoryNameInvalidateCacheBadRequest", 400)
}

func (o *PostRepositoriesByRepositoryNameInvalidateCacheBadRequest) String() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/invalidate-cache][%d] postRepositoriesByRepositoryNameInvalidateCacheBadRequest", 400)
}

func (o *PostRepositoriesByRepositoryNameInvalidateCacheBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesByRepositoryNameInvalidateCacheUnauthorized creates a PostRepositoriesByRepositoryNameInvalidateCacheUnauthorized with default headers values
func NewPostRepositoriesByRepositoryNameInvalidateCacheUnauthorized() *PostRepositoriesByRepositoryNameInvalidateCacheUnauthorized {
	return &PostRepositoriesByRepositoryNameInvalidateCacheUnauthorized{}
}

/*
PostRepositoriesByRepositoryNameInvalidateCacheUnauthorized describes a response with status code 401, with default header values.

Authentication required
*/
type PostRepositoriesByRepositoryNameInvalidateCacheUnauthorized struct {
}

// IsSuccess returns true when this post repositories by repository name invalidate cache unauthorized response has a 2xx status code
func (o *PostRepositoriesByRepositoryNameInvalidateCacheUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories by repository name invalidate cache unauthorized response has a 3xx status code
func (o *PostRepositoriesByRepositoryNameInvalidateCacheUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories by repository name invalidate cache unauthorized response has a 4xx status code
func (o *PostRepositoriesByRepositoryNameInvalidateCacheUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories by repository name invalidate cache unauthorized response has a 5xx status code
func (o *PostRepositoriesByRepositoryNameInvalidateCacheUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories by repository name invalidate cache unauthorized response a status code equal to that given
func (o *PostRepositoriesByRepositoryNameInvalidateCacheUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post repositories by repository name invalidate cache unauthorized response
func (o *PostRepositoriesByRepositoryNameInvalidateCacheUnauthorized) Code() int {
	return 401
}

func (o *PostRepositoriesByRepositoryNameInvalidateCacheUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/invalidate-cache][%d] postRepositoriesByRepositoryNameInvalidateCacheUnauthorized", 401)
}

func (o *PostRepositoriesByRepositoryNameInvalidateCacheUnauthorized) String() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/invalidate-cache][%d] postRepositoriesByRepositoryNameInvalidateCacheUnauthorized", 401)
}

func (o *PostRepositoriesByRepositoryNameInvalidateCacheUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesByRepositoryNameInvalidateCacheForbidden creates a PostRepositoriesByRepositoryNameInvalidateCacheForbidden with default headers values
func NewPostRepositoriesByRepositoryNameInvalidateCacheForbidden() *PostRepositoriesByRepositoryNameInvalidateCacheForbidden {
	return &PostRepositoriesByRepositoryNameInvalidateCacheForbidden{}
}

/*
PostRepositoriesByRepositoryNameInvalidateCacheForbidden describes a response with status code 403, with default header values.

Insufficient permissions
*/
type PostRepositoriesByRepositoryNameInvalidateCacheForbidden struct {
}

// IsSuccess returns true when this post repositories by repository name invalidate cache forbidden response has a 2xx status code
func (o *PostRepositoriesByRepositoryNameInvalidateCacheForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories by repository name invalidate cache forbidden response has a 3xx status code
func (o *PostRepositoriesByRepositoryNameInvalidateCacheForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories by repository name invalidate cache forbidden response has a 4xx status code
func (o *PostRepositoriesByRepositoryNameInvalidateCacheForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories by repository name invalidate cache forbidden response has a 5xx status code
func (o *PostRepositoriesByRepositoryNameInvalidateCacheForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories by repository name invalidate cache forbidden response a status code equal to that given
func (o *PostRepositoriesByRepositoryNameInvalidateCacheForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the post repositories by repository name invalidate cache forbidden response
func (o *PostRepositoriesByRepositoryNameInvalidateCacheForbidden) Code() int {
	return 403
}

func (o *PostRepositoriesByRepositoryNameInvalidateCacheForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/invalidate-cache][%d] postRepositoriesByRepositoryNameInvalidateCacheForbidden", 403)
}

func (o *PostRepositoriesByRepositoryNameInvalidateCacheForbidden) String() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/invalidate-cache][%d] postRepositoriesByRepositoryNameInvalidateCacheForbidden", 403)
}

func (o *PostRepositoriesByRepositoryNameInvalidateCacheForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostRepositoriesByRepositoryNameInvalidateCacheNotFound creates a PostRepositoriesByRepositoryNameInvalidateCacheNotFound with default headers values
func NewPostRepositoriesByRepositoryNameInvalidateCacheNotFound() *PostRepositoriesByRepositoryNameInvalidateCacheNotFound {
	return &PostRepositoriesByRepositoryNameInvalidateCacheNotFound{}
}

/*
PostRepositoriesByRepositoryNameInvalidateCacheNotFound describes a response with status code 404, with default header values.

Repository not found
*/
type PostRepositoriesByRepositoryNameInvalidateCacheNotFound struct {
}

// IsSuccess returns true when this post repositories by repository name invalidate cache not found response has a 2xx status code
func (o *PostRepositoriesByRepositoryNameInvalidateCacheNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post repositories by repository name invalidate cache not found response has a 3xx status code
func (o *PostRepositoriesByRepositoryNameInvalidateCacheNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post repositories by repository name invalidate cache not found response has a 4xx status code
func (o *PostRepositoriesByRepositoryNameInvalidateCacheNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post repositories by repository name invalidate cache not found response has a 5xx status code
func (o *PostRepositoriesByRepositoryNameInvalidateCacheNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post repositories by repository name invalidate cache not found response a status code equal to that given
func (o *PostRepositoriesByRepositoryNameInvalidateCacheNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the post repositories by repository name invalidate cache not found response
func (o *PostRepositoriesByRepositoryNameInvalidateCacheNotFound) Code() int {
	return 404
}

func (o *PostRepositoriesByRepositoryNameInvalidateCacheNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/invalidate-cache][%d] postRepositoriesByRepositoryNameInvalidateCacheNotFound", 404)
}

func (o *PostRepositoriesByRepositoryNameInvalidateCacheNotFound) String() string {
	return fmt.Sprintf("[POST /v1/repositories/{repositoryName}/invalidate-cache][%d] postRepositoriesByRepositoryNameInvalidateCacheNotFound", 404)
}

func (o *PostRepositoriesByRepositoryNameInvalidateCacheNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
