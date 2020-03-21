// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "context"

// WorkbookRangeRequestBuilder is request builder for WorkbookRange
type WorkbookRangeRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookRangeRequest
func (b *WorkbookRangeRequestBuilder) Request() *WorkbookRangeRequest {
	return &WorkbookRangeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookRangeRequest is request for WorkbookRange
type WorkbookRangeRequest struct{ BaseRequest }

// Get performs GET request for WorkbookRange
func (r *WorkbookRangeRequest) Get(ctx context.Context) (resObj *WorkbookRange, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkbookRange
func (r *WorkbookRangeRequest) Update(ctx context.Context, reqObj *WorkbookRange) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkbookRange
func (r *WorkbookRangeRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// Format is navigation property
func (b *WorkbookRangeRequestBuilder) Format() *WorkbookRangeFormatRequestBuilder {
	bb := &WorkbookRangeFormatRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/format"
	return bb
}

// Sort is navigation property
func (b *WorkbookRangeRequestBuilder) Sort() *WorkbookRangeSortRequestBuilder {
	bb := &WorkbookRangeSortRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/sort"
	return bb
}

// Worksheet is navigation property
func (b *WorkbookRangeRequestBuilder) Worksheet() *WorkbookWorksheetRequestBuilder {
	bb := &WorkbookWorksheetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/worksheet"
	return bb
}