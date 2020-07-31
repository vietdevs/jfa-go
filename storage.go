package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type Storage struct {
	timePattern                                                                  string
	invite_path, emails_path, policy_path, configuration_path, displayprefs_path string
	invites                                                                      Invites
	emails, policy, configuration, displayprefs                                  map[string]interface{}
}

// timePattern: %Y-%m-%dT%H:%M:%S.%f

type Invite struct {
	Created       time.Time                  `json:"created"`
	NoLimit       bool                       `json:"no-limit"`
	RemainingUses int                        `json:"remaining-uses"`
	ValidTill     time.Time                  `json:"valid_till"`
	Email         string                     `json:"email"`
	UsedBy        [][]string                 `json:"used-by"`
	Notify        map[string]map[string]bool `json:"notify"`
}

type Invites map[string]Invite

func (st *Storage) loadInvites() error {
	return loadJSON(st.invite_path, &st.invites)
}

func (st *Storage) storeInvites() error {
	return storeJSON(st.invite_path, st.invites)
}

func (st *Storage) loadEmails() error {
	return loadJSON(st.emails_path, &st.emails)
}

func (st *Storage) storeEmails() error {
	return storeJSON(st.emails_path, st.emails)
}

func (st *Storage) loadPolicy() error {
	return loadJSON(st.policy_path, &st.policy)
}

func (st *Storage) storePolicy() error {
	return storeJSON(st.policy_path, st.policy)
}

func (st *Storage) loadConfiguration() error {
	return loadJSON(st.configuration_path, &st.configuration)
}

func (st *Storage) storeConfiguration() error {
	return storeJSON(st.configuration_path, st.configuration)
}

func (st *Storage) loadDisplayprefs() error {
	return loadJSON(st.displayprefs_path, &st.displayprefs)
}

func (st *Storage) storeDisplayprefs() error {
	return storeJSON(st.displayprefs_path, st.displayprefs)
}

func loadJSON(path string, obj interface{}) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(file, &obj)
	return err
}

func storeJSON(path string, obj interface{}) error {
	test := json.NewEncoder(os.Stdout)
	test.Encode(obj)
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(path, data, 0644)
	return err
}