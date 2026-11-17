// SPDX-FileCopyrightText: Contributors to the Gardener project
//
// SPDX-License-Identifier: Apache-2.0

package ssh_test

import (
	"crypto/rand"
	"crypto/rsa"
	"net"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/ssh"

	sshutils "github.com/gardener/gardener/pkg/utils/ssh"
)

var _ = Describe("Options", func() {
	Describe("#WithHostKey", func() {
		var (
			config    *sshutils.Config
			publicKey ssh.PublicKey
			authKey   []byte
		)

		BeforeEach(func() {
			config = sshutils.DefaultConfig()

			privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
			Expect(err).NotTo(HaveOccurred())
			publicKey, err = ssh.NewPublicKey(&privateKey.PublicKey)
			Expect(err).NotTo(HaveOccurred())
			authKey = ssh.MarshalAuthorizedKey(publicKey)
		})

		It("should accept the pinned host key", func() {
			Expect(sshutils.WithHostKey(authKey)(config)).To(Succeed())
			Expect(config.HostKeyCallback("host:22", &net.IPAddr{}, publicKey)).To(Succeed())
		})

		It("should reject a different host key", func() {
			otherPrivateKey, err := rsa.GenerateKey(rand.Reader, 2048)
			Expect(err).NotTo(HaveOccurred())
			otherPublicKey, err := ssh.NewPublicKey(&otherPrivateKey.PublicKey)
			Expect(err).NotTo(HaveOccurred())

			Expect(sshutils.WithHostKey(authKey)(config)).To(Succeed())
			Expect(config.HostKeyCallback("host:22", &net.IPAddr{}, otherPublicKey)).To(MatchError(ContainSubstring("host key mismatch")))
		})

		It("should fail for a malformed public key", func() {
			Expect(sshutils.WithHostKey([]byte("not-a-valid-key"))(config)).To(HaveOccurred())
		})
	})
})
