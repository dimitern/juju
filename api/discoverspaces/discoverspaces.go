// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package discoverspaces

import (
	"github.com/juju/errors"
	"github.com/juju/loggo"

	"github.com/juju/juju/api/base"
	"github.com/juju/juju/api/common"
	"github.com/juju/juju/apiserver/params"
)

var logger = loggo.GetLogger("juju.api.discoverspaces")

const discoverspacesFacade = "DiscoverSpaces"

// API provides access to the InstancePoller API facade.
type API struct {
	*common.EnvironWatcher
	facade base.FacadeCaller
}

// NewAPI creates a new facade.
func NewAPI(caller base.APICaller) *API {
	if caller == nil {
		panic("caller is nil")
	}
	facadeCaller := base.NewFacadeCaller(caller, discoverspacesFacade)
	return &API{
		EnvironWatcher: common.NewEnvironWatcher(facadeCaller),
		facade:         facadeCaller,
	}
}

func (api *API) ListSpaces() (params.DiscoverSpacesResults, error) {
	var result params.DiscoverSpacesResults
	if err := api.facade.FacadeCall("ListSpaces", nil, &result); err != nil {
		return result, errors.Trace(err)
	}
	return result, nil
}

func (api *API) AddSubnets(args params.AddSubnetsParams) (params.ErrorResults, error) {
	var result params.ErrorResults
	err := api.facade.FacadeCall("AddSubnets", args, &result)
	if err != nil {
		return result, errors.Trace(err)
	}
	return result, nil
}

func (api *API) CreateSpaces(args params.CreateSpacesParams) (results params.ErrorResults, err error) {
	var result params.ErrorResults
	err = api.facade.FacadeCall("CreateSpaces", args, &result)
	if err != nil {
		return result, errors.Trace(err)
	}
	return result, nil
}
