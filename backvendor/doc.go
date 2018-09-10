// Copyright (C) 2018 Tim Waugh
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package backvendor provides a way to represent Go source code in a
// filesystem, and taken from a source code repository. It allows
// mapping vendored packages back to the original versions they came
// from.
//
// A GoSource represents a filesystem tree containing Go source
// code. Create it using NewGoSource. The Project and VendoredProjects
// methods return information about the top-level project and the
// vendored projects it has.
//
//     src := backvendor.NewGoSource(path)
//     proj, perr := src.Project(importPath)
//     vendored, verr := src.VendoredProjects()
//
// Both of these methods use RepoRoot (from golang.org/x/tools/go/vcs)
// to describe the projects.
//
// The DescribeProject function takes a RepoRoot and returns a
// Representation, indicating the upstream version of the project or
// vendored project, e.g.
//
//     ref, rerr := backvendor.DescribeProject(proj, src.Path)
//
// It does this by comparing file hashes of the local files with those
// from commits in the upstream repository.
package backvendor
