package chroma

import (
	"encoding/json"
	"net/http"
	"strings"
)

func NewRawSql(rawSql string) *RawSql {
	return &RawSql{RawSql: rawSql}
}

func NewAddEmbedding(embeddings []interface{}) *AddEmbedding {
	return &AddEmbedding{Embeddings: embeddings, IncrementIndex: true}
}

func NewCreateCollection(name string) *CreateCollection {
	return &CreateCollection{Name: name, GetOrCreate: false}
}

func NewDeleteEmbedding() *DeleteEmbedding {
	return &DeleteEmbedding{}
}

func NewGetEmbedding() *GetEmbedding {
	return &GetEmbedding{Include: []string{"metadatas", "documents"}}
}

func NewQueryEmbedding(queryEmbeddings []interface{}) *QueryEmbedding {
	return &QueryEmbedding{
		QueryEmbeddings: queryEmbeddings,
		NResults:        10,
		Include:         []string{"metadatas", "documents", "distances"},
	}
}

func NewUpdateCollection() *UpdateCollection {
	return &UpdateCollection{}
}

func NewUpdateEmbedding(embeddings []interface{}) *UpdateEmbedding {
	return &UpdateEmbedding{Embeddings: embeddings, IncrementIndex: true}
}

func NewChromaClient(baseURL string) *ChromaClient {
	return &ChromaClient{
		BaseURL: baseURL,
		Client:  &http.Client{},
	}
}

type ChromaClient struct {
	BaseURL string
	Client  *http.Client
}

type RawSql struct {
	RawSql string `json:"raw_sql"`
}

type AddEmbedding struct {
	Embeddings     []interface{} `json:"embeddings"`
	Metadatas      interface{}   `json:"metadatas,omitempty"`
	Documents      interface{}   `json:"documents,omitempty"`
	Ids            interface{}   `json:"ids,omitempty"`
	IncrementIndex bool          `json:"increment_index,omitempty"`
}

type CreateCollection struct {
	Name        string      `json:"name"`
	Metadata    interface{} `json:"metadata,omitempty"`
	GetOrCreate bool        `json:"get_or_create,omitempty"`
}

type DeleteEmbedding struct {
	Ids           []interface{} `json:"ids,omitempty"`
	Where         interface{}   `json:"where,omitempty"`
	WhereDocument interface{}   `json:"where_document,omitempty"`
}

type GetEmbedding struct {
	Ids           []interface{} `json:"ids,omitempty"`
	Where         interface{}   `json:"where,omitempty"`
	WhereDocument interface{}   `json:"where_document,omitempty"`
	Sort          string        `json:"sort,omitempty"`
	Limit         int           `json:"limit,omitempty"`
	Offset        int           `json:"offset,omitempty"`
	Include       []string      `json:"include,omitempty"`
}

type HTTPValidationError struct {
	Detail []ValidationError `json:"detail"`
}

type QueryEmbedding struct {
	Where           interface{}   `json:"where,omitempty"`
	WhereDocument   interface{}   `json:"where_document,omitempty"`
	QueryEmbeddings []interface{} `json:"query_embeddings"`
	NResults        int           `json:"n_results,omitempty"`
	Include         []string      `json:"include,omitempty"`
}

type UpdateCollection struct {
	NewName     string      `json:"new_name,omitempty"`
	NewMetadata interface{} `json:"new_metadata,omitempty"`
}

type UpdateEmbedding struct {
	Embeddings     []interface{} `json:"embeddings"`
	Metadatas      interface{}   `json:"metadatas,omitempty"`
	Documents      interface{}   `json:"documents,omitempty"`
	Ids            interface{}   `json:"ids,omitempty"`
	IncrementIndex bool          `json:"increment_index,omitempty"`
}

type ValidationError struct {
	Loc  []interface{} `json:"loc"`
	Msg  string        `json:"msg"`
	Type string        `json:"type"`
}

func (c *ChromaClient) Root() (*http.Response, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/api/v1", nil)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *ChromaClient) Reset() (*http.Response, error) {
	req, err := http.NewRequest("POST", c.BaseURL+"/api/v1/reset", nil)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *ChromaClient) Version() (*http.Response, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/api/v1/version", nil)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *ChromaClient) Persist() (*http.Response, error) {
	req, err := http.NewRequest("POST", c.BaseURL+"/api/v1/persist", nil)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *ChromaClient) RawSql(rawSql *RawSql) (*http.Response, error) {
	reqBody, err := json.Marshal(rawSql)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.BaseURL+"/api/v1/raw_sql", strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	return c.Client.Do(req)
}

func (c *ChromaClient) ListCollections() (*http.Response, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/api/v1/collections", nil)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *ChromaClient) CreateCollection(createCollection *CreateCollection) (*http.Response, error) {
	reqBody, err := json.Marshal(createCollection)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.BaseURL+"/api/v1/collections", strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	return c.Client.Do(req)
}

func (c *ChromaClient) Add(collectionName string, addEmbedding *AddEmbedding) (*http.Response, error) {
	reqBody, err := json.Marshal(addEmbedding)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.BaseURL+"/api/v1/collections/"+collectionName+"/add", strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	return c.Client.Do(req)
}

func (c *ChromaClient) Update(collectionName string, updateEmbedding *UpdateEmbedding) (*http.Response, error) {
	reqBody, err := json.Marshal(updateEmbedding)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.BaseURL+"/api/v1/collections/"+collectionName+"/update", strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	return c.Client.Do(req)
}

func (c *ChromaClient) Get(collectionName string, getEmbedding *GetEmbedding) (*http.Response, error) {
	reqBody, err := json.Marshal(getEmbedding)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.BaseURL+"/api/v1/collections/"+collectionName+"/get", strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	return c.Client.Do(req)
}

func (c *ChromaClient) Delete(collectionName string, deleteEmbedding *DeleteEmbedding) (*http.Response, error) {
	reqBody, err := json.Marshal(deleteEmbedding)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.BaseURL+"/api/v1/collections/"+collectionName+"/delete", strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	return c.Client.Do(req)
}

func (c *ChromaClient) Count(collectionName string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/api/v1/collections/"+collectionName+"/count", nil)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *ChromaClient) GetNearestNeighbors(collectionName string, queryEmbedding *QueryEmbedding) (*http.Response, error) {
	reqBody, err := json.Marshal(queryEmbedding)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.BaseURL+"/api/v1/collections/"+collectionName+"/query", strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	return c.Client.Do(req)
}

func (c *ChromaClient) CreateIndex(collectionName string) (*http.Response, error) {
	req, err := http.NewRequest("POST", c.BaseURL+"/api/v1/collections/"+collectionName+"/create_index", nil)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *ChromaClient) GetCollection(collectionName string) (*http.Response, error) {
	req, err := http.NewRequest("GET", c.BaseURL+"/api/v1/collections/"+collectionName, nil)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *ChromaClient) UpdateCollection(collectionName string, updateCollection *UpdateCollection) (*http.Response, error) {
	reqBody, err := json.Marshal(updateCollection)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", c.BaseURL+"/api/v1/collections/"+collectionName, strings.NewReader(string(reqBody)))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	return c.Client.Do(req)
}

func (c *ChromaClient) DeleteCollection(collectionName string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", c.BaseURL+"/api/v1/collections/"+collectionName, nil)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}
