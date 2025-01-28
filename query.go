package quickbooks

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

var (
	InvalidQueryResponse = errors.New("Invalid query response")
)

const (
	queryResponse = "QueryResponse"
	startPosition = "startPosition"
	maxResults    = "maxResults"
)

type Query[E Entity] struct {
	service *Service

	// TODO: support pagination
	startPosition uint
	maxResults    uint
}

func NewQuery[E Entity](service *Service) *Query[E] {
	return &Query[E]{
		service:       service,
		startPosition: 0,
		maxResults:    0,
	}
}

func (q *Query[E]) All() ([]E, error) {
	name := (*new(E)).GetName()

	path := fmt.Sprintf("company/%s/query", q.service.realmID)
	query := fmt.Sprintf("select * from %s", name)

	resp, err := q.service.request("POST", path, nil, "application/text", bytes.NewBufferString(query))
	if err != nil {
		return nil, err
	}

	parsed, err := gabs.ParseJSON(resp)
	if err != nil {
		return nil, err
	}

	queryResponse := parsed.Search(queryResponse)
	parsedStartPosition, ok := queryResponse.Search(startPosition).Data().(float64)
	if !ok {
		return nil, fmt.Errorf("%w: missing %s", InvalidQueryResponse, startPosition)
	}
	q.startPosition = uint(parsedStartPosition)

	parsedMaxResults, ok := queryResponse.Search(maxResults).Data().(float64)
	if !ok {
		return nil, fmt.Errorf("%w: missing %s", InvalidQueryResponse, maxResults)
	}
	q.maxResults = uint(parsedMaxResults)

	children := queryResponse.Search(name).Children()
	result := make([]E, len(children))

	for i, child := range children {
		data := child.String()
		err := json.Unmarshal([]byte(data), &result[i])
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}
