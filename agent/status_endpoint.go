/*
 * Copyright (C) 2016-2018. ActionTech.
 * Based on: github.com/hashicorp/nomad, github.com/github/gh-ost .
 * License: MPL version 2: https://www.mozilla.org/en-US/MPL/2.0 .
 */

package agent

import (
	"net/http"

	"github.com/actiontech/udup/internal/models"
)

func (s *HTTPServer) StatusLeaderRequest(resp http.ResponseWriter, req *http.Request) (interface{}, error) {
	if req.Method != "GET" {
		return nil, CodedError(405, ErrInvalidMethod)
	}

	var args models.GenericRequest
	if s.parse(resp, req, &args.Region, &args.QueryOptions) {
		return nil, nil
	}

	var leader string
	if err := s.agent.RPC("Status.Leader", &args, &leader); err != nil {
		return nil, err
	}
	return leader, nil
}

func (s *HTTPServer) StatusPeersRequest(resp http.ResponseWriter, req *http.Request) (interface{}, error) {
	if req.Method != "GET" {
		return nil, CodedError(405, ErrInvalidMethod)
	}

	var args models.GenericRequest
	if s.parse(resp, req, &args.Region, &args.QueryOptions) {
		return nil, nil
	}

	var peers []string
	if err := s.agent.RPC("Status.Peers", &args, &peers); err != nil {
		return nil, err
	}
	if len(peers) == 0 {
		peers = make([]string, 0)
	}
	return peers, nil
}

func (s *HTTPServer) RegionListRequest(resp http.ResponseWriter, req *http.Request) (interface{}, error) {
	if req.Method != "GET" {
		return nil, CodedError(405, ErrInvalidMethod)
	}
	args := &models.GenericRequest{}
	if s.parse(resp, req, &args.Region, &args.QueryOptions) {
		return nil, nil
	}

	var regions []string
	if err := s.agent.RPC("Status.RegionList", &args, &regions); err != nil {
		return nil, err
	}
	if len(regions) == 0 {
		regions = make([]string, 0)
	}
	return regions, nil
}
