package models

import (
	"regexp"
	"strings"

	"github.com/astaxie/beego/orm"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db orm.Ormer
)

type MovieInfo struct {
	Id                   int64
	Movie_id             int64
	Movie_name           string
	Movie_pic            string
	Movie_director       string
	Movie_writer         string
	Movie_country        string
	Movie_language       string
	Movie_main_character string
	Movie_type           string
	Movie_on_time        string
	Movie_span           string
	Movie_grade          string
}

func init() {
	orm.Debug = true // 是否开启调试模式 调试模式下会打印出sql语句
	// 参数四 设置最大的空闲连接 注册一个别名为default 的数据库 ，作为默认使用
	orm.RegisterDataBase("default", "mysql", "root:root123@tcp(127.0.0.1)/test?charset=utf8", 30)
	orm.RegisterModel(new(MovieInfo))
	db = orm.NewOrm()
}

func GetMovieName(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<span\s*property="v:itemreviewed">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	if len(result) == 0 {
		return ""
	}
	//fmt.Printf("%T", result[0][1])
	//return string(result[0][1])
	return result[0][1]
}

func GetMovieDirector(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	// MustCompile解析并返回一个正则表达式。如果成功返回，该Regexp就可用于匹配文本
	reg := regexp.MustCompile(`<a.*?rel="v:directedBy">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}

func GetMovieMainCharacters(movieHtml string) string {
	if movieHtml == "" {
		return ""
	}
	reg := regexp.MustCompile(`<a.*?rel="v:starring">(.*?)</a>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)
	if len(result) == 0 {
		return ""
	}
	mainCharacters := ""
	for _, v := range result {
		//fmt.Println("======>", v[1])
		mainCharacters += v[1] + "/"
	}

	return strings.Trim(mainCharacters, "/")
}

func GetMovieGenre(movieHtml string) string {
	reg := regexp.MustCompile(`<span.*?property="v:genre">(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0 {
		return ""
	}

	movieGenre := ""
	for _, v := range result {
		movieGenre += v[1] + "/"
	}
	return strings.Trim(movieGenre, "/")
}

func GetMovieOnTime(movieHtml string) string {
	reg := regexp.MustCompile(`<span.*?property="v:initialReleaseDate".*?>(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0 {
		return ""
	}

	return string(result[0][1])
}

func GetMovieRunningTime(movieHtml string) string {
	reg := regexp.MustCompile(`<span.*?property="v:runtime".*?>(.*?)</span>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0 {
		return ""
	}

	return string(result[0][1])
}

func GetMovieUrls(movieHtml string) []string {
	reg := regexp.MustCompile(`<a.*?href="(https://movie.douban.com/.*?)"`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	var movieSets []string
	for _, v := range result {
		movieSets = append(movieSets, v[1])
	}

	return movieSets
}

func AddMovie(movie_info *MovieInfo) (int64, error) {
	movie_info.Id = 0
	id, err := db.Insert(movie_info)
	return id, err
}

func GetMovieGrade(movieHtml string) string {
	reg := regexp.MustCompile(`<strong.*?property="v:average">(.*?)</strong>`)
	result := reg.FindAllStringSubmatch(movieHtml, -1)

	if len(result) == 0 {
		return ""
	}
	return string(result[0][1])
}
