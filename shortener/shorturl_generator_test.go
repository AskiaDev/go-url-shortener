package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


const UserId = "XBXAzzAWQrluFZJcImTQCCm1eLEXzuLq8hdDJc+AtweSo6B6y0I8D7qLgH7lOFzx9kmpbMmnVNnO5pQDCmGH2A=="


func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://www.guru3d.com/news-story/spotted-ryzen-threadripper-pro-3995wx-processor-with-8-channel-ddr4,2.html"
	shortLink_1 := GenerateShortLink(initialLink_1, UserId)

	initialLink_2 := "https://www.eddywm.com/lets-build-a-url-shortener-in-go-with-redis-part-2-storage-layer/"
	shortLink_2 := GenerateShortLink(initialLink_2, UserId)

	initialLink_3 := "https://spectrum.ieee.org/automaton/robotics/home-robots/hello-robots-stretch-mobile-manipulator"
	shortLink_3 := GenerateShortLink(initialLink_3, UserId)


	assert.Equal(t, shortLink_1, "g6WpKevc")
	assert.Equal(t, shortLink_2, "2cBPu9Fs")
	assert.Equal(t, shortLink_3, "dhZTayYQ")
}