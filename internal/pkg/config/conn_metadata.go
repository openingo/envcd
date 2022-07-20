/*
 * Copyright (c) 2022, AcmeStack
 * All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package config

import (
	"fmt"
	"net/url"

	"github.com/acmestack/godkits/log"
)

// ConnMetadata with standard Url: etcd://user:123@localhost:123 metadata
type ConnMetadata struct {
	Type     string // url schema
	UserName string
	Password string
	Host     string
	Hostname string
	Port     string
}

func parser(connection string) *ConnMetadata {
	u, err := url.Parse(connection)
	if err != nil {
		log.Error(" parser connection metadata error %s\n", err)
	}
	metadata := &ConnMetadata{}
	metadata.Type = u.Scheme
	metadata.UserName = u.User.Username()
	password, _ := u.User.Password()
	metadata.Password = password
	metadata.Host = u.Host
	metadata.Hostname = u.Hostname()
	// todo port to int?
	metadata.Port = u.Port()
	return metadata
}

func (connMedata *ConnMetadata) information(t string) {
	// todo logging
	log.Info(fmt.Sprintf("ConnectionMetadata For %v", t))
	log.Info(fmt.Sprintf("Type: %v", connMedata.Type))
	log.Info(fmt.Sprintf("UserName: %v", connMedata.UserName))
	log.Info(fmt.Sprintf("Host: %v", connMedata.Host))
	log.Info(fmt.Sprintf("Hostname: %v", connMedata.Hostname))
	log.Info(fmt.Sprintf("Port: %v", connMedata.Port))
	log.Info("--")
}
