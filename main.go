package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/Anderson-Lu/gofasion/gofasion"
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"
	"time"
)

var versionsJson []string
var server int
var (
	h   bool
	l   bool
	f   int
	p   int
	vid string
)

func urlReplace(url string, from int) string {
	switch from {
	case 0:
		break
	case 1:
		r := regexp.MustCompile(`https*://launchermeta\.mojang\.com/|https*://launcher\.mojang\.com/`)
		url = r.ReplaceAllString(url, "https://bmclapi2.bangbang93.com/")
		url = strings.ReplaceAll(url, "http://resources.download.minecraft.net", "https://bmclapi2.bangbang93.com/assets")
		url = strings.ReplaceAll(url, "https://libraries.minecraft.net/", "https://bmclapi2.bangbang93.com/maven")
		url = strings.ReplaceAll(url, "https://files.minecraftforge.net/maven", "https://bmclapi2.bangbang93.com/maven")
		url = strings.ReplaceAll(url, "http://dl.liteloader.com/versions/versions.json", "https://bmclapi.bangbang93.com/maven/com/mumfrey/liteloader/versions.json")
		url = strings.ReplaceAll(url, "https://authlib-injector.yushi.moe", "https://bmclapi2.bangbang93.com/mirrors/authlib-injector")
		url = strings.ReplaceAll(url, "https://meta.fabricmc.net", "https://bmclapi2.bangbang93.com/fabric-meta")
		url = strings.ReplaceAll(url, "https://maven.fabricmc.net", "https://bmclapi2.bangbang93.com/maven")
	case 2:
		r := regexp.MustCompile(`https*://launchermeta\.mojang\.com/|https*://launcher\.mojang\.com/`)
		url = r.ReplaceAllString(url, "https://download.mcbbs.net/")
		url = strings.ReplaceAll(url, "http://resources.download.minecraft.net", "https://download.mcbbs.net/assets")
		url = strings.ReplaceAll(url, "https://libraries.minecraft.net/", "https://download.mcbbs.net/maven")
		url = strings.ReplaceAll(url, "https://files.minecraftforge.net/maven", "https://download.mcbbs.net/maven")
		url = strings.ReplaceAll(url, "http://dl.liteloader.com/versions/versions.json", "https://download.mcbbs.net/maven/com/mumfrey/liteloader/versions.json")
		url = strings.ReplaceAll(url, "https://authlib-injector.yushi.moe", "https://download.mcbbs.net/mirrors/authlib-injector")
		url = strings.ReplaceAll(url, "https://meta.fabricmc.net", "https://download.mcbbs.net/fabric-meta")
		url = strings.ReplaceAll(url, "https://maven.fabricmc.net", "https://download.mcbbs.net/maven")
	}
	return url
}
func downloadFileUrl(hash string, from int) (string, string) {

	head := hash[0:2]

	url := fmt.Sprintf(urlReplace("http://resources.download.minecraft.net/%s/%s", server), head, hash)
	path := fmt.Sprintf("%s/%s", head, hash)
	return url, path

}
func init() {

	flag.BoolVar(&h, "h", false, "帮助")
	flag.BoolVar(&l, "l", false, "列出所有版本 JSON格式")
	flag.IntVar(&f, "f", 1, "指定下载源[1.官方下载源 2.bmcl下载源 3.mcbbs下载源] 默认官方")
	flag.IntVar(&p, "p", 50, "下载线程")
	flag.StringVar(&vid, "d", "", "下载指定版本 例如 -v 20w22a")
	flag.Usage = usage
}
func usage() {
	fmt.Fprintf(os.Stderr, `MC Download Version: 0.0.1 beta
Usage: mcdownload [-hl] [-d versionId] [-f 1] [-p 50]

Options:
`)
	flag.PrintDefaults()
}
func lists() interface{} {

	rep, err := http.Get("http://launchermeta.mojang.com/mc/game/version_manifest.json")
	if err != nil {
		fmt.Println(err.Error())
	}
	var versions_json []map[string]string

	b, _ := ioutil.ReadAll(rep.Body)
	versions := gofasion.NewFasion(string(b))
	versions.Get("versions").ArrayForEach(func(i int, version *gofasion.Fasion) {
		_, id := version.Get("id").ValStr()
		_, r_type := version.Get("type").ValStr()
		_, releaseTime := version.Get("releaseTime").ValStr()

		localtime,err:= time.ParseInLocation(time.RFC3339,releaseTime,time.Local)
		if err != nil {
			fmt.Println(err.Error())
		}
		releaseTime = localtime.Format("2006-01-02 15:04:05")
		tmp := map[string]string{
			"releaseTime":releaseTime,
			"versionId":id,
			"type":r_type,
		}
		versions_json = append(versions_json, tmp)
	})
	return versions_json
}
func main() {
	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")
	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 868,
		Title:  "我的世界客户端下载",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(lists)
	app.Run()
	args := os.Args
	if len(args) >= 2 {
		flag.Parse()
		f -= 1
		if h {
			flag.Usage()
			return
		}
		if l {
			rep, err := http.Get(urlReplace("http://launchermeta.mojang.com/mc/game/version_manifest.json", f))
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			var versions_json []map[string]string

			b, _ := ioutil.ReadAll(rep.Body)
			versions := gofasion.NewFasion(string(b))
			versions.Get("versions").ArrayForEach(func(i int, version *gofasion.Fasion) {
				_, id := version.Get("id").ValStr()
				_, r_type := version.Get("type").ValStr()
				_, releaseTime := version.Get("releaseTime").ValStr()

				localtime, err := time.ParseInLocation(time.RFC3339, releaseTime, time.Local)
				if err != nil {
					fmt.Println(err.Error())
				}
				releaseTime = localtime.Format("2006-01-02 15:04:05")
				tmp := map[string]string{
					"releaseTime": releaseTime,
					"versionId":   id,
					"type":        r_type,
				}
				versions_json = append(versions_json, tmp)
			})
			vstr, _ := json.Marshal(versions_json)
			fmt.Println(string(vstr))
			return
		}
		if vid != "" {
			rep, err := http.Get(urlReplace("http://launchermeta.mojang.com/mc/game/version_manifest.json", f))
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			b, _ := ioutil.ReadAll(rep.Body)
			versions := gofasion.NewFasion(string(b))
			versions.Get("versions").ArrayForEach(func(i int, version *gofasion.Fasion) {
				_, id := version.Get("id").ValStr()
				if id == vid {
					_, url := version.Get("url").ValStr()
					url = urlReplace(url, f)
					rep, err = http.Get(url)

					if err != nil {
						fmt.Println(err.Error())
						return
					}
					b, _ = ioutil.ReadAll(rep.Body)
					downloadPlans := DownloadPlans{}
					version := gofasion.NewFasion(string(b))
					_, assetsIndex := version.Get("assetIndex").Get("url").ValStr()
					_, assetsIndexId := version.Get("assetIndex").Get("id").ValStr()
					_, versionIdName := version.Get("id").ValStr()
					err = WriteToFile(fmt.Sprintf("./.minecraft/versions/%s/%s.json", versionIdName, versionIdName), b)
					if err != nil {
						fmt.Println(err.Error())
						return
					}
					_, jarDownloadUrl := version.Get("downloads").Get("client").Get("url").ValStr()
					jarDownload := DownloadPlan{
						Url:      jarDownloadUrl,
						Path:     fmt.Sprintf("./.minecraft/versions/%s/%s.jar", versionIdName, versionIdName),
						Finished: false,
					}
					downloadPlans.addPlan(jarDownload)

					rep, err = http.Get(urlReplace(assetsIndex, f))
					if err != nil {
						fmt.Println(err.Error())
						return
					}
					b, _ = ioutil.ReadAll(rep.Body)
					err = WriteToFile("./.minecraft/assets/indexes/"+assetsIndexId+".json", b)
					if err != nil {
						fmt.Println(err.Error())
						return
					}
					assets := gofasion.NewFasion(string(b))
					objs := assets.Get("objects")
					libs := version.Get("libraries")
					libs.ArrayForEach(func(i int, lib *gofasion.Fasion) {
						lib = lib.Get("downloads")
						_, path := lib.Get("artifact").Get("path").ValStr()
						_, url := lib.Get("artifact").Get("url").ValStr()
						path = "./.minecraft/libraries/" + path
						tmp := DownloadPlan{
							Url:      url,
							Path:     path,
							Finished: false,
						}
						downloadPlans.addPlan(tmp)
						if lib.HasKey("classifiers") {
							_, path = lib.Get("classifiers").Get("natives-windows").Get("path").ValStr()
							_, url = lib.Get("classifiers").Get("natives-windows").Get("url").ValStr()
							path = "./.minecraft/libraries/" + path
							tmp = DownloadPlan{
								Url:      url,
								Path:     path,
								Finished: false,
							}
							downloadPlans.addPlan(tmp)
						}
					})
					for _, m := range objs.Keys() {
						_, hash := objs.Get(m).Get("hash").ValStr()
						url, path := downloadFileUrl(hash, f)
						path = "./.minecraft/assets/objects/" + path
						tmp := DownloadPlan{
							Url:      url,
							Path:     path,
							Finished: false,
						}

						downloadPlans.addPlan(tmp)
						//downloadPlans.downloadPlan = append(downloadPlans.downloadPlan, tmp)

						fmt.Println("加入下载队列", m)

					}
					fmt.Println("下载队列准备完成！")
					fmt.Println("开始下载")
					downloadPlans.startDownloadPlan(p)
					fmt.Println("下载完成")
					return
				}
			})
			fmt.Println("错误版本未找到!")
			return
		}
		fmt.Println("键入 -h 获取帮助!")
		return
	}
	fmt.Println("Minecraft 各版本下载")
	fmt.Println("1. 打印各版本，并按序号下载")
	fmt.Println("2. 直接输入版本号下载")
	sel := 0
	fmt.Print("请输入序号!")
	fmt.Scan(&sel)
	switch sel {
	default:
		fmt.Println("序号输入错误，退出!")
		os.Exit(0)
	case 1:
		fmt.Print("请输入下载源(1.官方下载源 2.bmcl下载源 3.mcbbs下载源): ")
		fmt.Scan(&server)
		server -= 1
		rep, err := http.Get(urlReplace("http://launchermeta.mojang.com/mc/game/version_manifest.json", server))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		downloadPlans := DownloadPlans{}

		b, _ := ioutil.ReadAll(rep.Body)
		versions := gofasion.NewFasion(string(b))
		versions.Get("versions").ArrayForEach(func(i int, version *gofasion.Fasion) {
			_, id := version.Get("id").ValStr()
			_, r_type := version.Get("type").ValStr()
			_, releaseTime := version.Get("releaseTime").ValStr()

			localtime, err := time.ParseInLocation(time.RFC3339, releaseTime, time.Local)
			if err != nil {
				fmt.Println(err.Error())
			}
			releaseTime = localtime.Format("2006-01-02 15:04:05")
			_, url := version.Get("url").ValStr()
			url = urlReplace(url, server)
			fmt.Println(i+1, id, r_type, releaseTime)
			versionsJson = append(versionsJson, url)
		})
		versionId := 0
		fmt.Print("请发送版本序号开始下载!")
		fmt.Scan(&versionId)
		rep, err = http.Get(versionsJson[versionId-1])

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		b, _ = ioutil.ReadAll(rep.Body)

		version := gofasion.NewFasion(string(b))
		_, assetsIndex := version.Get("assetIndex").Get("url").ValStr()
		_, assetsIndexId := version.Get("assetIndex").Get("id").ValStr()
		_, versionIdName := version.Get("id").ValStr()
		err = WriteToFile(fmt.Sprintf("./.minecraft/versions/%s/%s.json", versionIdName, versionIdName), b)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		_, jarDownloadUrl := version.Get("downloads").Get("client").Get("url").ValStr()
		jarDownload := DownloadPlan{
			Url:      jarDownloadUrl,
			Path:     fmt.Sprintf("./.minecraft/versions/%s/%s.jar", versionIdName, versionIdName),
			Finished: false,
		}
		downloadPlans.addPlan(jarDownload)

		rep, err = http.Get(urlReplace(assetsIndex, server))
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		b, _ = ioutil.ReadAll(rep.Body)
		err = WriteToFile("./.minecraft/assets/indexes/"+assetsIndexId+".json", b)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		assets := gofasion.NewFasion(string(b))
		objs := assets.Get("objects")
		libs := version.Get("libraries")
		libs.ArrayForEach(func(i int, lib *gofasion.Fasion) {
			lib = lib.Get("downloads")
			_, path := lib.Get("artifact").Get("path").ValStr()
			_, url := lib.Get("artifact").Get("url").ValStr()
			path = "./.minecraft/libraries/" + path
			tmp := DownloadPlan{
				Url:      url,
				Path:     path,
				Finished: false,
			}
			downloadPlans.addPlan(tmp)
			if lib.HasKey("classifiers") {
				_, path = lib.Get("classifiers").Get("natives-windows").Get("path").ValStr()
				_, url = lib.Get("classifiers").Get("natives-windows").Get("url").ValStr()
				path = "./.minecraft/libraries/" + path
				tmp = DownloadPlan{
					Url:      url,
					Path:     path,
					Finished: false,
				}
				downloadPlans.addPlan(tmp)
			}
		})
		for _, m := range objs.Keys() {
			_, hash := objs.Get(m).Get("hash").ValStr()
			url, path := downloadFileUrl(hash, server)
			path = "./.minecraft/assets/objects/" + path
			tmp := DownloadPlan{
				Url:      url,
				Path:     path,
				Finished: false,
			}

			downloadPlans.addPlan(tmp)
			//downloadPlans.downloadPlan = append(downloadPlans.downloadPlan, tmp)

			fmt.Println("加入下载队列", m)

		}
		fmt.Print("下载队列准备完成！请输入下载线程(推荐30,越高越快,量力而为): ")
		process := 10
		fmt.Scan(&process)
		if process > downloadPlans.Len {
			process = downloadPlans.Len
		}
		fmt.Println("开始下载")
		downloadPlans.startDownloadPlan(process)
		fmt.Println("下载完成")

	case 2:

	}

}

type DownloadPlans struct {
	DownloadPlan []DownloadPlan
	Len          int
	Plans        chan int
	Finished     chan int
	ErrorPlans   chan int
}
type DownloadPlan struct {
	Url      string
	Path     string
	Finished bool
}

func (plans *DownloadPlans) addPlan(plan DownloadPlan) {
	plans.DownloadPlan = append(plans.DownloadPlan, plan)
	plans.Len += 1
}
func (plans *DownloadPlans) startDownloadPlan(process int) {
	plans.Plans = make(chan int, plans.Len)
	plans.Finished = make(chan int, plans.Len)
	plans.ErrorPlans = make(chan int, plans.Len)
	for i := 0; i < plans.Len; i++ {
		plans.Plans <- i
	}
	b, err := json.Marshal(&plans)
	if err == nil {
		ioutil.WriteFile("./mcdownload.json", b, os.ModePerm)
	}
	wg := sync.WaitGroup{}
	for i := 0; i < process; i++ {
		wg.Add(1)
		go plans.downloadFile(&wg)
	}
	wg.Wait()
	for {
		select {
		case v := <-plans.ErrorPlans:
			fmt.Printf("下载错误任务ID（%d），任务下载文件[%s]\n", v, plans.DownloadPlan[v].Url)
		default:
			return
		}

	}

	//a1 := plans.Len
	//num := int(math.Ceil(float64(a1) / float64(process)))
	//all := make(map[int][]int)
	//for i := 0; i < process; i++ {
	//	for j := 0; j < num; j++ {
	//		if (i*num + j) < a1 {
	//			all[i] = append(all[i], i*num+j)
	//		} else {
	//			break
	//		}
	//	}
	//}

	//wg := sync.WaitGroup{}
	//for _, v := range all {
	//	wg.Add(1)
	//	go plans.downloadFile(&wg, v)
	//
	//}
	//wg.Wait()
}
func (plans *DownloadPlans) downloadFile(wg *sync.WaitGroup) {
	for {
		select {
		case v := <-plans.Plans:
			fmt.Printf("ID(%d)开始下载！\n", v)
			url := plans.DownloadPlan[v].Url
			if url == "" {
				plans.Finished <- v
				break
			}
			rep, err := http.Get(url)
			if err != nil {
				fmt.Println(err.Error())
				plans.ErrorPlans <- v
				break
			}
			bt, err := ioutil.ReadAll(rep.Body)
			if err != nil {
				fmt.Println(err.Error())
				plans.ErrorPlans <- v
				break
			}
			err = WriteToFile(plans.DownloadPlan[v].Path, bt)
			if err != nil {
				fmt.Println(err.Error())
				plans.ErrorPlans <- v
				break
			}
			//ioutil.WriteFile(plans.DownloadPlan[v].Path,,os.ModePerm)
			fmt.Printf("ID(%d)下载成功\n", v)
			//fmt.Println("下载成功", plans.DownloadPlan[v].Url, "to", plans.DownloadPlan[v].Path)
			plans.DownloadPlan[v].Finished = true
			plans.Finished <- v
		default:
			wg.Done()
			return
		}
	}

	//
	//for _, v := range planId {
	//	rep,err := http.Get(plans.DownloadPlan[v].Url)
	//	if err != nil{
	//		fmt.Println(err.Error())
	//		return
	//	}
	//	bt,err := ioutil.ReadAll(rep.Body)
	//	if err != nil{
	//		fmt.Println(err.Error())
	//		return
	//	}
	//	err = WriteToFile(plans.DownloadPlan[v].Path,bt)
	//	if err != nil{
	//		fmt.Println(err.Error())
	//		return
	//	}
	//	//ioutil.WriteFile(plans.DownloadPlan[v].Path,,os.ModePerm)
	//
	//	fmt.Println("Download", plans.DownloadPlan[v].Url, "to", plans.DownloadPlan[v].Path)
	//	plans.DownloadPlan[v].Finished = true
	//}
	//

}
func WriteToFile(fileName string, content []byte) error {
	os.MkdirAll(path.Dir(fileName), 0644)
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("错误信息 " + err.Error())
	} else {
		// offset
		//os.Truncate(filename, 0) //clear
		n, _ := f.Seek(0, os.SEEK_END)
		_, err = f.WriteAt(content, n)
		defer f.Close()
	}
	return err
}
