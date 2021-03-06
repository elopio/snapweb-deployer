/*
 * Copyright (C) 2016 Canonical Ltd
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License version 3 as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package snapcraft_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/elopio/snapweb-deployer/snapcraft"
)

func Test(t *testing.T) {
	repoDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("error creating temp dir: %v", err)
	}
	yamlFile, err := os.Create(path.Join(repoDir, "snapcraft.yaml"))
	if err != nil {
		t.Fatalf("error creating yaml file: %v", err)
	}
	defer yamlFile.Close()
	yamlFile.WriteString(fmt.Sprintf(`name: testsnap
version: testver
summary: dummy
description: dummy
architectures: [all]

parts:
  dummy-part:
    plugin: nil
`))

	s := snapcraft.Snapcraft{}
	fmt.Println(repoDir)
	snap, err := s.Snap(repoDir)
	if err != nil {
		t.Fatalf("error on Snap: %v", err)
	}
	if snap != path.Join(repoDir, "testsnap_testver_all.snap") {
		t.Fatalf("wrong snap: %v", snap)
	}
}
