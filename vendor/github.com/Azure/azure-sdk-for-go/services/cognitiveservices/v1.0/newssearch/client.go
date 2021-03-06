// Package newssearch implements the Azure ARM Newssearch service API version 1.0.
//
// The News Search API lets you send a search query to Bing and get back a list of news that are relevant to the search
// query. This section provides technical details about the query parameters and headers that you use to request news
// and the JSON response objects that contain them. For examples that show how to make requests, see [Searching the web
// for news](https://docs.microsoft.com/en-us/azure/cognitive-services/bing-news-search/search-the-web).
package newssearch

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
)

const (
	// DefaultBaseURI is the default URI used for the service Newssearch
	DefaultBaseURI = "https://api.cognitive.microsoft.com/bing/v7.0"
)

// BaseClient is the base client for Newssearch.
type BaseClient struct {
	autorest.Client
	BaseURI string
}

// New creates an instance of the BaseClient client.
func New() BaseClient {
	return NewWithBaseURI(DefaultBaseURI)
}

// NewWithBaseURI creates an instance of the BaseClient client.
func NewWithBaseURI(baseURI string) BaseClient {
	return BaseClient{
		Client:  autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI: baseURI,
	}
}
