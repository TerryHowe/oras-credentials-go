/*
Copyright The ORAS Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package credentials

import (
	"context"

	"oras.land/oras-go/v2/registry/remote/auth"
)

// Store is the interface that any credentials store must implement.
type Store interface {
	// Get retrieves credentials from the store for the given server address.
	Get(ctx context.Context, serverAddress string) (auth.Credential, error)
	// Put saves credentials into the store for the given server address.
	Put(ctx context.Context, serverAddress string, cred auth.Credential) error
	// Delete removes credentials from the store for the given server address.
	Delete(ctx context.Context, serverAddress string) error
}
