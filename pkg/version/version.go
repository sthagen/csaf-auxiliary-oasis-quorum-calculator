// This file is Free Software under the Apache-2.0 License
// without warranty, see README.md and LICENSE for details.
//
// SPDX-License-Identifier: Apache-2.0
//
// SPDX-FileCopyrightText: 2024 German Federal Office for Information Security (BSI) <https://www.bsi.bund.de>
// Software-Engineering: 2024 Intevation GmbH <https://intevation.de>

// Package version implements the burned-in version of the binaries.
package version

// SemVersion the version in semver.org format, MUST be overwritten during
// the linking stage of the build process
var SemVersion = "0.0.0"
