# Modified the basic makefiles referred to from the
# Go home page.
#
# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=container
GOFILES=\
        container.go\

include $(GOROOT)/src/Make.pkg

bucket:
	cd bucket && gomake

bucketclean:
	cd bucket && gomake clean

bucketinstall:
	cd bucket && gomake install

bucketnuke:
	cd bucket && gomake nuke
