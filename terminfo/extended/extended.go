// Copyright 2020 The TCell Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package extended contains an extended set of terminal descriptions.
// Applications desiring to have a better chance of Just Working by
// default should include this package.  This will significantly increase
// the size of the program.
package extended

import (
	// The following imports just register themselves --
	// these are the terminal types we aggregate in this package.
	_ "github.com/tagirmukail/tcell/terminfo/a/aixterm"
	_ "github.com/tagirmukail/tcell/terminfo/a/alacritty"
	_ "github.com/tagirmukail/tcell/terminfo/a/ansi"
	_ "github.com/tagirmukail/tcell/terminfo/b/beterm"
	_ "github.com/tagirmukail/tcell/terminfo/c/cygwin"
	_ "github.com/tagirmukail/tcell/terminfo/d/dtterm"
	_ "github.com/tagirmukail/tcell/terminfo/e/emacs"
	_ "github.com/tagirmukail/tcell/terminfo/f/foot"
	_ "github.com/tagirmukail/tcell/terminfo/g/gnome"
	_ "github.com/tagirmukail/tcell/terminfo/h/hpterm"
	_ "github.com/tagirmukail/tcell/terminfo/k/konsole"
	_ "github.com/tagirmukail/tcell/terminfo/k/kterm"
	_ "github.com/tagirmukail/tcell/terminfo/l/linux"
	_ "github.com/tagirmukail/tcell/terminfo/p/pcansi"
	_ "github.com/tagirmukail/tcell/terminfo/r/rxvt"
	_ "github.com/tagirmukail/tcell/terminfo/s/screen"
	_ "github.com/tagirmukail/tcell/terminfo/s/simpleterm"
	_ "github.com/tagirmukail/tcell/terminfo/s/sun"
	_ "github.com/tagirmukail/tcell/terminfo/t/termite"
	_ "github.com/tagirmukail/tcell/terminfo/t/tmux"
	_ "github.com/tagirmukail/tcell/terminfo/v/vt100"
	_ "github.com/tagirmukail/tcell/terminfo/v/vt102"
	_ "github.com/tagirmukail/tcell/terminfo/v/vt220"
	_ "github.com/tagirmukail/tcell/terminfo/v/vt320"
	_ "github.com/tagirmukail/tcell/terminfo/v/vt400"
	_ "github.com/tagirmukail/tcell/terminfo/v/vt420"
	_ "github.com/tagirmukail/tcell/terminfo/v/vt52"
	_ "github.com/tagirmukail/tcell/terminfo/w/wy50"
	_ "github.com/tagirmukail/tcell/terminfo/w/wy60"
	_ "github.com/tagirmukail/tcell/terminfo/w/wy99_ansi"
	_ "github.com/tagirmukail/tcell/terminfo/x/xfce"
	_ "github.com/tagirmukail/tcell/terminfo/x/xterm"
	_ "github.com/tagirmukail/tcell/terminfo/x/xterm_kitty"
	_ "github.com/tagirmukail/tcell/terminfo/x/xterm_termite"
)
