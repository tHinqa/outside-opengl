// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package opengl

import "testing"

func TestInit(t *testing.T) {
	//TODO(t): Why ""?
	t.Log("Vendor:", GetString(Vendor))
	t.Log("Renderer:", GetString(Renderer))
	t.Log("Version:", GetString(Version))
	t.Log("Extensions:", GetString(Extensions))
}
