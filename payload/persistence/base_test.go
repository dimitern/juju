// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package persistence

import (
	"fmt"

	gitjujutesting "github.com/juju/testing"
	"github.com/juju/utils"
	gc "gopkg.in/check.v1"
	"gopkg.in/juju/charm.v6-unstable"

	"github.com/juju/juju/payload"
	"github.com/juju/juju/testing"
)

type BaseSuite struct {
	testing.BaseSuite

	Stub    *gitjujutesting.Stub
	State   *fakeStatePersistence
	Unit    string
	Machine string
}

func (s *BaseSuite) SetUpTest(c *gc.C) {
	s.BaseSuite.SetUpTest(c)

	s.Stub = &gitjujutesting.Stub{}
	s.State = &fakeStatePersistence{Stub: s.Stub}
	s.Unit = "a-unit/0"
	s.Machine = "0"
}

type PayloadDoc payloadDoc

func (doc PayloadDoc) convert() *payloadDoc {
	return (*payloadDoc)(&doc)
}

func (s *BaseSuite) NewDoc(id string, pl payload.FullPayloadInfo) *payloadDoc {
	return &payloadDoc{
		DocID:     "payload#" + s.Unit + "#" + pl.Name,
		UnitID:    s.Unit,
		Name:      pl.Name,
		MachineID: pl.Machine,
		StateID:   id,
		Type:      pl.Type,
		State:     pl.Status,
		Labels:    append([]string{}, pl.Labels...),
		RawID:     pl.ID,
	}
}

func (s *BaseSuite) SetDoc(id string, pl payload.FullPayloadInfo) *payloadDoc {
	payloadDoc := s.NewDoc(id, pl)
	s.State.SetDocs(payloadDoc)
	return payloadDoc
}

func (s *BaseSuite) RemoveDoc(name string) {
	docID := "payload#" + s.Unit + "#" + name
	delete(s.State.docs, docID)
}

func (s *BaseSuite) NewPersistence() *Persistence {
	return NewPersistence(s.State)
}

func (s *BaseSuite) SetUnit(id string) {
	s.Unit = id
}

func (s *BaseSuite) NewPayloads(pType string, ids ...string) []payload.FullPayloadInfo {
	var payloads []payload.FullPayloadInfo
	for _, id := range ids {
		pl := s.NewPayload(pType, id)
		payloads = append(payloads, pl)
	}
	return payloads
}

func (s *BaseSuite) NewPayload(pType string, id string) payload.FullPayloadInfo {
	name, pluginID := payload.ParseID(id)
	if pluginID == "" {
		pluginID = fmt.Sprintf("%s-%s", name, utils.MustNewUUID())
	}

	return payload.FullPayloadInfo{
		Payload: payload.Payload{
			PayloadClass: charm.PayloadClass{
				Name: name,
				Type: pType,
			},
			ID:     pluginID,
			Status: "running",
			Unit:   s.Unit,
		},
		Machine: s.Machine,
	}
}
