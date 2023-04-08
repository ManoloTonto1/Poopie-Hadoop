package collectors

import (
	"net/http"
	"time"

	"github.com/gocolly/colly"
)

var userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/111.0"

var limitRules = &colly.LimitRule{
	RandomDelay: 2 * time.Second,
	Parallelism: 4,
}

var decathlonHeaders = http.Header{
	"User-Agent":      []string{userAgent},
	"Accept":          []string{"text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8"},
	"Accept-Language": []string{"en-US,en;q=0.5"},
	"Accept-Encoding": []string{"gzip, deflate, br"},
	"Content-Type":    []string{"application/json"},
	"Referer":         []string{"YOUR_REFERER_URL"},
	"Origin":          []string{"https://www.decathlon.nl"},
	"Sec-Fetch-Dest":  []string{"empty"},
	"Sec-Fetch-Mode":  []string{"cors"},
	"Sec-Fetch-Site":  []string{"same-origin"},
	"Connection":      []string{"keep-alive"},
	"TE":              []string{"Trailers"},
}
var decathlonCookieString = "NFS_USER_ID=bb3f91e0-0d06-4f50-b430-2b4f493d7009; nfs-traffic-split-nl=leadtime-cart%3Dcontrol%2Csustainable-delivery-options%3Dvariation-04; visid_incap_989872=P4L+Ho7SSsSFmrDfRWMv1gCKMWQAAAAAQUIPAAAAAADUjICeCwcbRave49Kl9Xgl; ABTasty=uid=d4vbjn1qy15xf68w&fst=1680968193428&pst=1680968193428&cst=1680987098041&ns=2&pvt=31&pvis=11&th=; DKT_SESSION=p7TArcYOOLfD/h+pCgOtimpeqsjQDdXu8kS+nopTp2wdknYAa1QAc/GAQU7G56epgGefEOPt9zEroNQWgDKbYX19b4DyYOT4ErLZEg2at/ancnL1jaq8QrsxehFszdnB/JpDtUzkb8ZgGGHGlJ7ZlO+/25I1Ei0ysVtmQY7uP4k=; didomiVendorsConsent=c:ysance-3xwFx9e7,c:flowbox-qMePtMmf,c:impactrad-hfjEpR2d,c:facebooka-QyMDLGJQ,c:pinterest-VJGF92Fa,c:mopiniong-XKtap2Up,c:rtbhouse-Bft2p4g2,c:criteo-jT4wQwwT,c:mpulse-TiGTen8F,c:baqend-WizGUBcV,c:googleads-Ak32TPDA,c:bing-nLiQLACi,c:luckyoran-rWnAdXQU,c:retailroc-jbyDr6ZY,c:abtasty-kBHy79W8,c:salesforc-MrCedCaG,c:youtube-XejLUqmf,c:dynamicyi-ETeaxHck,c:cookiesan-aY7WekKr,c:zendesk-a8ypMXZT,c:medallia-wmD8Zf9h,; USER_TOKEN_FOR_ANALYTICS=bb3f91e0-0d06-4f50-b430-2b4f493d7009; baqend-speedkit-user-id=eYkBSfMwLrf3oWdkcF3bRxAu7; mPulse_activated=true; _gcl_au=1.1.687846696.1680968196; luckyorange_activated=false; didomi_token=eyJ1c2VyX2lkIjoiMTdkY2NlYTYtYTM1Zi02NWIwLWJjZWItNGY0YzNlZTJmMWVmIiwiY3JlYXRlZCI6IjIwMjMtMDQtMDhUMTU6MzY6MzYuMzcwWiIsInVwZGF0ZWQiOiIyMDIzLTA0LTA4VDE1OjM2OjM2LjM3MFoiLCJ2ZXJzaW9uIjoyLCJwdXJwb3NlcyI6eyJlbmFibGVkIjpbInBlcnNvbmFsaXMtdGY4cFpUVkgiLCJtYXJrZXRpbmctTVJaVnByZWEiLCJjb250ZW50cGUtMmJqTnduOWciXX0sInZlbmRvcnMiOnsiZW5hYmxlZCI6WyJjOnlzYW5jZS0zeHdGeDllNyIsImM6Zmxvd2JveC1xTWVQdE1tZiIsImM6c3dvZ28tYWt6Q2hLS0oiLCJjOmltcGFjdHJhZC1oZmpFcFIyZCIsImM6emVuZGVzay1hOHlwTVhaVCIsImM6YmFxZW5kLVdpekdVQmNWIiwiYzpmYWNlYm9va2EtUXlNRExHSlEiLCJjOm1vcGluaW9uZy1YS3RhcDJVcCIsImM6cnRiaG91c2UtQmZ0MnA0ZzIiLCJjOmNyaXRlby1qVDR3UXd3VCIsImM6bXB1bHNlLVRpR1RlbjhGIiwiYzpnb29nbGVhZHMtQWszMlRQREEiLCJjOmJpbmctbkxpUUxBQ2kiLCJjOmx1Y2t5b3Jhbi1yV25BZFhRVSIsImM6cmV0YWlscm9jLWpieURyNlpZIiwiYzpkeW5hbWljeWktRVRlYXhIY2siLCJjOmFidGFzdHkta0JIeTc5VzgiLCJjOnNhbGVzZm9yYy1NckNlZENhRyIsImM6Y29va2llc2FuLWFZN1dla0tyIiwiYzpwaW50ZXJlc3QtVkpHRjkyRmEiLCJjOnlvdXR1YmUtWGVqTFVxbWYiLCJjOm1lZGFsbGlhLXdtRDhaZjloIl19fQ==; euconsent-v2=CPp5tEAPp5tEAAHABBENC-CgAAAAAAAAAAAAAAAAAAAA.YAAAAAAAAAAA; __ywtfpcvuid=25966058001680968198027; IR_PI=27112e0c-d623-11ed-91dd-91276d944408%7C1681054596530; cto_bundle=qTsXAV84RTFudTBHazRRV0VYOWF3TUFTeDNiM3hDS1BuTG9IaUhnRld4TktHaWNub1VSYUY3Yml1Ylc3MENOc1ZoUWJPOWltZiUyRnpWJTJCQ0xRUkRLUjgxdGtSWVRkalRyTXJORnFjZUFxckNIQWRZa1ExRlZMVlhqUzBNNjNvaHp2dklIdHA4UkZ3cDNROThVb1VEMmxCS0ljYVFRJTNEJTNE; RT='sl=8&ss=1680987096588&tt=32714&obo=0&bcn=%2F%2F02179914.akstat.io%2F&sh=1680990668976%3D8%3A0%3A32714%2C1680990639923%3D7%3A0%3A27116%2C1680989281235%3D6%3A0%3A21846%2C1680989129934%3D5%3A0%3A19764%2C1680989003923%3D4%3A0%3A17308&dm=decathlon.nl&si=a0973731-4066-4938-bf94-c0037cf0d01c&r=https%3A%2F%2Fwww.decathlon.nl%2Fp%2Flage-wandelschoenen-voor-kinderen-roze-maat-30-tot-38%2F_%2FR-p-X8746569&ul=1680991312280&hd=1680991312586'; mdLogger=false; kampyle_userid=2037-eb48-91ff-9177-3260-1c3d-d189-450a; kampyleUserSession=1680990669006; kampyleSessionPageCounter=1; kampyleUserSessionsCount=21; rrpvid=493348360213265; rcuid=64318a09cf40efef14f6181f; _ga_4WBL5RY5C5,G-KHS6MFWBPD=GS1.1.1680988843.2.1.1680991312.0.0.0; _dyid_server=76562148829006348; CUSTOMER_LVP=4605538; rrlevt=1680990669514; _dyjsession=rgjmocrheuphpvckaklo3fd8k8m4r7sp; PLAY_LANG=nl; incap_ses_282_989872=HaaiAP4jo3mS7bsHb97pA6rhMWQAAAAAGUMBP6vgdQcezOnXvVwuUA==; PLAY_SESSION=270dc265e16be16ad5dc5fb6c39f0aa53455fc20-JSESSIONID=UkRi2YdmVAMWhjDydBSfBH91fat70g3GaaAo4MdfTQ6B590zmLbx%212062747088&APISERVER=API03; nlbi_989872=Nih8LleYvXPrH0WcUDnf7wAAAADAVGKr3jJHH3kBz8PJYnzo; incap_ses_764_989872=ToGRaWdDkTRl4tP9NEaaCsPhMWQAAAAAhvwsuMqNOVp9U88oDpFZ2w==; ABTastySession=mrasn=&lp=https%253A%252F%252Fwww.decathlon.nl%252Fr%252Flage-wandelschoenen-voor-kinderen-roze-maat-30-tot-38%252F_%252FR-p-X8746569%253Fmc%253D8746569%2526c%253DROZE; nlbi_989872_2482950=TKd8TAOFcyldGRGDUDnf7wAAAAArXtK2Mo0S/ga26ytJ0xld; __ywtfpcsuid=37570781061680987099993; _ga_KHS6MFWBPD=GS1.1.1680987100.1.1.1680991312.60.0.0; _ga=GA1.1.dbfd25d1-6024-44e4-850b-37f96faadb7b; _ga_4WBL5RY5C5=GS1.1.1680987100.1.1.1680991312.0.0.0; IR_gbd=decathlon.nl; IR_5122=1680990666365%7C0%7C1680990666365%7C%7C; incap_ses_769_989872=CUhCXUi+dRcuDStBQwqsClviMWQAAAAAOkTqsY+DjLOLg9JXVscKlw==; incap_ses_450_989872=XMWGMHlV+0gLLOuAhrk+BqraMWQAAAAAdAXKgPfqXWM/mhAYv5aTyA==; incap_ses_763_989872=FS/sFwig20Usxv9UUrmWCqzaMWQAAAAAovOSj0FpFb9hT+JadvqcFw==; incap_ses_767_989872=qYssODpe2wjP7N9Vs+6kCsXhMWQAAAAAyC+0k95PFIzVcnU6pjIxoA==; incap_ses_766_989872=E1WEHNZkTCzN9/tH1WGhCqvhMWQAAAAAVPzROYlHeHoPdP9y9V12sw==; rrviewed=8746569; incap_ses_768_989872=tb3vJx2tLwRvxm0/NHyoCrXaMWQAAAAAqs57SjrscBN+d9muZNaJyQ==; incap_ses_770_989872=iTeWZWAxKzUYOwXjwZevCrXaMWQAAAAAlqpE70rHZY3Z5EAwWHYiCg==; ACTIVE_USER=y; _uetsid=25ef0980d62311edbb0ea9c85a0e76fd; _uetvid=119c18a0276c11edb2bdafc0c1b63704; incap_ses_451_989872=borOWfdG5jEqa7bY+0ZCBlLkMWQAAAAAUZgcFRqXB7l3zKzWlTcs4Q=="

type DecathlonPostData struct {
	Components []struct {
		ID    string `json:"id"`
		Input struct {
			AsyncRequest bool     `json:"asyncRequest"`
			Count        int      `json:"count"`
			Ids          []string `json:"ids"`
			Page         int      `json:"page"`
		} `json:"input"`
		Type string `json:"type"`
	} `json:"components"`
}
