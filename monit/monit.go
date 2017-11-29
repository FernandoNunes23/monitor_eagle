package monit

import (
	"types"
	"time"
	"utils"
)

func NewMonit(name string, config types.MonitConfig) (types.Monit, bool) {
	return newMonit(name, config)
}

func newMonit(name string, config types.MonitConfig) (types.Monit, bool) {
	m := types.Monit{}
	t := time.Now()
	datetime := t.Format("2006-01-02 15:04:05")

	uniqid, err := utils.GenerateUniqid(); if err != nil {
		return m, false
	}

	m.Id = uniqid
	m.Name = name
	m.Created = datetime
	m.Config = config

	Init("/tmp/eagle")
	if Create(m) == false {
		return m , false 
	}

	return m , true
}