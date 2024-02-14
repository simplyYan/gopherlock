package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var sites = map[string]string{
	"Facebook":             "https://www.facebook.com/%s",
	"Instagram":            "https://www.instagram.com/%s",
	"Twitter":              "https://twitter.com/%s",
	"LinkedIn":             "https://www.linkedin.com/in/%s",
	"GitHub":               "https://github.com/%s",
	"YouTube":              "https://www.youtube.com/user/%s",
	"Reddit":               "https://www.reddit.com/user/%s",
	"Tumblr":               "https://%s.tumblr.com",
	"Pinterest":            "https://www.pinterest.com/%s",
	"Snapchat":             "https://www.snapchat.com/add/%s",
	"WhatsApp":             "https://wa.me/%s",
	"Twitch":               "https://www.twitch.tv/%s",
	"SoundCloud":           "https://soundcloud.com/%s",
	"Flickr":               "https://www.flickr.com/photos/%s",
	"Vimeo":                "https://vimeo.com/%s",
	"Imgur":                "https://imgur.com/user/%s",
	"WordPress":            "https://%s.wordpress.com",
	"Blogger":              "https://%s.blogspot.com",
	"Medium":               "https://medium.com/@%s",
	"StackOverflow":        "https://stackoverflow.com/users/%s",
	"DeviantArt":           "https://%s.deviantart.com",
	"Dribbble":             "https://dribbble.com/%s",
	"Behance":              "https://www.behance.net/%s",
	"SlideShare":           "https://slideshare.net/%s",
	"Meetup":               "https://www.meetup.com/members/%s",
	"Goodreads":            "https://www.goodreads.com/%s",
	"Last.fm":              "https://www.last.fm/user/%s",
	"Spotify":              "https://open.spotify.com/user/%s",
	"Bandcamp":             "https://www.bandcamp.com/%s",
	"Wikipedia":            "https://en.wikipedia.org/wiki/User:%s",
	"IMDb":                 "https://www.imdb.com/user/%s",
	"Etsy":                 "https://www.etsy.com/people/%s",
	"Amazon":               "https://www.amazon.com/gp/profile/amzn1.account.%s",
	"EBay":                 "https://www.ebay.com/usr/%s",
	"Alibaba":              "https://www.alibaba.com/member/%s.html",
	"AliExpress":           "https://www.aliexpress.com/store/%s",
	"PayPal":               "https://www.paypal.me/%s",
	"Venmo":                "https://venmo.com/%s",
	"CashApp":              "https://cash.app/$%s",
	"Zelle":                "https://www.zellepay.com/%s",
	"GooglePlus":           "https://plus.google.com/%s",
	"VK":                   "https://vk.com/%s",
	"TikTok":               "https://www.tiktok.com/@%s",
	"Weibo":                "https://weibo.com/%s",
	"QQ":                   "https://user.qzone.qq.com/%s",
	"Renren":               "https://www.renren.com/%s",
	"Baidu":                "https://www.baidu.com/home/center/profile/%s",
	"Xanga":                "https://%s.xanga.com",
	"MySpace":              "https://myspace.com/%s",
	"LiveJournal":          "https://%s.livejournal.com",
	"Orkut":                "https://www.orkut.com/%s",
	"Hi5":                  "https://www.hi5.com/%s",
	"Friendster":           "https://www.friendster.com/%s",
	"MeetMe":               "https://www.meetme.com/%s",
	"Tagged":               "https://www.tagged.com/%s",
	"Foursquare":           "https://foursquare.com/%s",
	"Yelp":                 "https://www.yelp.com/user_details?userid=%s",
	"TripAdvisor":          "https://www.tripadvisor.com/members/%s",
	"Airbnb":               "https://www.airbnb.com/users/show/%s",
	"CouchSurfing":         "https://www.couchsurfing.com/people/%s",
	"Fiverr":               "https://www.fiverr.com/%s",
	"Freelancer":           "https://www.freelancer.com/u/%s",
	"Upwork":               "https://www.upwork.com/freelancers/%s",
	"Guru":                 "https://www.guru.com/freelancers/%s",
	"PeoplePerHour":        "https://www.peopleperhour.com/freelancer/%s",
	"Toptal":               "https://www.toptal.com/%s",
	"HackerRank":           "https://www.hackerrank.com/%s",
	"Codecademy":           "https://www.codecademy.com/profiles/%s",
	"Kaggle":               "https://www.kaggle.com/%s",
	"LeetCode":             "https://leetcode.com/%s",
	"Hackerearth":          "https://www.hackerearth.com/@%s",
	"TopCoder":             "https://www.topcoder.com/members/%s",
	"Exercism":             "https://exercism.io/profiles/%s",
	"Codeforces":           "https://codeforces.com/profile/%s",
	"CoderByte":            "https://www.coderbyte.com/profile/%s",
	"GeeksforGeeks":        "https://auth.geeksforgeeks.org/user/%s",
	"SPOJ":                 "https://www.spoj.com/users/%s",
	"UVa Online Judge":     "https://uhunt.onlinejudge.org/id/%s",
	"XVideos":              "https://www.xvideos.com/profiles/%s",
	"Pornhub":              "https://www.pornhub.com/users/%s",
	"XNXX":                 "https://www.xnxx.com/profile/%s",
	"XHamster":             "https://xhamster.com/users/%s",
	"ROBLOX":               "https://www.roblox.com/user.aspx?username=%s",
	"Duolingo":             "https://www.duolingo.com/profile/%s",
	"Tinder":               "https://www.tinder.com/@%s",
	"Quora":                "https://www.quora.com/profile/%s",
	"XDA Developers":       "https://forum.xda-developers.com/member.php?username=%s",
	"Bodybuilding.com":     "https://forum.bodybuilding.com/member.php?username=%s",
	"SitePoint":            "https://www.sitepoint.com/forums/index.php?action=profile;user=%s",
	"DigitalPoint":         "https://forums.digitalpoint.com/members/%s/",
	"Warrior Forum":        "https://www.warriorforum.com/members/%s.html",
	"Black Hat World":      "https://www.blackhatworld.com/members/%s/",
	"WebmasterWorld":       "https://www.webmasterworld.com/home.htm?username=%s",
	"MoneyMakerDiscussion": "https://www.moneymakerdiscussion.com/members/%s.html",
	"V7N":                  "https://www.v7n.com/forums/members/%s.html",
	"BitcoinTalk":          "https://bitcointalk.org/index.php?action=profile;u=%s",
	"CryptoCompare":        "https://www.cryptocompare.com/profile/%s/",
	"HackForums":           "https://hackforums.net/member.php?action=profile&uid=%s",
	"Cracked.to":           "https://cracked.to/member.php?action=profile&uid=%s",
	"TheHackToday":         "https://forum.thehacktoday.com/member.php?action=profile&uid=%s",
	"WebHostingTalk":       "https://www.webhostingtalk.com/member.php?username=%s",
	"Steam":                "https://steamcommunity.com/id/%s",
	"Battle.net":           "https://battle.net/u/%s",
	"Ubisoft Connect":      "https://connect.ubisoft.com/user/%s",
	"Origin":               "https://www.origin.com/profile/%s",
}

func main() {
	fmt.Println(`
╔═╗┌─┐┌─┐┬ ┬┌─┐┬─┐┬  ┌─┐┌─┐┬┌─
║ ╦│ │├─┘├─┤├┤ ├┬┘│  │ ││  ├┴┐
╚═╝└─┘┴  ┴ ┴└─┘┴└─┴─┘└─┘└─┘┴ ┴	

Tool under BSD 3-Clause license. Misuse is your responsibility, the author bears no responsibility.
	`)

	if len(os.Args) < 2 {
		fmt.Println("Usage: <username>")
		os.Exit(1)
	}

	username := os.Args[1]

	for site, url := range sites {
		checkUsername(site, url, username)
	}
}

func checkUsername(site, url, username string) {
	resp, err := goquery.NewDocument(fmt.Sprintf(url, username))
	if err != nil {
		fmt.Printf("Erro ao acessar %s: %s\n", site, err)
		return
	}

	if resp.Find("body").Text() != "" {
		fmt.Printf("User found at %s: %s\n", site, fmt.Sprintf(url, username))
	} else {
		fmt.Printf("User not found in %s\n", site)
		tryVariations(site, url, username)
	}
}

func tryVariations(site, url, username string) {
	variations := []string{
		strings.ReplaceAll(username, " ", "_"),
		strings.ReplaceAll(username, " ", "."),
		username + "_",
		username + ".",
		username + "1",
	}

	for _, variation := range variations {
		resp, err := goquery.NewDocument(fmt.Sprintf(url, variation))
		if err != nil {
			fmt.Printf("Access error %s: %s\n", site, err)
			continue
		}

		if resp.Find("body").Text() != "" {
			fmt.Printf("User found in %s (with variation): %s\n", site, fmt.Sprintf(url, variation))
			return
		}
	}

	fmt.Printf("No user variations found in %s\n", site)
}
