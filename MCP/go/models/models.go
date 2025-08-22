package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Episode represents the Episode schema from the OpenAPI specification
type Episode struct {
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type,omitempty"` // The type of resource. This is always `episodes`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id,omitempty"`
}

// Image represents the Image schema from the OpenAPI specification
type Image struct {
	TypeField string `json:"type,omitempty"` // The type of resource. This is always `images`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id,omitempty"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
}

// Links represents the Links schema from the OpenAPI specification
type Links struct {
	First string `json:"first,omitempty"` // A URL to retrieve the first page of data keeping the provided page size
	Last string `json:"last,omitempty"` // A URL to retrieve the last page of data keeping the provided page size
	Next string `json:"next,omitempty"` // A URL to retrieve the next page of data keeping the provided page size
	Prev string `json:"prev,omitempty"` // A URL to retrieve the previous page of keeping using the provided page size
	Self string `json:"self,omitempty"` // A URL to retrieve the collection as the primary data
}

// MediaAsset represents the MediaAsset schema from the OpenAPI specification
type MediaAsset struct {
	TypeField string `json:"type,omitempty"` // The type of resource. This is always `media_assets`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id,omitempty"`
}

// Person represents the Person schema from the OpenAPI specification
type Person struct {
	Id string `json:"id,omitempty"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type,omitempty"` // The type of resource. This is always `people`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// ResourceIdentifier represents the ResourceIdentifier schema from the OpenAPI specification
type ResourceIdentifier struct {
	Id string `json:"id,omitempty"` // The ID of the resource
	TypeField string `json:"type,omitempty"` // The type of the resource
}

// Classification represents the Classification schema from the OpenAPI specification
type Classification struct {
	TypeField string `json:"type,omitempty"` // The type of resource. This is always `classifications`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id,omitempty"`
}

// Network represents the Network schema from the OpenAPI specification
type Network struct {
	Id string `json:"id,omitempty"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type,omitempty"` // The type of resource. This is always `networks`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// ClassificationInclusion represents the ClassificationInclusion schema from the OpenAPI specification
type ClassificationInclusion struct {
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type,omitempty"` // The type of resource. This is always `classification_inclusion`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id,omitempty"`
}

// Credit represents the Credit schema from the OpenAPI specification
type Credit struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id,omitempty"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type,omitempty"` // The type of resource. This is always `credits`
}

// ResourceLink represents the ResourceLink schema from the OpenAPI specification
type ResourceLink struct {
	Self string `json:"self,omitempty"` // A URL to retrieve the resource as the primary data
}

// Season represents the Season schema from the OpenAPI specification
type Season struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id,omitempty"`
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type,omitempty"` // The type of resource. This is always `seasons`
}

// Series represents the Series schema from the OpenAPI specification
type Series struct {
	Relationships map[string]interface{} `json:"relationships,omitempty"`
	TypeField string `json:"type,omitempty"` // The type of resource. This is always `series`
	Attributes map[string]interface{} `json:"attributes,omitempty"`
	Id string `json:"id,omitempty"`
}
