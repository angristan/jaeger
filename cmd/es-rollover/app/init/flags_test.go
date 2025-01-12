// Copyright (c) 2021 The Jaeger Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package init

import (
	"flag"
	"testing"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBindFlags(t *testing.T) {
	v := viper.New()
	c := &Config{}
	command := cobra.Command{}
	flags := &flag.FlagSet{}
	c.AddFlags(flags)
	command.PersistentFlags().AddGoFlagSet(flags)
	v.BindPFlags(command.PersistentFlags())

	err := command.ParseFlags([]string{
		"--shards=8",
		"--replicas=16",
		"--priority-span-template=300",
		"--priority-service-template=301",
		"--priority-dependencies-template=302",
	})
	require.NoError(t, err)

	c.InitFromViper(v)
	assert.Equal(t, 8, c.Shards)
	assert.Equal(t, 16, c.Replicas)
	assert.Equal(t, 300, c.PrioritySpanTemplate)
	assert.Equal(t, 301, c.PriorityServiceTemplate)
	assert.Equal(t, 302, c.PriorityDependenciesTemplate)
}
