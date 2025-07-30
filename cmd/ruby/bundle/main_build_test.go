// Copyright 2025 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"testing"
	"os"

	buildpacktest "github.com/GoogleCloudPlatform/buildpacks/internal/buildpacktest"
	"github.com/GoogleCloudPlatform/buildpacks/internal/mockprocess"
	"github.com/GoogleCloudPlatform/buildpacks/pkg/ruby"
)

func TestBuild(t *testing.T) {
	testCases := []struct {
		name              string
		rubyVersion       string
		wantCommands      []*mockprocess.ExecRe
	}{
		{
			name:        "ruby 3.4.0 installs bundled gems",
			rubyVersion: "3.4.0",
			wantCommands: []*mockprocess.ExecRe{
				mockprocess.NewExecRe(t, `gem install bigdecimal`),
				mockprocess.NewExecRe(t, `gem install cgi`),
				mockprocess.NewExecRe(t, `gem install csv`),
				mockprocess.NewExecRe(t, `gem install drb`),
				mockprocess.NewExecRe(t, `gem install fcntl`),
				mockprocess.NewExecRe(t, `gem install fileutils`),
				mockprocess.NewExecRe(t, `gem install find`),
				mockprocess.NewExecRe(t, `gem install ftools`),
				mockprocess.NewExecRe(t, `gem install getoptlong`),
				mockprocess.NewExecRe(t, `gem install io-console`),
				mockprocess.NewExecRe(t, `gem install io-nonblock`),
				mockprocess.NewExecRe(t, `gem install io-wait`),
				mockprocess.NewExecRe(t, `gem install irb`),
				mockprocess.NewExecRe(t, `gem install logger`),
				mockprocess.NewExecRe(t, `gem install mutex_m`),
				mockprocess.NewExecRe(t, `gem install net-ftp`),
				mockprocess.NewExecRe(t, `gem install net-http`),
				mockprocess.NewExecRe(t, `gem install net-imap`),
				mockprocess.NewExecRe(t, `gem install net-pop`),
				mockprocess.NewExecRe(t, `gem install net-protocol`),
				mockprocess.NewExecRe(t, `gem install open-uri`),
				mockprocess.NewExecRe(t, `gem install optparse`),
				mockprocess.NewExecRe(t, `gem install pp`),
				mockprocess.NewExecRe(t, `gem install prettyprint`),
				mockprocess.NewExecRe(t, `gem install rdoc`),
				mockprocess.NewExecRe(t, `gem install readline`),
				mockprocess.NewExecRe(t, `gem install reline`),
				mockprocess.NewExecRe(t, `gem install resolv`),
				mockprocess.NewExecRe(t, `gem install rinda`),
				mockprocess.NewExecRe(t, `gem install securerandom`),
				mockprocess.NewExecRe(t, `gem install set`),
				mockprocess.NewExecRe(t, `gem install tempfile`),
				mockprocess.NewExecRe(t, `gem install time`),
				mockprocess.NewExecRe(t, `gem install tmpdir`),
				mockprocess.NewExecRe(t, `gem install tracer`),
				mockprocess.NewExecRe(t, `gem install un`),
				mockprocess.NewExecRe(t, `gem install uri`),
				mockprocess.NewExecRe(t, `gem install weakref`),
				mockprocess.NewExecRe(t, `gem install win32ole`),
				mockprocess.NewExecRe(t, `gem install yaml`),
			},
		},
		{
			name:        "ruby 3.3.0 does not install bundled gems",
			rubyVersion: "3.3.0",
			wantCommands:      []*mockprocess.ExecRe{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			os.Setenv(ruby.RubyVersionKey, tc.rubyVersion)
			defer os.Unsetenv(ruby.RubyVersionKey)

			opts := []buildpacktest.Option{
				buildpacktest.WithTestName(tc.name),
				buildpacktest.WithExecMock(tc.wantCommands...),
			}

			buildpacktest.TestBuild(t, buildFn, opts...)
		})
	}
}
