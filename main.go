package main

import (
	"fmt"
	"net/http"
	"os"
	"io"
	"time"
	"strconv"
)

func loop () {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
}

func main() {
	loop()
	loop()

	videoUrl := "https://rr2---sn-oguesnzl.googlevideo.com/videoplayback?expire=1649801897&ei=SaZVYqWTNpTLigT36pCADA&ip=209.107.196.171&id=o-AGYcm0XC8gicBDG6231nuv59neZi9-25iAzvU5b6Yx0v&itag=18&source=youtube&requiressl=yes&spc=4ocVC6W4HJYGI-sxUgbRWXOl-LFF&vprv=1&mime=video%2Fmp4&ns=4DfjDNBdBWtdJa-xVCYcSAMG&gir=yes&clen=9846375&ratebypass=yes&dur=191.518&lmt=1649732627949957&fexp=24001373,24007246&c=WEB&txp=5438434&n=gVmPgC4Swbfedg&sparams=expire%2Cei%2Cip%2Cid%2Citag%2Csource%2Crequiressl%2Cspc%2Cvprv%2Cmime%2Cns%2Cgir%2Cclen%2Cratebypass%2Cdur%2Clmt&sig=AOq0QJ8wRQIgF8uaUEz-VwnPe8HpM9bPp4B7rS9FOlSF81GtNIgJc8sCIQCQYDb3syrbMdnR_aKQ_qNkELDgZDNHNPLGkdMaVn30ZQ%3D%3D&redirect_counter=1&rm=sn-5uayk7z&req_id=3968379c5ed6a3ee&cms_redirect=yes&ipbypass=yes&mh=jW&mip=18.183.79.3&mm=31&mn=sn-oguesnzl&ms=au&mt=1649783357&mv=m&mvi=2&pl=18&lsparams=ipbypass,mh,mip,mm,mn,ms,mv,mvi,pl&lsig=AG3C_xAwRAIgE4iwgOPucyvuCYS0cJmzNnHXljDldJLYB1B7ODx1fzcCIHqs_j87lIni7nus7S8x7mjKvmibVezOcS7Xuh3nCNX7"
	timeUnix := time.Now().Unix()
	timeUnixString := strconv.FormatInt(timeUnix, 10)

	err := DownloadFile("test" + timeUnixString + ".mp4", videoUrl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Downloaded: " + videoUrl)
}

func DownloadFile(filePath string, url string) error {

	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, response.Body)
	return err
}