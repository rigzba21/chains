// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build darwin && amd64 && !go1.21

package waf

import _ "embed" // Needed for go:embed

//go:embed lib/darwin-amd64/_libddwaf.dylib
var libddwaf []byte
