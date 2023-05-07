package security

import (
	"log"
	"testing"
)

func TestCheckUrl(t *testing.T) {

	if CheckUrl("//ass") {
		log.Fatalln("//ass bypass")
	}
	if CheckUrl("/\\ass") {
		log.Fatalln("/\\ass bypass")
	}
	if CheckUrl("/\\\ass") {
		log.Fatalln("/\\\ass bypass")
	}
	if CheckUrl("/\\\\ass") {
		log.Fatalln("/\\\\ass bypass")
	}
	if CheckUrl("/\\\\ass//a.com") {
		log.Fatalln("/\\\\ass bypass")
	}
	if CheckUrl("http://s.com") {
		log.Fatalln("http://s.com bypass")
	}
	if CheckUrl("127.0.0.1/a//a.com") {
		log.Fatalln("bai.com bypass")
	}
	if !CheckUrl("/login") {
		log.Fatalln("/login bypass")
	}
	if !CheckUrl("localhost") {
		log.Fatalln("localhost bypass")
	}
	if !CheckUrl("http://127.0.0.1/a//a.com") {
		log.Fatalln("bai.com bypass")
	}
	if !CheckUrl("www.baidu.com") {
		log.Fatalln("bai.com bypass")
	}

}
