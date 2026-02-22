package e2e_test

import (
	"net/url"
	"os"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"

	"github.com/circa10a/shoutrrr/internal/testutils"
	"github.com/circa10a/shoutrrr/pkg/services/chat/discord"
	"github.com/circa10a/shoutrrr/pkg/types"
)

var _ = ginkgo.Describe("Discord E2E Embed Author Test", func() {
	ginkgo.When("running e2e tests", func() {
		ginkgo.It("should send embed with author information", func() {
			envURL := os.Getenv("SHOUTRRR_DISCORD_URL")
			if envURL == "" {
				ginkgo.Skip("SHOUTRRR_DISCORD_URL not set, skipping author info test")

				return
			}

			serviceURL, _ := url.Parse(envURL)
			service := &discord.Service{}
			err := service.Initialize(serviceURL, testutils.TestLogger())
			gomega.Expect(err).NotTo(gomega.HaveOccurred())

			items := []types.MessageItem{
				{
					Text: "E2E Test: Embed with author information",
					Fields: []types.Field{
						{Key: "embed_author_name", Value: "Shoutrrr E2E Test"},
						{
							Key:   "embed_author_url",
							Value: "https://github.com/circa10a/shoutrrr",
						},
						{
							Key:   "embed_author_icon_url",
							Value: "https://raw.githubusercontent.com/circa10a/shoutrrr/master/docs/assets/media/shoutrrr-180px.png",
						},
					},
				},
			}
			err = service.SendItems(items, nil)
			gomega.Expect(err).NotTo(gomega.HaveOccurred())
		})
	})
})
