/*
 * Copyright (c) 2024 Yunshan Networks
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package log_data

import (
	"fmt"
	"net"
	"strings"
)

func IPIntToString(ipInt uint32) string {
	return net.IPv4(byte(ipInt>>24), byte(ipInt>>16), byte(ipInt>>8), byte(ipInt)).String()
}

// eg. url=http://nacos:8848/nacos/v1/ns/instance/list, parse return `/nacos/v1/ns/instance/list`
func ParseUrlPath(rawURL string) (string, error) {
	parts := strings.SplitN(rawURL, "://", 2)
	if len(parts) != 2 || parts[1] == "" {
		return "", fmt.Errorf("invalid URL format")
	}
	pathStart := strings.Index(parts[1], "/")
	if pathStart == -1 {
		return "/", nil
	}

	return parts[1][pathStart:], nil
}
